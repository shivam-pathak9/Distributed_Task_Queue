package queue

import (
	"Distributed_Task_Queue/internals/task"
	"container/heap"
	"sync"
	"time"
)

type TaskItem struct {
	Task     *task.Task
	Priority int
	Index    int // required for updating/removal
	Created  time.Time
}

type PriorityQueue []*TaskItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].Priority == pq[j].Priority {
		return pq[i].Created.Before(pq[j].Created)
	}
	return pq[i].Priority < pq[j].Priority // lower is higher priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*TaskItem)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for GC
	*pq = old[0 : n-1]
	return item
}

type SafeQueue struct {
	queue PriorityQueue
	lock  sync.Mutex
	cond  *sync.Cond
}

func NewSafeQueue() *SafeQueue {
	pq := make(PriorityQueue, 0)
	sq := &SafeQueue{
		queue: pq,
	}
	sq.cond = sync.NewCond(&sq.lock)
	heap.Init(&sq.queue)
	return sq
}

func (sq *SafeQueue) Enqueue(t *task.Task, priority int) {
	sq.lock.Lock()
	defer sq.lock.Unlock()

	item := &TaskItem{
		Task:     t,
		Priority: priority,
		Created:  time.Now(),
	}
	heap.Push(&sq.queue, item)
	sq.cond.Signal() // notify Dequeue
}

func (sq *SafeQueue) Dequeue() *task.Task {
	sq.lock.Lock()
	defer sq.lock.Unlock()

	for sq.queue.Len() == 0 {
		sq.cond.Wait() // block until task is available
	}

	item := heap.Pop(&sq.queue).(*TaskItem)
	return item.Task
}
