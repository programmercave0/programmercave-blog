---
date: "2025-09-19T00:00:00Z"
description:
header-img:
keywords: Go programming language, Golang speed, compiled languages, concurrency in Go, goroutines, Go garbage collection, high-performance applications, Go interview preparation, native compilation, modern hardware optimization, fast compile times, efficient coding practices, Golang features
tags:
- Go
title: 'Golang - Data Structures - The Empty Struct'
toc: true
---

### Introduction

The empty struct, written as `struct{}`, is a struct with no fields. It is a peculiar but surprisingly useful tool in Go. Its power comes from one key property: **it consumes zero memory**.

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s struct{}
	fmt.Println("Size of empty struct:", unsafe.Sizeof(s)) // Output: 0
}
```

Because all instances of an empty struct are identical and have no data, the Go compiler can optimize them to not consume any memory. This makes them perfect for situations where you need to signal something's existence without storing any actual information.

---

### Use Case 1: Implementing a Set

The most common use case for an empty struct is to implement a **set** data structure using a map. A set is a collection of unique items.

In a map, you need a key and a value. If you only care about the uniqueness of the keys (the items in your set), what do you use for the value? An empty struct is the perfect choice because it adds no memory overhead for the value.

```go
package main

import "fmt"

func main() {
	// A set of strings
	set := make(map[string]struct{})

	// Add items to the set
	set["apple"] = struct{}{}
	set["banana"] = struct{}{}
	set["apple"] = struct{}{} // Adding a duplicate does nothing

	// Check for existence
	if _, exists := set["banana"]; exists {
		fmt.Println("banana is in the set")
	}

	// How many unique items?
	fmt.Println("Number of unique items:", len(set)) // 2
}
```
Using `map[string]bool` would also work, but it's less idiomatic and uses at least one byte per value. `map[string]struct{}` clearly signals that the value is irrelevant and is the most memory-efficient solution.

---

### Use Case 2: Signaling on a Channel

Channels are used to send data between goroutines. But what if you don't need to send any data? What if you just need to send a **signal** to notify another goroutine that an event has occurred?

You can send an empty struct. This is the most efficient way to signal because you are not sending any actual data, just the event itself.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// A channel that only carries signals, not data
	done := make(chan struct{})

	go func() {
		fmt.Println("Worker goroutine is running...")
		time.Sleep(2 * time.Second)
		fmt.Println("Worker is done.")

		// Signal that the work is finished by sending an empty struct
		done <- struct{}{}
	}()

	// Block until a signal is received from the 'done' channel
	<-done
	fmt.Println("Main goroutine received signal and is exiting.")
}
```

A common and even more concise way to do this is to `close` the channel. A receive operation on a closed channel always returns immediately with a zero value, making it a perfect signaling mechanism.

```go
// In the goroutine:
close(done)

// In the main function, the receive works the same way:
<-done
```

---

### Use Case 3: A Method Receiver for Grouping

Sometimes you want to group a set of related functions together, but they don't need to operate on any specific state. You can define them as methods on an empty struct to provide a clear, namespaced API.

```go
package main

import "fmt"

// No state is needed, just a place to hang the methods.
type Greeter struct{}

func (g Greeter) SayHello() {
	fmt.Println("Hello!")
}

func (g Greeter) SayGoodbye() {
	fmt.Println("Goodbye!")
}

func main() {
	var g Greeter
	g.SayHello()
	g.SayGoodbye()
}
```
This can be a clean way to organize helper or utility functions without needing to create an unnecessary object with state.

---

### Summary Table

| Use Case                               | Why `struct{}`?                                                              | Example                               |
| -------------------------------------- | ---------------------------------------------------------------------------- | ------------------------------------- |
| **Set Implementation**                 | Zero memory cost for map values. Clearly signals "presence is what matters". | `mySet := make(map[string]struct{})`  |
| **Channel Signaling**                  | Sends a notification without the overhead of sending actual data.            | `done <- struct{}{}` or `close(done)` |
| **Method Grouping (Stateless API)**    | Provides a namespace for a group of functions without any instance data.     | `type MyAPI struct{}`                 |
