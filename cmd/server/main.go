package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"serverfn/internal/api"
	"serverfn/internal/config"
	"serverfn/internal/task"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		logrus.Fatalf("Failed to load configuration: %v", err)
	}

	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)

	taskQueue := make(chan *task.Task, cfg.QueueSize)
	taskManager := task.NewManagerImpl(taskQueue, cfg.WorkerPoolSize, logger)
	server := api.NewServer(taskManager, logger)

	go taskManager.Start()

	srv := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: server.Router(),
	}

	go func() {
		logger.Infof("Starting server on %s", cfg.ServerAddress)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Could not listen on %s: %v", cfg.ServerAddress, err)
		}
	}()

	gracefulShutdown(srv, taskManager, logger)
}

func gracefulShutdown(srv *http.Server, tm task.Manager, logger *logrus.Logger) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}

	tm.Stop()

	logger.Info("Server exiting")
}
