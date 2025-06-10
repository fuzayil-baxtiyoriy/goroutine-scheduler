package scheduler

import (
	"sync"
	"time"
)

type Worker struct {
	id        int
	queue     []Job
	qMutex    sync.Mutex
	scheduler *Scheduler
}

func NewWorker(id int, scheduler *Scheduler) *Worker {
	return &Worker{
		id:        id,
		queue:     []Job{},
		scheduler: scheduler,
	}
}

func (w *Worker) Run() {
	for {
		job := w.getJob()
		if job == nil {
			job = w.stealJob()
			if job == nil {
				time.Sleep(100 * time.Microsecond)
			}
		}
		job.Run()
	}
}

func (w *Worker) AddJob(job Job) {
	w.qMutex.Lock()
	defer w.qMutex.Unlock()
	w.queue = append(w.queue, job)
}

func (w *Worker) getJob() *Job {
	w.qMutex.Lock()
	defer w.qMutex.Unlock()
	if len(w.queue) == 0 {
		return nil
	}
	job := w.queue[0]
	w.queue = w.queue[1:]
	return &job
}

func (w *Worker) stealJob() *Job {
	for _, other := range w.scheduler.workers {
		if other.id == w.id {
			continue
		}
		other.qMutex.Lock()
		if len(other.queue) == 0 {
			other.qMutex.Unlock()
			continue
		}
		job := other.queue[len(other.queue)-1]
		other.queue = other.queue[:len(other.queue)-1]
		other.qMutex.Unlock()
		return &job
	}
	return nil
}
