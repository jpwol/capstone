package engine

import "context"

type Worker struct {
	ID          int
	completions chan<- Completion
}

// TODO: workers need a loop to keep processing. consider a separate function Run() that will have this loop and call Execute()
//
// Also, this is a generic worker. We could have this worker operate with both a centralized and work-stealing scheduler.
// The work stealing scheduler will hold an array of deques for all workers, and handle the stealing.
// Alternatively, we can have separate worker structures for the separate schedulers, which will allow us to better tune the scheduling.
// somethin to think about

func (w *Worker) Execute(ctx context.Context, task *TaskNode) {
	task.State = Running
	err := task.Task.Run(ctx)

	w.completions <- Completion{
		ID:  task.ID,
		Err: err,
	}
}
