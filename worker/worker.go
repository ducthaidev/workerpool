package worker

import (
	"log"
	"time"
)

type worker struct {
	WorkerManager    *workerManager
	StateRunning     bool `default:"false"`
	StateFailed      bool `default:"false"`
	StateTimeout     bool `default:"false"`
	StateCompleted   bool `default:"false"`
	StateRetryFailed bool `default:"false"`
	RetryCounter     int  `default:"0"`
	job              *Job
}

type Job struct {
	State bool        `default:"false"`
	Work  func() bool `default:"false"`
}

func (w *worker) do() {
	state := w.job.Work()
	w.StateRunning = true
	if state {
		log.Println("JOB IS SUCCESS!")
		w.StateRunning = false
		w.job.State = true
		w.StateCompleted = true
	} else {
		log.Println("JOB FAILT", w.RetryCounter + 1, "TIMES")
		w.StateCompleted = true
		w.StateFailed = true
		w.retry()
	}
}

func (w *worker) retry() {
	w.RetryCounter++
	if w.RetryCounter > w.WorkerManager.RetryTimes {
		w.StateRetryFailed = true
		log.Println("JOB END")
		return
	}
	log.Println("RETRYING")
	time.Sleep(time.Second * time.Duration(w.WorkerManager.SecondDuration))
	w.do()
}
