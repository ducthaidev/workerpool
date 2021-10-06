package main

import (
	"fmt"
	"log"
	"thai/worker"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	wc := make(chan bool)
	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()

	manager := worker.WorkerManager(0)
	manager.Run()
	go manager.AddJob(&worker.Job{Work: func() bool {
		fmt.Println("test Job")
		time.Sleep(time.Second * 3)
		return true
	},
	})
	go manager.AddJob(&worker.Job{Work: func() bool {
		fmt.Println("Fail Job")
		return false
	},
	})
	go manager.AddJob(&worker.Job{Work: func() bool {
		fmt.Println("Success1 Job")
		return true
	},
	})
	go manager.AddJob(&worker.Job{Work: func() bool {
		fmt.Println("Success2 Job")
		time.Sleep(time.Second * 2)
		return true
	},
	})
	go manager.AddJob(&worker.Job{Work: func() bool {
		fmt.Println("Success3 Job")
		return true
	},
	})

	<-wc
}
