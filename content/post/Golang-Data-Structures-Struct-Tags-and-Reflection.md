---
date: "2025-09-20T00:00:00Z"
description:
header-img:
keywords: Go programming language, Golang speed, compiled languages, concurrency in Go, goroutines, Go garbage collection, high-performance applications, Go interview preparation, native compilation, modern hardware optimization, fast compile times, efficient coding practices, Golang features
tags:
- Go
title: 'Golang - Data Structures - Struct Tags and Reflection'
toc: true
---

### Introduction

Struct tags are small pieces of metadata that you can attach to the fields of a struct. They are just strings, but they provide a way to give instructions to other Go packages that inspect your structs at runtime.

The most common use for struct tags is to control how your data is encoded to or decoded from formats like JSON, XML, or database records. This is achieved through a mechanism called **reflection**.

---

### What are Struct Tags?

A struct tag is a raw string literal placed after the type in a struct field definition.

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    ID    int    `json:"-"` 
    // The dash means "ignore this field"
}
```

**Anatomy of a Tag:**
*   It's enclosed in backticks ``.
*   Inside, it's a space-separated list of `key:"value"` pairs.
*   The `key` (e.g., `json`, `xml`, `db`) indicates which package the tag is for.
*   The `value` (e.g., `"name"`, `"-"`) is the instruction for that package.

**Analogy:** Think of struct tags as **sticky notes** you put on the fields of a form. A note might say, "For the accounting department, call this field 'user_id'" or "For the public API, don't show this field at all." The struct itself doesn't care about the notes, but other programs (like the `encoding/json` package) can read them and act accordingly.

---

### Common Use Case: JSON Marshaling

The `encoding/json` package is the most common consumer of struct tags.

*   `json:"fieldName"`: Use a different name in the JSON output. This is the idiomatic way to have lowercase or snake_case field names in your JSON.
*   `json:"-"`: Exclude this field from the JSON output entirely.
*   `json:"fieldName,omitempty"`: If the field has its zero value (e.g., 0, "", nil), omit it from the JSON output.

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ProductID      int     `json:"id"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	Published      bool    `json:"is_published,omitempty"`
	internalSecret string  // Unexported fields are always ignored
	AdminNotes     string  `json:"-"`
}

func main() {
	p1 := Product{
		ProductID:      101,
		Name:           "Go T-Shirt",
		Price:          20.50,
		Published:      true,
		internalSecret: "secret",
		AdminNotes:     "notes",
	}

	p2 := Product{
		ProductID: 102,
		Name:      "Go Mug",
		Price:     15.00,
		// Published is false (the zero value), so it will be omitted
	}

	p1Json, _ := json.MarshalIndent(p1, "", "  ")
	p2Json, _ := json.MarshalIndent(p2, "", "  ")

	fmt.Println(string(p1Json))
	fmt.Println(string(p2Json))
}
```

**Output:**

```json
{
  "id": 101,
  "name": "Go T-Shirt",
  "price": 20.5,
  "is_published": true
}
{
  "id": 102,
  "name": "Go Mug",
  "price": 15
}
```
Notice how `AdminNotes` and `internalSecret` are missing, the field names match the tags, and `is_published` is omitted for `p2`.

---

### How Does It Work? Reflection

Struct tags are meaningless to the Go compiler itself. They are just strings. Their power comes from the `reflect` package.

**Reflection** is the ability of a program to inspect its own structure at runtime. The `encoding/json` package uses reflection to:
1.  Iterate over the fields of a struct it's given.
2.  For each field, check if it has a `json` tag.
3.  Read the value of the tag to decide the field's name or whether to ignore it.
4.  Read the actual value of the field.
5.  Construct the final JSON output.

Reflection is a powerful but advanced feature. It is generally slower than compile-time code, which is why it's often used in I/O-bound tasks like marshaling where the performance impact is less critical.

**Interview Question:** "How does JSON encoding work in Go? How can you control the output?"
**Answer:** "It works using the `reflect` package. When you call `json.Marshal`, it inspects the struct's fields at runtime. You can control the output by adding `json` struct tags to the fields to change their names, omit them if they're empty, or ignore them completely."

---

### Other Uses for Struct Tags

*   **Database ORMs:** `db:"user_id"` to map a struct field to a database column.
*   **Configuration Loading:** `env:"PORT"` to load a struct field from an environment variable.
*   **Validation:** `validate:"required,min=8"` to specify validation rules for a field.

Struct tags are a key part of what makes Go's standard library and third-party packages so flexible and powerful.
