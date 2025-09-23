---
date: "2025-09-22T00:00:00Z"
description:
header-img:
keywords: Go programming language, Golang speed, compiled languages, concurrency in Go, goroutines, Go garbage collection, high-performance applications, Go interview preparation, native compilation, modern hardware optimization, fast compile times, efficient coding practices, Golang features
tags:
- Go
title: 'Golang - Memory - Garbage Collector Explained'
toc: true
---

### Introduction

Go provides automatic memory management, which means you don't have to manually free the memory you allocate. This is the job of the **Garbage Collector (GC)**. The Go GC is a highly sophisticated piece of engineering designed to have a minimal impact on your application's performance.

**Analogy:** Think of the GC as a team of **diligent, efficient cleaners** working in the "heap" warehouse. Their job is to periodically walk through the warehouse, identify all the boxes that are no longer labeled for use by any part of the factory (your program), and haul them away to make space for new ones.

---

### The Core Algorithm: Concurrent Tri-Color Mark-Sweep

Go's GC uses a **concurrent, tri-color, mark-and-sweep** algorithm. Let's break down what that means.

*   **Concurrent:** This is the most important feature. The GC does most of its work *at the same time* as your application is running. This avoids long "stop-the-world" (STW) pauses where your entire application freezes. Go's GC is famous for its extremely short pause times (often sub-millisecond).
*   **Mark-and-Sweep:** This is the overall strategy.
    1.  **Mark Phase:** The GC finds all the objects on the heap that are still in use ("reachable").
    2.  **Sweep Phase:** The GC scans the entire heap and reclaims the memory of any object that was not marked.
*   **Tri-Color:** This is the specific algorithm used during the Mark phase to keep track of objects. It's a clever way to allow the marking to happen concurrently with the main program.

---

### How the Tri-Color Algorithm Works

The algorithm conceptually paints each object one of three colors:

1.  **White:** The initial state. A white object is a candidate for garbage collection. Think of it as "trash until proven otherwise."
2.  **Grey:** An object that is reachable (not trash), but its fields (pointers to other objects) have not yet been scanned. A grey object is on the "to-do" list.
3.  **Black:** An object that is reachable, and all of its fields have been scanned. A black object is "done and safe."

The process works like this:

1.  **Start:** The GC pauses the program for a very short time (the STW pause), scans for all the "roots" (global variables and active stack variables), and colors them **grey**.
2.  **Marking (Concurrent):** The program resumes. The GC workers now pick up any **grey** object. They scan it for pointers to other objects.
    *   When they find a pointer to a **white** object, they color that object **grey** (adding it to the "to-do" list).
    *   Once a **grey** object has been fully scanned, it is colored **black**.
3.  **Finish:** This continues until there are no more **grey** objects left. At this point, any object still **white** is unreachable and is garbage.
4.  **Sweeping (Concurrent):** The GC now sweeps through memory and reclaims all the **white** objects.

---

### The Write Barrier: A Critical Helper

There's a problem with the concurrent process described above. What if your program, while the GC is running, does this:
`blackObject.pointer = whiteObject`

A **black** object (which the GC thinks is "done") now points to a **white** object (which the GC thinks is "trash"). If the GC doesn't find out about this new pointer, it will incorrectly delete the white object!

This is solved by a **Write Barrier**. A write barrier is a small piece of code that the compiler inserts whenever your program writes a pointer. This barrier checks for the `black -> white` situation. If it happens, the barrier colors the **white** object **grey**, putting it back on the GC's "to-do" list and ensuring it won't be collected.

---

### Key Takeaways for Interviews

*   **What kind of GC does Go use?**
    *   A concurrent, tri-color, mark-and-sweep garbage collector.
*   **What is the main goal of Go's GC?**
    *   **Low latency.** The GC is optimized to minimize "stop-the-world" pause times, making Go suitable for responsive services.
*   **How does it achieve this?**
    *   By running most of its work **concurrently** with the main program.
*   **What is a write barrier?**
    *   It's a mechanism that ensures the GC doesn't miss any pointers that are created or moved while it's running. It prevents the GC from incorrectly collecting an object that is still in use.
*   **When does the GC run?**
    *   The GC is triggered automatically based on the amount of new heap memory allocated since the last run. The trigger point is controlled by the `GOGC` environment variable, which defaults to 100. `GOGC=100` means the next GC will start when the heap size is 100% larger than it was after the previous GC finished.
