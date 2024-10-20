package manager

import (
	"fmt"
	"github.com/MHG14/Containestrator/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]*task.Task
	EventDb       map[string][]*task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("i will select a worker")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("i will update tasks")
}

func (m *Manager) SendWork() {
	fmt.Println("i will send work to workers")
}
