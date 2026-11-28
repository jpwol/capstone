package engine

import "context"

type Scheduler interface {
	Schedule(task *TaskNode) error
	Next(ctx context.Context, workerID int) (*TaskNode, error)
	Shutdown()
}
