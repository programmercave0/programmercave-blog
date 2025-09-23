---
date: "2025-09-21T00:00:00Z"
description:
header-img:
keywords: Go programming language, Golang speed, compiled languages, concurrency in Go, goroutines, Go garbage collection, high-performance applications, Go interview preparation, native compilation, modern hardware optimization, fast compile times, efficient coding practices, Golang features
tags:
- Go
title: 'Golang - Memory - Pointers Explained'
toc: true
---

### Introduction

A pointer is a variable that stores the memory address of another variable. Instead of holding a value like `5` or `"hello"`, it holds the location where that value lives in your computer's memory.

**Analogy:** Think of a variable as a **house** with a specific address.
*   A normal variable (`var x int = 10`) is like knowing what's *inside* the house (the value `10`).
*   A pointer (`var p *int = &x`) is like having a piece of paper with the *address* of the house written on it. It doesn't tell you what's inside, but it tells you where to find it.

---

### The Two Key Pointer Operators

You only need to know two operators to work with pointers:

1.  **`&` (The "Address Of" Operator):**
    *   When placed before a variable, it returns that variable's memory address.
    *   `&x` means "give me the address of the house `x`."

2.  **`*` (The "Dereference" Operator):**
    *   When placed before a pointer variable, it gives you the value stored at that address. It "follows the address" to see what's inside the house.
    *   `*p` means "go to the address stored in `p` and give me the value you find there."

---

### Declaring and Using Pointers

```go
package main

import "fmt"

func main() {
	// 1. Declare a regular variable.
	// This is the "house" with the value 42 inside.
	answer := 42

	// 2. Declare a pointer variable.
	// `*int` means this pointer can only hold the address of an integer.
	// We use the `&` operator to get the address of `answer`.
	var pointerToAnswer *int = &answer

	// 3. Print the values.
	fmt.Println("Value of answer:", answer)
	fmt.Println("Address of answer:", &answer)
	fmt.Println("Value of pointerToAnswer (it's an address):", pointerToAnswer)

	// 4. Dereference the pointer to get the original value.
	// We use the `*` operator to "go to the address".
	valueAtPointer := *pointerToAnswer
	fmt.Println("Value at the address the pointer points to:", valueAtPointer)

	// 5. Modify the original value *through* the pointer.
	// This is the most powerful feature of pointers.
	*pointerToAnswer = 99 // Go to the house and change what's inside.
	fmt.Println("New value of answer (changed via pointer):", answer)
}
```

**Output:**
```
Value of answer: 42
Address of answer: 0x... (some memory address)
Value of pointerToAnswer (it's an address): 0x... (the same memory address)
Value at the address the pointer points to: 42
New value of answer (changed via pointer): 99
```

---

### Why Use Pointers?

Pointers are essential for two main reasons in Go:

#### 1. Modifying a Function's Arguments

Go passes all function arguments **by value**, meaning it passes a *copy*. If you want a function to modify the original variable, you must pass a pointer to it.

```go
func double(val int) {
	val = val * 2 // Modifies the copy, not the original
}

func doubleWithPointer(val *int) {
	*val = *val * 2 // Modifies the original value
}

func main() {
	num := 5
	double(num)
	fmt.Println("After double():", num) // Still 5

	doubleWithPointer(&num)
	fmt.Println("After doubleWithPointer():", num) // Now 10
}
```

#### 2. Efficiency with Large Structs

Copying large structs can be expensive. Passing a pointer to a struct is much cheaper, as you are only copying a memory address (typically 8 bytes), not the entire struct's data.

```go
type BigStruct struct {
    // ...imagine many fields here...
}

// Inefficient: copies the entire struct
func process(b BigStruct) { /* ... */ }

// Efficient: copies only a pointer
func processWithPointer(b *BigStruct) { /* ... */ }
```

---

### The Zero Value of a Pointer: `nil`

The zero value for a pointer is `nil`. A `nil` pointer does not point to any memory address. Attempting to dereference a `nil` pointer will cause a runtime **panic**.

```go
var p *int // p is nil
fmt.Println(p) // <nil>
// *p = 10 // This would PANIC!
```
Always check if a pointer is `nil` before trying to use it if you are not sure it's been initialized.
