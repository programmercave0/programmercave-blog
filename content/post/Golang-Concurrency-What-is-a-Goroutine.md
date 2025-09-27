---
date: "2025-09-27T00:00:00Z"
description:
header-img:
keywords: Go programming language, Golang speed, compiled languages, concurrency in Go, goroutines, Go garbage collection, high-performance applications, Go interview preparation, native compilation, modern hardware optimization, fast compile times, efficient coding practices, Golang features
tags:
- Go
title: 'Golang - Concurrency - What is a Goroutine'
toc: true
---

## The Elevator Pitch: What's a Goroutine?

A **goroutine** is a lightweight "thread" managed by the Go runtime. Think of it as a function that can run concurrently (alongside) other functions.

The magic is in the `go` keyword. When you put `go` in front of a function call, you're telling Go: "Run this function in the background, and don't wait for it to finish. Just move on to the next line of code."

```go
package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from the goroutine!")
}

func main() {
	go sayHello() // This runs in the background

	fmt.Println("Hello from the main function!")
	time.Sleep(100 * time.Millisecond) // Wait a bit to see the goroutine's output
}
```

---

## Goroutines vs. OS Threads: The Key Difference

This is a classic interview question!

| Feature           | Goroutines                                  | OS Threads                                    |
|-------------------|---------------------------------------------|-----------------------------------------------|
| **Managed By**    | The Go Runtime                              | The Operating System (OS)                     |
| **Size**          | Very small (starts at ~2KB)                 | Larger (usually 1MB or more)                  |
| **Creation Cost** | Very cheap and fast                         | Relatively expensive and slow                 |
| **Switching**     | Fast, happens in user space               | Slower, requires a context switch to the OS kernel |
| **How Many?**     | You can have hundreds of thousands, even millions. | You're limited to a few thousand.             |

**Analogy:**

*   **OS Threads** are like having a separate, heavy-duty vehicle for every task. They're powerful but expensive to create and maintain.
*   **Goroutines** are like having a fleet of lightweight delivery drones. You can launch thousands of them cheaply and quickly to get many small jobs done.

---

## The Main Goroutine and Why Your Program Exits

Every Go program starts with one goroutine, called the **main goroutine**. It's the one that runs your `main()` function.

**Here's the critical rule:** When the main goroutine finishes, your entire program exits, even if other goroutines are still running.

That's why in the first example, we needed `time.Sleep()`. Without it, `main()` would finish before `sayHello()` had a chance to run, and we'd never see its output.

---

## How to Properly Wait for Goroutines: `sync.WaitGroup`

Using `time.Sleep()` is a hack. The right way to wait for goroutines is with a `sync.WaitGroup`. It's a simple counter that lets you coordinate your goroutines.

**How it works:**

1.  `wg.Add(1)`: Before you start a goroutine, you tell the `WaitGroup` that you're adding one task to the group.
2.  `defer wg.Done()`: Inside the goroutine, you use `defer` to tell the `WaitGroup` that this task is done when the function exits.
3.  `wg.Wait()`: In your main function, `wg.Wait()` will pause and wait until all the tasks have called `Done()`.

### Code Example:

```go
package main

import (
	"fmt"
	"sync"
	time "time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // This will be called when the function returns

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate some work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the counter
		go worker(i, &wg)
	}

	wg.Wait() // Block until the counter is zero

	fmt.Println("All workers have finished.")
}
```

---

## Key Takeaways for Your Interview

*   Goroutines are lightweight, concurrent execution units managed by the Go runtime.
*   They are much cheaper than OS threads, and you can run thousands of them.
*   The `go` keyword starts a new goroutine.
*   When the `main` goroutine exits, the program exits.
*   Use `sync.WaitGroup` to wait for goroutines to finish, not `time.Sleep()`.