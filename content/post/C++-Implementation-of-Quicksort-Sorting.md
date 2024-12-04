---
author: Programmercave
date: "2017-07-16T00:00:00Z"
description: The algorithm selects an element as pivot. The input array is divided
  into two subarrays. All elements in left subarray are less than pivot and all elements
  in right subarray are greater than pivot. These two subarrays are sorted independently
  and then merged to form a single sorted array.
header-img: /assets/images/quicksort.png
tags:
- Cpp
- Algorithm
- Sorting
title: Quicksort | C++ Implementation
---



Like mergesort, **quicksort** also follows *divide-and-conquer* approach. The algorithm selects an element as *pivot*. The input array is **divided** into two subarrays. All elements in left subarray are less than pivot and all elements in right subarray are greater than pivot. These two subarrays are **sorted** independently and then **merged** to form a single sorted array.

**Related :** [Merge Sort](/C-Implementation-of-Merge-Sort)

![Quicksort](/assets/images/quicksort.png)

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

![Quicksort](/assets/images/quicksortsplit.png)

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

<br/>

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

Reference:<br/>
[Introduction to Algorithms](https://amzn.to/2OarGBs)<br/>
[The Algorithm Design Manual](https://amzn.to/2CH9h9Z)<br/>
[Data Structures and Algorithms Made Easy](https://amzn.to/2NLM0dd)<br/>
Competitive Programmer’s Handbook - Antti Laaksonen<br/>

 <input type="hidden" name="IL_IN_ARTICLE"> 
<h3>You may also like</h3>

[Selection sort](/C-Selection-sort-using-STL)<br/>
[Insertion Sort](/C-Insertion-Sort-using-STL-Sorting)<br/>
[Heapsort](/C-Implementation-of-Heapsort-Sorting)<br/>




