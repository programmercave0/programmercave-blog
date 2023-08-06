---
layout: post
title: "Roy and Code Streak | HackerEarth Challenge"
description: "Roy is working on HackerEarth Profile. Right now he is working on User Statistics.
One of the statistics data (Code Streak) is as follows:"
author: "Programmercave"
header-img: "/assets/HERoy.png"
tags:  [Data-Structure, Algorithm, Competitive-Programming, Hackerearth, Cpp]
date: 2019-10-24
---
* toc
{:toc}

This is the HackerEarth challenge of easy level. Problem Link[ Roy and Code Streak](https://www.hackerearth.com/problem/algorithm/roy-and-code-streak/).

# Problem:

Roy is working on HackerEarth Profile. Right now he is working on User Statistics.
One of the statistics data (Code Streak) is as follows:

Given the User Activity Data, find the maximum number of continuous correct solutions submitted by any user.<br/>
Seems easy eh? Here's the catch! In order to maximize this number a user could have submitted a correct answer to the same problem which he has already solved. Now such cases cannot be tolerated. (See test case for clarification). 
Also in all the continuous correct submissions multiple correct submissions to same problem should be counted only once.

User Activity Data will be as follows:<br/>
Total number of submissions by the user - N<br/>
For each submission its Submission ID - S and Result - R (Solved or Unsolved) will be given.<br/>
Solved is represented by integer 1 and Unsolved by 0. <br/>

**Input:**

First line contains integer **T** - number of test cases. T test cases follow. First line of each test case will contain **N**. Next N lines each will contain two space separated integers **S** and **R**.

**Output:**

For each test case output in a single line the required answer.

**Constraints:**

1 <= T <= 1000<br/>
1 <= N <= 1000000 (10<sup>6</sup>)<br/>
1 <= S <= 1000000 (10<sup>6</sup>)<br/>
0 <= R <= 1

Note: Sum of N over all the test cases in each file does not exceed 1000000 (10<sup>6</sup>) 

![Sample]({{ site.url }}/assets/HERoy.png){:class="img-responsive"}

# Solution:

For fast I/O in C++, use `std::ios::sync_with_stdio(false)` and `std::cin.tie(nullptr)` at beginning in the main function.

`std::ios::sync_with_stdio(false)` disables the synchronization between the C and C++ standard streams. By default all standard streams are synchronized, which in practice allow mixing C and C++ I/O streams and get expected results. Refer [std::ios_base::sync_with_stdio](https://en.cppreference.com/w/cpp/io/ios_base/sync_with_stdio)

`std::cin.tie(nullptr)` unties cin from cout. By default they are tied and flushing of `std::cout` (visible on console) occurs before `std::cin` accepts an input. After untie we can get to input before getting the output.

We have used unordered___set to store submission id. `std::unordered_set<unsigned int>` submissions; We will insert a value in submissions only if the result for that submission id is 1(i.e. successfully submitted) and count of that submission id is 0 in the submissions(unordered_set) means that submission id is not submitted before.

If encounter result is 0 then we have to reset count because we want a streak of successful submissions.

**Member functions used:**

`clear()` : It erases all elements from the unordered_set. After this call, size() returns zero. <br/>
`insert()` : It insert value in unordered_set, if it does not contain thar value.<br/>
`count()` : It returns the number of occurence of that element in the unordered_set, which is 0 or 1.<br/>
`max()` : It is defined in header <algorithm> and it returns the greater of the two values.<br/>
    
Refer [std::max](https://en.cppreference.com/w/cpp/algorithm/max) and [std::unordered_set](https://en.cppreference.com/w/cpp/container/unordered_set)

{% include ads.html %}
**C++ implementation**

```cpp
#include <iostream>
#include <algorithm>
#include <unordered_set>

int main() 
{
    std::ios::sync_with_stdio(false);
    std::cin.tie(nullptr);

    unsigned int turns, no_of_submissions, s_id, result;
    std::unordered_set<unsigned int> submissions;

    std::cin >> turns;
    while (turns--) 
    {
        unsigned int count = 0, max_count = 0;

        submissions.clear();

        std::cin >> no_of_submissions;

        for (unsigned int i = 0; i < no_of_submissions; ++i) 
        {
            std::cin >> s_id >> result;

            if (result == 1 && submissions.count(s_id) == 0)
            {
                count++;
                submissions.insert(s_id);
            }
            else if (result == 0)
            {
                count = 0;
            }
            max_count = std::max(count, max_count);
        }
        std::cout << max_count << "\n";
    }
}
```

View this code on [Github](https://github.com/{{site.github_username}}/Competitive-Programming/blob/master/Hackerearth/Roy_and_Code_Streak.cpp)

Get this post in [pdf form](https://www.file-up.org/bqcqt9dxgdgp).

Reference:<br/>
[Why does the following snippet speed up the code?](https://stackoverflow.com/questions/48367983/why-does-the-following-snippet-speed-up-the-code)<br/>
[Significance of ios_base::sync_with_stdio(false); cin.tie(NULL);](https://stackoverflow.com/questions/31162367/significance-of-ios-basesync-with-stdiofalse-cin-tienull)<br/>
[Fast I/O for Competitive Programming](https://www.geeksforgeeks.org/fast-io-for-competitive-programming/)<br/>
