---
layout: post
title: "Left Rotation - HackerRank | C++ Implementation"
subtitle: "A left rotation operation on an array of size n shifts each of the array's elements 1 unit to the left. For example, if 2 left rotations are performed on array [1, 2, 3, 4, 5], then the array would become  [3, 4, 5, 1, 2]. Given an array of n integers and a number, d, perform d left rotations on the array. Then print the updated array as a single line of space-separated integers."
author: "Programmercave"
header-img: "/assets/2020-04-25-Left-Rotation-HackerRank/HR_repeated_string.jpg"
tags:  [Cpp, Algorithm, Competitive-Programming, Hackerrank, Data-Structure]
date: 2020-04-25
---

<h1>Problem:</h1>

A left rotation operation on an array of size *n* shifts each of the array's elements *1* unit to the left. For example, if *2* left rotations are performed on array `[1, 2, 3, 4, 5]`, then the array would become  `[3, 4, 5, 1, 2]`.

Given an array of *n* integers and a number, *d*, perform *d* left rotations on the array. Then print the updated array as a single line of space-separated integers.

Read the full problem here : [Left Rotation](https://www.hackerrank.com/challenges/array-left-rotation/copy-from/155235597)

<h1>Solution:</h1>

![Left Rotation HackerRank]({{ site.url }}/assets/2020-04-25-Left-Rotation-HackerRank/HR_repeated_string.jpg){:class="img-responsive"}

This is a very easy problem. In the above figure we see that after two rotations, the element at index 2 is the first element in the new array. Thus after *d* rotations, the element at index *d* is the first element. And all elements after *d*<sup>th</sup> elements with follow.

```cpp
for (std::size_t i = rotations; i < arr.size(); ++i)
{
    rotated_array.push_back(arr[i]);
}
```

Now, elements from 0 to *d*-1 indices are added to the new rotated array.

```cpp
for (std::size_t i = 0; i < rotations; ++i)
{
    rotated_array.push_back(arr[i]);
}
```

{% include ads.html %}<br/>

<h3>C++ Implementation</h3>

```cpp
#include <iostream>
#include <vector>

std::vector<int> left_rotation(std::vector<int>& arr, int rotations)
{
    std::vector<int> rotated_array;

    for (std::size_t i = rotations; i < arr.size(); ++i)
    {
        rotated_array.push_back(arr[i]);
    }
    for (std::size_t i = 0; i < rotations; ++i)
    {
        rotated_array.push_back(arr[i]);
    }
    return rotated_array;
}

int main()
{
    int num_elements, num_left_rotations;

    std::cin >> num_elements;
    std::cin >> num_left_rotations;

    std::vector<int> input_array(num_elements);
    for (int i = 0; i < num_elements; ++i)
    {
        std::cin >> input_array[i];
    }

    std::vector<int> result_array = left_rotation(input_array, num_left_rotations);

    for (int i = 0; i < num_elements; ++i)
    {
        std::cout << result_array[i] <<" ";
    }
    std::cout << "\n";
}

```

View this on [Github](https://github.com/{{site.github_username}}/Competitive-Programming/tree/master/Hackerrank)

If you have another solution then please share it in the comments below.

<h3>Other Competitive Programming Problems and Solutions</h3>
[Repeated String]({{ site.url }}/blog/2020/04/24/Repeated-String-HackerRank-Challenge-Cpp-Implementation)<br/>
[Picking Numbers]({{ site.url }}/blog/2020/03/28/Picking-Numbers-HackerRank-Challenge-Cpp-Implementation)<br/>





