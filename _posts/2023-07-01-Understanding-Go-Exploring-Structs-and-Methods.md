---
layout: post
title: "Understanding Go: Exploring Structs and Methods"
subtitle: "Go is a statically typed, compiled programming language that offers a unique approach to object-oriented programming. Unlike languages such as C++ or Java, Go doesn't have traditional classes. Instead, it utilizes structs and methods to achieve similar functionality. In this blog post, we'll delve into a code snippet written in Go and explain how it works, step by step."
author: "Programmercave"
header-img: "/assets/Understanding-Go-Exploring-Structs-and-Methods/Understanding-Go-Exploring-Structs-and-Methods.png"
tags:  [Go]
date: 2023-07-01
---

# Introduction:

Go is a statically typed, compiled programming language that offers a unique approach to object-oriented programming. Unlike languages such as C++ or Java, Go doesn't have traditional classes. Instead, it utilizes structs and methods to achieve similar functionality. In this blog post, we'll delve into a code snippet written in Go and explain how it works, step by step.

![Understanding Go: Exploring Structs and Methods]({{ site.url }}/assets/Understanding-Go-Exploring-Structs-and-Methods/Understanding-Go-Exploring-Structs-and-Methods.png){:class="img-responsive"}

# Code Explanation:

Let's dissect the provided Go code snippet and understand its various components.

```go
type Person struct {
	Name    string
	Surname string
	Hobbies []string
	id      string
}

func (person *Person) GetFullName() string {
	return fmt.Sprintf("%s %s", person.Name, person.Surname)
}
func main() {
	p := Person{
		Name:    "Mario",
		Surname: "Castro",
		Hobbies: []string{"cycling", "electronics", "planes"},
		id:      "sa3-223-asd",
	}

	fmt.Printf("%s likes %s, %s and %s\n", p.GetFullName(), p.Hobbies[0],
		p.Hobbies[1], p.Hobbies[2])
}
```

The above example is taken from book : [Go Design Patterns](https://amzn.to/3NVFsb5)

## Defining the Person Struct:

The code begins with the declaration of a struct named `Person`. A struct is a composite data type that allows you to group together values of different types. In this case, the `Person` struct has four fields:

`Name` (string): Represents the person's first name.
`Surname` (string): Represents the person's last name.
`Hobbies` ([]string): Represents a list of hobbies the person has. It is a slice of strings, allowing multiple hobbies to be stored.
`id` (string): Represents a unique identifier for the person.

## Creating a Method on the Person Struct:

Next, we define a method named `GetFullName()` on the `Person` struct. In Go, a method is associated with a struct by declaring the method with a receiver type. In this case, the receiver type is a pointer to the `Person` struct `(*Person)`. The purpose of this method is to concatenate the `Name` and `Surname` fields of a `Person` instance and return the full name as a string.

## The main() Function:

Moving on to the `main()` function, we see the following steps:

Creating a `Person` Instance: We create a new instance of the `Person` struct named p using a struct literal. The values for the `Name`, `Surname`, `Hobbies`, and `id` fields are provided.

Calling the `GetFullName()` Method: Using the instance `p`, we call the `GetFullName()` method defined on the `Person` struct. This is achieved by using the dot notation (`p.GetFullName()`).

Printing the Result: The `fmt.Printf()` function is used to print a formatted string. The string includes the full name obtained from `p.GetFullName()` and the first three hobbies from the `Hobbies` field.

## Output:

The final output of the program is the following string: "Mario Castro likes cycling, electronics, and planes". This demonstrates that the code successfully retrieves the full name and prints it alongside the person's hobbies.

# Conclusion:

In conclusion, Go provides an alternative approach to object-oriented programming by utilizing structs and methods. Although Go doesn't have classes like other languages, it allows you to associate methods with structs by defining a receiver type. This code snippet showcases the usage of structs and methods in Go, demonstrating how methods can be called on instances of a struct to perform specific operations. By understanding these concepts, developers can leverage Go's unique approach to build efficient and concise code.