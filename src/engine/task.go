package engine

import "context"

type TaskState int

const (
	Pending TaskState = iota
	Ready
	Running
	Completed
	Failed
	Blocked
)

type Completion struct {
	ID NodeId
	Err    error
}

type TaskFunc func(context.Context) error

type Task struct {
	Name string
	Run  TaskFunc
}

type TaskNode struct {
	ID                    NodeId
	Task                  Task
	State                 TaskState
	remainingDependencies int
}
