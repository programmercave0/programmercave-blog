---
date: "2025-03-11T00:00:00Z"
description: Learn how to use Go's iota for creating constants efficiently. Understand basic to advanced examples, including enums, bitmask flags, and arithmetic sequences with clear explanations for your software engineering interview.
header-img: "/assets/images/Golang/What-is-iota-in-Go.png"
keywords: Go iota, constant declarations, Go enums, bitmask flags, Go constants tutorial, iota examples, iota bitmask, Go interview questions, constant sequences, Go programming
tags:
- Go
title: 'iota in Golang'
toc: true
---

## **Key Takeaways:**
- `iota` automatically assigns incrementing integer values to constants.
- Use `_` to skip values or start sequences from non-zero numbers.
- Combine `iota` with arithmetic/bit shifts for enums, flags, or memory units.
- Each `const` block resets `iota` to 0.
- String methods can make `iota`-based enums human-readable.

![What is iota in Golang](/assets/images/Golang/What-is-iota-in-Go.png)

---

## **What is `iota` in Go?**  
`iota` is a built-in Go tool that simplifies creating sequences of related constants. It starts at 0 in each `const` block and increments by 1 for each subsequent constant.  

**Example:**  
```go
const (
    Red = iota // 0
    Green      // 1
    Blue       // 2
)
```

---

## **Common Use Cases**  

### **1. Skipping Values**  
Use `_` to skip the first value:  
```go
const (
    _ = iota // Skip 0
    One      // 1
    Two      // 2
)
```

### **2. Starting from Non-Zero**  
Add offsets to `iota`:  
```go
const (
    A = iota + 1 // 1
    B            // 2
    C            // 3
)
```

### **3. Bitmask/Flags**  
Create binary flags using shifts:  
```go
const (
    Read = 1 << iota // 1 (0001)
    Write            // 2 (0010)
    Execute          // 4 (0100)
)
```
**Use Case:** Combine flags with bitwise operations (e.g., `Read | Write`).

---

### **4. Enums with String Values**  
Attach human-readable names to constants:  
```go
type Direction int

const (
    North Direction = iota
    East
    South
    West
)

func (d Direction) String() string {
    return [...]string{"North", "East", "South", "West"}[d]
}
```
**Output:**  
```go
fmt.Println(North) // "North"
```

---

### **5. Complex Sequences**  
Use arithmetic or exponents:  
```go
const (
    KB = 1 << (10 * iota) // 1024 (2^10)
    MB                     // 1,048,576 (2^20)
    GB                     // 1,073,741,824 (2^30)
)
```

---

## **Important Rules**  
| Rule | Example | Behavior |
|------|---------|----------|
| **Reset on New Block** | Two separate `const` blocks | `iota` restarts at 0 |
| **Same Line, Same Value** | `Active, Running = iota, iota + 100` | Both use the same `iota` (0) |

---

## **Frequently Confused Scenarios**  

### **Multiple Assignments in One Line**  
```go
const (
    Active, Running = iota, iota + 100 // Active=0, Running=100
    Passive, Stopped                   // Passive=1, Stopped=101
)
```
**Why?** `iota` increments per line, not per constant.

---

### **Resetting `iota`**  
Each `const` block resets `iota`:  
```go
const (
    FirstA = iota // 0
    FirstB        // 1
)

const (
    SecondA = iota // 0 (reset)
    SecondB        // 1
)
```

---

## **Why Use `iota`?**  
- **Readable Code:** Avoids magic numbers (e.g., `Read: 1 << 0`).  
- **Maintainable:** Adding new constants doesn’t require manual numbering.  
- **Efficient:** Reduces errors in large enum lists.