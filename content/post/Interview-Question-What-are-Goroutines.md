---
date: "2025-02-11T00:00:00Z"
description: Learn Goroutines in Golang for software engineering interviews. Understand concurrency, M:N scheduling, sync.WaitGroup, and avoid common pitfalls. Master Go routines today!
header-img: "/assets/images/Golang/Goroutines.png"
keywords: Goroutines, Golang concurrency, Go routines, Software engineering interview, Go programming, Concurrency in Go, sync.WaitGroup, M:N scheduler, Fork-Join model, Preemptive scheduling
tags:
- Go
- Interview-Questions
title: 'Interview Question: What are Goroutines?'
toc: true
---

Concurrency is a cornerstone of modern software engineering, and Go (Golang) makes it accessible with **Goroutines**. If you’re preparing for a software engineering interview, understanding Goroutines is crucial. This guide breaks down Goroutines in simple terms, complete with examples and diagrams.  

![What are Goroutines?](/assets/images/Golang/Goroutines.png)

## What Are Goroutines?  

A **Goroutine** is a lightweight execution thread managed by the Go runtime. Think of it as a function that runs independently and concurrently with other Goroutines.  

### Key Features:  
- **Lightweight**: Goroutines use minimal memory (as low as 2KB) compared to OS threads (1-2MB).  
- **Simple Syntax**: Prefix a function with `go` to run it concurrently:  
  ```go
  go myFunction() // Runs myFunction in a new Goroutine
  ```  
- **Managed by Go Runtime**: The runtime handles scheduling, suspension, and resumption automatically.  

### The Main Goroutine  
Every Go program starts with a **main Goroutine**. If it exits, all other Goroutines terminate immediately.  

## Goroutines vs. Threads  

| **Feature**       | **Goroutines**              | **OS Threads**          |  
|--------------------|-----------------------------|-------------------------|  
| **Memory**         | 2KB                        | 1-2MB                  |  
| **Management**     | Go runtime                 | Operating System       |  
| **Scheduling**     | M:N model (see below)      | 1:1 with OS threads    |  

[What are the differences between Goroutines and Threads?](/Interview-Question-What-are-difference-between-Goroutines-and-Threads/)

## Concurrency vs. Parallelism  

- **Concurrency**: Multiple tasks progress *simultaneously* (e.g., handling web requests).  
- **Parallelism**: Tasks run *at the same time* on multiple CPU cores.  

Goroutines enable **concurrency**, but parallelism depends on available CPU cores.  

## The Fork-Join Model  

Go uses the **fork-join concurrency model**:  
- **Fork**: Start a child Goroutine.  
- **Join**: Wait for child Goroutines to finish.  

Example without a Join Point:  
```go  
func main() {  
    go fmt.Println("Hello") // Fork  
    // Main Goroutine exits before "Hello" prints  
}  
```  
Here, `"Hello"` might never print because the main Goroutine exits too quickly.  

## Fixing Race Conditions with sync.WaitGroup  

Use `sync.WaitGroup` to create a **join point** and ensure Goroutines complete:  

```go  
var wg sync.WaitGroup  

func main() {  
    wg.Add(1) // Add 1 Goroutine to wait for  
    go sayHello()  
    wg.Wait() // Join point: wait until Done() is called  
}  

func sayHello() {  
    defer wg.Done() // Signal completion  
    fmt.Println("Hello")  
}  
```  

## Goroutine Scheduling: Preemptive or Nonpreemptive?  

Goroutines are **nonpreemptive** but appear preemptive because the Go runtime:  
1. Observes blocking operations (e.g., I/O, channel operations).  
2. Automatically suspends/resumes Goroutines.  

This hybrid model simplifies concurrency without manual yield points.  
[Goroutine Scheduling: Preemptive or Nonpreemptive?](/Goroutine-Scheduling-Preemptive-or-Nonpreemptive/)

[Can You Explain the Lifecycle and Scheduling of a Goroutine?](/Interview-Question-Can-You-Explain-the-Lifecycle-and-Scheduling-of-a-Goroutine/)

## Summary  

- **Goroutines** are lightweight, easy-to-use concurrency primitives in Go.  
- Use **sync.WaitGroup** for synchronization, not `time.Sleep`.  
- The **M:N scheduler** enables efficient multitasking.  

Practice these concepts to ace concurrency questions in your software engineering interview!  