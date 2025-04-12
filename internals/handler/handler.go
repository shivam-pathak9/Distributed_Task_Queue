package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"Distributed_Task_Queue/internals/broker"
	"Distributed_Task_Queue/internals/task"
)

type TaskRequest struct {
	Type       string            `json:"type"`
	Payload    map[string]string `json:"payload"`
	MaxRetries int               `json:"max_retries"`
}

// APIHandler wraps broker so we can use it in routes
type APIHandler struct {
	Broker *broker.Broker
}

func NewAPIHandler(b *broker.Broker) *APIHandler {
	return &APIHandler{Broker: b}
}

// SubmitTaskHandler handles POST /task
func (h *APIHandler) SubmitTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Type == "" {
		http.Error(w, "Missing task type", http.StatusBadRequest)
		return
	}

	if req.MaxRetries < 0 {
		req.MaxRetries = 0
	}

	t := task.NewTask(req.Type, req.Payload, req.MaxRetries)
	log.Printf("[API] Received Task: %+v\n", t)

	h.Broker.Enqueue(t)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Task enqueued",
		"id":      t.ID,
	})
}
