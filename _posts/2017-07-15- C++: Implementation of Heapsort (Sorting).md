---
layout: post
title: "Heapsort | C++ Implementation"
subtitle: "Heapsort is implemented using heap data structure. Heap helps us to represent binary tree without using any pointers. Using heap an array can be viewed as a binary tree and each node of the tree stores an element of the array."
author: "Programmercave"
header-img: "/assets/heapsort.png"
tags:  [Cpp, Algorithm, Sorting]
date: 2017-07-15
---

**Heapsort** is implemented using *heap* data structure. Heap helps us to represent binary tree without using any *pointers*. Using heap an array can be viewed as a *binary tree* and each node of the tree stores an element of the array.

There are two kinds of binary heaps: max-heaps and min-heaps. In *max-heap*, the value stored at the parent node is greater than the value stored at its children nodes. Thus in a max-heap, root node contains the largest element. In *min-heap*, the value stored at the parent node is smaller than the value stored at its children nodes. Thus in a min-heap, root node contains the smallest element.

![Heapsort]({{ site.url }}/assets/binaryheap.png){:class="img-responsive"}

Max-heap is used in heapsort algorithm and min-heap is used in priority queues.

![Heapsort]({{ site.url }}/assets/heapsort.png){:class="img-responsive"}

When `arr[i] = parent`, then `left_child = 2*i + 1` and `right_child = 2*i + 2`.

<h1>Implementation</h1>

`max_heapify` maintains the max-heap property of the heap. The input array, index of the element and size of the array is passed as an argument. 

```cpp
void max_heapify(std::vector<int>& arr, int i, int size_)
{
    int largest, l = (2*i) + 1, r = l + 1;

    if(l < size_ && arr[l] > arr[i])
        largest = l;
    else
        largest = i;

    if(r < size_ && arr[r] > arr[largest])
        largest = r;

    if(largest != i)
    {
        std::swap(arr[i], arr[largest]);
        max_heapify(arr, largest, size_);
    }
}
```

If `arr[i]` is `largest`, then subtree rooted at node `i` is a max-heap and function terminates. Otherwise, `largest` stores the index of child whose value is greatest of the three elements and `arr[i]` is swapped with `arr[largest]` and thus max-heap property is satisfied at node `i`. Then `max_heapify` with node indexed by `largest` is called to satisfy max-heap property at node `largest`.

`build_max_heap` produces a max-heap from an input array.

```cpp
void build_max_heap(std::vector<int>& arr)
{
    for(int i = (arr.size() / 2); i >= 0; i--)
    max_heapify(arr, i, arr.size());
}
```

![Heapsort]({{ site.url }}/assets/heapsort1.png){:class="img-responsive"}

`heapsort` sorts an array in-place.

```cpp
void heap_sort(std::vector<int>& arr)
{
   build_max_heap(arr);
   int sz = arr.size();
   for(int i = arr.size() - 1; i > 0; i--)
   {
        std::swap(arr[0], arr[i]);
        sz--;
        max_heapify(arr, 0, sz);
    }
}
```
{% include ads.html %}<br/>

`heapsort` starts with `build_max_heap` and now largest element of the array is at index 0. So the first value is  swapped with the last value and then the node with largest value is removed from the tree and new max-heap is created with `max_heapify`.

<h3>C++ Implementation of Heapsort</h3>

```cpp
#include <iostream>
#include <vector>
#include <algorithm>

void max_heapify(std::vector<int>& arr, int i, int size_)
{
    int largest, l = (2*i) + 1, r = l + 1;

    if(l < size_ && arr[l] > arr[i])
        largest = l;
    else
        largest = i;

    if(r < size_ && arr[r] > arr[largest])
        largest = r;

    if(largest != i)
    {
        std::swap(arr[i], arr[largest]);
        max_heapify(arr, largest, size_);
    }
}

void build_max_heap(std::vector<int>& arr)
{
    for(int i = (arr.size() / 2); i >= 0; i--)
    max_heapify(arr, i, arr.size());
}

void heap_sort(std::vector<int>& arr)
{
   build_max_heap(arr);
   int sz = arr.size();
   for(int i = arr.size() - 1; i > 0; i--)
   {
        std::swap(arr[0], arr[i]);
        sz--;
        max_heapify(arr, 0, sz);
    }
}

int main()
{
    std::vector<int> arr = {4, 1, 3, 2, 16, 9, 10, 14, 8, 7};
    heap_sort(arr);
    
    for(int i = 0; i < arr.size(); i++)
    {
         std::cout << arr[i] << " ";
    }
    std::cout << "\n";
    return 0;
}
```

View this code on [Github](https://github.com/{{site.github_username}}/Algo-Data-Structure/blob/master/Heapsort/C++/heapsort.cpp)

Get this post in pdf - [Heapsort](https://www.file-up.org/5etqv0d4uih4)

Reference:<br/>
[Introduction to Algorithms](https://amzn.to/2OarGBs)<br/>
[The Algorithm Design Manual](https://amzn.to/2CH9h9Z)<br/>
[Data Structures and Algorithms Made Easy](https://amzn.to/2NLM0dd)<br/>
Competitive Programmer’s Handbook - Antti Laaksonen<br/>

 <input type="hidden" name="IL_IN_ARTICLE"> 
<h3>You may also like</h3>

[Selection sort]({{ site.url }}/blog/2017/08/29/C++-Selection-sort-using-STL)<br/>
[Merge Sort]({{ site.url }}/blog/2017/08/24/C++-Implementation-of-Merge-Sort)<br/>
[Insertion Sort]({{ site.url }}/blog/2017/08/20/C++-Insertion-Sort-using-STL-(Sorting))<br/>
[Quicksort]({{ site.url }}/blog/2017/07/16/C++-Implementation-of-Quicksort-(Sorting))<br/>

