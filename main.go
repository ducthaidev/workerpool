package main

import (
	"fmt"
	"thai/worker"
)

func main() {
	manager := worker.WorkerManager(1)
	manager.Run()
	manager.AddJob(func() { fmt.Println("hello") })
}
