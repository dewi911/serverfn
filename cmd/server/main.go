package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dewi911/cruda-app/pkg/database"
	_ "github.com/dewi911/serverfn/docs"
	"github.com/dewi911/serverfn/internal/config"
	"github.com/dewi911/serverfn/internal/models"
	"github.com/dewi911/serverfn/internal/repository/psql"
	"github.com/dewi911/serverfn/internal/service"
	"github.com/dewi911/serverfn/internal/transport/rest"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "main"
)

// @title Server API
// @version 1.0
// @description API server with task management.
// @host localhost:8080
// @BasePath /
func main() {
	cfg, err := config.New(CONFIG_DIR, CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		DBName:   cfg.DB.Name,
		SSLMode:  cfg.DB.SSLMode,
		Password: cfg.DB.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//init deps
	repo := psql.NewTasksRepository(db)
	tm := service.NewTaskManager(cfg.QueueCapacity, cfg.WorkerPoolSize, repo, log.New())
	services := service.NewTasksService(repo, tm)
	handler := rest.NewHandler(services)

	//start manager
	go tm.Start()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: handler.InitRouter(),
	}

	go func() {
		log.Infof("Starting server on %d", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v", cfg.Server, cfg.Server.Port)
		}
	}()

	gracefulShutdown(srv, tm, db, log.New())
	log.Info("Main: exiting")
}

func gracefulShutdown(srv *http.Server, tm models.TaskManager, db *sql.DB, logger *log.Logger) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	logger.Info("Initiating graceful shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	logger.Info("Shutting down HTTP server...")
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}
	logger.Info("HTTP server stopped")

	logger.Info("Stopping task manager...")
	tm.Stop()
	logger.Info("Task manager stopped")

	logger.Info("Closing database connection...")
	if err := db.Close(); err != nil {
		logger.Fatalf("Error closing database connection: %v", err)
	}
	logger.Info("Database connection closed")

	logger.Info("Graceful shutdown completed")
}
