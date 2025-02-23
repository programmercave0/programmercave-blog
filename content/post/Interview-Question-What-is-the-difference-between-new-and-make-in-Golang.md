---
date: "2025-02-22T00:00:00Z"
description: Understand the difference between Golang's `new()` and `make()` with simple examples. Learn when to use each for slices, maps, and pointers, and prepare effectively for software engineering interviews.
header-img: "/assets/images/Golang/new-make-golang.png"
keywords: new() vs make(), memory allocation, Go programming, software engineering interview, Go memory management, slices maps channels, Go pointers, Go language, coding interview prep.
tags:
- Go
- Interview-Questions
title: 'Interview Question: What is the difference between new() and make() in Golang?'
toc: true
---

## Key Takeaways
- **new()** allocates memory and returns a pointer to zeroed memory (e.g., `*int`, `*struct`).  
- **make()** initializes slices, maps, or channels and returns a ready-to-use value (not a pointer).  
- Use **new()** for basic types (int, bool) or structs when you need a pointer.  
- Use **make()** for slices, maps, or channels to avoid runtime errors like nil pointers.  
- Both functions handle memory automatically—no manual cleanup is required.  

![What is the difference between new and make in Golang](/assets/images/Golang/new-make-golang.png)

## Why Understanding `new()` and `make()` Matters in Go  

In Go, memory management is handled differently compared to languages like C or Java. Misusing `new()` or `make()` can lead to runtime errors such as nil pointer dereferences or uninitialized data structures. Knowing the difference between these two functions is a common topic in software engineering interviews.  

This guide will explain everything you need to know about `new()` and `make()` in Go, including their differences, use cases, and common mistakes to avoid.

## What Does `new()` Do?  

The `new()` function in Go allocates memory for a given type and returns a pointer to it. The allocated memory is initialized to the type’s zero value (e.g., `0` for integers, `""` for strings, `false` for booleans).  

### Syntax  
```go
pointer := new(Type) // Returns *Type
```

### Example  
```go
// Allocating a pointer to an int
numPtr := new(int)
fmt.Println(*numPtr) // Output: 0 (zero value of int)

// Allocating a pointer to a struct
type Person struct { Name string; Age int }
personPtr := new(Person)
fmt.Println(personPtr) // Output: &{ 0 }
```

### When to Use `new()`  
- You need a pointer to modify a value across functions.  
- Working with basic types (e.g., `int`, `string`) or structs.  

## What Does `make()` Do?  

The `make()` function in Go is used to initialize slices, maps, and channels so they are ready to use. Unlike `new()`, `make()` returns the actual value (not a pointer).  

### Syntax  
```go
slice := make([]Type, length, capacity)  
m := make(map[KeyType]ValueType)  
ch := make(chan Type)
```

### Example  
```go
// Slice with length 3
names := make([]string, 3)
names[0] = "Alice" // Works (initialized)
fmt.Println(names) // Output: [Alice  ]

// Map
scores := make(map[string]int)
scores["Bob"] = 85 // No error
```

### When to Use `make()`  
- Creating slices, maps, or channels.  
- Avoiding runtime errors like “nil map” or “slice out of bounds”.  

## `new()` vs `make()`: Key Differences  

Here’s a table summarizing the key differences between `new()` and `make()`:  

| Feature         | `new()`                          | `make()`                          |  
|--|-|--|  
| **Returns**     | Pointer (`*Type`)                | Initialized value (`Type`)        |  
| **Used For**    | Any type (int, struct, etc.)     | Only slices, maps, channels       |  
| **Initializes** | Zero value (0, "", false)        | Ready-to-use structure            |  

## Memory Management  

Go simplifies memory management by handling it automatically:  
- Both `new()` and `make()` allocate memory on the **heap**.  
- Go’s garbage collector automatically frees memory when it’s no longer used.  
- You don’t need to manually delete pointers or values.  

## Conclusion

Understanding the differences between new() and make() is essential for writing effective Go code. Using these functions correctly will help you allocate memory properly, initialize data structures, and create efficient, idiomatic Go programs. With this knowledge, you will be well-prepared for your software engineering interview.