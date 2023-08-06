---
layout: post
title: "How to Manipulate Bits in C and C++"
description: "All data in computer is represented in binary i.e. in 0 or 1. Computers or machines do not understand our languages, they understand bits. Generally programmer do not care about operations at the bit level. But sometimes a programmer has to dive in a deeper level and work on bits."
author: "Programmercave"
header-img: "/assets/bitmtable1.png"
tags:  [Cpp, Competitive-Programming, Bitmasking]
date: 2019-10-19
---

All data in computer is represented in binary i.e. in 0 or 1. Computers or machines do not understand our languages, they understand bits. Generally programmer do not care about operations at the bit level. But sometimes a programmer has to dive in a deeper level and work on bits.

# Bits representation
In programming, an n bit integer is stored as a binary number that consists of n bits. So a 32-bit integer consists of 32 bits and 64 bit integer consists of 64 bits. In C++ programming language int data type is 16-bit, 32-bit and 64-bit type. [see here](https://en.cppreference.com/w/cpp/language/types)

Here is the bit representation of 32 bit int number 10:<br/>
	00000000000000000000000000001010

In C++, `int` is either *signed* or *unsigned* and so a bit representation is either *signed* or *unsigned*. 

![Range]({{ site.url }}/assets/bitmtable1.png){:class="img-responsive"}

In a signed representation, first bit represents the sign of a number (0 for positive and 1 for negative), and remaining n-1 bits contains the magnitude of the number.
 <br/><input type="hidden" name="IL_IN_ARTICLE"> <br/>
There is a connection between a signed and an unsigned representation. A signed number `-x` equals an unsigned number 2<sup>n</sup> – x.<br/>
  -x (signed) = 2<sup>n</sup> - x (unsigned)

```
int a = -10;
unsigned int b = a;
std::cout << a << "\n"; /* -10 */
std::cout << b << "\n"; /* 4294967286 */
```

In a signed representation, the next number after 2<sup>n – 1</sup> – 1 is -2<sup>n</sup> – 1, and in an unsigned representation, the next number after 2<sup>n</sup> – 1 is `0`.

# Bit Operations

![Operations]({{ site.url }}/assets/bitmtable2.png){:class="img-responsive"}


We can use & operator to check if a number is even or odd. If `x & 1 = 0` then `x` is even and if `x & 1 = 1` then `x` is odd. We can also say that, `x` is divisible by 2<sup>k</sup> exactly when x & (2<sup>k</sup> – 1) = 0.

`x<<k` corresponds to multiplying `x` by 2<sup>k</sup>, and `x>>k` corresponds to dividing `x` by 2<sup>k</sup> rounded down to an integer.

# Common Bit Tasks

**Binary representation of _unsigned int_:**
```
void binary(unsigned int num)
{
	for(int i = 256; i > 0; i = i/2) {
		if(num & i) 
			std::cout << "1 ";
		else
			std::cout << "0 ";
	}
	std::cout << std::endl;
}
```

**Setting Bit at *position*:**
```
int set_bit(int num, int position)
{
	int mask = 1 << position;
	return num | mask;
}
```
 <br/><input type="hidden" name="IL_IN_ARTICLE"> <br/>
**Getting Bit at *position*:**
```
bool get_bit(int num, int position)
{
	bool bit = num & (1 << position);
	return bit;
}
```
**Clearing Bit at *position*:**
```
int clear_bit(int num, int position)
{
	int mask = 1 << position;
	return num & ~mask;
}
```

{% include ads.html %}
# Representing Sets

Bits representation of an integer are 0-indexed and the index starts from right side i.e. least significant bit. So we can represent every subset of the set `{0, 1, 2, ..., n-1}` as an *n* bit integer and whose bits indicate which element belongs to the subset. If bit at index 3 is 1 and at index 4 is 0 in binary representation of a number, then 3 belongs to the subset and 4  does not belong to the subset.

For a 32-bit integer, the set is {0, 1, 2,..., 31} and the subset is {1, 3, 4, 8}. The binary representation of the set is:
	00000000000000000000000100011010<br/>
and decimal representation is 2<sup>8</sup> + 2<sup>4</sup> + 2<sup>3</sup> + 2<sup>1</sup> = 282.

Code to form subset and add elements to it:
```
int add_elements_to_subset()
{
	int subset = 0;
	subset = subset | (1 << 1);
	subset = subset | (1 << 3);
	subset = subset | (1 << 4);
	subset = subset | (1 << 8);
	return subset;
}
```

Code to print elements of the subset:
```
void printing_subset(int subset)
{
	for (int i = 0; i < 32; i++) 
	{
		if (subset & (1 << i)) std::cout << i << " ";
	}
}
```

# Additional functions

The g++ compiler provides the following functions for counting bits:<br/>
	•` __builtin_clz(x)`: the number of zeros at the beginning of the number<br/>
	• `__builtin_ctz(x)`: the number of zeros at the end of the number<br/>
	• `__builtin_popcount(x)`: the number of ones in the number<br/>
	• `__builtin_parity(x)`: the parity (even or odd) of the number of ones<br/>

  
Get this post in [pdf format](https://www.file-up.org/uqqr7dbnhy8w)

Reference: <br/>
	Competitive Programmer’s Handbook - Antti Laaksonen

