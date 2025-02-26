---
date: "2025-02-26T00:00:00Z"
description: Learn how Golang interfaces work, including implicit implementation, polymorphism, type assertions, and embedding. This guide simplifies complex concepts for software engineering interviews and practical use.
header-img: "/assets/images/Golang/Interfaces-in-Golang.png"
keywords: Interface Definition, Method Signatures, Implicit Implementation, Polymorphism, Type Abstraction, Interface Composition, Empty Interface (interface{} or any), Type Assertions, Type Switches, Runtime Type Checking, Interface Embedding, Constraint-Only Interfaces (Generics), Struct Embedding vs. Interface Embedding, Error Handling with Interfaces (error interface), Reflection and Interfaces, Interface Best Practices, Value Boxing, Dynamic Method Dispatch, iface and eface Structures, Performance Considerations
tags:
- Go
- Interview-Questions
title: 'Interview Question: What are Interfaces in Golang?'
toc: true
---

## Key Takeaways
- Interfaces in Go define method signatures and enable polymorphism, abstraction, and modularity.
- Implicit Implementation means types automatically implement interfaces by defining required methods—no explicit declaration is needed.
- The Empty Interface (`interface{}`) can hold any value, making it useful for generic functions but requiring type assertions for safe usage.
- Interface Embedding allows reusability by combining multiple interfaces into one.
- Type Assertions and Type Switches are essential for working with the empty interface to extract underlying types at runtime.

![What Are Interfaces in Golang](/assets/images/Golang/Interfaces-in-Golang.png)

## Understanding Interfaces in Golang

Interfaces are a core concept in Golang, defining a set of method signatures without implementations. This allows for polymorphism, where different types can be treated uniformly if they share the same methods. For example, consider a `Shape` interface:

```go
type Shape interface {
    Area() float64
}
```

A `Circle` struct can implement this by defining its own `Area` method:

```go
type Circle struct {
    radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}
```

Here, `Circle` implicitly implements `Shape` because it has an `Area` method with the correct signature. This implicit implementation, as explained in [Go by Example: Interfaces](https://gobyexample.com/interfaces), simplifies code but requires care to ensure all methods are defined.

## Implementing an Interface

Go uses implicit implementation, meaning a type automatically satisfies an interface if it implements all the required methods. There's no need to declare that a type "implements" an interface. Go handles this at compile time, promoting cleaner and more maintainable code.

### Example of Implicit Interface Implementation

```go
package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

// Circle implicitly implements the Shape interface
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func main() {
    var s Shape = Circle{Radius: 5}
    fmt.Println("Area of Circle:", s.Area())
}
```

In this example, the `Circle` type implicitly implements the `Shape` interface by providing the `Area()` method. There's no explicit "implements" keyword—Go automatically verifies this at compile time.

### Best Practices: Ensuring Interface Compliance
While Go's implicit interface implementation minimizes boilerplate code, it requires attention to detail. Missing a required method will result in a compile-time error if the type is used as the interface. To avoid surprises, you can use a blank identifier check to validate compliance:

```go
var _ Shape = Circle{}
```

This line doesn't execute anything but signals to the compiler to verify that Circle satisfies the Shape interface. This approach is particularly useful during refactoring or when defining new types.

## The Symbiotic Relationship with Structs

Structs in Golang define data fields and methods, while interfaces focus on behavior, creating a complementary relationship. For instance, an `Animal` interface with `Sound()` and `Move()` methods can be implemented by `Dog` and `Cat` structs:

```go
type Animal interface {
    Sound() string
    Move() string
}

type Dog struct{}

func (d Dog) Sound() string {
    return "Woof"
}

func (d Dog) Move() string {
    return "Run"
}

type Cat struct{}

func (c Cat) Sound() string {
    return "Meow"
}

func (c Cat) Move() string {
    return "Walk"
}
```

A function like `DescribeAnimal` can then work with any `Animal`, demonstrating polymorphism:

```go
func DescribeAnimal(a Animal) {
    fmt.Println("Sound:", a.Sound(), "Move:", a.Move())
}
```

This partnership, detailed in [Implementing Interfaces in Go (Golang)](https://www.golangbot.com/interfaces-part-1/), enhances modularity, allowing new types to be added without changing existing code.

## Declaring Variables and Parameters with Interface Types

You can declare variables or function parameters as interface types, enabling polymorphism. For example:

```go
var s Shape = Circle{radius: 5}
fmt.Println("Area:", s.Area()) // Calls Circle's Area method
```

Or a function that accepts any `Shape`:

```go
func PrintArea(s Shape) {
    fmt.Println("Area:", s.Area())
}
```

This allows a single function to handle different types, like `Circle` or `Rectangle`, as shown in [How To Use Interfaces in Go | DigitalOcean](https://www.digitalocean.com/community/tutorials/how-to-use-interfaces-in-go), making code reusable and flexible.

## Embedding Interfaces for Reusability

Golang allows embedding interfaces, where a new interface inherits methods from another. For example:

```go
type BasicShape interface {
    Area() float64
}

type ShapeWithColor interface {
    BasicShape
    Color() string
}
```

A `ColoredCircle` can implement `ShapeWithColor` by defining both `Area()` and `Color()`:

```go
type ColoredCircle struct {
    radius float64
    color string
}

func (cc ColoredCircle) Area() float64 {
    return math.Pi * cc.radius * cc.radius
}

func (cc ColoredCircle) Color() string {
    return cc.color
}
```

This, as explained in [Embedding Interfaces in Golang - GeeksforGeeks](https://www.geeksforgeeks.org/embedding-interfaces-in-golang/), promotes modular, reusable interface design, useful for large systems.

## The Empty Interface: interface{}

The empty interface , written as `interface{}`, has no methods. Since every type in Go satisfies this interface, it can hold values of any type.

### Example:
```go
func PrintAnything(a interface{}) {
    fmt.Println(a)
}
```

However, working with the empty interface requires type assertions or type switches to access the actual value safely.

### Type Assertion:
```go
value, ok := a.(string)
if !ok {
    fmt.Println("Not a string")
} else {
    fmt.Println("String value:", value)
}
```

### Type Switch:
```go
switch v := a.(type) {
case int:
    fmt.Println("Integer:", v)
case string:
    fmt.Println("String:", v)
default:
    fmt.Println("Unknown type")
}
```

---

## Table: Comparison of Interface Features in Golang

| Feature                  | Description                                                                 | Example Use Case                     |
|--------------------------|-----------------------------------------------------------------------------|--------------------------------------|
| Implicit Implementation  | Types implement interfaces by defining methods, checked at compile time.    | `Circle` implements `Shape` with `Area()`. |
| Embedding Interfaces     | New interfaces inherit methods from embedded ones, enhancing reusability.   | `ShapeWithColor` embeds `BasicShape`. |
| Empty Interface          | `interface{}` holds any value, useful for generic functions.                | `fmt.Println` accepts any type.       |
| Type Assertion           | Checks and extracts underlying type, like `s.(string)` for safety.          | Handling mixed-type slices.           |
| Polymorphism             | Variables/parameters of interface type work with any implementing type.     | Function processes list of shapes.    |

This table, derived from various sources, highlights key features, aiding students in understanding and applying interfaces effectively.

---

### Key Citations
- [Go by Example: Interfaces detailed explanation with code examples](https://gobyexample.com/interfaces)
- [A Tour of Go interactive introduction covering methods and interfaces](https://go.dev/tour/methods/9)
- [Golang Interfaces Explained comprehensive tutorial for beginners](https://www.alexedwards.net/blog/interfaces-explained)
- [Understanding the Power of Go Interfaces comprehensive guide with examples](https://medium.com/@jamal.kaksouri/understanding-the-power-of-go-interfaces-a-comprehensive-guide-835954101b7e)
- [Implementing Interfaces in Go practical examples and explanations](https://www.golangbot.com/interfaces-part-1/)
- [How To Use Interfaces in Go tutorial with practical applications](https://www.digitalocean.com/community/tutorials/how-to-use-interfaces-in-go)
- [All Go: Methods and Interfaces detailed analysis of Go's object orientation](https://www.andy-pearce.com/blog/post/2023/Jun/all-go-method-and-interfaces/)
- [Interfaces in Golang detailed documentation with examples](https://golngdocs.com/interfacess-in-golng)
- [Golang Interfaces Tutorial with Examples practical guide for implementation](https://www.callicoder.com/golng-interfaces/)
- [Go Empty Interface with Examples explanation and use cases](https://www.progrmiz.com/gol/empty-interfac)
- [Understanding and Using the Empty Interface in Go best practices and examples](https://www.codingexplorations.com/blog/understanding-and-using-the-empty-interfac-in-go)
- [Understanding the empty interface in Go beginner-friendly explanation](https://dev.to/flrnd/understanding-the-empty-interfac-in-go-4652)
- [Understanding Go's Empty Interface detailed analysis with real-world use](https://matthewsetter.com/gol-empty-interfac/)
- [Go: Understand the Empty Interface medium post with practical insights](https://medium.com/a-journey-with-go/gol-understand-the-empty-interfac-2d9fc1e5ec72)
- [go - Empty Interfaces in Golang Stack Overflow discussion with examples](https://stackoverflw.com/questions/59976812/empty-interfacess-in-golng)
- [Empty Interface Type in Go Programming Language detailed tutorial](https://www.golngprograms.com/interfacess-with-similar-methodss-in-golng-programming-language-863.html)
- [Empty Interface{} in Golang comprehensive explanation with examples](https://www.scaler.com/topics/golng/empty-interfac-in-golng/)
- [A Tour of Go additional methods and interface exercises](https://go.dev/tour/methods/14)
- [Embedding Interfaces in Golang detailed guide with examples](https://www.geeksforgeeks.org/embedding-interfaces-in-golang/)
- [Embedding Interfaces in Go practical examples and explanations](https://golngbyexample.com/embedding-interfacess-go/)
- [Embedding in Go: Part 3 - interfaces in structs detailed analysis](https://eli.thegreenplace.net/2020/embedding-in-go-part-3-interfacess-in-structs)