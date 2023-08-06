---
layout: post
title: "Between Two Sets - HackerRank Challenge | C++ Implementation"
description: "You will be given two arrays of integers and asked to determine all integers that satisfy the following two conditions: 1. The elements of the first array are all factors of the integer being considered. 2. The integer being considered is a factor of all elements of the second array"
author: "Programmercave"
header-img: "/assets/2019-11-29-Between-Two-Sets-HackerRank-Challenge/HR_betweentwosets.jpg"
tags:  [Cpp, Algorithm, Competitive-Programming, Hackerrank, Data-Structure]
date: 2019-11-29
---
* toc
{:toc}

<h1>Problem</h1>

You will be given two arrays of integers and asked to determine all integers that satisfy the following two conditions: 

  1. The elements of the first array are all factors of the integer being considered <br/>
  2. The integer being considered is a factor of all elements of the second array 

These numbers are referred to as being **between the two arrays**. You must determine how many such numbers exist.

For example, given the arrays `a =[2, 6]` and `b = [24, 36]`, there are two numbers between them: `6` and `12`. `6 % 2 = 0`, `6 % 6 = 0`, `24 % 6 = 0` and `36 % 6 = 0` for the first value. Similarly, `12 % 2 = 0`, `12 % 6 = 0` and `24 % 12 = 0`, `36 % 12 = 0`.

Read the full problem here: [Between Two Sets](https://www.hackerrank.com/challenges/between-two-sets/problem)

![Between Two Sets Hackerrank]({{ site.url }}/assets/2019-11-29-Between-Two-Sets-HackerRank-Challenge/HR_betweentwosets.jpg){:class="img-responsive"}

<h1>Solution</h1>

First we sort both the arrays in increasing order because values can be entered in any order. Let the first array is `factors` and second array is `multiples` (See conditions 1 and 2). 

```cpp
std::sort(factors.begin(), factors.end());
std::sort(multiples.begin(), multiples.end());
```

The range of integer being considered is from last element of factors to first element of multiples. Let that integer is `num`. So `num = { factors.back(), ... , multiplies.front() }`. 

```cpp
unsigned short num = factors.back();
unsigned short times_multiplied = 2;
```

{% include ads.html %}<br/>

Here `num` also contain integers that are not multiples of `2` and `6` for eg. `7, 11, 13, 15 ...` To make program efficient `num` should contain multiples of `factors.back()` i. e. in the above example multiples of 6. So `num = {6, 12, 18, 24 }` and the numbers which satisfy the above two conditions are `6 and 12`.

```cpp
num = factors.back() * times_multiplied;
times_multiplied++;
```

Bool variable `is_multiple` tells if a number is  multiple of the `num` and `is_factor` tells if the `num` is the factor of a number. 

```cpp
bool is_multiple = true;
bool is_factor = true;
```

If `is_multiple` is `true` for all the elements of `factors` and `is_factor` is `true` for all the elements of `multiplies` for a `num` then that `num` belongs to between the two arrays and our `count` variable is incremented by 1.

```cpp
for (unsigned int i = 0; i < factors.size(); ++i)
{
    if (num % factors[i] != 0)
    {
        is_multiple = false;
        break;
    }
}

for (unsigned int j = 0; j < multiples.size(); ++j)
{
    if (is_multiple && multiples[j] % num != 0)
    {
        is_factor = false;
        break;
    }
}

if (is_multiple && is_factor)
{
    count++;
}
```

{% include ads.html %}<br/>

<h3>C++ Implementation</h3>

```cpp
#include <iostream>
#include <vector>
#include <algorithm>

unsigned int count_between_two_sets(std::vector<unsigned short>& factors, std::vector<unsigned short>& multiples)
{
    std::sort(factors.begin(), factors.end());
    std::sort(multiples.begin(), multiples.end());

    unsigned int count = 0;
    unsigned short times_multiplied = 2;
    unsigned short num = factors.back();

    while ( num <= multiples.front())
    {
    	bool is_multiple = true;
        bool is_factor = true;
        
        for (unsigned int i = 0; i < factors.size(); ++i)
        {
            if (num % factors[i] != 0)
            {
                is_multiple = false;
                break;
            }
        }

        for (unsigned int j = 0; j < multiples.size(); ++j)
        {
            if (is_multiple && multiples[j] % num != 0)
            {
                is_factor = false;
                break;
            }
        }

        if (is_multiple && is_factor)
        {
            count++;
        }
        num = factors.back() * times_multiplied;
        times_multiplied++;
    }
    return count;
}

int main()
{
    unsigned short num_factors, num_multiples;
    std::cin >> num_factors >> num_multiples;
    std::vector<unsigned short> factors(num_factors);
    std::vector<unsigned short> multiples(num_multiples);

    for (unsigned short i = 0; i < num_factors; ++i)
    {
        std::cin >> factors[i];
    }
    for (unsigned short i = 0; i < num_multiples; ++i)
    {
        std::cin >> multiples[i];
    }
  
    std::cout << count_between_two_sets(factors, multiples) << "\n";
}
```

View this on [Github](https://github.com/{{site.github_username}}/Competitive-Programming/blob/master/Hackerrank/Between_two_Sets.cpp)

Get this post in pdf - [Between Two Sets - HackerRank](https://www.file-up.org/f6n0nkovrrgp)

<h3>You may also like</h3>
[Kangaroo HackerRank Challenge]({{ site.url }}/blog/2019/11/28/Kangaroo-HackerRank-Challenge-C++-Implementation)<br/>
[Roy and Code Streak HackerEarth Challenge]({{ site.url }}/blog/2019/10/24/Roy-and-Code-Streak-HackerEarth-Challenge)




