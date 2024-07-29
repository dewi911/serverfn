package api

//
//import (
//	"github.com/gorilla/mux"
//	"github.com/sirupsen/logrus"
//	"serverfn/internal/task"
//)
//
//type Server struct {
//	taskManager task.Manager
//	logger      *logrus.Logger
//}
//
//func NewServer(taskManager task.Manager, logger *logrus.Logger) *Server {
//	return &Server{
//		taskManager: taskManager,
//		logger:      logger,
//	}
//}
//
//func (s *Server) Router() *mux.Router {
//	r := mux.NewRouter()
//	r.HandleFunc("/task", s.CreateTask).Methods("POST")
//	r.HandleFunc("/task/{taskId}", s.GetTaskStatus).Methods("GET")
//	return r
//}
