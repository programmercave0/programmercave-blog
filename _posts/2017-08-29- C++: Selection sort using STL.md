---
layout: post
title: "Selection sort | C++ Implementation"
description: "Selection sort is an in-place sorting algorithm. In the input array there is a sorted portion and an unsorted portion. The algorithm repeatedly finds the smallest element in the unsorted portion of the array and puts it at the end of the sorted portion of the array."
author: "Programmercave"
header-img: "/assets/selectionsort.png"
tags:  [Cpp, Algorithm, Sorting]
date: 2017-08-29
---
* toc
{:toc}

**Selection sort** is an in-place sorting algorithm. In the input array there is a sorted portion and an unsorted portion. The algorithm repeatedly finds the smallest element in the unsorted portion of the array and puts it at the end of the sorted portion of the array.

![Selectionsort]({{ site.url }}/assets/selectionsort.png){:class="img-responsive"}

First the algorithm finds the smallest element in the array which is 1 and it is added to the sorted array and then the algorithm finds smallest element in the remaining array and so on.

<h1>Implementation</h1>

Here is the implementation of `selection_sort` function.

```cpp
void selection_sort(int array[], int size)
{
    int start_index = 0;
    while(start_index < size)
    {
        int min_index = start_index;
        for(int i = start_index+1; i < size; i++)
        {
            if(array[i] < array[min_index])
            {
                min_index = i;
            }
        }
        std::swap(array[start_index], array[min_index]);
        start_index++;
    }
}
```

The above function works only for integer value. Here is the generic function which works for all data types.

```cpp
template<typename T>
void selection_sort(std::vector<T>& array)
{
    typedef typename std::vector<T>::iterator Itr;
    Itr itr = array.begin();
    while(itr != array.end())
    {
        Itr itr_min = itr;
        for(Itr i = itr + 1; i != array.end(); i++)
        {
            if(*i < *itr_min)
            {
                itr_min = i;
            }
        }
        std::iter_swap(itr, itr_min);
        itr++;
    }
}
```

The function can take a vector of integers or vector of characters or a vector of strings, it sorts all of them. `itr` is the iterator, initially which points to the first element of the array. `itr_min` is the iterator which points to the smallest element in the range `[i, array.end())`. `std::iter_swap` method is used to swap `itr` and `itr_min` iterator and it puts the smallest element in the range, at the end of the sorted portion in the array.

`std::swap` swap the values whereas `std::iter_swap` swap the iterators.

{% include ads.html %}<br/>

Here is the output after running generic function.

![Selectionsort]({{ site.url }}/assets/selectionsortout.png){:class="img-responsive"}

<h3>C++ Implementation of Selection Sort</h3>

```cpp
#include <iostream>
#include <vector>

template<typename T>
void selection_sort(std::vector<T>& array)
{
    typedef typename std::vector<T>::iterator Itr;
    Itr itr = array.begin();
    while(itr != array.end())
    {
        Itr itr_min = itr;
        for(Itr i = itr + 1; i != array.end(); i++)
        {
            if(*i < *itr_min)
            {
                itr_min = i;
            }
        }
        std::iter_swap(itr, itr_min);
        itr++;
    }
}

template<typename T>
void print(const std::vector<T>& array)
{
    for(auto itr = array.begin(); itr != array.end(); itr++)
    {
       std::cout << *itr << " ";
    }
    std::cout << '\n';
}

int main()
{
    std::vector<int> v({5, 3, 12, 2, 8});
    std::cout << "Original Array :";
    print(v);
    selection_sort(v);
    std::cout <<"Sorted Array :";
    print(v);
    std::cout << '\n';

    std::vector<char> c({'t', 'q', 'a', 'r', 'p'});
    std::cout << "Original Array :";
    print(c);
    selection_sort(c);
    std::cout <<"Sorted Array :";
    print(c);
    std::cout << '\n';

    std::vector<std::string> str({"code", "live", "love", "sing", "create"});
    std::cout << "Original Array :";
    print(str);
    selection_sort(str);
    std::cout <<"Sorted Array :";
    print(str);
    std::cout << '\n';
}
```

You can view this code on [Github](https://github.com/{{site.github_username}}/Algo-Data-Structure/blob/master/Selection%20Sort/C++/selectionsort.cpp)

Get this post in pdf - [Selection Sort](https://www.file-up.org/34e0v6l3iomj)

Reference:<br/>
[Introduction to Algorithms](https://amzn.to/2OarGBs)<br/>
[The Algorithm Design Manual](https://amzn.to/2CH9h9Z)<br/>
[Data Structures and Algorithms Made Easy](https://amzn.to/2NLM0dd)<br/>
Competitive Programmer’s Handbook - Antti Laaksonen<br/>

 <input type="hidden" name="IL_IN_ARTICLE"> 
<h3>You may also like</h3>

[Merge Sort]({{ site.url }}/blog/2017/08/24/C++-Implementation-of-Merge-Sort)<br/>
[Insertion Sort]({{ site.url }}/blog/2017/08/20/C++-Insertion-Sort-using-STL-(Sorting))<br/>
[Quicksort]({{ site.url }}/blog/2017/07/16/C++-Implementation-of-Quicksort-(Sorting))<br/>
[Heapsor]({{ site.url }}/blog/2017/07/15/C++-Implementation-of-Heapsort-(Sorting))<br/>



