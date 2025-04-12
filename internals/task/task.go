package task

import (
	"time"

	"github.com/google/uuid"
)

// Status of task
type Status string

const (
	Pending  Status = "PENDING"
	Running  Status = "RUNNING"
	Success  Status = "SUCCESS"
	Failed   Status = "FAILED"
	Retrying Status = "RETRYING"
)

// Task represents a single unit of work
type Task struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	Payload    map[string]string `json:"payload"` // Can be JSON key-values
	RetryCount int               `json:"retry_count"`
	MaxRetries int               `json:"max_retries"`
	Status     Status            `json:"status"`
	CreatedAt  time.Time         `json:"created_at"`
	Priority   int               `json:"priority"` // Priority for task execution
}

// NewTask creates a new task with a generated ID
func NewTask(taskType string, payload map[string]string, maxRetries int, priority int) *Task {
	return &Task{
		ID:         uuid.NewString(),
		Type:       taskType,
		Payload:    payload,
		RetryCount: 0,
		MaxRetries: maxRetries,
		Status:     Pending,
		CreatedAt:  time.Now(),
		Priority:   priority,
	}
}
