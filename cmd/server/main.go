package main

import (
	"log"
	"net/http"

	"Distributed_Task_Queue/internals/broker"
	"Distributed_Task_Queue/internals/handler"
	"Distributed_Task_Queue/internals/worker"
)

func main() {
	// Config
	numWorkers := 5 // Total number which can serve the request at a time
	queueSize := 5  // Total request which can handle by worker at a time

	// Setup broker
	br := broker.NewBroker(queueSize)

	// Setup worker pool
	wp := worker.NewWorkerPool("default", numWorkers, br.Dequeue())
	wp.Start()

	// Setup API handler
	h := handler.NewAPIHandler(br)

	// Routes
	http.HandleFunc("/task", h.SubmitTaskHandler)

	log.Println("[Main] Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	// Optional: block on worker pool if we wanted to wait
	// wp.Wait()
}
