---
date: "2025-09-07T00:00:00Z"
description:
header-img:
keywords: Go programming language, Golang speed, compiled languages, concurrency in Go, goroutines, Go garbage collection, high-performance applications, Go interview preparation, native compilation, modern hardware optimization, fast compile times, efficient coding practices, Golang features
tags:
- Go
title: 'Golang - Data Structures - Slice Internals'
toc: true
---

# Introduction

Understanding how slices behave when they share a backing array is critical for avoiding bugs. The `append` function, in particular, has behavior that can be surprising if you don't understand the underlying mechanics of length and capacity.

---

### Gotcha #1: Shared Backing Arrays and Unintended Modifications

When one slice is created from another, they **share the same backing array**. Modifying the elements of one slice will affect the other.

```go
package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[1:4] // s2 is [2, 3, 4]

	fmt.Println("Before modification:")
	fmt.Printf("s1: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))

	// Modify an element in s2
	s2[0] = 99

	fmt.Println("\nAfter modification:")
	fmt.Printf("s1: %v\n", s1) // s1 is also changed!
	fmt.Printf("s2: %v\n", s2)
}
```

**Output:**

```
Before modification:
s1: [1 2 3 4 5], len: 5, cap: 5
s2: [2 3 4], len: 3, cap: 4

After modification:
s1: [1 99 3 4 5]
s2: [99 3 4]
```

This happens because both `s1` and `s2`'s pointers refer to the same underlying data. `s2[0]` is the same memory location as `s1[1]`.

---

### Gotcha #2: The `append` Function's Surprising Behavior

The `append` function is where most slice-related confusion comes from. Here's the rule:

1.  If the slice has **enough capacity** for the new elements, `append` will reuse the existing backing array. The original slice will be modified.
2.  If the slice does **not have enough capacity**, `append` will allocate a **new, larger backing array**, copy the elements over, and return a slice pointing to this new array.

This distinction is the source of many bugs.

#### Case 1: `append` with enough capacity

When `append` reuses the array, it can overwrite elements that are part of the original slice but outside the view of the new slice.

```go
package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[:3] // s2 is [1, 2, 3], but has a capacity of 5

	fmt.Printf("s1 before append: %v\n", s1)
	fmt.Printf("s2 before append: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))

	// s2 has enough capacity, so this reuses the backing array
	s2 = append(s2, 99)

	fmt.Println("\nAfter append:")
	fmt.Printf("s1 after append: %v\n", s1) // s1 is modified!
	fmt.Printf("s2 after append: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))
}
```

**Output:**

```
s1 before append: [1 2 3 4 5]
s2 before append: [1 2 3], len: 3, cap: 5

After append:
s1 after append: [1 2 3 99 5]
s2 after append: [1 2 3 99], len: 4, cap: 5
```
The `append` operation placed `99` at index 3 of the backing array, overwriting the original `4`.

#### Case 2: `append` without enough capacity

When `append` allocates a new array, the connection between the two slices is broken.

```go
package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := s1

	fmt.Printf("s1 before append: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))

	// s2 does NOT have enough capacity. A new array is allocated.
	s2 = append(s2, 4)

	// To prove they are separate, let's modify s2 again.
	s2[0] = 99

	fmt.Println("\nAfter append and modification:")
	fmt.Printf("s1: %v\n", s1) // s1 is unchanged!
	fmt.Printf("s2: %v\n", s2)
}
```

**Output:**

```
s1 before append: [1 2 3], len: 3, cap: 3

After append and modification:
s1: [1 2 3]
s2: [99 2 3 4]
```
Because `append` created a new backing array for `s2`, the link to `s1` was severed. Subsequent modifications to `s2` do not affect `s1`.

---

### How to Prevent These Gotchas

**Interview Question:** "How can you append to a slice without modifying the original?"

**Answer:** "You need to ensure that the new slice has its own backing array. The best way is to use the `copy` function."

1.  **Use `copy` for safe modifications:** If you want to modify a slice without affecting the original, create a new slice and `copy` the data into it.

    ```go
    s1 := []int{1, 2, 3}
    s2 := make([]int, len(s1))
    copy(s2, s1) // s2 now has its own backing array

    s2[0] = 99 // This will not affect s1
    ```

2.  **Be careful with function arguments:** When a function receives a slice and appends to it, it might be modifying the caller's slice. Functions that modify a slice's length or capacity should always return the new slice.

    ```go
    // This is the idiomatic way to write a function that appends
    func addElement(s []int, value int) []int {
        return append(s, value)
    }

    func main() {
        mySlice := []int{1, 2}
        mySlice = addElement(mySlice, 3) // Always re-assign the result
    }
    ```
    This pattern works correctly regardless of whether `append` allocates a new array or not.

	[Common Slice Operations](/Golang-Data-Structures-Common-Slice-Operations/)