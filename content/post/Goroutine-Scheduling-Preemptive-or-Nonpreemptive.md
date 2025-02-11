---
date: "2025-02-10T00:00:00Z"
description: Understanding whether Goroutines are preemptively or nonpreemptively scheduled is critical for writing efficient concurrent code. Let’s break down how Go’s runtime manages Goroutines and why this hybrid model is unique.
header-img: "/assets/images/Golang/Groutines-Scheduling.png"
keywords: Goroutines, Golang concurrency, Go routines, Software engineering interview, Go programming, Concurrency in Go, sync.WaitGroup, M:N scheduler, Fork-Join model, Preemptive scheduling
tags:
- Go
- Interview-Questions
title: 'Goroutine Scheduling: Preemptive or Nonpreemptive?'
toc: true
---
Understanding whether Goroutines are preemptively or nonpreemptively scheduled is critical for writing efficient concurrent code. Let’s break down how Go’s runtime manages Goroutines and why this hybrid model is unique.  

![Goroutine Scheduling: Preemptive or Nonpreemptive?](/assets/images/Golang/Groutines-Scheduling.png)

### **Preemptive vs. Nonpreemptive Scheduling**  

1. **Preemptive Scheduling**:  
   - The scheduler can **forcefully pause** a running task (e.g., an OS thread) at **any time** to allow other tasks to run.  
   - Common in OS-level thread scheduling.  
   - Example: A thread running a loop can be interrupted mid-execution.  

2. **Nonpreemptive (Cooperative) Scheduling**:  
   - Tasks **voluntarily yield control** at defined points (e.g., during I/O waits).  
   - Requires tasks to explicitly release resources.  
   - Example: A function must call `yield()` to let another task run.  

### **Goroutines: A Hybrid Approach**  

Goroutines blend **nonpreemptive** and **preemptive** behaviors. Here’s how it works:  

#### 1. **Default Behavior (Cooperative)**  
Goroutines **voluntarily yield control** when they hit **blocking operations** like:  
- I/O operations (e.g., reading a file).  
- Channel sends/receives.  
- System calls (e.g., network requests).  
- Sleep (`time.Sleep`).  

**Example**:  
```go  
func main() {  
    go func() {  
        fmt.Println("Goroutine 1")  
        time.Sleep(100 * time.Millisecond) // Yield point  
    }()  
    go func() { fmt.Println("Goroutine 2") }()  
    time.Sleep(time.Second)  
}  
```  
Here, `time.Sleep` creates a yield point, allowing the scheduler to run other Goroutines.  

#### 2. **Preemptive-Like Behavior (Since Go 1.14)**  
For **CPU-bound tasks** (e.g., loops without blocking calls), Go introduced **asynchronous preemption** to prevent starvation:  
- The runtime uses OS signals (e.g., `SIGURG` on Unix) to interrupt long-running Goroutines.  
- This allows the scheduler to fairly allocate CPU time across all Goroutines.  

**Example**:  
```go  
func main() {  
    go func() {  
        for {  
            // CPU-bound loop (no blocking calls)  
        }  
    }()  
    go func() { fmt.Println("Goroutine 2") }()  
    time.Sleep(time.Second)  
}  
```  
Before Go 1.14, the infinite loop could block other Goroutines. Now, the scheduler preempts it. 

### **Key Implications**  

1. **Developers Don’t Need to Manually Yield**  
   - The Go runtime handles suspension/resumption automatically.  
   - Blocking operations or long CPU tasks no longer freeze the program.  

2. **Safe Points for Preemption**  
   - Preemption occurs at **function prologues** (entry points of functions).  
   - The compiler inserts checks to ensure safe interruption.  

3. **Fairness**  
   - Ensures no single Goroutine monopolizes the CPU.  

### **Summary**  

- Goroutines use **nonpreemptive scheduling** by default (yield at blocking points).  
- **Preemptive-like behavior** was added in Go 1.14 for CPU-bound tasks.  
- The hybrid model simplifies concurrency: no manual yielding, no starvation.  
