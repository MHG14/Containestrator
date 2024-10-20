package worker

import (
	"fmt"
	"github.com/MHG14/Containestrator/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

func (w *Worker) Start() {
	fmt.Println("i will start a task")
}

func (w *Worker) Stop() {
	fmt.Println("i will stop a task")
}

func (w *Worker) RunTask() {
	fmt.Println("i will run a task")
}

func (w *Worker) CollectStats() {
	fmt.Println("i will collect stats")
}
