package worker

import (
	"errors"
	"math/rand"
	"time"

	"Distributed_Task_Queue/internals/task"
)

// Simulates task execution
func executeTask(t *task.Task) error {
	time.Sleep(time.Duration(rand.Intn(500)+200) * time.Millisecond)

	// Simulate failure for "email" task randomly
	if t.Type == "email" && rand.Float32() < 0.3 {
		return errors.New("simulated email failure")
	}

	// Simulate success for others
	return nil
}
