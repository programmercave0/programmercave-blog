---
layout: post
title: "Repeated String - HackerRank | C++ Implementation"
description: "Lilah has a string, s, of lowercase English letters that she repeated infinitely many times. Given an integer, n, find and print the number of letter a's in the first n letters of Lilah's infinite string.For example, if the string s = ‘abcac’ and n = 10, the substring we consider is , abcacabcac the first 10 characters of her infinite string. There are 4 occurrences of a in the substring. "
author: "Programmercave"
header-img: "/assets/images/2020-04-24-Repeated-String-HackerRank/HR_repeated_strings-2.jpg"
tags:  [Cpp, Algorithm, Competitive-Programming, Hackerrank, Data-Structure]
date: 2020-04-24
---



<h1>Problem:</h1>

Lilah has a string, `s`, of lowercase English letters that she repeated infinitely many times.

Given an integer, `n`, find and print the number of letter *a*'s in the first `n` letters of Lilah's infinite string.

For example, if the string `s = ‘abcac’` and `n = 10`, the substring we consider is , `abcacabcac` the first 10 characters of her infinite string. There are 4 occurrences of *a* in the substring. 

Read the full problem here: [Repeated String](https://www.hackerrank.com/challenges/repeated-string/copy-from/124938661?isFullScreen=false)

<h1>Solution:</h1>

Let us consider that `sub_str` is the input string and `str_len` is the length of the infinite string we are considering.

When the `str_len` is less than the length of `sub_str`, we can easily calculate the number of *a*’s by iterating over the string.

![Repeated String HackerRank]({{ site.url }}/assets/images/2020-04-24-Repeated-String-HackerRank/HR_repeated_strings-1.jpg){:class="img-responsive"}

```cpp
long i = 0;
while (i < str_len)
{
    if (sub_str[i] == 'a')
    {
        count++;
    }
    i++;
}
```

Now, when the `str_len` is greater than the length of the `sub_str` length.

![Repeated String HackerRank]({{ site.url }}/assets/images/2020-04-24-Repeated-String-HackerRank/HR_repeated_strings-2.jpg){:class="img-responsive"}

First, we will find the number of sub strings (`sub_str`) in the string we are considering.

```cpp
long num_sub_str = str_len / sub_str.size();
```

In many cases, the size of combined sub strings is not the multiple of `str_len`, so few characters are left out and these characters may contain ‘*a*’. We have to count these characters.

```cpp
long remaining_str_len = str_len % sub_str.size();
```

Now, we need to find the number of *a*’s in the `sub_str`. Let `count` will store number of *a*’s in the `sub_str`.

```cpp
long i = 0;
while (i < sub_str.size())
{
    if (sub_str[i] == 'a')
    {
        count++;
    }
    i++;
}
```

Multiplying `count` with the `num_sub_str` will give us number of *a*’s in the combined sub strings whose size is less than or equal to `str_len` and multiple of the size of `sub_str`. But there are few remaining characters whose count is stored in `remaining_str_len`. We have to find the occurrence of ‘*a8’ in these characters and increment the count if it is found.

```cpp
count = count * num_sub_str;

i = 0;
while (i < remaining_str_len)
{
    if (sub_str[i] == 'a')
    {
        count++;
    }
    i++;
}
```

{% include ads.html %}<br/>

<h3>C++ Implementation</h3>

```cpp
#include <iostream>
#include <string>

long repeated_string(std::string sub_str, long str_len) {
    long count = 0;
    if (str_len > sub_str.size())
    {
        long num_sub_str = str_len / sub_str.size();
        long remaining_str_len = str_len % sub_str.size();

        long i = 0;
        while (i < sub_str.size())
        {
            if (sub_str[i] == 'a')
            {
                count++;
            }
            i++;
        }

        count = count * num_sub_str;

        i = 0;
        while (i < remaining_str_len)
        {
            if (sub_str[i] == 'a')
            {
                count++;
            }
            i++;
        }
    }
    else
    {
        long i = 0;
        while (i < str_len)
        {
            if (sub_str[i] == 'a')
            {
                count++;
            }
            i++;
        }
    }
    return count;
}

int main()
{
    std::string input_str;
    long final_str_length;

    std::cin >> input_str;
    std::cin >> final_str_length;

    std::cout << repeated_string(input_str, final_str_length) << "\n";

    return 0;
}

```

View this on [Github](https://github.com/{{site.github_username}}/Competitive-Programming/blob/master/Hackerrank/Repeated_String.cpp)

If you have another solution then please share it in the comments below.

<h3>Other Competitive Programming Problems and Solutions</h3>
[Picking Numbers]({{ site.url }}/blog/2020/03/28/Picking-Numbers-HackerRank-Challenge-Cpp-Implementation)<br/>
[Library Fine]({{ site.url }}/blog/2019/12/20/Library-Fine-HackerRank)<br/>




