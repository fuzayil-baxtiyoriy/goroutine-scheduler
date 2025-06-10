package main

import (
	"fmt"
	"goroutine-scheduler/scheduler"
	"time"
)

func main() {
	s := scheduler.NewScheduler(3) // 3 workers

	for i := 1; i <= 10; i++ {
		id := i
		s.AddJob(scheduler.Job{
			ID: id,
			Work: func() {
				fmt.Printf("Executing work for job #%d\n", id)
				time.Sleep(1 * time.Second)
			},
		})
	}
	fmt.Println("Starting Work-Stealing Scheduler...")
	s.Start()
	time.Sleep(10 * time.Second) // let workers finish
}
