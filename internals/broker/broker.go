package broker

import (
	"Distributed_Task_Queue/internals/task"
)

type Broker struct {
	queue chan *task.Task
}

// NewBroker creates a broker with a fixed-size channel
func NewBroker(bufferSize int) *Broker {
	return &Broker{
		queue: make(chan *task.Task, bufferSize),
	}
}

// Enqueue adds a task to the broker queuesssss
func (b *Broker) Enqueue(t *task.Task) {
	b.queue <- t
}

// Dequeue returns a channel to listen for tasks
func (b *Broker) Dequeue() <-chan *task.Task {
	return b.queue
}
