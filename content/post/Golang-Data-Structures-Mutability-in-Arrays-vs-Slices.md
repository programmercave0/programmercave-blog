---
date: "2025-09-10T00:00:00Z"
description:
header-img:
keywords: Go programming language, Golang speed, compiled languages, concurrency in Go, goroutines, Go garbage collection, high-performance applications, Go interview preparation, native compilation, modern hardware optimization, fast compile times, efficient coding practices, Golang features
tags:
- Go
title: 'Golang - Data Structures - Mutability in Arrays vs. Slices'
toc: true
---

### Introduction

Both arrays and slices in Go allow you to change their elements after creation—this is what we mean by "mutability." But how changes affect other parts of your code differs a lot between the two. This note builds on the basics of arrays and slices by focusing on mutability in practice, including how to control it safely in larger programs. We'll cover advanced tips useful for experienced developers, like handling shared changes in concurrent code or designing functions that avoid surprises.

***

### The Core of Mutability: Changes and Sharing

*   **Mutability basics:** You can always update elements in both arrays and slices using index assignment (like `a = newValue`). The big difference is in how those changes "spread" when you have multiple variables or functions involved.
*   **Array mutability:** Changes are isolated because arrays copy everything. This makes them safe but can be slow for big data.
*   **Slice mutability:** Changes are often shared because slices point to the same hidden storage. This is efficient but can lead to unexpected updates if not managed.

**Analogy:** Updating an array is like editing a personal notebook—your changes stay in your copy. Updating a slice is like editing a shared document—everyone with access sees the changes instantly, unless someone makes a private copy.

***

### Controlling Slice Mutability with Capacity Limits

Slices can share storage, but you can limit this sharing by controlling how much room a slice has to grow without creating new storage.

*   **Full slice expression:** Use `s[low:high:max]` to set a custom limit on growth space (capacity = max - low). This forces any growth to create fresh storage, preventing unwanted shared changes.
*   **Why it helps mutability:** It stops one slice from accidentally overwriting another's data during growth.

#### Code Example: Limiting Shared Changes

```go
package main

import "fmt"

func main() {
    parent := []int{10, 20, 30, 40, 50}

    // Unlimited growth space: shares with parent
    unlimited := parent[1:3] // [20 30], can grow into parent's space

    // Limited growth space: forces new storage on growth
    limited := parent[1:3:3] // [20 30], max=3 so capacity=2 (can't grow without new storage)

    unlimited = append(unlimited, 99) // Might change parent
    limited = append(limited, 99)     // Always creates new storage, safe from changing parent

    fmt.Println("Parent after unlimited append:", parent) // Could be changed
    fmt.Println("Parent after limited append:", parent)   // Unchanged
}
```

**Output (may vary, but shows the idea):**
```
Parent after unlimited append: [10 20 30 99 50] // Shared change happened
Parent after limited append: [10 20 30 99 50]   // No further change
```

This technique is great for handing out "read-only views" that can't accidentally mutate the original through growth.

***

### Safe Copying to Break Mutability Sharing

To make sure changes don't affect the original, create a full copy with its own storage.

*   **Using slices.Clone (Go 1.21+):** This built-in function makes a new slice with copied elements and independent storage. It's clearer than older tricks.
*   **Why for mutability:** It gives you a "private" version where changes stay isolated, like making a photocopy before marking it up.

#### Code Example: Cloning for Independent Changes

```go
package main

import (
    "fmt"
    "slices" // Need this for slices.Clone
)

func main() {
    original := []string{"apple", "banana", "cherry"}

    cloned := slices.Clone(original) // New storage, safe to change

    cloned[0] = "avocado" // Only affects cloned

    fmt.Println("Original:", original) // Unchanged
    fmt.Println("Cloned:", cloned)     // Updated
}
```

**Output:**
```
Original: [apple banana cherry]
Cloned: [avocado banana cherry]
```

For older Go versions, use `append([]T(nil), original...)` instead—it's the same idea but less readable.

***

### Mutability in Concurrent Code

When multiple parts of your program run at the same time (using goroutines), mutability can cause "data races"—unpredictable bugs from unsynchronized changes.

*   **Safe concurrent writes:** You can update different positions in a pre-allocated slice from different goroutines without locks, as long as no one changes the slice's size or creates sub-views during the updates.
*   **Risks to avoid:** Don't mix growth or sub-view creation with concurrent updates—these can cause crashes or wrong data. Use tools like the race detector to catch issues.

**Analogy:** Think of concurrent updates like multiple chefs chopping veggies on separate cutting boards (different positions). It's fine until someone rearranges the boards mid-chop.

#### Code Example: Concurrent Updates Without Races

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    s := make([]int, 5) // Pre-allocate to avoid growth
    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(idx int) {
            defer wg.Done()
            s[idx] = idx * 10 // Update own position safely
        }(i)
    }

    wg.Wait()
    fmt.Println(s) // Output like [0 10 20 30 40] (order may vary but correct)
}
```

Run with `go run -race main.go` to check for races—none here because positions are separate and no size changes.

***

### Mutability in Methods and Functions

How you define methods affects whether changes stick to the original data.

*   **Array methods:** Use a pointer receiver (like `func (a *[N]T) Method()`) to make changes visible outside the method, since non-pointer receivers work on copies.
*   **Slice methods:** Non-pointer receivers can still update elements (shared storage), but return the slice if you change its size to update the caller's view.

This is key for designing types that wrap arrays or slices, ensuring changes behave as expected.

***

### Interview Question and Answer

**Interview Question:** "How do you prevent a sub-slice from mutating the parent's data during append?"

**Answer:** "Use the full slice expression like `parent[low:high:max]` to limit capacity, forcing append to allocate new storage. Or clone the sub-slice with `slices.Clone` for full independence."

This note adds practical ways to manage mutability without repeating basics, helping you build reliable code in teams or high-performance systems. For 7+ years engineers: Remember, mutability control is about ownership—decide early if data should be shared or copied to avoid debugging nightmares in production.