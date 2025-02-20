---
date: "2025-02-20T00:00:00Z"
description: Learn how Go's garbage collector works using a concurrent, tri-color mark-sweep algorithm. Simple explanations for software engineering interview prep.
header-img: "/assets/images/Golang/Garbage-Collection-in-Go.png"
keywords: Go garbage collection, tri-color algorithm, mark-sweep, write barriers, concurrent GC, memory management, interview preparation, software engineer interview
tags:
- Go
- Interview-Questions
title: 'Interview Question: How Garbage Collection in Go works?'
toc: true
---

## Key Takeaways
- **Concurrent Execution**: Go’s garbage collector runs in the background, minimizing pauses.
- **Tri-Color Marking**: Uses white, grey, and black labels to identify unused memory efficiently.
- **Low Latency**: Short pauses make Go ideal for real-time applications like servers.
- **Write Barriers**: Prevent accidental deletion of active objects during garbage collection.
- **Interview-Ready**: Understanding these concepts helps you answer memory management questions confidently.

![How Garbage Collection in Go works?](/assets/images/Golang/Garbage-Collection-in-Go.png)

## What is Garbage Collection?

Garbage collection (GC) is an automatic process that frees up memory used by programs. In Go, it acts like a background cleaner that identifies and deletes unused objects while your application runs. This ensures that your program doesn't run out of memory and maintains optimal performance.

## How Go’s Garbage Collector Works: A Step-by-Step Guide

### 1. Concurrent Execution

#### Problem with Traditional GC
- Traditional garbage collectors stop the entire program temporarily (stop-the-world), causing delays.

#### Go’s Solution
- Go’s garbage collector runs concurrently alongside your code, minimizing pauses.

#### Example
- Imagine cleaning a room while people are still working in it, instead of waiting for everyone to leave.

### 2. Tri-Color Marking Algorithm

Go uses a tri-color marking algorithm to label objects and decide what to delete:

| Color | Meaning |
|-------|---------|
| White | Not checked yet (might be garbage). |
| Grey  | Checked but needs further inspection. |
| Black | Confirmed as active (not garbage). |

#### How It Works
- **Start**: All objects are initially white.
- **Mark Roots**: Global variables, active functions, etc., are marked grey.
- **Process Grey Objects**:
  - Inspect references in a grey object.
  - Turn its referenced white objects grey.
  - Turn the original object black.
- **Repeat** until no grey objects remain.
- **Delete**: All remaining white objects are deleted.

#### Why It’s Efficient
- No long pauses.
- Only scans active objects.

### 3. Mark-Sweep Phases

- **Mark Phase**: Identifies active objects using the tri-color marking.
- **Sweep Phase**: Deletes white (unused) objects.
- **Concurrent Execution**: Both phases run in the background, ensuring minimal disruption.

## Write Barriers: Keeping Things Safe

When the garbage collector and your application run simultaneously, issues can arise. For example:
- If a black object (marked as active) suddenly references a white object (unchecked), the white object might get deleted by mistake.

#### How Write Barriers Help
- **How**: If a black object references a white object, the write barrier turns the white object grey.
- **Result**: The garbage collector rechecks the white object, ensuring it’s not deleted prematurely.

## Summary

Go’s garbage collector is designed for performance and simplicity. By combining concurrent execution, tri-color marking, and write barriers, it keeps applications fast and memory-efficient. For interviews, focus on explaining these concepts with real-world examples (e.g., servers handling thousands of requests).