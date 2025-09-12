---
date: "2025-09-08T00:00:00Z"
description:
header-img:
keywords: Go programming language, Golang speed, compiled languages, concurrency in Go, goroutines, Go garbage collection, high-performance applications, Go interview preparation, native compilation, modern hardware optimization, fast compile times, efficient coding practices, Golang features
tags:
- Go
title: 'Golang - Data Structures - Common Slice Operations'
toc: true
---

### Introduction

Beyond simple indexing and appending, Go slices can be used to implement a variety of common data structure operations. Because Go doesn't provide these out of the box, knowing these idioms is a sign of a proficient Go developer.

---

### 1. Deleting an Element from a Slice

There are two main ways to delete an element from a slice, depending on whether you need to preserve the original order.

#### a) Deleting Without Preserving Order (Fastest)

If the order of elements doesn't matter, the most efficient way to delete is to swap the element you want to remove with the last element, and then shrink the slice. This avoids shifting all subsequent elements.

```go
// Let's remove the element at index i
s[i] = s[len(s)-1] // Replace element at i with the last element
s = s[:len(s)-1]   // Shrink the slice by one
```

**Example:**

```go
package main

import "fmt"

func main() {
	s := []string{"A", "B", "C", "D", "E"}
	indexToRemove := 2 // Remove "C"

	// Swap with the last element
	s[indexToRemove] = s[len(s)-1]
	// Shrink the slice
	s = s[:len(s)-1]

	fmt.Println(s) // Output: [A B E D] - Order is not preserved
}
```

#### b) Deleting While Preserving Order (Slower)

If you must preserve the order, you need to shift all the elements after the deleted one to the left. This is less performant for large slices. The `append` function provides a concise way to do this.

```go
// Let's remove the element at index i
s = append(s[:i], s[i+1:]...)
```

**Example:**

```go
package main

import "fmt"

func main() {
	s := []string{"A", "B", "C", "D", "E"}
	indexToRemove := 2 // Remove "C"

	s = append(s[:indexToRemove], s[indexToRemove+1:]...)

	fmt.Println(s) // Output: [A B D E] - Order is preserved
}
```

---

### 2. Inserting an Element into a Slice

To insert an element, you need to make space for it by shifting elements to the right. The `append` function is again the perfect tool for this.

```go
// Let's insert `value` at index `i`
s = append(s[:i], append([]T{value}, s[i:]...)...)
// A more efficient way for single element is:
s = append(s, T{}) // Make room for one more element
copy(s[i+1:], s[i:]) // Shift elements to the right
s[i] = value // Insert the new value
```

**Example:**

```go
package main

import "fmt"

func main() {
	s := []string{"A", "B", "D", "E"}
	valueToInsert := "C"
	index := 2

	// Using the more efficient copy method
	s = append(s, "") // Use the zero value for the type
	copy(s[index+1:], s[index:])
	s[index] = valueToInsert

	fmt.Println(s) // Output: [A B C D E]
}
```

[Slice Performance](/Golang-Data-Structures-Slice-Performance/)

---

### 3. Filtering a Slice In-Place

A common task is to filter a slice based on a condition (i.e., create a new slice containing only the elements that satisfy the predicate). The most memory-efficient way to do this is to reuse the original slice's backing array.

This idiom uses two indices: `i` to iterate over the original slice, and `j` to keep track of the position of the next element to keep.

```go
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Filter: keep only the even numbers
	n := 0
	for _, x := range s {
		if x%2 == 0 { // Our filter condition
			s[n] = x
			n++
		}
	}
	s = s[:n] // Shrink the slice to the new size

	fmt.Println(s) // Output: [2 4 6 8 10]
}
```
This approach is efficient because it avoids allocating a new backing array. It modifies the slice in-place.

---

### 4. Using a Slice as a Stack (LIFO)

Slices are perfect for implementing a stack (Last-In, First-Out).

*   **Push:** `stack = append(stack, value)`
*   **Pop:** `value := stack[len(stack)-1]; stack = stack[:len(stack)-1]`

```go
package main

import "fmt"

func main() {
	var stack []string

	// Push
	stack = append(stack, "A")
	stack = append(stack, "B")
	stack = append(stack, "C")
	fmt.Println("Stack:", stack)

	// Pop
	var value string
	value = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println("Popped:", value)
	fmt.Println("Stack:", stack)
}
```
This is both simple and efficient.
