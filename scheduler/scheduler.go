package scheduler

import (
	"fmt"
	"sync"
)

type Scheduler struct {
	workers     []*Worker
	numWorkers  int
	assignIndex int
	mu          sync.Mutex
}

func NewScheduler(numWorkers int) *Scheduler {
	s := &Scheduler{
		numWorkers: numWorkers,
		workers:    make([]*Worker, 0, numWorkers),
		
	}
	for i := 0; i < numWorkers; i++ {
		w := NewWorker(i, s)
		s.workers = append(s.workers, w)
	}
	return s
}

func (s *Scheduler) Start() {
	for _, w := range s.workers {
		go w.Run()
	}
}

func (s *Scheduler) AddJob(job Job) {
	// Round-robin job assignment
	s.mu.Lock()
	target := s.workers[s.assignIndex%len(s.workers)]
	fmt.Println("target", target)
	s.assignIndex++
	s.mu.Unlock()
	target.AddJob(job)
}
