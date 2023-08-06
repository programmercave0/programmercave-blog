---
layout: post
title: "Pointers in C++"
description: "A pointer is a special type of variable which holds the address of a value. A pointer is declared using * (asterisk).A constant pointer only points to the object it was initialized to.In a pointer to a constant, the pointer can not be used to modify the value at which it points."
author: "Programmercave"
header-img: "/assets/2019-12-02-Pointers-in-C++-Explained/ptr.png"
tags:  [Cpp]
date: 2019-12-02
---
* toc
{:toc}

A **pointer** is a special type of variable which holds the address of a value. A pointer is declared using * (asterisk).

```cpp
int a = 5;  //an integer variable
int * ptr;  // an integer pointer or a pointer to an int
ptr = &a; // assign address of a to ptr
```

![Pointers]({{ site.url }}/assets/2019-12-02-Pointers-in-C++-Explained/ptr.png){:class="img-responsive"}

In the above figure (a), `a` and `ptr` are name of location in the memory and value stored at `a` is `5` and value stored at `ptr` is `0x41a`. The address of a is `0x41a` and address of ptr is `0x021`. So ptr stores the address of `a` and we can also define another pointer which stores the address of `ptr` i.e. `0x021`.

`ptr` will output value stored in it and *ptr will output the value stored at address stored in `ptr`.

```cpp
std::cout << ptr << "\n"; // 0x41a
std::cout << *ptr << "\n"; // 5
std::cout << &a << "\n"; //0x41a
```

Since `ptr` is pointing at `a`, value stored in a can be modified by `*ptr`.

```cpp
*ptr = 10;
std::cout << *ptr << "\n"; // 10
std::cout << a << "\n"; // 10
```

Similarly, a pointer to char or a pointer to double is declared.

<h1>Pointers and Constants</h1>

<h3>A Constant Pointer</h3>

A **constant pointer** only points to the object it was initialized to. If it was initialized to point to an integer variable `a`, then it cannot be changed to point to another integer variable `b`. This means that the address stored in a constant pointer cannot be changed.

```cpp
int a = 5;
int * const ptr = &a; // a constant pointer
```

A constant pointer must be initialized when it is declared. The following lines of code will produce an error.

```cpp
int * const ptr; // Error
ptr = &b //Error
```

{% include ads.html %}<br/>

The value stored in a can be changed. It can be changed by modifying a or by modifying `*ptr.`

```cpp
a = a + 5;
std::cout << a << "\n"; // 10
*ptr = *ptr + 5;
std::cout << *ptr << "\n"; // 15
```

<h3>A Pointer to a Constant</h3>

In a **pointer to a constant**, the pointer can not be used to modify the value at which it points.

```cpp
int a = 5;
const int * ptr = &a;
a = a + 5;
std::cout << a << "\n"; // 10
*ptr = *ptr + 5; // Error
```

<h1>Pointer Arithmetic</h1>

When 1 is added to an integer variable, the value stores in it increases by 1. Similarly when a pointer is incremented by 1 it points to next address.

If it is a pointer to an int, then its the value in it is increased by 4 bytes and if it is a pointer to a char, then the value in it is increased by 1 byte.

```cpp
int arr[] = {1, 3, 5};
int * ptr = &arr[0];
std::cout << *ptr << "\n"; // 1
ptr++;
std::cout << *ptr << "\n"; // 3
ptr++;
std::cout << *ptr << "\n"; // 5
```

Get this post in pdf : [Pointers in C++](https://www.file-up.org/gxmotwacbs8x)

Watch [Pointers in C++ explanation in Hindi on Youtube](https://youtu.be/pQ82tDhTX5g).

Practice the programs to have a good grasp on Pointers.

[Move all Odd numbers after Even numbers in Singly Linked List]({{ site.url }}/blog/2018/02/08/C++-Move-all-Even-numbers-before-Odd-numbers-in-Singly-Linked-List-(Using-STL))<br/>
[Merge two sorted Linked List (in-place)]({{ site.url }}/blog/2018/02/06/C++-Merge-two-sorted-Linked-List-(in-place))<br/>
[Split Singly Circular Linked List]({{ site.url }}/blog/2018/02/04/C++-Split-Singly-Circular-Linked-List-program)<br/>
[Reverse the Linked List]({{ site.url }}/blog/2018/01/23/C++-Reverse-the-Linked-List-(Iterative-Method)-program)<br/>
[Finding Length of Loop in Linked List]({{ site.url }}/blog/2018/01/20/C++-Linked-List-containing-Loop-(Floyd-Cycle-finding-Algorithm)-program)<br/>







