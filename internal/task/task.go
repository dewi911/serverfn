//package task
//
//
//import "net/http"
//
//type TaskStatus string
//
//const (
//	TaskStatusNew       TaskStatus = "new"
//	TaskStatusInProcess TaskStatus = "in_process"
//	TaskStatusDone      TaskStatus = "done"
//	TaskStatusError     TaskStatus = "error"
//)
//
//type Task struct {
//	ID              string
//	Status          TaskStatus
//	Method          string
//	URL             string
//	Headers         map[string]string
//	HTTPStatusCode  int
//	ResponseHeaders http.Header
//	ResponseLength  int64
//	Error           string
//}
