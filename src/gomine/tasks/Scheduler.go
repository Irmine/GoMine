package tasks

type Scheduler struct {
	delayedTasks map[int]Task
	repeatingTasks map[int]Task
}

func NewScheduler() Scheduler {
	return Scheduler{make(map[int]Task), make(map[int]Task)}
}

func (scheduler *Scheduler) DoTick() bool {

}

func (scheduler *Scheduler) scheduleDelayedTask(task Task, ticksDelay int) bool {

}
