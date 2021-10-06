package worker

type workerManager struct {
	workerNumber int
	jobs         chan func()
}

func WorkerManager(Numb int) *workerManager {
	return &workerManager{
		workerNumber: Numb,
	}
}

func (wm *workerManager) AddJob(job func()) {
	wm.jobs <- job
}

func (wm *workerManager) Run() {
	for i := 0; i < wm.workerNumber; i++ {
		go func() {
			for jb := range wm.jobs {
				wk := &worker{job: jb}
				wk.do()
			}
		}()
	}
}
