package broker

import (
	"Distributed_Task_Queue/internals/queue"
	"Distributed_Task_Queue/internals/task"
)

type Broker struct {
	queue *queue.SafeQueue
}

// NewBroker creates a broker with a fixed-size channel
func NewBroker(bufferSize int) *Broker {
	return &Broker{
		queue: queue.NewSafeQueue(),
	}
}

// Enqueue adds a task to the broker queuesssss
func (b *Broker) Enqueue(t *task.Task, priority int) {
	b.queue.Enqueue(t, priority)
}

// Dequeue returns a channel to listen for tasks
func (b *Broker) Dequeue() *task.Task {
	return b.queue.Dequeue()
}
