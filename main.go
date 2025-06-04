package main

import (
	"fmt"
	"goroutine-scheduler/scheduler"
	"time"
)

func main() {
	s := scheduler.NewScheduler(2 * time.Second)

	for i := 1; i <= 5; i++ {
		id := i
		s.AddJob(scheduler.Job{
			ID: id,
			Work: func() {
				fmt.Printf("Executing work for job #%d\n", id)
				time.Sleep(1 * time.Second)
			},
		})
	}
	fmt.Println("Starting Round-Robin Scheduler...")
	s.StartRoundRobin()
}
