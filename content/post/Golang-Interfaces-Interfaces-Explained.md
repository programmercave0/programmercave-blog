---
date: "2025-09-24T00:00:00Z"
description:
header-img:
keywords: Go programming language, Golang speed, compiled languages, concurrency in Go, goroutines, Go garbage collection, high-performance applications, Go interview preparation, native compilation, modern hardware optimization, fast compile times, efficient coding practices, Golang features
tags:
- Go
title: 'Golang - Interfaces - Interfaces Explained'
toc: true
---

### Introduction

An interface in Go is a type that defines a set of method signatures. It's a contract that specifies what a type can *do*. If a type provides definitions for all the methods in an interface, it is said to **satisfy** that interface.

**Analogy:** Think of an interface as a **job description**.
*   The job description for a `Writer` might say, "Must be able to `Write(data []byte)`."
*   Any person (or type) that has this exact skill (method) can perform the job of a `Writer`. It doesn't matter if they are a `File`, a `NetworkConnection`, or a `Terminal`. As long as they can `Write`, they fit the contract.

---

### Implicit Implementation: The Core of Go Interfaces

This is the most important concept to understand. In many other languages, a class must explicitly declare that it implements an interface (e.g., `class MyWriter implements Writer`).

In Go, the implementation is **implicit**. If a type has all the methods required by an interface, it automatically satisfies that interface. There is no `implements` keyword.

**Why is this so powerful?**
*   **Decoupling:** You can write code that depends on an interface without knowing about the specific types that implement it. This makes your code more modular and reusable.
*   **Flexibility:** New types can satisfy existing interfaces without any modification to the original code.

#### Example: A `Shape` Interface

Let's define our own interface called `Shape`. A shape is anything that has an area.

```go
// Shape is an interface for anything that has an area.
type Shape interface {
    Area() float64
}
```

Now, any type that has an `Area() float64` method will automatically satisfy this interface.

```go
package main

import (
	"fmt"
	"math"
)

// Shape is an interface for anything that has an area.
type Shape interface {
	Area() float64
}

// Circle is a struct representing a circle.
type Circle struct {
	Radius float64
}

// By implementing this method, Circle now satisfies the Shape interface.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Rectangle is a completely different type.
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle also satisfies the Shape interface.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	myCircle := Circle{Radius: 5}
	myRectangle := Rectangle{Width: 10, Height: 4}

	// We can use a function that accepts any Shape.
	printArea(myCircle)
	printArea(myRectangle)
}

// This function works with ANY type that satisfies the Shape interface.
func printArea(s Shape) {
	fmt.Printf("The area of the shape is %0.2f\n", s.Area())
}
```

---

### Polymorphism in Go

Interfaces are Go's primary tool for polymorphism. You can write a single function that operates on an interface type, and that function can then accept values of any concrete type that satisfies the interface.

```go
// This function can take any type that satisfies the Shape interface.
func printArea(s Shape) {
    fmt.Printf("The area of the shape is %0.2f\n", s.Area())
}

func main() {
    myCircle := Circle{Radius: 5}
    myRectangle := Rectangle{Width: 10, Height: 4}

    printArea(myCircle)    // Works with a Circle
    printArea(myRectangle) // Works with a Rectangle
}
```

---

### Key Takeaways for Interviews

*   **What is an interface?** It's a contract—a collection of method signatures. It defines behavior, not data.
*   **How is it implemented?** Implicitly. If a type has the methods, it satisfies the interface. No `implements` keyword.
*   **Why is this useful?** It promotes decoupling and allows for polymorphism, making code flexible and reusable. You can write functions that operate on behaviors (interfaces) rather than on specific types.
