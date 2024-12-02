---
layout: post
title: "Concurrency Patterns in Go: Fan-In and Fan-Out Pattern"
description: "Learn about the Fan-In and Fan-Out concurrency patterns in Go, their real-world use cases, and complete implementation examples to simplify task aggregation and distribution in parallel systems."
keywords: "Go concurrency, Fan-In pattern, Fan-Out pattern, Go goroutines, concurrency patterns in Go, task distribution, data aggregation, Go channels, parallel processing in Go, scalable systems"
header-img: ""
tags:  [Go]
date: 2024-11-22
toc: true
---

Concurrency is one of Go’s standout features, and two of the most popular concurrency patterns are **Fan-In** and **Fan-Out**. These patterns help manage data flow and task distribution efficiently, enabling parallelism and better resource utilization.

---

### **Fan-In Pattern**  
The **Fan-In** pattern consolidates data from multiple input sources into a single output channel. This is particularly useful when multiple goroutines produce data that needs to be aggregated for further processing.

#### **How Fan-In Works**  
1. **Producers**: Multiple goroutines generate data independently and send it to their respective channels.  
2. **Fan-In Function**: Merges data from multiple input channels into a single output channel.  
3. **Consumer**: Processes data from the single output channel.  

#### **Use Case**  
For example, consider a scenario where you collect data from multiple APIs concurrently and merge the results into one stream for processing.

---

#### **Code: Fan-In Implementation**  
```go
func producer(id int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 3; i++ {
			ch <- id*10 + i // Produce unique values
		}
	}()
	return ch
}

func fanIn(inputs ...<-chan int) <-chan int {
	output := make(chan int)
	var wg sync.WaitGroup

	// Function to copy data from input channel to output
	copy := func(ch <-chan int) {
		defer wg.Done()
		for val := range ch {
			output <- val
		}
	}

	wg.Add(len(inputs))
	for _, ch := range inputs {
		go copy(ch) // Start a goroutine for each input channel
	}

	go func() {
		wg.Wait()
		close(output) // Close output channel once all inputs are done
	}()

	return output
}

func main() {
	// Two producers generating data
	ch1 := producer(1)
	ch2 := producer(2)

	// Fan-In merges the data
	merged := fanIn(ch1, ch2)

	// Consume the merged data
	for val := range merged {
		fmt.Println("Received:", val)
	}
}
```

---

#### **Explanation**  
1. **Producer**: Creates a channel and sends three unique values.  
2. **Fan-In Function**:  
   - Takes multiple input channels.  
   - Launches goroutines to copy data from each input channel into the output channel.  
   - Uses a `sync.WaitGroup` to ensure all input channels finish before closing the output channel.  
3. **Consumer**: Reads from the merged channel and processes the data.

**Output Example**:  
```plaintext
Received: 10  
Received: 11  
Received: 12  
Received: 20  
Received: 21  
Received: 22  
```

---

### **Fan-Out Pattern**  
The **Fan-Out** pattern distributes tasks from a single source channel to multiple worker goroutines. Each worker processes tasks independently, enabling parallel task execution.

#### **How Fan-Out Works**  
1. **Task Source**: A single channel produces tasks to be processed.  
2. **Workers**: Multiple goroutines (workers) consume tasks from the source channel.  
3. **Results Collection**: Processed results are collected from workers into a results channel.

#### **Use Case**  
Imagine a scenario where a web scraper fetches URLs from a task queue and distributes them to multiple workers for scraping.

---

#### **Code: Fan-Out Implementation**  
```go
func source(taskQueue chan<- int, numTasks int) {
	for i := 1; i <= numTasks; i++ {
		taskQueue <- i // Add tasks to the queue
	}
	close(taskQueue) // Close the channel after sending all tasks
}

func worker(id int, taskQueue <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskQueue {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(time.Millisecond * 100) // Simulating work
		results <- task * 2 // Send processed result
	}
}

func main() {
	const numWorkers = 3
	const numTasks = 10

	taskQueue := make(chan int, numTasks)
	results := make(chan int, numTasks)
	var wg sync.WaitGroup

	// Create tasks
	go source(taskQueue, numTasks)

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, taskQueue, results, &wg)
	}

	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(results) // Close results channel after all workers are done
	}()

	// Collect results
	for result := range results {
		fmt.Println("Result:", result)
	}
}
```

---

#### **Explanation**  
1. **Task Source**: Produces `numTasks` and sends them into `taskQueue`.  
2. **Workers**:  
   - Consume tasks from `taskQueue`.  
   - Simulate processing by multiplying the task by 2.  
3. **Results Channel**: Collects processed results from all workers.  
4. **Concurrency**: `sync.WaitGroup` ensures that all workers finish before results processing starts.

**Output Example**:  
```plaintext
Worker 1 processing task 1  
Worker 2 processing task 2  
Worker 3 processing task 3  
Result: 2  
Result: 4  
Result: 6  
...
```

### **Conclusion**  
The **Fan-In** and **Fan-Out** patterns are essential tools for building scalable and efficient concurrent systems in Go. Whether consolidating data streams or distributing workloads, these patterns help simplify and optimize your application’s concurrency model.  

Feel free to experiment with these patterns to suit your application’s specific needs!