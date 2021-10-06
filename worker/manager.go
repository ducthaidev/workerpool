package worker

import (
	"os"
	"strconv"
)

type workerManager struct {
	SecondDuration int
	RetryTimes     int
	workerNumber   int
	jobs           chan *Job
}

func WorkerManager(WorkerNumb int) *workerManager {
	return &workerManager{
		workerNumber: WorkerNumb,
		jobs:         make(chan *Job),
	}
}

func (wm *workerManager) AddJob(job *Job) {
	wm.jobs <- job
}

func (wm *workerManager) Run() {
	rtrT, _ := strconv.ParseInt(os.Getenv("RETRY_TIMES"), 10, 32)
	scndD, _ := strconv.ParseInt(os.Getenv("SECOND_DURATION"), 10, 32)
	wm.RetryTimes = int(rtrT)
	wm.SecondDuration = int(scndD)

	for i := 0; i < wm.workerNumber; i++ {
		go func() {
			for jb := range wm.jobs {
				wk := &worker{job: jb, WorkerManager: wm}
				wk.do()
			}
		}()
	}
}
