package models

type TaskManager interface {
	Start()
	Stop()
	CreateTask(task *Task)
}
