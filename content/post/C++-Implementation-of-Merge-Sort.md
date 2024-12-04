---
author: Programmercave
date: "2017-08-24T00:00:00Z"
description: Merge sort follows *divide-and-conquer* approach. It divides an array
  of n elements into two subarrays of n/2 elements each. Then it sort the two subarrays
  recursively using merge sort. And then these subarrays are merged to produce a single
  sorted array.
header-img: /assets/images/mergesort.png
tags:
- Cpp
- Algorithm
- Sorting
title: Merge Sort | C++ Implementation
---



**Merge sort** follows *divide-and-conquer* approach. It **divides** an array of *n* elements into two subarrays of *n*/2 elements each. Then it **sort** the two subarrays recursively using merge sort. And then these subarrays are **merged** to produce a single sorted array.

![Mergesort](/assets/images/mergesort.png)

If the size of the array is even then the size of subarrays is equal and if it is odd then first array has one element more than the second array. The division of array stops when subarrays contain only one element and then merging of subaarrays starts. The arrays are merged in sorted order.

<h1>Implementation</h1>

Here is the generic implementation of `merge` and `merge_sort` function. It sorts array of all datatypes.

```cpp
template<class T>
void merge(std::vector<T>& v, int p, int q, int r)
{
    int size1 = q-p+1;
    int size2 = r-q;
    std::vector<T> L(size1);
    std::vector<T> R(size2);

    for(int i = 0; i < size1; i++)
    {
        L[i] = v[p+i];
    }
    for(int j = 0; j < size2; j++)
    {
        R[j]=v[q+j+1];
    }

    int i=0,j=0;
    int k;
    for(k = p; k <= r && i < size1 && j < size2; k++)
    {
        if(L[i] <= R[j])
        {
            v[k] = L[i];
            i++;
        }
        else
        {
            v[k] = R[j];
            j++;
        }
    }
    for(i = i; i < size1; ++i)
    {
        v[k] = L[i];
        k++;
    }

    for(j = j; j < size2; j++)
    {
        v[k] = R[j];
        k++;
    }
}
```

When `merge` function is called, two auxiliary vectors `L` and `R` are created. `L` is initialised with elements of left subarray and `R` is initialised with elements of right subarray. Both subarrays contain one element each. For eg. from the above fig. `L` contains 5 and `R` contains 2.

<br/>

When we merge these subarrays, the smallest element from these subarray is selected and the element is entered in the input array `v` at index 0. `k` maintains the index of input array and `i` and `j` maintains the index of `L` and `R` subarrayarray respectively.

If all the elements of one of the subarray are entered then remaining elements of the another subarrays are entered in the input array since all elements are greater than the last input value. This is done with the help of last two `for` loops in the above function.

```cpp
template<class T>
void merge_sort(std::vector<T>& v, int p, int r)
{
    if(p < r)
    {
        int q = (p+r)/2;
        merge_sort(v, p, q);
        merge_sort(v, q+1, r);
        merge(v, p, q, r);
    }
}
```

`merge_sort` recursively calls itself until `r` is smaller than or equal to `p`.

Here is the output of merge sort with generic functions.

![Mergesort](/assets/images/mergesortout.png)

<h3>C++ Implementation of Merge Sort</h3>
    
 ```cpp
 #include <iostream>
#include <vector>

template<class T>
void merge(std::vector<T>& v, int p, int q, int r)
{
    int size1 = q-p+1;
    int size2 = r-q;
    std::vector<T> L(size1);
    std::vector<T> R(size2);

    for(int i = 0; i < size1; i++)
    {
        L[i] = v[p+i];
    }
    for(int j = 0; j < size2; j++)
    {
        R[j]=v[q+j+1];
    }

    int i=0,j=0;
    int k;
    for(k = p; k <= r && i < size1 && j < size2; k++)
    {
        if(L[i] <= R[j])
        {
            v[k] = L[i];
            i++;
        }
        else
        {
            v[k] = R[j];
            j++;
        }
    }
    for(i = i; i < size1; ++i)
    {
        v[k] = L[i];
        k++;
    }

    for(j = j; j < size2; j++)
    {
        v[k] = R[j];
        k++;
    }
}

template<class T>
void merge_sort(std::vector<T>& v, int p, int r)
{
    if(p < r)
    {
        int q = (p+r)/2;
        merge_sort(v, p, q);
        merge_sort(v, q+1, r);
        merge(v, p, q, r);
    }
}

int main()
{
    std::vector<int>vec;
    vec.push_back(13);
    vec.push_back(5);
    vec.push_back(7);
    vec.push_back(7);
    vec.push_back(4);
    vec.push_back(2);
    vec.push_back(10);
    int sz = vec.size();
    std::cout << "Entered array : ";
    for(int n = 0; n < sz; n++)
    {
        std::cout << vec[n] <<" ";
    }
    std::cout << "\n";
    std::cout << "Sorted array : ";
    merge_sort(vec, 0, sz-1);
    for(int n = 0; n < sz; n++)
    {
        std::cout << vec[n] << " ";
    }
    std::cout << "\n\n";

    std::vector<char> c;
    c.push_back('d');
    c.push_back('y');
    c.push_back('h');
    c.push_back('l');
    c.push_back('e');
    c.push_back('a');
    int sz1 = c.size();
    std::cout << "Entered array : ";
    for(int n = 0; n < sz1; n++)
    {
        std::cout << c[n] <<" ";
    }
    std::cout << "\n";
    std::cout << "Sorted array : ";
    merge_sort(c, 0, sz1-1);
    for(int n = 0; n < sz1; n++)
    {
        std::cout << c[n] << " ";
    }
    std::cout << "\n\n";

    std::vector<std::string> str;
    str.push_back("car");
    str.push_back("dog");
    str.push_back("apple");
    str.push_back("ball");
    str.push_back("tree");
    int sz2 = str.size();

    std::cout << "Entered array : ";
    for(int n = 0; n < sz2; n++)
    {
        std::cout << str[n] <<" ";
    }
    std::cout << "\n";
    std::cout << "Sorted array : ";
    merge_sort(str, 0, sz2-1);
    for(int n = 0; n < sz2; n++)
    {
        std::cout << str[n] << " ";
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
[Quicksort](/C-Implementation-of-Quicksort-Sorting)<br/>
[Heapsort](/C-Implementation-of-Heapsort-Sorting)<br/>
