---
date: "2025-09-05T00:00:00Z"
description:
header-img:
keywords: Go programming language, Golang speed, compiled languages, concurrency in Go, goroutines, Go garbage collection, high-performance applications, Go interview preparation, native compilation, modern hardware optimization, fast compile times, efficient coding practices, Golang features
tags:
- Go
title: 'Golang - Data Structures - Arrays vs Slices'
toc: true
---

### Introduction

Arrays and Slices are both used for storing sequences of data, but they have fundamental differences in how they work. 

---

### The Core Difference: Fixed vs. Dynamic

*   **Array:** A fixed-size collection of elements of the same type. The size is part of its type definition.
*   **Slice:** A dynamic, flexible view into the elements of an underlying array.

**Analogy:** An **array** is like a **shipping carton with a fixed number of slots**. If you order a carton for 6 items, you get exactly 6 slots. You can't add a 7th. A **slice** is like a **movable, resizable window** that you can place over a set of items in a warehouse (the underlying array).

---

### Feature Comparison: Array vs. Slice

| Feature               | Array                                           | Slice                                                              |
| --------------------- | ----------------------------------------------- | ------------------------------------------------------------------ |
| **Size**              | **Fixed**. The length is part of its type.      | **Dynamic**. Can grow and shrink.                                  |
| **Type Declaration**  | `[5]int` (length is specified)                  | `[]int` (length is omitted)                                        |
| **Memory**            | Is a value type. Contains the data directly.    | Is a reference type. Contains a pointer to an underlying array.    |
| **Assignment**        | **Copies the entire array.**                    | **Copies the slice header.** Both slices point to the *same* data. |
| **Function Passing**  | Passed **by value** (the whole array is copied). | Passed **by value** (the slice header is copied, but not the data). |
| **Use Case**          | Rarely used directly. Good for fixed data sets. | The idiomatic choice for almost all collections in Go.             |

---

### 1. Arrays: The Rigid Foundation

An array's type is defined by both its element type and its size. This means `[3]int` and `[4]int` are completely different, incompatible types.

#### Code Example: Arrays are Value Types

When you assign or pass an array, you are making a full copy of its data.

```go
package main

import "fmt"

func main() {
	// arr1 is an array of 3 integers
	arr1 := [3]int{10, 20, 30}

	// arr2 is a COPY of arr1
	arr2 := arr1

	// Modifying arr2 does NOT affect arr1
	arr2[0] = 999

	fmt.Println("arr1:", arr1) // Output: arr1: [10 20 30]
	fmt.Println("arr2:", arr2) // Output: arr2: [999 20 30]
}
```

This copy-on-assignment behavior makes arrays inefficient for passing around in your application.

---

### 2. Slices: The Flexible View

A slice is a lightweight data structure that "describes" a contiguous section of an underlying array. It doesn't store any data itself.

#### Code Example: Slices are Reference Types

When you assign or pass a slice, you copy the slice header, but both the original and the copy point to the **same underlying array**.

```go
package main

import "fmt"

func main() {
	// slice1 is a slice of integers
	slice1 := []int{10, 20, 30}

	// slice2 refers to the SAME underlying data as slice1
	slice2 := slice1

	// Modifying slice2 WILL affect slice1
	slice2[0] = 999

	fmt.Println("slice1:", slice1) // Output: slice1: [999 20 30]
	fmt.Println("slice2:", slice2) // Output: slice2: [999 20 30]
}
```

This behavior makes slices very efficient to work with. Modifying a slice inside a function will be visible to the caller, which is usually the desired behavior.

[Slices In-Depth](/Golang-Data-Structures-Slices-In-Depth/)

[Mutability in Arrays vs. Slices](/Golang-Data-Structures-Mutability-in-Arrays-vs-Slices/)

---

### How to Choose?

*   **Always prefer slices.** They are the idiomatic and flexible way to handle collections in Go.
*   Use an array only when you have a very specific reason to, such as:
    *   You need a truly fixed-size collection and want the compiler to enforce it.
    *   You are working with low-level code where the precise memory layout is critical (e.g., interacting with C libraries).

**Interview Question:** "What's the difference between an array and a slice?"
**Answer:** "An array has a fixed size that's part of its type, and it's a value type, meaning it's copied on assignment. A slice is a dynamic, reference type that provides a flexible view over an underlying array. Slices are far more common and idiomatic in Go."
