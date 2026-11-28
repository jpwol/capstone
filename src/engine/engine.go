package engine


type Engine struct {
	graph *Graph
	scheduler Scheduler
	completions chan Completion
}

func NewEngine(graph *Graph, scheduler Scheduler) *Engine {
	return &Engine{
		graph: graph,
		scheduler: scheduler,
		completions: make(chan Completion),
	}
}

func (e *Engine) processCompletions() {
	for completion := range e.completions {
		if completion.Err != nil {
			e.graph.FailTask(completion.ID, completion.Err)
			continue
		}

		ready, err := e.graph.CompleteTask(completion.ID)
		if err != nil {
			// handle failed task here
			continue
		}

		for _, task := range ready {
			_ = task
			// schedule ready tasks here
		}
	}
}
