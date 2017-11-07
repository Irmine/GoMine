package tasks

type ITask interface {
	OnRun() // Every task MUST implement the OnRun function.
	OnInit()
}

type Task struct {
	executed bool
}

func (task Task) isExecuted() bool {
	return task.executed
}