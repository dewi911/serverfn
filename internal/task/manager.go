package task

//
//
//import (
//	"sync"
//
//	"github.com/google/uuid"
//	"github.com/sirupsen/logrus"
//)
//
//type Manager interface {
//	Start()
//	Stop()
//	CreateTask(method, url string, headers map[string]string) *Task
//	GetTask(id string) (*Task, bool)
//	RemoveTask(id string)
//}
//
//type manager struct {
//	tasks     map[string]*Task
//	mutex     sync.RWMutex
//	taskQueue chan *Task
//	workers   []*Worker
//	logger    *logrus.Logger
//	wg        sync.WaitGroup
//}
//
//func NewManagerImpl(taskQueue chan *Task, workerCount int, logger *logrus.Logger) Manager {
//	m := &manager{
//		tasks:     make(map[string]*Task),
//		taskQueue: taskQueue,
//		logger:    logger,
//	}
//
//	m.workers = make([]*Worker, workerCount)
//	for i := 0; i < workerCount; i++ {
//		m.workers[i] = NewWorker(m.taskQueue, logger)
//	}
//
//	return m
//}
//
//func (m *manager) Start() {
//	for _, worker := range m.workers {
//		m.wg.Add(1)
//		go func(w *Worker) {
//			defer m.wg.Done()
//			w.Start()
//		}(worker)
//	}
//}
//
//func (m *manager) Stop() {
//	for _, worker := range m.workers {
//		worker.Stop()
//	}
//	m.wg.Wait()
//	m.logger.Info("All workers have stopped")
//}
//
//func (m *manager) CreateTask(method, url string, headers map[string]string) *Task {
//	task := &Task{
//		ID:      uuid.New().String(),
//		Status:  TaskStatusNew,
//		Method:  method,
//		URL:     url,
//		Headers: headers,
//	}
//
//	m.mutex.Lock()
//	m.tasks[task.ID] = task
//	m.mutex.Unlock()
//
//	m.taskQueue <- task
//
//	return task
//}
//
//func (m *manager) GetTask(id string) (*Task, bool) {
//	m.mutex.RLock()
//	defer m.mutex.RUnlock()
//	task, exists := m.tasks[id]
//	return task, exists
//}
//
//func (m *manager) RemoveTask(id string) {
//	m.mutex.Lock()
//	defer m.mutex.Unlock()
//	delete(m.tasks, id)
//}
