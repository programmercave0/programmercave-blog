---
layout: post
title: "Migratory Birds - HackerRank Challenge | C++ Implementation"
description: "You have been asked to help study the population of birds migrating across the continent. Each type of bird you are interested in will be identified by an integer value. Each time a particular kind of bird is spotted, its id number will be added to your array of sightings. You would like to be able to find out which type of bird is most common given a list of sightings. Your task is to print the type number of that bird and if two or more types of birds are equally common, choose the type with the smallest ID number."
author: "Programmercave"
header-img: "/assets/images/2019-12-01-Migratory-Birds-HackerRank-Challenge-C++-Implementation/HR_migratory-birds.jpg"
tags:  [Cpp, Algorithm, Data-Structure, Competitive-Programming, Hackerrank]
date: 2019-12-01
---



<h1>Problem :</h1>

You have been asked to help study the population of birds migrating across the continent. Each type of bird you are interested in will be identified by an integer value. Each time a particular kind of bird is spotted, its id number will be added to your array of sightings. You would like to be able to find out which type of bird is most common given a list of sightings. Your task is to print the type number of that bird and if two or more types of birds are equally common, choose the type with the smallest ID number.

For example, assume your bird sightings are of types `arr = [1, 1, 2, 2, 3]`. There are two each of types `1` and `2`, and one sighting of type `3`. Pick the lower of the two types seen twice: type `1`. 

Read the full problem here : [Migratory Birds](https://www.hackerrank.com/challenges/migratory-birds/problem)

<h1>Solution :</h1>

To solve the problem first the input array, `types` storing types of birds, must be sorted. An array `type_count` stores count of bird of each type.

```cpp
std::sort(types.begin(), types.end());
std::vector<unsigned int> type_count(6, 0);
```

Since there are only 5 types of birds with ids 1, 2, 3, 4 and 5 so `type_count[0] = 0.`

{% include ads.html %}<br/>

We will start iteration on types from index 1 and check if id of bird at that index and id of bird at previous index is equal. If it is equal i.e. `types[i] == types[i-1]`, integer variable `count` is incremented. `count` counts the number of birds with that id(`type[-1]`).

If `types[i]` is not equal to `types[i-1]` then `type_count` for id `types[i-1]` is updated with the value stored in `count`.

```cpp
unsigned int count = 1;
for (unsigned int i = 1; i < types.size(); ++i)
{
    if (types[i] == types[i-1])
    {
        count++;
    }
    else if (types[i] != types[i-1])
    {
        type_count[ types[i-1] ] = count;
        count = 1;
    }
}
```

When the loop is executed and we are out of the `for` loop then value stored in `count` is assigned to `type_count` at index which is equal to largest id number. This index is accessed by `types.back()` because `types` is sorted and `types.back()` returns last element which is the largest id of the birds.

```cpp
type_count[ types.back() ] = count;
```

The program should return the id of birds which has the highest number of `count` and if two ids have same `count` then lowest id is returned. In other words, index of the highest element of `type_count` is returned and if it contains two or more occurrence of highest value then lower index of these highest value is returned.

{% include ads.html %}<br/>

Integer variable `max_val` stores the maximum number of birds with same type and `min_index` stores the minimum index of the `max_val` in `type_count`.

```cpp
unsigned int max_val = type_count[0];
unsigned short  min_index = 0;

for (unsigned short i = 1; i < type_count.size(); ++i)
{
    if (type_count[i] > max_val)
    {
        max_val = type_count[i];
        min_index = i;
    }
}
```

**Related:** [Kangaroo HackerRank Challenge]({{ site.url }}/blog/2019/11/28/Kangaroo-HackerRank-Challenge-C++-Implementation)

<h3>C++ Implementation</h3>

```cpp
#include <iostream>
#include <vector>
#include <algorithm>

unsigned short common_type(std::vector<unsigned short>& types)
{
    std::sort(types.begin(), types.end());
    std::vector<unsigned int> type_count(6, 0); //stores count of each bird whose id == index

    unsigned int count = 1;

    for (unsigned int i = 1; i < types.size(); ++i)
    {
        if (types[i] == types[i-1])
        {
            count++;
        }
        else if (types[i] != types[i-1])
        {
            type_count[ types[i-1] ] = count;
            count = 1;
        }
    }
    type_count[ types.back() ] = count;

    unsigned int max_val = type_count[0];
    unsigned short  min_index = 0;

    for (unsigned short i = 1; i < type_count.size(); ++i)
    {
        if (type_count[i] > max_val)
        {
            max_val = type_count[i];
            min_index = i;
        }
    }
    return min_index;
}

int main()
{
    unsigned int num_of_birds;
    std::cin >> num_of_birds;
    std::vector<unsigned short> types(num_of_birds);

    for (unsigned int i = 0; i < num_of_birds; ++i)
    {
        std::cin >> types[i];
    }

    std::cout << common_type(types) << "\n";
}
```

View this code on [Github](https://github.com/{{site.github_username}}/Competitive-Programming/blob/master/Hackerrank/Migratory_Birds.cpp)

<h3>You may also like</h3>
[Roy and Code Streak HackerEarth Challenge]({{ site.url }}/blog/2019/10/24/Roy-and-Code-Streak-HackerEarth-Challenge)



