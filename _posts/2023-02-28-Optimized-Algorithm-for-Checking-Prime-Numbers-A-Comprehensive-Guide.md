---
layout: post
title: "Optimized Algorithm for Checking Prime Numbers: A Comprehensive Guide"
description: "In this blog, we will discuss different methods to check if a number is a prime number or not, and analyze their time complexity. We will start with a basic algorithm and then move on to more advanced algorithms.."
author: "Programmercave"
header-img: "/assets/Prime-Numbers/prime_number2.png"
tags:  [Cpp, Algorithm, Data-Structures, Competitive-Programming, Mathematics]
date: 2023-02-28
---
* toc
{:toc}

Prime numbers are a fundamental concept in mathematics and computer science. A prime number is a positive integer greater than 1 that has no positive integer divisors other than 1 and itself. In other words, a prime number is a number that is only divisible by 1 and itself.

In this blog, we will discuss different methods to check if a number is a prime number or not, and analyze their time complexity. We will start with a basic algorithm and then move on to more advanced algorithms.

## Basic Algorithm

The most basic algorithm to check if a number is a prime number or not is to iterate over all positive integers less than the number, and check if any of them divide the number evenly. If none of them divide the number evenly, then the number is a prime number. 

![Prime Numbers]({{ site.url }}/assets/Prime-Numbers/prime_number1.png){:class="img-responsive"}

Here's a simple implementation of this algorithm in C++:

```cpp
bool is_prime(int n) {
    if (n < 2) return false;
    for (int i = 2; i < n; i++) {
        if (n % i == 0) return false;
    }
    return true;
}
```

This algorithm has a time complexity of **O(n)**, since it iterates over all positive integers less than n. Therefore, this algorithm is not suitable for large values of n.

## Optimized Algorithm

The optimized algorithm for checking if a number is prime or not works by reducing the number of iterations in the loop. All non-prime numbers can be expressed as the product of two numbers less than or equal to its square root, so we only need to check for divisors up to the square root of n. This is because if a number has a divisor greater than its square root, then it must also have a corresponding factor that is less than its square root. For example, if 16 has a factor greater than 4, then it must also have a factor less than 4.

![Prime Numbers]({{ site.url }}/assets/Prime-Numbers/prime_number2.png){:class="img-responsive"}

Let's consider an example where n = 16. To check if 16 is a prime number, we start by finding the square root of n, which is 4. We then iterate over all numbers from 2 to 4 and check if they divide n evenly. When a = 2, then b = 8. If `n % a == 0 (16 % 2 == 0)`, this also means `n % b == 0 (16 % 8 == 0)`. If we traverse beyond the square root of n, the values of a and b will be interchanged, and we will end up calculating what we have already calculated. Therefore, we don't want i > sqrt(n). If we find a divisor, then we know that n is not a prime number, and we can return false. If we reach the end of the loop without finding a divisor, then we know that n is a prime number, and we can return true.

Here's the implementation of the optimized algorithm in C++:

```cpp
bool is_prime(int n) {
    if (n < 2) return false;
    for (int i = 2; i*i <= n; i++) {
        if (n % i == 0) return false;
    }
    return true;
}
```

In this implementation, we use the condition `i*i <= n` instead of `i <= sqrt(n)` to avoid the overhead of the square root operation. This is because calculating the square root of a number is a relatively expensive operation, especially for large numbers.

In the case of n = 16, the loop will iterate over the numbers 2, 3, and 4. The first iteration checks if 2 divides 16 evenly, which it does. Therefore, we return false and terminate the function. If we had continued iterating, we would have found that 4 also divides 16 evenly, which would have also led to a false return value.

By iterating only up to the square root of n, we can significantly reduce the number of iterations in the loop and improve the performance of the algorithm. This is especially important for large values of n, where the basic algorithm would be prohibitively slow.

This algorithm has a time complexity of **O(sqrt(n))**, which is a significant improvement over the basic algorithm.

If you'd like to see the complete code for checking prime numbers using an optimized algorithm, you can find it on my Github repository [here](https://github.com/{{site.github_username}}/Algo-Data-Structure/blob/master/Maths/check_prime.cpp). 

In our next post we will discuss how to find all prime numbers till N using [Sieve of Eratosthenes]({{site.url}}/blog/2023/03/02/Efficiently-Find-Prime-Numbers-Till-N-Basic-vs-Sieve-of-Eratosthenes).

## Conclusion

In conclusion, checking if a number is a prime number or not is a fundamental problem in mathematics and computer science. We have discussed several algorithms for solving this problem, ranging from the basic algorithm with a time complexity of O(n), to the optimized algorithm with a time complexity of O(sqrt(n)).

If you're interested in checking out some of my code related to algorithms and data structures, be sure to visit [Algo-Data-Structure](https://github.com/{{site.github_username}}/Algo-Data-Structure) on Github. For my solutions to problems from competitive programming sites, you can find them in [Competitive-Programming](https://github.com/{{site.github_username}}/Competitive-Programming).

---
* toc
{:toc}

By the way, if you're a teacher or parent looking for resources to help your child get ready for school, you might be interested in these fun and informative workbooks developed by a pre-school teacher. Covering all the basic skills needed for school-readiness, they're perfect for the pre-school education niche. Check them out here: [WORKSHEETS FOR PRESCHOOL](https://ce8977zhz1vrft28uay3ofipe9.hop.clickbank.net/?cbpage=wfpaffiliate)
