package worker

type worker struct {
	StateRunning     bool `default:"false"`
	StateFailed      bool `default:"false"`
	StateTimeout     bool `default:"false"`
	StateCompleted   bool `default:"false"`
	StateRetryFailed bool `default:"false"`
	job func()
}

func (w *worker) do() {
	w.job()
}