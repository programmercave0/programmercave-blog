---
date: "2025-09-25T00:00:00Z"
description:
header-img:
keywords: Go programming language, Golang speed, compiled languages, concurrency in Go, goroutines, Go garbage collection, high-performance applications, Go interview preparation, native compilation, modern hardware optimization, fast compile times, efficient coding practices, Golang features
tags:
- Go
title: 'Golang - Interfaces - Composition and Embedding'
toc: true
---

### Introduction

Go does not have classes or inheritance in the traditional object-oriented sense. Instead, it favours a simpler and more flexible approach to code reuse and building complex types: **composition**.

The guiding principle is **"Composition over Inheritance."** Instead of a type *being* another type (inheritance), a type *has* another type (composition).

---

### 1. Composition (The "Has-A" Relationship)

This is the most basic form of composition. You simply include one struct as a field within another.

**Analogy:** A `Car` doesn't *inherit* from an `Engine`; a `Car` *has an* `Engine`.

```go
package main

import "fmt"

type Engine struct {
	Horsepower int
}

func (e *Engine) Start() {
	fmt.Println("Engine started.")
}

type Car struct {
	Make   string
	Model  string
	Engine Engine 
	// The Car "has an" Engine
}

func main() {
	myCar := Car{
		Make:  "Ford",
		Model: "Mustang",
		Engine: Engine{
			Horsepower: 450,
		},
	}

	// To access the engine's fields or methods, you go through the `Engine` field.
	fmt.Printf("My %s %s has %d HP.\n", myCar.Make, myCar.Model, myCar.Engine.Horsepower)
	myCar.Engine.Start()
}
```
This is simple, explicit, and easy to understand.

---

### 2. Embedding (Go's "Inheritance")

Embedding is a special, more convenient form of composition. When you embed a type, its fields and methods are **promoted** to the containing struct, so you can access them directly without going through the embedded type's field name.

You embed a type by declaring it in a struct without giving it a field name. The type's name is used implicitly.

**Analogy:** Embedding is like hiring a **specialist contractor**. If you hire a `Manager` for your `Team`, you can give instructions directly to the `Manager` (`team.Manage()`), and it feels like the `Team` itself has that skill.

#### a) Struct Embedding

```go
package main

import "fmt"

type Report struct {
	Content string
}

func (r *Report) Generate() {
	fmt.Println("Generating report:", r.Content)
}

type FinancialReport struct {
	Report // Embed the Report struct
	Amount float64
}

func main() {
	finReport := FinancialReport{
		Report: Report{Content: "Q1 Earnings"},
		Amount: 100000.0,
	}

	// Because `Report` is embedded, its fields and methods are "promoted".
	// You can access them directly on `FinancialReport`.
	fmt.Println(finReport.Content) // Instead of finReport.Report.Content
	finReport.Generate()           // Instead of finReport.Report.Generate()

	// You can still access the embedded struct explicitly if needed.
	finReport.Report.Content = "Q2 Earnings"
	fmt.Println(finReport.Content)
}
```
This gives you the convenience of inheritance-like syntax, but it's still composition under the hood.

**Method "Overriding":** Go doesn't have overriding, but if the outer struct defines a method with the same name as an embedded one, the outer method will be promoted. The embedded method is still accessible via the explicit field name.

```go
func (fr *FinancialReport) Generate() {
    fmt.Println("Generating FINANCIAL report:", fr.Content)
}

// finReport.Generate() now calls the FinancialReport's method.
// finReport.Report.Generate() would call the original embedded method.
```

#### b) Interface Embedding

You can also embed interfaces within other interfaces. This is a powerful way to build up larger contracts from smaller, reusable pieces.

```go
// A small, focused interface
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Another small, focused interface
type Writer interface {
    Write(p []byte) (n int, err error)
}

// A larger interface composed by embedding smaller ones.
// A type must satisfy both Read() and Write() to be a ReadWriter.
type ReadWriter interface {
    Reader
    Writer
}
```
This is used extensively in the Go standard library (e.g., `io.ReadWriter`).

---

### Key Takeaways for Interviews

*   **Go's Philosophy:** Go prefers composition over inheritance.
*   **Composition vs. Embedding:** Composition is the general "has-a" relationship (`Car` has an `Engine`). Embedding is a special kind of composition that promotes the fields and methods of the embedded type for convenience.
*   **No True Inheritance:** Embedding is not inheritance. There is no "is-a" relationship in the classical OOP sense. A `FinancialReport` *is not* a `Report`; it *contains* a `Report`.
*   **Benefits:** This approach avoids the "fragile base class" problem and the complex hierarchies that can arise from deep inheritance chains, leading to more flexible and maintainable code.