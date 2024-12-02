---
layout: post
title: "Concurrency Patterns in Go: Worker Pool Pattern"
description: "Learn about the Worker Pool pattern in Go, a concurrency design pattern for efficiently processing tasks using a pool of goroutines. This blog explains use cases, benefits, and provides a complete implementation with code examples to help you build scalable and resource-efficient applications."
keywords: "Go concurrency, Worker Pool in Go, Go concurrency patterns, Worker Pool example, Go goroutines, concurrent task processing, Go worker pool implementation, Go scalability, Go programming tutorial, resource-efficient concurrency"
header-img: ""
tags:  [Go]
date: 2024-11-21
toc: true
---

The Worker Pool pattern is a powerful concurrency design pattern that helps process tasks efficiently using a pool of worker goroutines. Let's explore how it works step-by-step.

---

### **Step 1: Defining the Job**

A **Job** is a unit of work that the worker goroutines will process. In our implementation, the `Job` struct contains an `ID` field to identify the task.

```go
// Job represents the task to be executed by a worker
type Job struct {
	ID int
}
```

---

### **Step 2: Designing the Worker Pool**

The `WorkerPool` struct manages the workers, job queue, and results channel. It uses a **WaitGroup** to wait for all workers to finish their tasks.

```go
type WorkerPool struct {
	numWorkers int
	jobQueue   chan Job
	results    chan int
	wg         sync.WaitGroup
}

// NewWorkerPool creates a new worker pool with the specified number of workers
func NewWorkerPool(numWorkers, jobQueueSize int) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		jobQueue:   make(chan Job, jobQueueSize),
		results:    make(chan int, jobQueueSize),
	}
}
```

- **`numWorkers`**: Number of worker goroutines.
- **`jobQueue`**: A buffered channel to hold jobs.
- **`results`**: A buffered channel to store results of completed jobs.

---

### **Step 3: Implementing the Worker Function**

Each worker goroutine picks tasks from the `jobQueue`, processes them, and sends the result to the `results` channel.

```go
func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for job := range wp.jobQueue {
		fmt.Printf("Worker %d started job %d\n", id, job.ID)
		time.Sleep(time.Second) // Simulating work
		fmt.Printf("Worker %d finished job %d\n", id, job.ID)
		wp.results <- job.ID
	}
}
```

- **`id`**: Identifies the worker.
- Workers keep processing jobs until the `jobQueue` is closed.

---

### **Step 4: Starting the Worker Pool**

The `Start` method creates the worker goroutines and assigns them to process jobs from the `jobQueue`.

```go
func (wp *WorkerPool) Start() {
	for i := 1; i <= wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}
```

---

### **Step 5: Adding Jobs and Collecting Results**

Jobs are added to the `jobQueue` using `AddJob`. After all jobs are added, we close the queue to signal workers to stop.

```go
func (wp *WorkerPool) AddJob(job Job) {
	wp.jobQueue <- job
}

func (wp *WorkerPool) CollectResults() {
	for result := range wp.results {
		fmt.Printf("Result received for job %d\n", result)
	}
}
```

---

### **Step 6: Waiting for Workers to Complete**

The `Wait` method waits for all workers to finish their tasks and then closes the `results` channel.

```go
func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
	close(wp.results)
}
```

---

### **Putting It All Together**

Finally, we bring everything together in the `main` function. We create a worker pool, add jobs, and process the results.

```go
func main() {
	numWorkers := 3
	numJobs := 10

	workerPool := NewWorkerPool(numWorkers, numJobs)

	// Adding jobs to the Job Queue
	for i := 1; i <= numJobs; i++ {
		workerPool.AddJob(Job{ID: i})
	}
	close(workerPool.jobQueue)

	workerPool.Start()
	workerPool.Wait()
	workerPool.CollectResults()
}
```

---

### **Output Example**

```plaintext
Worker 1 started job 1
Worker 2 started job 2
Worker 3 started job 3
Worker 1 finished job 1
Worker 2 finished job 2
Worker 3 finished job 3
Worker 1 started job 4
Worker 2 started job 5
...
Result received for job 1
Result received for job 2
Result received for job 3
...
```

---

### **Key Takeaways**

1. **Controlled Concurrency**: The number of workers (`numWorkers`) controls the concurrency level.
2. **Efficient Resource Usage**: Worker Pools ensure optimal resource utilization by preventing an excessive number of goroutines.
3. **Scalability**: Easy to scale by adjusting the number of workers and the size of the job queue.

---

### **Extensions**
- **Error Handling**: Add an error channel to handle failures.
- **Metrics**: Track the time taken to process each job.
- **Dynamic Scaling**: Adjust the number of workers based on the workload.

This implementation demonstrates the simplicity and power of the Worker Pool pattern for managing concurrent tasks in Go.