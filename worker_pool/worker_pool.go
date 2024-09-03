package main

import (
	"sync"
)

type Task interface {
	process()
}

type WorkerPool struct {
	Task        []Task
	concurrency int
	taskChain   chan Task
	wg          *sync.WaitGroup
}

func (wp *WorkerPool) worker() {
	for task := range wp.taskChain {
		task.process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	wp.wg = &sync.WaitGroup{}
	wp.taskChain = make(chan Task, len(wp.Task))

	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}

	for _, task := range wp.Task {
		wp.wg.Add(1)
		wp.taskChain <- task
	}
	close(wp.taskChain)

	wp.wg.Wait()
}
