# Distributed Task Queue

A distributed task queue system that allows clients to submit tasks via a REST API, processes tasks using a worker pool, and supports features like task prioritization, retries, and concurrency.

---

## Features

| Feature                     | Included in MVP | Description                                                                 |
|-----------------------------|-----------------|-----------------------------------------------------------------------------|
| REST API to submit tasks    | ✅              | Enables external clients to submit tasks to the queue.                     |
| In-memory broker            | ✅              | Keeps implementation simple without requiring a database.                  |
| Worker pool                 | ✅              | Processes tasks concurrently using a pool of workers.                      |
| Retry once on failure       | ✅              | Provides basic fault tolerance by retrying failed tasks once.              |
| CLI logging for monitoring  | ✅              | Logs task submissions and processing for easy debugging and monitoring.    |
| JSON payload + task type    | ✅              | Supports different task behaviors through JSON payloads and task types.    |
| Task prioritization         | ✅              | Allows tasks to be processed based on priority levels.                     |
| Concurrency in worker pool  | ✅              | Handles multiple tasks in parallel using a configurable number of workers. |

---

## High-Level Design

The system consists of the following components:

1. **REST API**:
   - Exposes an endpoint (`/task`) to submit tasks.
   - Accepts a JSON payload with task details, including type, payload, retries, and priority.

2. **In-Memory Broker**:
   - Acts as a queue to store tasks.
   - Supports task prioritization to ensure higher-priority tasks are processed first.

3. **Worker Pool**:
   - A pool of workers that process tasks concurrently.
   - Configurable number of workers and queue size.

4. **Task Retry**:
   - Automatically retries a task once if it fails during processing.

5. **Logging**:
   - Logs task submissions, processing, and errors to the CLI for monitoring.

---

## API Endpoints

### Submit Task
**Endpoint**: `POST /task`

**Request Body**:
```json
{
    "type": "example_task",
    "payload": {
        "key": "value"
    },
    "max_retries": 3,
    "priority": 2
}
```
**Request Body**:
```json
{
    "message": "Task enqueued",
    "id": "unique-task-id"
}
```
---
## Description:

type: The type of task to be processed.
payload: A map of key-value pairs containing task-specific data.
max_retries: The maximum number of retries allowed for the task.
priority: The priority of the task (higher values indicate higher priority).

---
## Configuration
The system can be configured in the main.go file:

Number of Workers: Set the numWorkers variable to define the number of workers in the pool.
Queue Size: Set the queueSize variable to define the maximum number of tasks the broker can hold.

numWorkers := 5  // Total number of workers
queueSize := 5   // Maximum number of tasks in the queue


## How It Works
    1. Task Submission:
        Clients submit tasks via the /task endpoint.
        The task is validated and enqueued in the broker with its priority.

    2. Task Processing:
        Workers in the pool fetch tasks from the broker based on priority.
        Each worker processes tasks concurrently.

    3. Retries:
        If a task fails, it is retried once before being discarded.

    4. Logging:
        Logs are printed to the CLI for task submissions, processing, and errors.

##Future Enhancements
    Add persistent storage for tasks using a database.
    Implement advanced retry policies with exponential backoff.
    Add support for distributed worker pools across multiple nodes.
    Introduce authentication and authorization for the REST API.
