---
date: "2025-09-18T00:00:00Z"
description:
header-img:
keywords: Go programming language, Golang speed, compiled languages, concurrency in Go, goroutines, Go garbage collection, high-performance applications, Go interview preparation, native compilation, modern hardware optimization, fast compile times, efficient coding practices, Golang features
tags:
- Go
title: 'Golang - Data Structures - Structs'
toc: true
---

### Introduction

A `struct` is a composite data type that groups together zero or more fields of arbitrary types into a single unit. If you're coming from another language, you can think of a struct as a lightweight `class` (without methods attached directly to it), a `record`, or a `struct` from C.

Structs are the primary way you create complex, custom data types in Go. They are fundamental to organizing and representing data.

**Analogy:** A `struct` is like a **blueprint for a form or a template**. For example, a "Person" template might have fields for "Name," "Age," and "City." Each individual person you create using this template is an *instance* of the struct.

---

### Defining and Initializing Structs

#### 1. Definition

You define a struct using the `type` and `struct` keywords.

```go
// Defines a blueprint for a Person
type Person struct {
    Name string
    Age  int
}
```

#### 2. Initialization

There are several ways to create an instance of a struct.

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// 1. Using the zero value
	// All fields are initialized to their zero value ("" for string, 0 for int)
	var p1 Person
	fmt.Printf("p1: %+v\n", p1) // p1: {Name: Age:0}

	// 2. Using a struct literal with field names (most readable and robust)
	// You can omit fields, and they will be zero-valued.
	p2 := Person{
		Name: "Alice",
		Age:  30,
	}
	fmt.Printf("p2: %+v\n", p2) // p2: {Name:Alice Age:30}

	// 3. Using a struct literal without field names (order matters!)
	// You must provide all fields in the exact order they are defined.
	// This is brittle; if you add a new field to the struct, this code will break.
	p3 := Person{"Bob", 25}
	fmt.Printf("p3: %+v\n", p3) // p3: {Name:Bob Age:25}

	// 4. Using the `new` keyword
	// `new` allocates memory for the struct, initializes it to zero, and returns a POINTER to it.
	p4 := new(Person)
	fmt.Printf("p4: %+v\n", p4) // p4: &{Name: Age:0}
	p4.Name = "Charlie" // You can access fields on the pointer directly
}
```
**Best Practice:** Use the struct literal with field names (`Person{Name: "Alice", ...}`). It's the most readable and the least likely to break if the struct definition changes.

---

### Structs are Value Types

This is a critical concept. Like `int` or `[5]int`, a `struct` is a **value type**. When you assign a struct to a new variable or pass it to a function, the **entire struct is copied**.

```go
package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	p1 := Point{X: 10, Y: 20}
	p2 := p1 // p2 is a complete copy of p1

	p2.X = 99 // Modifying the copy

	fmt.Println("p1:", p1) // p1: {10 20}
	fmt.Println("p2:", p2) // p2: {99 20}
}
```

Because of this, it is very common to pass **pointers to structs** (`*Point`) to functions, especially if the struct is large or if the function needs to modify the original struct. This avoids the expensive copy and allows for modification.

```go
func modifyPoint(p *Point) {
    p.X = 100 // Modifies the original struct
}

func main() {
    p1 := &Point{X: 10, Y: 20} // p1 is now a pointer to a Point
    modifyPoint(p1)
    fmt.Println("p1:", *p1) // p1: {100 20}
}
```

---

### Anonymous and Embedded Structs

#### Anonymous Structs
You can declare a struct type on the fly without giving it a name. This is useful for short-lived, simple data structures.

```go
var user struct {
    Name string
    ID   int
}
user.Name = "Admin"
user.ID = 1
```

#### Embedded Structs (Composition)
Go does not have inheritance. Instead, it favors **composition** through struct embedding. You can "embed" one struct inside another, and the fields of the inner struct are promoted to the outer one.

```go
type Employee struct {
	Person // Embedded field (type name acts as the field name)
	Salary int
}

func main() {
	emp := Employee{
		Person: Person{Name: "David", Age: 40},
		Salary: 80000,
	}

	// You can access the embedded struct's fields directly
	fmt.Println(emp.Name)   // "David"
	fmt.Println(emp.Age)    // 40
	fmt.Println(emp.Salary) // 80000

	// You can also access the embedded struct itself
	fmt.Println(emp.Person) // {David 40}
}
```
This is Go's idiomatic way to build complex types by combining simpler ones.