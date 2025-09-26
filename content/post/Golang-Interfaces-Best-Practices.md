---
date: "2025-09-26T00:00:00Z"
description:
header-img:
keywords: Go programming language, Golang speed, compiled languages, concurrency in Go, goroutines, Go garbage collection, high-performance applications, Go interview preparation, native compilation, modern hardware optimization, fast compile times, efficient coding practices, Golang features
tags:
- Go
title: 'Golang - Interfaces - Best Practices'
toc: true
---

### Introduction

How you design and use interfaces can have a significant impact on the quality, flexibility, and maintainability of your Go code. The Go community has developed a set of idiomatic best practices for working with interfaces.

---

### 1. Prefer Small, Focused Interfaces

The single most important rule of interface design in Go is to **keep interfaces small**. An interface should typically have only one or two methods. This is the essence of the "I" in the SOLID principles (Interface Segregation Principle).

The `io.Reader` and `io.Writer` interfaces are the canonical examples of this. They each have only one method, which makes them incredibly versatile and easy to implement.

**Why are small interfaces better?**
*   **Easier to Satisfy:** A type only needs to implement one or two methods to satisfy the interface, making it easy to adapt new types to your code.
*   **More Reusable:** A function that accepts a small interface is more reusable because it asks for the minimum behaviour it needs.
*   **Clearer Intent:** A small interface clearly documents the exact behaviour a function requires.

**Example:**

```go
// BAD: A large, monolithic interface
type Animal interface {
    Eat()
    Sleep()
    Walk()
    Fly()
    Swim()
}

// GOOD: Small, composable interfaces
type Eater interface {
    Eat()
}
type Walker interface {
    Walk()
}
type Flyer interface {
    Fly()
}

// A function should only ask for what it needs.
// This function doesn't care if the animal can fly or swim.
func takeForAWalk(w Walker) {
    w.Walk()
}
```

---

### 2. "Accept Interfaces, Return Structs"

This is a famous and highly valuable Go proverb.

#### a) Accept Interfaces

When writing a function, the parameters it accepts should be interfaces whenever possible. This makes the function more general and decoupled from specific implementations.

```go
// This function can open a door from any type that knows how to `Unlock`.
// It doesn't need to know if it's a `CarKey`, a `HouseKey`, or a `DigitalKeycard`.
type Unlocker interface {
    Unlock() error
}

func openDoor(u Unlocker) {
    u.Unlock()
    // ...
}
```

#### b) Return Structs

When a function creates and returns a value, it should typically return a concrete type (like a pointer to a struct).

```go
// This function returns a concrete type, not an interface.
func NewCarKey() *CarKey {
    return &CarKey{...}
}
```

**Why?**
*   **Flexibility for the Caller:** The caller receives a concrete type with all of its fields and methods available. They are not limited to the small subset of methods defined by an interface.
*   **Avoids Unnecessary Abstraction:** Returning an interface can hide useful information and force the caller to use a type assertion if they need access to the concrete type's full capabilities.
*   **The caller can always assign the concrete type to an interface variable if they wish.**

```go
// The caller gets a concrete *CarKey...
key := NewCarKey()

// ...and can choose to store it in an interface variable if that's all they need.
var unlocker Unlocker = key
openDoor(unlocker)
```
This pattern provides the maximum flexibility to the consumer of your API.

---

### 3. Define Interfaces Where They Are Used

In many object-oriented languages, it's common to define interfaces in the same package as the types that implement them.

In Go, the idiomatic approach is often to **define an interface in the package that *uses* it**, not the package that implements it.

**Example:**
Imagine you have a `user` package and a `storage` package. The `storage` package needs to save users.

*   **The `storage` package (the consumer) defines the interface it needs:**
    ```go
    // in package storage
    package storage

    // Storable is the interface this package needs to be able to save something.
    type Storable interface {
        ID() string
        Data() []byte
    }

    func Save(s Storable) error {
        // ... logic to save the data ...
    }
    ```

*   **The `user` package (the producer) implements the interface:**
    ```go
    // in package user
    package user

    type User struct { ... }

    func (u *User) ID() string { ... }
    func (u *User) Data() []byte { ... }
    ```

**Why is this better?**
*   **No Dependency:** The `user` package does not need to import the `storage` package. This avoids circular dependencies and keeps your code decoupled. The `storage` package can save *any* type from *any* package, as long as it satisfies the `Storable` contract.
*   **Clearer API:** The `storage` package's API is self-documenting. It clearly states, "I can work with anything that looks like a `Storable`."
