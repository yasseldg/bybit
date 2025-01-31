package common

import (
	"sync"

	"github.com/yasseldg/go-simple/logs/sLog"
)

type workerPool struct {
	workers int
	jobs    chan []byte
	process func([]byte)
	scale   chan int
	stop    chan struct{}
	wg      sync.WaitGroup
}

func newWorkerPool(workers int, jobs chan []byte, process func([]byte)) *workerPool {
	return &workerPool{
		workers: workers,
		jobs:    jobs,
		process: process,
		scale:   make(chan int),
		stop:    make(chan struct{}),
	}
}

func (wp *workerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}

	sLog.Info("WebSocket starting with %d workers", wp.workers)

	go wp.scalingManager()
}

// Escala la cantidad de workers
func (wp *workerPool) Scale(n int) {
	wp.scale <- n
}

// Cierra el pool
func (wp *workerPool) Shutdown() {
	close(wp.jobs)
	close(wp.scale)
	wp.wg.Wait()
}

// private methods

func (wp *workerPool) worker(id int) {
	defer wp.wg.Done()
	for {
		select {
		case job, ok := <-wp.jobs:
			if !ok {
				sLog.Info("Worker %d stopping (channel closed)", id)
				return
			}
			wp.process(job)

		case <-wp.stop:
			sLog.Info("Worker %d stopping (signal received)", id)
			return
		}
	}
}

func (wp *workerPool) scalingManager() {
	for newWorkerCount := range wp.scale {
		if newWorkerCount > wp.workers {
			// Agregar más workers
			diff := newWorkerCount - wp.workers
			for i := 0; i < diff; i++ {
				wp.wg.Add(1)
				go wp.worker(wp.workers + i)
			}
		} else if newWorkerCount < wp.workers {
			// Reducir workers enviando señales de parada
			diff := wp.workers - newWorkerCount
			for i := 0; i < diff; i++ {
				wp.stop <- struct{}{}
			}
		}
		wp.workers = newWorkerCount

		sLog.Info("Scaling to %d workers", newWorkerCount)
	}
}
