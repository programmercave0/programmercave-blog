---
date: "2025-09-24T00:00:00Z"
description:
header-img:
keywords: Go programming language, Golang speed, compiled languages, concurrency in Go, goroutines, Go garbage collection, high-performance applications, Go interview preparation, native compilation, modern hardware optimization, fast compile times, efficient coding practices, Golang features
tags:
- Go
title: 'Golang - Interfaces - The Empty Interface and any'
toc: true
---

### Introduction

The empty interface is an interface that has **zero methods**. It is written as `interface{}`.

Since every concrete type has zero or more methods, **every type satisfies the empty interface**. This means you can store a value of any type in a variable of type `interface{}`.

**In Go 1.18 and later, `any` was introduced as an alias for `interface{}`.** You should prefer using `any` in modern Go code as it is more readable and concise.

**Analogy:** An `any` (`interface{}`) variable is like a **box that can hold anything**. You can put an integer, a string, a struct, or any other value into it. The box doesn't know or care what's inside.

---

### Why Use `any`?

The primary use case for `any` is to write functions that can accept arguments of any type. The most famous example is `fmt.Println`:

```go
// The signature is effectively:
func Println(a ...any) (n int, err error)
```

This allows `fmt.Println` to accept any number and combination of types.

```go
package main

import "fmt"

func main() {
	var a any // `any` is an alias for `interface{}`

	a = 42
	fmt.Println(a) // 42

	a = "hello"
	fmt.Println(a) // hello

	a = struct{ Name string }{"Alice"}
	fmt.Println(a) // {Alice}
}
```

---


### The Problem with `any`: Loss of Type Safety

While `any` is flexible, it comes at a cost: you lose static type safety. When you put a value into an `any` variable, the compiler "forgets" its original type. You can't perform type-specific operations on it directly.

```go
var a any = "hello"

// This will not compile!
// The compiler only knows the type is `any`, not `string`.
// It doesn't know if `any` has a `len()` function.
// fmt.Println(len(a)) // Error: invalid argument: a (variable of type any) has no len
```

To use the value, you must get its underlying concrete type back. This is done with a **type assertion**.

---


### Getting the Type Back: Type Assertions

A type assertion is an operation that checks the underlying type of an interface variable and extracts it.

```go
value, ok := a.(T)
```
*   `a`: The variable of type `any` (or another interface).
*   `T`: The concrete type you *think* is inside.
*   `value`: If the assertion is correct, `value` will hold the value with type `T`.
*   `ok`: A boolean that is `true` if the assertion succeeded and `false` otherwise.

This "comma, ok" idiom is the safe way to perform a type assertion, as it prevents a **panic** if the type is not what you expected.

```go
package main

import "fmt"

func describe(a any) {
	// Safely assert the type to string
	s, ok := a.(string)
	if ok {
		fmt.Printf("It's a string: '%s'\n", s)
		return
	}

	// Safely assert the type to int
	i, ok := a.(int)
	if ok {
		fmt.Printf("It's an int: %d\n", i)
		return
	}

	fmt.Println("It's some other type.")
}

func main() {
	describe("hello") // It's a string: 'hello'
	describe(123)     // It's an int: 123
	describe(3.14)    // It's some other type.
}
```

---


### When to Use `any` (and When Not To)

**Use `any` when:**
*   You are writing truly generic code that needs to operate on values of unknown types (e.g., encoding/json, fmt.Println, database drivers).
*   You need a collection of mixed types (though this is often a sign that a better data structure could be used).

**Avoid `any` when:**
*   You can use a more specific interface type. If you only need a `Write` method, accept an `io.Writer`, not `any`.
*   You can use **Go Generics** (Go 1.18+). Generics provide a way to write type-safe functions and data structures that work with a set of types, without the runtime overhead and loss of type safety associated with `any`.

**Interview Question:** "What is an empty interface or `any`? When should you use it?"
**Answer:** "`any` is an interface with no methods, so all types satisfy it. It's used to hold a value of any type. However, you lose static type safety and must use a type assertion to get the original type back. It's useful for functions like `fmt.Println`, but for your own code, you should prefer specific interfaces or, in modern Go, generics, as they are more type-safe."