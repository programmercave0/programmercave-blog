---
date: "2025-09-27T00:00:00Z"
description:
header-img:
keywords: Go programming language, Golang speed, compiled languages, concurrency in Go, goroutines, Go garbage collection, high-performance applications, Go interview preparation, native compilation, modern hardware optimization, fast compile times, efficient coding practices, Golang features
tags:
- Go
title: 'Golang - Concurrency - Concurrency vs Parallelism'
toc: true
---

## The Core Idea: What's the Difference?

Imagine you're a chef in a kitchen.

*   **Concurrency** is like juggling multiple tasks. You start cooking the soup, then while it simmers, you chop vegetables for the salad. You switch between tasks, making progress on all of them. You're *handling* multiple things at once, but not necessarily *doing* them at the exact same instant.
*   **Parallelism** is like having multiple chefs in the kitchen. One chef is making the soup, and at the same time, another chef is making the salad. Multiple tasks are happening simultaneously.

In short:

*   **Concurrency:** Dealing with multiple things at once.
*   **Parallelism:** Doing multiple things at once.

---

## Concurrency in Go

Go is designed for concurrency. It has first-class support for `goroutines`, which are lightweight threads.

*   **What it is:** You can have thousands of goroutines running, and Go will switch between them efficiently on a single core.
*   **Why it's useful:** It allows you to write clean, readable code that can handle many tasks without getting bogged down. For example, a web server can handle multiple incoming requests concurrently.

### Code Example:

This code runs two tasks concurrently. They take turns running on a single CPU.

```go
package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go task("Task A") // Start a goroutine
	go task("Task B") // Start another goroutine

	// Wait for the goroutines to finish
	time.Sleep(500 * time.Millisecond)
	fmt.Println("All tasks done!")
}
```

---

## Parallelism in Go

If you have a multi-core processor, Go can run your concurrent code in parallel.

*   **How it works:** The Go runtime can assign different goroutines to run on different CPU cores.
*   **The benefit:** You get a significant performance boost for CPU-intensive tasks because you're using the full power of your hardware.

### Code Example:

By default, Go will use all available CPU cores. You can control this with `runtime.GOMAXPROCS`.

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func intensiveTask(name string) {
	// A simple, CPU-bound task
	sum := 0
	for i := 0; i < 1_000_000_000; i++ {
		sum += i
	}
	fmt.Printf("%s finished with sum: %d\n", name, sum)
}

func main() {
	// Use all available CPU cores
	fmt.Printf("Using %d CPU cores\n", runtime.NumCPU())

	go intensiveTask("Task 1")
	go intensiveTask("Task 2")

	// Give the tasks time to complete
	time.Sleep(5 * time.Second)
}
```

---

## Key Takeaways for Your Interview

| Feature       | Concurrency                               | Parallelism                               |
|---------------|-------------------------------------------|-------------------------------------------|
| **Analogy**   | One person juggling multiple tasks.       | Multiple people working on tasks.         |
| **Goal**      | To structure a program to handle many tasks.| To speed up a program by doing work at the same time.|
| **Requires**  | A way to switch between tasks (like goroutines). | A multi-core processor.                   |
| **In Go**     | Achieved with goroutines, even on one core. | Achieved when goroutines run on multiple cores. |

**The most important thing to remember:** You write *concurrent* code in Go. The Go runtime will execute it in *parallel* if you have the hardware for it.
