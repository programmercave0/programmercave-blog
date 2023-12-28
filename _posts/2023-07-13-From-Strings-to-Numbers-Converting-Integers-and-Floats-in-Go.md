---
layout: post
title: "From Strings to Numbers: Converting Integers and Floats in Go"
description: "In Go, working with integer and floating-point data types is fundamental for various programming tasks. This article will explore the integer and floating-point types available in Go, along with detailed instructions on how to convert strings to integers and floats using practical examples."
author: "Programmercave"
header-img: "/assets/images/From-Strings-to-Numbers-Converting-Integers-and-Floats-in-Go/go-string-to-int.png"
tags:  [Go]
date: 2023-07-13
toc: true
---
# Introduction:

In Go, working with integer and floating-point data types is fundamental for various programming tasks. This article will explore the integer and floating-point types available in Go, along with detailed instructions on how to convert strings to integers and floats using practical examples.

![From Strings to Numbers: Converting Integers and Floats in Go]({{ site.url }}/assets/images/From-Strings-to-Numbers-Converting-Integers-and-Floats-in-Go/go-string-to-int.png){:class="img-responsive"}

---



# Go Integer Data Types

Go provides a variety of integer data types, each with a different range and bit size. The following integer types are available:

1. `int8`: 8-bit signed integer (-128 to 127)
2. `int16`: 16-bit signed integer (-32768 to 32767)
3. `int32`: 32-bit signed integer (-2147483648 to 2147483647)
4. `int64`: 64-bit signed integer (-9223372036854775808 to 9223372036854775807)
5. `int`: Architecture-specific integer, either 32 or 64 bits
6. `uint8`: 8-bit unsigned integer (0 to 255)
7. `uint16`: 16-bit unsigned integer (0 to 65535)
8. `uint32`: 32-bit unsigned integer (0 to 4294967295)
9. `uint64`: 64-bit unsigned integer (0 to 18446744073709551615)
10. `uintptr`: Unsigned integer type capable of holding a pointer value

Additionally, Go introduces two type aliases:
- `byte`: Alias for `uint8`
- `rune`: Alias for `int32`, representing a Unicode code point

---



# Converting a String to Integer in Go

To convert a string to an integer in Go, you can utilize the `strconv` package. Here's an example code snippet demonstrating the conversion process for various integer types:

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "42" // Sample string representing a number

	// Convert to int8
	numInt8, _ := strconv.ParseInt(numStr, 10, 8)
	fmt.Println("int8:", int8(numInt8))

	// Convert to int16
	numInt16, _ := strconv.ParseInt(numStr, 10, 16)
	fmt.Println("int16:", int16(numInt16))

	// Convert to int32
	numInt32, _ := strconv.ParseInt(numStr, 10, 32)
	fmt.Println("int32:", int32(numInt32))

	// Convert to int64
	numInt64, _ := strconv.ParseInt(numStr, 10, 64)
	fmt.Println("int64:", numInt64)

	// Convert to int (architecture-specific)
	numInt, _ := strconv.Atoi(numStr)
	fmt.Println("int:", numInt)

	// Convert to uint8
	numUint8, _ := strconv.ParseUint(numStr, 10, 8)
	fmt.Println("uint8:", uint8(numUint8))

	// Convert to uint16
	numUint16, _ := strconv.ParseUint(numStr, 10, 16)
	fmt.Println("uint16:", uint16(numUint16))

	// Convert to uint32
	numUint32, _ := strconv.ParseUint(numStr, 10, 32)
	fmt.Println("uint32:", uint32(numUint32))

	// Convert to uint64
	numUint64, _ := strconv.ParseUint(numStr, 10, 64)
	fmt.Println("uint64:", numUint64)

	// Convert to uintptr
	numUintptr, _ := strconv.ParseUint(numStr, 10, 64)
	fmt.Println("uintptr:", uintptr(numUintptr))
}
```

In this code, we utilize the `strconv` package to convert the string `numStr` to different integer types. By using functions like `ParseInt` and `ParseUint`, we parse the string with a specified base (10 for decimal) and desired bit size. Then, through type casting, we convert the parsed value to the appropriate integer type.

Please note that this example omits error handling for brevity, but it's crucial to handle errors returned by the `strconv` functions in real-world scenarios.

---



# Go Floating-Point Data Types

Go offers two floating-point data types:

1. `float32`: 32-bit floating-point number, also known as single-precision float.
2. `float64`: 64-bit floating-point number, also known as double-precision float.

By default, Go assumes the `float64` type if you declare a floating-point variable without explicitly specifying the type. For instance:

```go
var x = 3.14 // x is inferred as float64
```

---



# Converting a String to Float in Go

To convert a string to a float in Go, you can use the `strconv` package. Here's an example code snippet demonstrating the process:

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "3.14159" // Sample string representing a floating-point number

	// Convert to float32
	numFloat32, _ := strconv.ParseFloat(numStr, 32)
	fmt.Println("float32:", float32(numFloat32))

	// Convert to float64 (default float type)
	numFloat64, _ := strconv.ParseFloat(numStr, 64)
	fmt.Println("float64:", numFloat64)
}
```

In the above code, we use the `strconv` package's `ParseFloat` function to convert the string `numStr` to floating-point values. We specify the desired bit size (32 for `float32` and 64 for `float64`) and then convert the parsed value to the appropriate float type using type casting.

Remember, error handling is excluded in this example for brevity, but it's essential to handle errors returned by `strconv.ParseFloat` in real-world applications.

---



# Conclusion:

This article covered the various integer and floating-point data types available in Go, providing insights into their ranges and bit sizes. Additionally, it demonstrated how to convert strings to different integer and floating-point types using the `strconv` package. By following the examples and guidelines outlined here, you can confidently convert strings to integers and floats in your Go programs.