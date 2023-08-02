---
layout: post
title: "Insertion Sort | C++ Implementation"
subtitle: "Insertion sort is an efficient algorithm for sorting a small number of elements. The algorithm selects an element from the unsorted array and put it in the proper position in the sorted. This process is repeated until all elements in the array are sorted.  The sorting is in-place means array consists of sorted portion and unsorted portion in it."
author: "Programmercave"
header-img: "/assets/insertionsort.png"
tags:  [Cpp, Algorithm, Sorting]
date: 2017-08-20
---

**Insertion sort** is an efficient algorithm for sorting a small number of elements. The algorithm selects an element from the unsorted array and put it in the proper position in the sorted. This process is repeated until all elements in the array are sorted.  The sorting is in-place means array consists of sorted portion and unsorted portion in it.

![Insertion Sort]({{ site.url }}/assets/insertionsort.png){:class="img-responsive"}

The index of the *key* starts from 1. The algorithm finds the correct position of the *key* in the array and put the *key* at that position and then the element with next index becomes *key*.

In fig. (d) the index of *key* is 4 and value is 1. Since 1 is the smallest element in the array so far, so it is shifted to index 0.

<h1>Implementation</h1>

Here is the `insertion_sort` and `insertion_sort_stl` function.

```cpp
void insertion_sort(std::vector<int>& vec)
{
    for(std::size_t j = 1; j < vec.size(); j++)
    {
      int key = vec[j];
      int i = j-1;

      while(i >= 0 && vec[i] > key)
      {
         vec[i+1] = vec[i];
         i--;
      }
      vec[i+1] = key;
    }
}
```

We can also use functions from C++ Standard Template Library for insertion sort.

```cpp
void insertion_sort_stl(std::vector<int>& vec)
{
	for(auto it = vec.begin(); it != vec.end(); it++)
 	{
   		// Search
   		auto const insertion_point = std::upper_bound(vec.begin(), it, *it);

   		//insert
   		std::rotate(insertion_point, it, it+1);
 	} 
}
```
`upper_bound` returns an iterator pointing to the first element in the range \[first, last) whose value is larger than the *x*. Here *x* is `*it`.

`rotate` perfoms left rotation. It swaps the elements in the range \[first,last), in such a way that the element pointed by middle becomes the new first element and element pointed by middle-1 becomes last element.

{% include ads.html %}<br/>

For eg. Given array is { 5, 2, 4, 6, 1, 3}  and iterator middle points to 4. After rotation array becomes {4, 6, 1, 3, 5, 2} which is left rotation 2 times.

Here the middle element is `*it` and pointed by iterator `it`.

So when array is `{ 2, 5, 4, 6, 1, 3 }` and in the range \[2, 4), 5 is the greatest so `insertion_point` points to 5 and it points to 4. `rotate` is applied for the range \[5, 6) and it is the middle element so 4 becomes the first element in the range and array becomes `{ 2, 4, 5, 6, 1, 3 }`

Here is the position of elements in the array after each iteration.

![Insertion Sort]({{ site.url }}/assets/insertionsort1.png){:class="img-responsive"}

<h3>C++ implementation of Insertion Sort</h3>

```cpp
#include <iostream>
#include <vector>

void insertion_sort(std::vector<int>& vec)
{
    for(std::size_t j = 1; j < vec.size(); j++)
    {
      int key = vec[j];
      int i = j-1;

      while(i >= 0 && vec[i] > key)
      {
         vec[i+1] = vec[i];
         i--;
      }
      vec[i+1] = key;
    }
}

void print(std::vector<int>& vec) 
{
    for(unsigned i = 0; i < vec.size(); i++)
    {
        std::cout << vec[i] << " ";
    }
    std::cout << "\n";
}

int main()
{
    std::vector<int> arr = {5, 2, 4, 6, 1, 3};
    insertion_sort(arr);
    print(arr);
}
```

View this code on [Github](https://github.com/{{site.github_username}}/Algo-Data-Structure/blob/master/Insertion%20Sort/C++/insertionsort.cpp)

Get this post in pdf - [Insertion Sort](https://www.file-up.org/31lsofzj6c9k)

Reference:<br/>
[Introduction to Algorithms](https://amzn.to/2OarGBs)<br/>
[The Algorithm Design Manual](https://amzn.to/2CH9h9Z)<br/>
[Data Structures and Algorithms Made Easy](https://amzn.to/2NLM0dd)<br/>
Competitive Programmer’s Handbook - Antti Laaksonen<br/>

 <input type="hidden" name="IL_IN_ARTICLE"> 
<h3>You may also like</h3><br/>
[Selection sort]({{ site.url }}/blog/2017/08/29/C++-Selection-sort-using-STL)<br/>
[Merge Sort]({{ site.url }}/blog/2017/08/24/C++-Implementation-of-Merge-Sort)<br/>
[Insertion Sort]({{ site.url }}/blog/2017/08/20/C++-Insertion-Sort-using-STL-(Sorting))<br/>
[Quicksort]({{ site.url }}/blog/2017/07/16/C++-Implementation-of-Quicksort-(Sorting))<br/>
[Heapsort]({{ site.url }}/blog/2017/07/15/C++-Implementation-of-Heapsort-(Sorting))<br/>
