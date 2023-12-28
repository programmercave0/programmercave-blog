---
layout: post
title: "Kefa and First Steps - CodeForces | C++ Implementation"
description: "Kefa decided to make some money doing business on the Internet for exactly *n* days. He knows that on the *i*-th day (1≤*i*≤n) he makes *a<sub>i</sub>* money. Kefa loves progress, that's why he wants to know the length of the maximum non-decreasing subsegment in sequence *a<sub>i</sub>*. Let us remind you that the subsegment of the sequence is its continuous fragment. A subsegment of numbers is called non-decreasing if all numbers in it follow in the non-decreasing order."
author: "Programmercave"
header-img: ""
tags:  [Cpp, Competitive-Programming, CodeForces, Dynamic-Programming]
date: 2021-06-28
toc: true
---
<h1>Problem:</h1>

Kefa decided to make some money doing business on the Internet for exactly *n* days. He knows that on the *i*-th day (1≤*i*≤n) he makes *a<sub>i</sub>* money. Kefa loves progress, that's why he wants to know the length of the maximum non-decreasing subsegment in sequence *a<sub>i</sub>*. Let us remind you that the subsegment of the sequence is its continuous fragment. A subsegment of numbers is called non-decreasing if all numbers in it follow in the non-decreasing order.

<h3>Input</h3>
The first line contains integer *n*.

The second line contains *n* integers *a<sub>1</sub>*, *a<sub>2</sub>*, ..., *a<sub>n</sub>*

<h3>Output</h3>

Print a single integer — the length of the maximum non-decreasing subsegment of sequence *a*.

Read full problem here : [Kefa and First Steps](https://codeforces.com/problemset/problem/580/A)

<h1>Solution:</h1>
 
Let `earnings` is an integer vector of size `n` which stores amount earned on each day. 

To find the length of the maximum increasing subsegment in a given sequence, you can use the following algorithm:

1. Initialize three integer variables: `current_earning`, `final_count`, and `local_count`. `current_earning` should be set to the first element of the `earnings` vector, `final_count` should be set to 1, and `local_count` should be set to 1.

	```
	int current_earning = earnings[0]; // stores highest earning of non-decreasing subsegment
	int final_count = 1; // stores length of longest non-decreasing subsegment
	int local_count = 1; // stores length of current subsegment in process
	```

2. Iterate over the `earnings` vector, starting at the second element (index 1). For each element `earnings[i]`, do the following:

	a. If `earnings[i] > current_earning`, increment `local_count` by 1 and set `current_earning` to `earnings[i]`.

	b. If `earnings[i] <= current_earning`, set `local_count` to 1 and set `current_earning` to `earnings[i]`.

	c. If `local_count > final_count`, set `final_count` to `local_count`.

	```
	for (int i = 1; i < n; i++) 
	{
    	if (earnings[i] > current_earning) 
    	{
      		// if current earning is greater than previous earning, increment local count by 1
      		local_count++;
      		current_earning = earnings[i];
    	} 
    	else 
    	{
      		// if current earning is not greater than previous earning, reset local count to 1
      		local_count = 1;
      		current_earning = earnings[i];
    	}

    	// update final count if local count is greater than final count
    	if (local_count > final_count) 
    	{
      		final_count = local_count;
    	}
	}
	```

3. After the loop finishes, final_count will contain the length of the maximum increasing subsegment.
{% include ads.html %}<br/>

<h3>C++ Implementation</h3>

```cpp
#include <iostream>
#include <vector>

int main()
{
	int n;
	std::cin >> n;

	std::vector<int> earnings(n);

	int current_earning = earnings[0]; // stores highest earning of non-decreasing subsegment
 	int final_count = 1; // stores length of longest non-decreasing subsegment
  	int local_count = 1; // stores length of current subsegment in process

	for (int i = 1; i < n; i++) 
	{
    	if (earnings[i] > current_earning) 
    	{
      		// if current earning is greater than previous earning, increment local count by 1
      		local_count++;
      		current_earning = earnings[i];
    	} 
    	else 
    	{
      		// if current earning is not greater than previous earning, reset local count to 1
      		local_count = 1;
      		current_earning = earnings[i];
    	}

    	// update final count if local count is greater than final count
    	if (local_count > final_count) 
    	{
      		final_count = local_count;
    	}
	}
	std::cout << final_count << "\n";
}
```

Check out this on [Github](https://github.com/{{site.github_username}}/Competitive-Programming/edit/master/Codeforces/Kefa_and_first_steps.cpp)

If you have a different solution for finding the length of the maximum increasing subsegment in a given sequence, please share it in the comments below.