---
layout: post
title: "Picking Numbers - HackerRank | C++ Implementation"
description: "Given an array of integers, find and print the maximum number of integers you can select from the array such that the absolute difference between any two of the chosen integers is less than or equal to 1. For example, if your array is a = [1, 1, 2, 2, 4, 4, 5, 5, 5], you can create two subarrays meeting the criterion: [1, 1, 2, 2] and [4, 4, 5, 5, 5]. The maximum length subarray has 5 elements."
author: "Programmercave"
header-img: "/assets/2020-03-28-Picking-Numbers-Hacker-Challenge/HR_picking_numbers.jpg"
tags:  [Cpp, Algorithm, Competitive-Programming, Hackerrank, Data-Structure]
date: 2020-03-28
---
* toc
{:toc}

<h1>Problem:</h1>

Given an array of integers, find and print the maximum number of integers you can select from the array such that the absolute difference between any two of the chosen integers is less than or equal to 1. 

For example, if your array is a = [1, 1, 2, 2, 4, 4, 5, 5, 5], you can create two subarrays meeting the criterion: [1, 1, 2, 2] and [4, 4, 5, 5, 5]. The maximum length subarray has 5 elements.

Read the full problem here: [Picking Numbers](https://www.hackerrank.com/challenges/picking-numbers/problem)

<h1>Solution:</h1>

To find the subarrays which satisfy the above conditions, the input array must be sorted. So our first statement in the function is

```cpp
std::sort(array.begin(), array.end());
```

Integer variable `result` will store the length of the subarray with maximum size, `count` will store the length of the subarray being processed and `subarray_first_element` will store the first element of the subarray being processed.

`result` is initialised with `0`, `subarray_first_element` with the first element of the array and `count` with `1`.

```cpp
int result = 0;
int count = 1;
int subarray_first_element = array[0];
```

`count` is initialised with `1` because we iterate from the index 1 of the array. The element being considered is checked with `subarray_first_element` and if the condition is `true`, `count` is incremented and we have two elements in the subarray.

![Picking Numbers HackerRank]({{ site.url }}/assets/2020-03-28-Picking-Numbers-Hacker-Challenge/HR_picking_numbers.jpg){:class="img-responsive"}

If the condition is `false`, `result` is updated if the `count` is greater than the `result` (or length of this subarray is greater than the previous subarray). `count` is initialised with `1` and `subarray_first_element` with `array[i]`.

```cpp
for (int i = 1; i < array.size(); ++i)
{
	if (subarray_first_element == array[i] || subarray_first_element + 1 == array[i])
	{
		count++;
	}
	else
	{
		if (count > result)
		{
			result = count;
		}
		count = 1;
		subarray_first_element = array[i];
	}
}
```

Once the loop is executed again we check if the `count` is greater than the `result`, because if the last two elements of the array satisfies the above conditions then the control did not enter in the else scope in the for loop and `result` is not updated with the correct `count`.

```cpp
if (count > result)
{
	result = count;
}
return result;
```

{% include ads.html %}<br/>

<h3>C++ Implementation</h3>

```cpp
#include <iostream>
#include <algorithm>
#include <vector>

int picking_numbers(std::vector<int>& array)
{
	std::sort(array.begin(), array.end());

	int result = 0;
	int count = 1;
	int subarray_first_element = array[0];

	for (int i = 1; i < array.size(); ++i)
	{
		if (subarray_first_element == array[i] || subarray_first_element + 1 == array[i])
		{
			count++;
		}
		else
		{
			if (count > result)
			{
				result = count;
			}
			count = 1;
			subarray_first_element = array[i];
		}
	}

	if (count > result)
	{
		result = count;
	}
	return result;
}

int main()
{
	int size_of_array;
	std::cin >> size_of_array;
	std::vector<int> array(size_of_array);

	for (int i = 0; i < size_of_array; ++i)
	{
		std::cin >> array[i];
	}

	int max_count = picking_numbers(array);
	std::cout << max_count << "\n";
}
```

View this on [Github](https://github.com/{{site.github_username}}/Competitive-Programming/blob/master/Hackerrank/Picking_Numbers.cpp)

If you have another solution then please share it in the comments below.

<h3>Other Competitive Programming Problems and Solutions</h3>
[Library Fine]({{ site.url }}/blog/2019/12/20/Library-Fine-HackerRank)<br/>
[Sherlock and Squares]({{ site.url }}/blog/2019/12/19/Sherlock-and-Squares-HackerRank)<br/>



