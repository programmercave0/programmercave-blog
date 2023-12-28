---
layout: post
title: "Quicksort | C++ Implementation"
description: "The algorithm selects an element as pivot. The input array is divided into two subarrays. All elements in left subarray are less than pivot and all elements in right subarray are greater than pivot. These two subarrays are sorted independently and then merged to form a single sorted array."
author: "Programmercave"
header-img: "/assets/images/quicksort.png"
tags:  [Cpp, Algorithm, Sorting]
date: 2017-07-16
---



Like mergesort, **quicksort** also follows *divide-and-conquer* approach. The algorithm selects an element as *pivot*. The input array is **divided** into two subarrays. All elements in left subarray are less than pivot and all elements in right subarray are greater than pivot. These two subarrays are **sorted** independently and then **merged** to form a single sorted array.

**Related :** [Merge Sort]({{ site.url }}/blog/2017/08/24/C++-Implementation-of-Merge-Sort)

![Quicksort]({{ site.url }}/assets/images/quicksort.png){:class="img-responsive"}

In fig. (a) 4 is selected as the pivot and in fig. (b) all elements left to the 4 are smaller than it and all elements right to the 4 are greater than it and then these two subarrays are solved independently.

The pivot element 4 in fig. (b) ends up in the correct position in the array and so on all pivot elements ends up in the correct position.

<h1>Implementation</h1>

`split` function returns the index of pivot element.

```cpp
int split(int a[], int start_index, int end_index)
{
    int x = a[end_index];
    int i = start_index - 1;
    for (int j = start_index; j < end_index; j++)
    {
        if (a[j] <= x)
        {
            i++;
            std::swap(a[i], a[j]);//swap(a[i],a[j])
        }
    }
    std::swap(a[i+1], a[end_index]);//swap(a[i+1],a[end_index])

    return i+1;
}
```

Here is the operation of `split` for the first time.

![Quicksort]({{ site.url }}/assets/images/quicksortsplit.png){:class="img-responsive"}

```cpp
void quicksort(int a[], int start_index, int end_index)
{
    if (start_index < end_index)
    {
        int mid_index = split(a, start_index, end_index);
                        quicksort(a, start_index, mid_index - 1);
                        quicksort(a, mid_index + 1, end_index);
    }
}
```

{% include ads.html %}<br/>

<h3>C++ Implementation of Quicksort</h3>

```cpp
#include <iostream>
#include <algorithm>

int split(int a[], int start_index, int end_index)
{
    int x = a[end_index];
    int i = start_index - 1;
    for (int j = start_index; j < end_index; j++)
    {
        if (a[j] <= x)
        {
            i++;
            std::swap(a[i], a[j]);//swap(a[i],a[j])
        }
    }
    std::swap(a[i+1], a[end_index]);//swap(a[i+1],a[end_index])

    return i+1;
}

void quicksort(int a[], int start_index, int end_index)
{
    if (start_index < end_index)
    {
        int mid_index = split(a, start_index, end_index);
                        quicksort(a, start_index, mid_index - 1);
                        quicksort(a, mid_index + 1, end_index);
    }
}

int main()
{
    int arr[8] = {2, 8, 7, 1, 3, 5, 6, 4};
    quicksort(arr, 0, 7);
    for (int i = 0; i < 8; i++)
    {
        std::cout << arr[i] << " ";
    }
    std::cout << "\n";
    return 0;
}
```

View this code on [Github](https://github.com/{{site.github_username}}/Algo-Data-Structure/blob/master/Quick%20Sort/C++/quicksort.cpp)

Get this post in pdf - [Quicksort](https://www.file-up.org/hj7tdj61r4sj)

Reference:<br/>
[Introduction to Algorithms](https://amzn.to/2OarGBs)<br/>
[The Algorithm Design Manual](https://amzn.to/2CH9h9Z)<br/>
[Data Structures and Algorithms Made Easy](https://amzn.to/2NLM0dd)<br/>
Competitive Programmer’s Handbook - Antti Laaksonen<br/>

 <input type="hidden" name="IL_IN_ARTICLE"> 
<h3>You may also like</h3>

[Selection sort]({{ site.url }}/blog/2017/08/29/C++-Selection-sort-using-STL)<br/>
[Insertion Sort]({{ site.url }}/blog/2017/08/20/C++-Insertion-Sort-using-STL-(Sorting))<br/>
[Heapsort]({{ site.url }}/blog/2017/07/15/C++-Implementation-of-Heapsort-(Sorting))<br/>




