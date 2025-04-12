# Distributed_Task_Queue


High level design for the MVP 1.

+-----------+          +-----------+          +---------------+
| Producer  |  ----->  |  Broker   |  <-----  |  Worker       |
| (REST API)|          | (In-Mem)  |          | (Worker Pool) |
+-----------+          +-----------+          +---------------+

 
 MVP 1 have ::

Feature	                    Included in MVP	                 Reason
REST API to submit tasks	✅	                 Enables external clients to use the queue
In-memory broker	        ✅	                 Keeps implementation simple (no DB needed)
Worker pool	                ✅	                 Simulates real task processing
Retry once on failure	    ✅	                 Introduces basic fault tolerance
CLI logging for monitoring	✅	                 Easy way to debug and observe system
JSON payload + task type	✅	                 Lets us simulate different task behaviors
Concurrency in worker pool	✅	                 Basic parallelism handling



Broker working::


                   +-------------------+
                   |    API Handler    |
                   +--------+----------+
                            |
                            | Enqueue(task)
                            v
                    +-------+--------+
                    |     Broker     |
                    | (chan *Task)   |
                    +-------+--------+
                            |
                            | Dequeue() (read-only channel)
                            v
                   +--------+--------+
                   |     Worker      |
                   +-----------------+
