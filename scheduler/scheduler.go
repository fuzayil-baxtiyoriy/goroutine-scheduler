package scheduler

import (
	"fmt"
	"time"
)

type Scheduler struct {
	Queue    []Job
	TimeSlot time.Duration
}

func NewScheduler(slot time.Duration) *Scheduler {
	return &Scheduler{
		Queue:    []Job{},
		TimeSlot: slot,
	}
}

func (s *Scheduler) AddJob(job Job) {
	s.Queue = append(s.Queue, job)
}

func (s *Scheduler) StartRoundRobin() {
	for len(s.Queue) > 0 {
		job := s.Queue[0]
		s.Queue = s.Queue[1:] // dequeue

		go job.Run()
		time.Sleep(s.TimeSlot)
		fmt.Printf("Finished job #%d\n", job.ID)
	}
}
