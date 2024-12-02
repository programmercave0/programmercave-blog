---
layout: post
title: "Interview Question: Thread-Safe Implementation of Map in Go"
description: "Explore two approaches to implementing thread-safe maps in Go: using `sync.Mutex` for control and `sync.Map` for simplicity. Learn their differences, use cases, and examples for concurrent programming."
keywords: "thread-safe map, Go programming, sync.Mutex, sync.Map, concurrent programming, Golang map, multi-threading, safe map implementation, Go synchronization, map concurrency, interview question"
header-img: ""
tags:  [Go, Interview-Questions]
date: 2024-11-20
toc: true
---

This blog demonstrates two approaches to implement a **thread-safe map** in Go: using `sync.Mutex` and the built-in `sync.Map`. These techniques ensure safe concurrent access to maps in multi-threaded environments.  

_This question was asked to me by Radisys for the role of Golang Developer._

---

### **Approach 1: SafeMap with `sync.Mutex`**  

This implementation uses `sync.Mutex` to synchronize access to a map.  

### **Structure and Methods**  

#### **1. SafeMap Struct**  
Contains:  
- `mu sync.Mutex`: The lock to synchronize map access.  
- `data map[string]int`: The underlying map to store key-value pairs.  

#### **2. Methods**  

- **NewSafeMap()**:  
  Creates and returns a new instance of `SafeMap` with an initialized map.  

- **Set(key string, value int)**:  
  Adds or updates a key-value pair in the map:  
  1. Acquires a lock (`mu.Lock()`).  
  2. Modifies the map.  
  3. Releases the lock (`mu.Unlock()`).  

- **Get(key string) (int, bool)**:  
  Retrieves the value for a given key:  
  1. Acquires a lock.  
  2. Checks if the key exists.  
  3. Returns the value and a boolean indicating the key's existence.  
  4. Releases the lock.  

---

#### **Code**  
```go
package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu   sync.Mutex
	data map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (m *SafeMap) Set(key string, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *SafeMap) Get(key string) (int, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	value, ok := m.data[key]
	return value, ok
}

func main() {
	m := NewSafeMap()
	m.Set("a", 42)
	value, ok := m.Get("a")
	if ok {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Key not found")
	}
}
```

---

### **Approach 2: Using `sync.Map`**  

The `sync.Map` is a concurrent map provided by Go's `sync` package, optimized for high-concurrency use cases. It eliminates the need for manual locking.  

#### **Characteristics**  
1. **Thread-safe**: Multiple goroutines can access it safely.  
2. **No Generics**: Stores `interface{}` types, requiring type casting.  
3. **Optimized for Writes**: Performs better in high-update scenarios.  
4. **Iterate with Care**: Iteration doesn't guarantee order or consistency during modifications.  

#### **Common Methods**  
- `Store(key, value)`: Adds or updates a key-value pair.  
- `Load(key)`: Retrieves the value for a key.  
- `Delete(key)`: Removes a key-value pair.  
- `LoadOrStore(key, value)`: Loads an existing value or stores and returns a new one.  
- `Range(func(key, value any) bool)`: Iterates over all entries.  

#### **Code Example**  
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map

	// Store values
	sm.Store("a", 1)
	sm.Store("b", 2)

	// Load values
	if value, ok := sm.Load("a"); ok {
		fmt.Println("Key 'a':", value)
	}

	// Iterate
	sm.Range(func(key, value any) bool {
		fmt.Println("Key:", key, "Value:", value)
		return true // Continue iteration
	})

	// Delete a key
	sm.Delete("b")
}
```

---

### **Use Cases**  

1. **`sync.Mutex` + Map**:  
   - When you need full control over locking.  
   - For simpler, low-concurrency use cases.  

2. **`sync.Map`**:  
   - High-concurrency applications with frequent reads and writes.  
   - Caches or shared state in concurrent systems.  

---

### **Conclusion**  

Both `sync.Mutex` and `sync.Map` are excellent tools for creating thread-safe maps in Go. Choose `sync.Mutex` for full control and type safety or `sync.Map` for simplicity and performance in high-concurrency scenarios.  

  