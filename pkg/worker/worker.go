//package worker
//
//import (
//	log "github.com/sirupsen/logrus"
//	"io"
//	"net/http"
//	"serverfn/internal/domain"
//	"time"
//)
//
//type Worker struct {
//	taskQueue <-chan *domain.Task
//	quit      chan bool
//	logger    *log.Logger
//}
//
//func NewWorker(taskQueue <-chan *domain.Task, logger *log.Logger) *Worker {
//	return &Worker{
//		taskQueue: taskQueue,
//		quit:      make(chan bool),
//		logger:    logger,
//	}
//}
//
//func (w *Worker) Start() {
//	for {
//		select {
//		case task := <-w.taskQueue:
//			w.processTask(task)
//		case <-w.quit:
//			return
//		}
//	}
//}
//
//func (w *Worker) Stop() {
//	w.quit <- true
//}
//
//func (w *Worker) processTask(task *domain.Task) {
//	task.Status = domain.TaskStatusInProcess
//
//	client := &http.Client{
//		Timeout: time.Second * 30,
//	}
//
//	req, err := http.NewRequest(task.Method, task.URL, nil)
//	if err != nil {
//		w.logger.WithError(err).Error("Error creating request")
//		task.Status = domain.TaskStatusError
//		task.Error = err.Error()
//		return
//	}
//
//	for key, value := range task.Headers {
//		req.Header.Set(key, value)
//	}
//
//	resp, err := client.Do(req)
//	if err != nil {
//		w.logger.WithError(err).Error("Error executing request")
//		task.Status = domain.TaskStatusError
//		task.Error = err.Error()
//		return
//	}
//	defer resp.Body.Close()
//
//	task.HTTPStatusCode = resp.StatusCode
//	task.ResponseHeaders = resp.Header
//	task.ResponseLength = resp.ContentLength
//
//	_, err = io.Copy(io.Discard, resp.Body)
//	if err != nil {
//		w.logger.WithError(err).Warn("Error reading response body")
//	}
//
//	task.Status = domain.TaskStatusDone
//}
