package worker

import (
	"log"
	"sync"
	"time"

	"Distributed_Task_Queue/internals/task"
)

// WorkerPool manages multiple workers to process tasks
type WorkerPool struct {
	ID          string
	NumWorkers  int
	TaskChannel <-chan *task.Task
	wg          sync.WaitGroup
}

// NewWorkerPool initializes a new worker pool
func NewWorkerPool(id string, numWorkers int, taskChan <-chan *task.Task) *WorkerPool {
	return &WorkerPool{
		ID:          id,
		NumWorkers:  numWorkers,
		TaskChannel: taskChan,
	}
}

// Start begins the worker pool
func (wp *WorkerPool) Start() {
	log.Printf("[WorkerPool-%s] Starting %d workers", wp.ID, wp.NumWorkers)

	for i := 0; i < wp.NumWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

// Wait blocks until all workers are done
func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}

// Actual worker function
func (wp *WorkerPool) worker(workerID int) {
	defer wp.wg.Done()
	log.Printf("[Worker-%d] Started", workerID)

	for t := range wp.TaskChannel {
		log.Printf("[Worker-%d] Processing task: %s", workerID, t.ID)
		t.Status = task.Running

		err := executeTask(t)
		if err != nil {
			log.Printf("[Worker-%d] Task %s failed: %v", workerID, t.ID, err)
			t.RetryCount++
			if t.RetryCount <= t.MaxRetries {
				log.Printf("[Worker-%d] Retrying task %s (%d/%d)", workerID, t.ID, t.RetryCount, t.MaxRetries)
				t.Status = task.Retrying
				time.Sleep(1 * time.Second) // basic retry delay
				// re-queueing could be handled via broker.Enqueue again (we'll wire this in main)
			} else {
				t.Status = task.Failed
				log.Printf("[Worker-%d] Task %s permanently failed", workerID, t.ID)
			}
		} else {
			t.Status = task.Success
			log.Printf("[Worker-%d] Task %s completed successfully", workerID, t.ID)
		}
	}
}
