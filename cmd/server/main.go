package main

import (
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"serverfn/internal/config"
	"serverfn/internal/repository/psql"
	"serverfn/internal/service"
	"serverfn/internal/transport/rest"
	"serverfn/pkg/database"
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

func main() {
	//cfg, err := config.Load()
	//if err != nil {
	//	logrus.Fatalf("Failed to load configuration: %v", err)
	//}
	cfg1, err := config.New(CONFIG_DIR, CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     cfg1.DB.Host,
		Port:     cfg1.DB.Port,
		Username: cfg1.DB.Username,
		DBName:   cfg1.DB.Name,
		SSLMode:  cfg1.DB.SSLMode,
		Password: cfg1.DB.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//init deps

	repo := psql.NewRepositories(db)
	services := service.NewServices(repo)
	handler := rest.NewHandler(services)
	//taskRepo := psql.NewTasksRepository(db)
	//taskService := service.NewTask(taskRepo)

	//handler := rest.NewHandler(taskService)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg1.Server.Port),
		Handler: handler.InitRouter(),
	}

	log.Info("Listening on port 8080")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	//taskQueue := make(chan *task.Task, cfg.QueueSize)
	//taskManager := task.NewManagerImpl(taskQueue, cfg.WorkerPoolSize, logger)
	//server := api.NewServer(taskManager, logger)
	//
	//go taskManager.Start()
	//
	//srv := &http.Server{
	//	Addr:    cfg.ServerAddress,
	//	Handler: server.Router(),
	//}
	//
	//go func() {
	//	logger.Infof("Starting server on %s", cfg.ServerAddress)
	//	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//		logger.Fatalf("Could not listen on %s: %v", cfg.ServerAddress, err)
	//	}
	//}()
	//
	//gracefulShutdown(srv, taskManager, logger)
}

//func gracefulShutdown(srv *http.Server, tm task.Manager, logger *logrus.Logger) {
//	stop := make(chan os.Signal, 1)
//	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
//	<-stop
//
//	logger.Info("Shutting down server...")
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	if err := srv.Shutdown(ctx); err != nil {
//		logger.Fatalf("Server forced to shutdown: %v", err)
//	}
//
//	tm.Stop()
//
//	logger.Info("Server exiting")
//}
