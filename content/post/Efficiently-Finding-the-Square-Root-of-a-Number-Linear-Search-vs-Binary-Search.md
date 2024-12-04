---
author: Programmercave
date: "2023-03-03T00:00:00Z"
description: Finding the square root of a number is a common problem in mathematics
  and computer science. In this blog post, we will focus on the linear search and
  binary search methods for finding the square root of a number, and provide an implementation
  in C++ for each method.
header-img: /assets/images/Square-Root/sq_root1.png
tags:
- Cpp
- Algorithm
- Data-Structures
- Competitive-Programming
- Mathematics
title: 'Efficiently Finding the Square Root of a Number: Linear Search vs Binary Search'
toc: true
---

## Introduction

Finding the square root of a number is a common problem in mathematics and computer science. In this blog post, we will focus on the linear search and binary search methods for finding the square root of a number, and provide an implementation in C++ for each method.

## Linear Search Method

The linear search method is a simple algorithm that iteratively checks each integer number from 1 to `n` to see if its square is equal to the input number n. If a number `i` is found such that `i*i = n`, then `i` is returned as the square root of `n`. If no such number is found, then the algorithm returns nothing.

Here is the implementation of the linear search method in C++:

```cpp
int linear_search_mtd(int n)
{
    for (int i = 1; i <= n; ++i)
    {
        if (i*i == n)
            return i;
    }
}
```

## Binary Search Method

The binary search method is a more efficient algorithm for finding the square root of a number. It works by repeatedly dividing the search interval in half and checking if the middle number squared is equal to `n`. If the middle number squared is less than n, then the search interval is updated to the right half of the interval. Otherwise, the search interval is updated to the left half of the interval. This process continues until the square root of n is found.

![Square Root of Number](/assets/images/Square-Root/sq_root1.png)
![Square Root of Number](/assets/images/Square-Root/sq_root2.png)

Here is the implementation of the binary search method in C++:

```cpp
int binary_search_mtd(int n)
{
    int l = 1, r = n;

    while (l <= r)
    {
        int mid = (l+r)/2;

        if (mid*mid == n)
            return mid;
        else if (mid*mid < n)
            l = mid+1;
        else 
            r = mid-1;
    }
}
```

## Comparison between Linear Search and Binary Search Methods

The binary search method is more efficient than the linear search method for finding the square root of a number. This is because the binary search method has a time complexity of **O(log n)**, while the linear search method has a time complexity of **O(n)**. In other words, the binary search method can find the square root of a number much faster than the linear search method, especially for large values of n.

If you'd like to see the complete code for checking prime numbers using an optimized algorithm, you can find it on [Github](https://github.com/{{< param "github_username" >}}/Algo-Data-Structure/blob/master/Maths/find_sq_root.cpp)
.

## Conclusion

In conclusion, finding the square root of a number is a common problem in mathematics and computer science. The linear search method and binary search method are two algorithms that can be used to solve this problem. The binary search method is more efficient than the linear search method, especially for large values of n. Both methods have been implemented in C++ and can be used depending on the specific requirements of the problem at hand.

---



By the way, if you're a teacher or parent looking for resources to help your child get ready for school, you might be interested in these fun and informative workbooks developed by a pre-school teacher. Covering all the basic skills needed for school-readiness, they're perfect for the pre-school education niche. Check them out here: [WORKSHEETS FOR PRESCHOOL](https://ce8977zhz1vrft28uay3ofipe9.hop.clickbank.net/?cbpage=wfpaffiliate)
