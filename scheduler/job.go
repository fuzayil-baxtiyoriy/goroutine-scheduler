package scheduler

import "fmt"

type Job struct {
	ID   int
	Work func()
}

func (j Job) Run() {
	fmt.Printf("Running job #%d\n", j.ID)
	j.Work()
}
