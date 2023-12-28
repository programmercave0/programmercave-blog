---
layout: post
title: "Prime Numbers and Sieve of Eratosthenes"
description: "A prime number is an integer n > 1 which is only divisible by 1 and n(itself). For example, 2, 3, 5 are prime numbers, but 6, 8, 9 are not prime numbers."
author: "Programmercave"
tags:  [Number-Theory, Cpp, Competitive-Programming]
date: 2019-11-12
---



A **prime number** is an integer *n* > 1 which is only divisible by 1 and *n*(itself). For example, 2, 3, 5 are prime numbers, but 6, 8, 9 are not prime numbers.

Every integer *n* > 1 can be uniquely expressed as the product of prime numbers. For example 8 = 2 * 2 * 2.

<h2>Some facts about prime numbers:</h2>
   1. Except 2, all prime numbers are odd.
   2. All prime numbers greater than 3 can be written in the form 6*k* ± 1.
   3. 1 is not a prime number.

A **perfect number** is a number *n* if *n* is equal to the sum of its factors between 1 and *n* – 1. For example 28 is a perfect number because 28 = 1 + 2 + 4 + 7 + 14.

<h1>Conjectures</h1>

   1. **Goldbach’s conjecture:** Each even integer *n* > 2 can be represented as a sum *n* = *a* + *b* such that both *a* and *b* are prime numbers. For example 12 = 5 + 7.

   2. **Twin prime conjecture:** There is an infinite number of pairs of the form {*p*, *p* + 2}, where both *p* and *p* + 2 are prime numbers. For example {5, 7}, {11, 13}.

   3. **Legendre’s conjecture:** There is always a prime between numbers n<sup>2</sup> and (n + 1)<sup>2</sup>, where *n* is any positive integer. For example 5 and 7 between 2<sup>2</sup> and 3<sup>2</sup>.

<h1>Basic Algorithm</h1>

If a number *n* > 1 is not a prime number then it has a factor between 2 and floor (√n). For example 27 = 3*9 and  √27 = 5.196 then 27 has a factor between 2 and 5 which is 3. So to test if a number is prime number we have to iterate to √n in a loop.

Here is a function `is_prime` which checks if the given number is prime.

```cpp
bool is_prime(int n)
{
	if (n < 2)
	{
		return false;
	}

	for (int i = 2; i*i <= n; ++i)
	{
		if (n % i == 0)
		{
			return false;
		}
	}
	return true;
}
```

Here is the function `prime_factorization` which returns an array of prime factors of the number.

```cpp
std::vector<int> prime_factorization(int n)
{
	std::vector<int> factors;

	for (int i = 2; i*i <= n; ++i)
	{
		while (n % i == 0)
		{
			factors.push_back(i);
			n = n / i;
		}
	}

	if (n > 1)
	{
		factors.push_back(n);
	}
	return factors;
}
```

{% include ads.html %}<br/>
<h1>Sieve of Eratosthenes</h1>

It is an efficient algorithm to find prime numbers upto a limit (say 10,000,000).

In this algorithm, first we mark all the multiples of 2 upto n, then all the multiples of 3, then 5 and so on. In the end we have all the prime numbers upto n.

For example, to find prime numbers less tha or equal to 20.

2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 

Now we mark all the multiples of 2 except 2 with red.

2 3 <span style="color:red">4</span> 5 <span style="color:red">6</span> 7 <span style="color:red">8</span> 9 <span style="color:red">10</span> 11 <span style="color:red">12</span> 13 <span style="color:red">14</span> 15 <span style="color:red">16</span> 17 <span style="color:red">18</span> 19 <span style="color:red">20</span> 

Now we mark all the multiples of 3 except 3 with red.

2 3 <span style="color:red">4</span> 5 <span style="color:red">6</span> 7 <span style="color:red">8</span> <span style="color:red">9</span> <span style="color:red">10</span> 11 <span style="color:red">12</span> 13 <span style="color:red">14</span> <span style="color:red">15</span> <span style="color:red">16</span> 17 <span style="color:red">18</span> 19 <span style="color:red">20</span> 

The remaining integers in black are prime numbers less than 20.

Here is the function `sieve_of_erastosthenes` which returns an array whose index can tell us the prime numbers less than or equal to *n*.

```cpp
std::vector<bool> sieve_of_erastosthenes(int n)
{
	std::vector<bool> is_prime(n+1, true);

	is_prime[0] = is_prime[1] = false;

	for (int i = 2; i <= n; ++i)
	{
		if (is_prime[i])
		{
			for (int j = i*i; j <= n; j += i)
			{
				is_prime[j] = false;
			}
		}
	}
	return is_prime;
}
```

Reference - Competitive Programmer’s Handbook, Antti Laaksonen

Get this post in pdf - [Prime Numbers and Sieve of Eratosthenes](https://www.file-up.org/6hib9eu4og14)
