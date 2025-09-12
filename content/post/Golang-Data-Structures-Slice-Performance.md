---
date: "2025-09-09T00:00:00Z"
description:
header-img:
keywords: Go programming language, Golang speed, compiled languages, concurrency in Go, goroutines, Go garbage collection, high-performance applications, Go interview preparation, native compilation, modern hardware optimization, fast compile times, efficient coding practices, Golang features
tags:
- Go
title: 'Golang - Data Structures - Slice Performance'
toc: true
---

### Introduction

For experienced engineers, writing code that is not only correct but also performant is a key requirement. Understanding the performance implications of how you use slices can have a significant impact on your application's speed and memory usage.

---

### 1. The Cost of `append`: Growth Strategy

When a slice's capacity is exceeded, `append` allocates a new, larger backing array and copies the old elements to it. This copy operation is expensive. To minimize the number of allocations, Go's runtime uses a growth strategy.

**The Rule (as of Go 1.18):**
*   For slices with a capacity of less than 256 elements, the capacity is **doubled**.
*   For slices with a capacity of 256 or more, the capacity is grown by a factor of **1.25** (i.e., `newCap = oldCap + (oldCap+3*256)/4`).

This strategy tries to balance memory usage with the number of allocations. Doubling small slices is cheap and avoids frequent re-allocations. For larger slices, a smaller growth factor prevents wasting too much memory.

**Interview Question:** "What happens when you `append` to a full slice? How does Go decide the new capacity?"
**Answer:** "If a slice's capacity is exceeded, `append` allocates a new backing array. The runtime employs a growth strategy: for small slices (under 256 elements), it doubles the capacity. For larger slices, it increases it by a smaller factor (roughly 1.25x) to avoid excessive memory consumption."

---

### 2. Pre-allocation: The Power of `make`

If you know roughly how many elements you are going to put in a slice, you can and **should** pre-allocate it with the required capacity using `make`.

**Why is pre-allocation so important?**
*   **It avoids re-allocations:** By setting the capacity upfront, you can avoid the expensive cycle of re-allocation and copying that happens when you `append` to a zero-length slice in a loop.
*   **It reduces memory fragmentation:** A single large allocation is often better for the memory manager than many small ones.

#### Scenario: No Pre-allocation (Bad)

```go
// BAD: This will cause 
// multiple re-allocations
var s []int
for i := 0; i < 1000; i++ {
    s = append(s, i)
}
```
In this example, the slice `s` starts with a capacity of 0. As the loop runs, it will be re-allocated and copied multiple times (e.g., capacity goes 0 -> 1 -> 2 -> 4 -> 8 -> 16...).

#### Scenario: With Pre-allocation (Good)

```go
// GOOD: One allocation
// no copying
s := make([]int, 0, 1000) 
// Length 0, Capacity 1000
for i := 0; i < 1000; i++ {
    s = append(s, i)
}
```
Here, we create a slice with a length of 0 but a capacity of 1000. The loop can add 1000 elements without a single re-allocation, making it significantly faster and more memory-efficient.

**Rule of Thumb:** Whenever you have a loop that appends to a slice, and you know the number of iterations, pre-allocate the slice's capacity.

---

### 3. "Leaky" Slices and the `copy` function

As discussed in the "Gotchas" note, creating a sub-slice can lead to memory leaks if the original backing array is large. The small sub-slice holds a reference to the entire large array, preventing it from being garbage collected.

**The Problem:**
```go
func getSmallSlice() []byte {
    largeData := make([]byte, 10*1024*1024) // 10 MB
    // ... fill largeData ...
    return largeData[:5] // Returns a small slice that pins 10MB in memory
}
```

**The Solution:**
When returning a small slice from a large one, always make a copy. This ensures the new slice has its own, tightly-sized backing array.

```go
func getSmallSliceFixed() []byte {
    largeData := make([]byte, 10*1024*1024) // 10 MB
    // ... fill largeData ...

    result := make([]byte, 5)
    copy(result, largeData[:5])
    return result // Returns a new slice with its own 5-byte backing array
}
```
This allows the `largeData` array to be garbage collected, saving a significant amount of memory. This is a critical optimization in memory-sensitive applications.
