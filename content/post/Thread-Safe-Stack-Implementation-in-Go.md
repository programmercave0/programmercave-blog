---
date: "2024-11-25T00:00:00Z"
description: Learn how to implement a thread-safe stack in Go using sync.RWMutex.
  This tutorial explains stack operations, ensuring data consistency and safety in
  multi-threaded environments with efficient read-write locking.
header-img: ""
keywords: thread-safe stack, Go programming, sync.RWMutex, concurrent programming,
  Golang stack, multi-threading, concurrent data structure, Go synchronization, stack
  implementation in Go, thread-safe operations
tags:
- Go
title: Thread-Safe Stack Implementation in Go
toc: true
---

This blog demonstrates how to implement a **thread-safe stack** in Go using `sync.RWMutex`. It ensures safe concurrent access in multi-threaded environments. Below is a detailed breakdown.  

---

### **Key Concepts**  

1. **Stack**: A data structure following the **Last-In-First-Out (LIFO)** principle, where items are added and removed from the top.  

2. **Thread-Safety**: Ensures multiple goroutines can perform operations on shared resources (the stack) without causing data corruption.  

3. **sync.RWMutex**:  
   - **Read Lock (`RLock`)**: Allows multiple readers concurrently.  
   - **Write Lock (`Lock`)**: Ensures exclusive access for operations like `Push` and `Pop`.  

This boosts performance by allowing concurrent reads but restricting writes.  

---

### **Structure and Methods**  

#### **1. ItemType**  
Defined as an empty interface (`interface{}`), it allows the stack to hold any type of data (e.g., int, string, structs).  

#### **2. Stack Struct**  
Contains:  
- `items []ItemType`: Holds the stack elements.  
- `rwLock sync.RWMutex`: Manages concurrent access.  

---

### **Methods**  

- **New()**:  
  Initializes and returns a new empty stack.  

- **Push(t ItemType)**:  
  Adds an item to the top:  
  1. Acquires a **write lock** (`Lock()`) to block other operations.  
  2. Appends the item.  
  3. Releases the lock (`Unlock()`).  

- **Pop()**:  
  Removes and returns the top item:  
  1. Acquires a write lock.  
  2. Removes the last item from the slice.  
  3. Resizes the slice.  
  4. Releases the lock.  

- **Size()**:  
  Returns the number of items in the stack:  
  1. Acquires a **read lock** (`RLock()`), allowing other reads simultaneously.  
  2. Returns the size.  
  3. Releases the lock (`RUnlock()`).  

- **All()**:  
  Returns all items in the stack with a read lock.  

- **IsEmpty()**:  
  Checks if the stack is empty, using a read lock.  

---

```go
package main

import (
	"fmt"
	"sync"
)

// ItemType - The type of item in the stack.
type ItemType interface{}

// Stack - Stack of items.
type Stack struct {

	// Slice of type ItemType, it holds items in the stack.
	items []ItemType

	// rwLock for handling concurrent operations on the stack.
	rwLock sync.RWMutex
}

// New - Creates a new Stack.
func (stack *Stack) New() *Stack {

	stack.items = []ItemType{}

	return stack
}

// Push - Adds an Item to the top of the stack
func (stack *Stack) Push(t ItemType) {

	//Initialize items slice if not initialized
	if stack.items == nil {
		stack.items = []ItemType{}
	}

	// Acquire read, write lock before inserting a new item in the stack.
	stack.rwLock.Lock()

	// Performs append operation.
	stack.items = append(stack.items, t)

	// This will release read, write lock
	stack.rwLock.Unlock()
}

// Pop removes an Item from the top of the stack
func (stack *Stack) Pop() *ItemType {

	// Acquire read, write lock as items are going to modify.
	stack.rwLock.Lock()

	// Popping item from items slice.
	item := stack.items[len(stack.items)-1]

	//Adjusting the item's length accordingly
	stack.items = stack.items[0 : len(stack.items)-1]

	// Release read-write lock.
	stack.rwLock.Unlock()

	// Return last popped item
	return &item
}

// Size return size i.e. number of items present in stack.
func (stack *Stack) Size() int {
	// Acquire read lock
	stack.rwLock.RLock()

	// defer operation of unlocking.
	defer stack.rwLock.RUnlock()

	// Return length of items slice.
	return len(stack.items)
}

// All - return all items present in stack
func (stack *Stack) All() []ItemType {
	// Acquire read lock
	stack.rwLock.RLock()

	// defer operation of unlocking.
	defer stack.rwLock.RUnlock()

	// Return items slice to the caller.
	return stack.items
}

// IsEmpty - Check is stack is empty or not.
func (stack *Stack) IsEmpty() bool {
	// Acquire read lock
	stack.rwLock.RLock()
	// defer operation of unlock.
	defer stack.rwLock.RUnlock()

	return len(stack.items) == 0
}

func main() {

	stack := Stack{}
	fmt.Println(stack.All())
	stack.Push(10)
	fmt.Println(stack.All())
	stack.Push(20)
	fmt.Println(stack.All())
	stack.Push(30)
	fmt.Println(stack.All())
	stack.Push(40)
	fmt.Println(stack.All())
	stack.Push(50)
	fmt.Println(stack.All())
	stack.Push(60)
	fmt.Println(stack.All())

	stack.Pop()
	fmt.Println(stack.All())
	stack.Pop()
	fmt.Println(stack.All())
	stack.Pop()
	fmt.Println(stack.All())
	stack.Pop()
	fmt.Println(stack.All())
	stack.Pop()
	fmt.Println(stack.All())
	stack.Pop()
	fmt.Println(stack.All())

}
```

### **Conclusion**  

The thread-safe stack in Go demonstrates the power of `sync.RWMutex` for managing shared resources. Its methods ensure **safe concurrent reads and writes**, making it a robust solution for handling stacks in multi-threaded programs.  

This implementation is essential knowledge for anyone working with concurrent programming in Go.