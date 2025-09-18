---
date: "2025-09-17T00:00:00Z"
description:
header-img:
keywords: Go programming language, Golang speed, compiled languages, concurrency in Go, goroutines, Go garbage collection, high-performance applications, Go interview preparation, native compilation, modern hardware optimization, fast compile times, efficient coding practices, Golang features
tags:
- Go
title: 'Golang - Data Structures - Concurrency with Maps'
toc: true
---

### Introduction

A critical point to understand about Go's built-in `map` type is that it is **not safe for concurrent use**. If one goroutine is writing to a map while another is reading from or writing to it, you will get a fatal runtime error. This is a common source of bugs in concurrent Go programs.

---

### The Problem: Race Conditions

Consider this simple program:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	m := make(map[int]int)

	// Goroutine 1: Writes to the map
	go func() {
		for i := 0; ; i++ {
			m[i] = i
		}
	}()

	// Goroutine 2: Reads from the map
	go func() {
		for i := 0; ; i++ {
			_ = m[i]
		}
	}()

	// Let it run for a bit
	time.Sleep(1 * time.Second)
	fmt.Println("Done (you will likely not see this)")
}
```

Running this code will almost immediately result in a panic:

```
fatal error: concurrent map read and map write
```

This happens because the operations are not atomic. A map write might be in the middle of resizing the map and moving buckets (an operation that is not instant) at the exact moment a read operation tries to access it. The Go runtime detects this unsafe situation and terminates the program.

---

### Solution 1: `sync.RWMutex` (The Standard Approach)

The most common way to make map access safe is to protect it with a `sync.RWMutex` (Read-Write Mutex).

A `sync.RWMutex` provides two types of locks:
1.  **Read Lock (`RLock`/`RUnlock`):** Allows multiple readers to access the map at the same time. As long as no one is writing, any number of goroutines can read.
2.  **Write Lock (`Lock`/`Unlock`):** Allows only a single writer. When a write lock is held, no other goroutines can read or write.

**Implementation:**
The standard practice is to create a custom struct that encapsulates the map and the mutex.

```go
package main

import (
	"fmt"
	"sync"
)

// SafeMap is a concurrency-safe map of string to int.
type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

// NewSafeMap creates a new SafeMap.
func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

// Set adds or updates a value in the map.
func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock() // Acquire a write lock
	defer sm.mu.Unlock()
	sm.data[key] = value
}

// Get retrieves a value from the map.
func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.RLock() // Acquire a read lock
	defer sm.mu.RUnlock()
	val, ok := sm.data[key]
	return val, ok
}

func main() {
	sm := NewSafeMap()
	var wg sync.WaitGroup

	// Start 100 writers
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n)
			sm.Set(key, n)
		}(i)
	}

	// Start 100 readers
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n)
			sm.Get(key)
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines finished without error.")
}
```

---

### Solution 2: `sync.Map` (For Specific Use Cases)

Go 1.9 introduced `sync.Map`, a concurrency-safe map implementation provided by the standard library.

**When should you use `sync.Map`?**
The official documentation states that `sync.Map` is optimized for two specific scenarios:
1.  When the key set is mostly stable, meaning keys are written once and then read many times.
2.  When multiple goroutines are reading, writing, and overwriting entries for disjoint sets of keys.

In these cases, `sync.Map` can be more performant than a regular map with a mutex because it can reduce lock contention.

**How it works:**
`sync.Map` works by having two internal maps: a read-only `read` map for fast lookups, and a `dirty` map for writes. Reads can happen lock-free on the `read` map. Writes require a lock but are stored in the `dirty` map. Periodically, the `dirty` map is promoted to be the new `read` map.

**Usage:**

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map
	var wg sync.WaitGroup

	// Store
	wg.Add(1)
	go func() {
		defer wg.Done()
		sm.Store("hello", "world")
	}()

	// Load
	wg.Add(1)
	go func() {
		defer wg.Done()
		val, ok := sm.Load("hello")
		if ok {
			fmt.Println("Loaded:", val)
		}
	}()

	wg.Wait()
}
```

**`sync.Map` vs. `map` with `RWMutex`**

| Feature             | `map` + `sync.RWMutex`                               | `sync.Map`                                                              |
| ------------------- | ---------------------------------------------------- | ----------------------------------------------------------------------- |
| **Type Safety**     | **Yes.** Keys and values are strongly typed.         | **No.** Keys and values are `interface{}`. Requires type assertions.    |
| **General Use**     | **The default choice.** Good for most scenarios.     | Optimized for specific read-mostly or disjoint key set workloads.       |
| **Performance**     | Can have lock contention if writes are frequent.     | Can be faster in its niche use cases, but may be slower otherwise.      |
| **Ease of Use**     | Requires creating a wrapper struct.                  | Simpler to use out of the box, but methods are different (`Store`, `Load`). |

**Interview Question:** "When would you use `sync.Map` over a map protected by a mutex?"
**Answer:** "`sync.Map` is a specialized tool. It's best used when you have a mostly-stable set of keys that are written once and read many times, or when you have high contention on disjoint key sets. For general-purpose concurrent map access, a standard map with a `sync.RWMutex` is often clearer, provides type safety, and performs well enough."
