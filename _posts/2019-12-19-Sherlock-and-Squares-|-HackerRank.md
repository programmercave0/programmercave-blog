---
layout: post
title: "Sherlock and Squares | HackerRank"
description: "Watson likes to challenge Sherlock's math ability. He will provide a starting and ending value describing a range of integers. Sherlock must determine the number of *square integers* within that range, inclusive of the endpoints."
author: "Programmercave"

tags:  [Cpp, Algorithm, Data-Structure, Competitive-Programming, Hackerrank]
date: 2019-12-19
---



This is an easy hackerrank challenge which will help you to become good at competitive programming. There are various competitive programming websites like [CodeChef](https://www.codechef.com/), [HackerEarth](https://www.hackerearth.com/challenges/), [Codeforces](https://codeforces.com/) where you can practice coding.

<h1>Problem : </h1>

Watson likes to challenge Sherlock's math ability. He will provide a starting and ending value describing a range of integers. Sherlock must determine the number of *square integers* within that range, inclusive of the endpoints.


**Note**: A square integer is an integer which is the square of an integer, e.g. *1, 4, 9, 16, 25*.
 
For example, the range is *a = 24* and *b = 49*, inclusive. There are three square integers in the range: *25, 36* and *49*.

Read full problem : [Sherlock and Squares](https://www.hackerrank.com/challenges/sherlock-and-squares/problem).

<h2>Solution : </h2>

Let the integer variables `first_num` store the first number of the range and `last_num` store the last number of the range.

To find number of square integers in the range `[first_num, last_num]`, first we calculate square root of the `first_num` and an integer variable `count` stores the number of square integers in the range.

```cpp
int sr = sqrt(first_num);
int count = 0;
```

For eg. in the range *[35, 70]*, variable `sr` is equal to *5* since *√35 = 5.916*. If square of `sr` is greater than or equal to the `first_num` and less than or equal to the `last_num` then the number (`sr * sr`) is in the range and is a square integer.

```cpp
while ((sr * sr) <= last_num)
{
    if ((sr * sr) >= first_num)
        count++;
    sr++;
}
```
{% include ads.html %}<br/>

<h2>C++ Implementation</h2>

```cpp
#include <iostream>
#include <cmath>

int num_of_squares(int first_num, int last_num) {
    int sr = sqrt(first_num);
    int count = 0;
    while ((sr * sr) <= last_num)
    {
        if ((sr * sr) >= first_num)
            count++;
        sr++;
    }
    return count;
}

int main()
{
    int test_cases;
    std::cin >> test_cases;

    while (test_cases--)
    {
        int first_num, last_num;
        std::cin >> first_num >> last_num;

        std::cout << num_of_squares(first_num, last_num) << "\n";
    }
}
```

View this code on [GitHub](https://github.com/{{site.github_username}}/Competitive-Programming/blob/master/Hackerrank/Sherlock_and_Squares.cpp).

If you have another solution then please share it in the comments below.

<h3>Other Competitive Programming Problems and Solutions</h3>
[Circular Array Rotation]({{ site.url }}/blog/2019/12/12/Circular-Array-Rotation-HackerRank)<br/>
[Drawing Book HackerRank Challenge]({{ site.url }}/blog/2019/12/11/Drawing-Book-HackerRank)<br/>


