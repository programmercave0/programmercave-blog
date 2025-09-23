---
date: "2025-09-23T00:00:00Z"
description:
header-img:
keywords: Go programming language, Golang speed, compiled languages, concurrency in Go, goroutines, Go garbage collection, high-performance applications, Go interview preparation, native compilation, modern hardware optimization, fast compile times, efficient coding practices, Golang features
tags:
- Go
title: 'Golang - Memory - Memory Leaks'
toc: true
---

### Introduction

Even though Go has a garbage collector, it is still possible to have memory leaks. A memory leak in Go doesn't mean the GC is broken; it means you have accidentally kept a reference to an object that you no longer need, preventing the GC from reclaiming its memory.

**Analogy:** A memory leak is like putting a box in a **long-term storage unit and then forgetting you have the key**. The storage company (the GC) sees that you are still the owner, so it will never throw the box away, even if you never intend to use its contents again. The box just sits there, taking up space you have to pay for.

---

### Common Causes of Memory Leaks in Go

#### 1. Leaking Goroutines

This is the most common cause of memory leaks in Go. A goroutine is a lightweight thread, but it still consumes memory (starting at 2KB for its stack, which can grow). If you start a goroutine that blocks forever and never terminates, its memory is never released.

**Example: A Blocked Channel Send/Receive**

```go
// A goroutine that waits to receive from a channel.
// If nothing is ever sent, it will block forever.
func worker(ch <-chan int) {
    for {
        val := <-ch // Blocks here
        fmt.Println(val)
    }
}

func main() {
    ch := make(chan int)
    go worker(ch) // This goroutine is now leaked.

    // The main goroutine exits, but in a long-running server,
    // this worker would be stuck in memory forever.
}
```
**How to prevent:**
*   Ensure every goroutine you start has a clear exit condition.
*   Use `context` cancellation to signal goroutines to shut down.
*   Use buffered channels or `select` statements with a `default` or `timeout` case to prevent indefinite blocking.

#### 2. Unbounded Global Collections (Maps and Slices)

If you have a global map or slice that you only ever add to, it will grow indefinitely, consuming more and more memory.

**Example: A Global Cache Without an Eviction Policy**

```go
// A simple cache that is never cleared.
var cache = make(map[string]string)

func handleRequest(key string) string {
    if val, ok := cache[key]; ok {
        return val
    }
    // ... calculate value ...
    val := "some calculated value"
    cache[key] = val // Add to cache, but never remove.
    return val
}
```
**How to prevent:**
*   Avoid global variables for request-scoped data.
*   If you need a cache, implement a bounding and eviction strategy (e.g., LRU - Least Recently Used, or TTL - Time To Live).

#### 3. Dangling Slices (Substring Memory Leak)

When you take a slice of a larger slice or array, the new slice's header still points to the original, large backing array. If the original large array is no longer needed *except* for that small slice, the GC cannot free the large array.

**Example:**

```go
// This function reads a 10MB file into memory, but only returns the first 5 bytes.
func getHeader() []byte {
    largeData, _ := ioutil.ReadFile("largefile.dat") // 10MB
    return largeData[:5] // This slice still holds a reference to the 10MB array!
}

func main() {
    // The header variable is small, but it's preventing 10MB from being garbage collected.
    header := getHeader()
    // ... use header ...
}
```
**How to prevent:**
*   When taking a small slice from a large one that you intend to keep, make an explicit copy.

```go
func getHeaderFixed() []byte {
    largeData, _ := ioutil.ReadFile("largefile.dat")
    header := make([]byte, 5)
    copy(header, largeData[:5]) // Make a copy
    return header // Now it has its own small backing array.
}
```

---

### Detecting Memory Leaks with `pprof`

Go comes with a powerful built-in profiling tool called `pprof`. It is the standard way to diagnose memory and CPU issues.

**How to use `pprof` for memory leaks:**

1.  **Expose the `pprof` endpoints:** Add the `net/http/pprof` package to your application. This will register HTTP handlers for profiling data on your debug server.

    ```go
    import (
        _ "net/http/pprof"
        "net/http"
        "log"
    )

    func main() {
        go func() {
            log.Println(http.ListenAndServe("localhost:6060", nil))
        }()
        // ... your application logic ...
    }
    ```

2.  **Capture a Heap Profile:** While your application is running (and potentially leaking), use the `go tool pprof` to look at the heap. The most useful profile is `inuse_space`, which shows memory that is currently allocated.

    ```bash
    # This command will open an interactive pprof session
    go tool pprof http://localhost:6060/debug/pprof/heap
    ```

3.  **Analyze the Profile:** Inside the `pprof` tool, you can use commands to analyze the data:
    *   `top`: Shows the functions that are holding the most memory. If you see a function at the top that you expect to be short-lived, it's a sign of a leak.
    *   `list <function_name>`: Shows the source code of the function, with memory usage annotated line-by-line.
    *   `web`: Generates a visual graph (in SVG format) of the memory usage, which can make it very easy to spot the source of a leak.

By taking heap profiles at different times, you can compare them to see where memory is growing unexpectedly, leading you directly to the source of the leak.
