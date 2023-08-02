---
layout: post
title: "Stock Exchange Losses - CodinGame | C++ Implementation"
subtitle: "A finance company is carrying out a study on the worst stock investments and would like to acquire a program to do so. The program must be able to analyze a chronological series of stock values in order to show the largest loss that it is possible to make by buying a share at a given time t0 and by selling it at a later date t1. The loss will be expressed as the difference in value between t0 and t1. If there is no loss, the loss will be worth 0."
author: "Programmercave"
header-img: "/assets/2021-03-17-Stock-Exchange-Losses-CodinGame/Stock-Exchange-Losses-img1.png"
tags:  [Cpp, Competitive-Programming, CodinGame]
date: 2021-03-17
---

The problem is from [CodinGame](https://www.codingame.com/home) with difficulty level Medium.

<h1>Problem:</h1>

A finance company is carrying out a study on the worst stock investments and would like to acquire a program to do so. The program must be able to analyze a chronological series of stock values in order to show the largest loss that it is possible to make by buying a share at a given time `t0` and by selling it at a later date `t1`. The loss will be expressed as the difference in value between `t0` and `t1`. If there is no loss, the loss will be worth `0`.

<h3>Input</h3>
Line 1: the number `n` of stock values available.

Line 2: the stock values arranged in order, from the date of their introduction on the stock market V<sub>1</sub> until the last known value V<sub>n</sub>. The values are integers.

<h3>Output</h3>

The maximal loss `p`, expressed negatively if there is a loss, otherwise `0`.

Read full problem here : [Stock Exchange Losses](https://www.codingame.com/ide/puzzle/stock-exchange-losses)

<h1>Solution:</h1>

The input values are the stock prices at various times. We store these stock prices in an integer vector called `values`. The index of the vector represents the time and the element at that index is the stock price at that time.

The three given test cases are:

```
Input
6
3 2 4 2 1 5

Output
-3
```

![Stock Exchange Losses]({{ site.url }}/assets/2021-03-17-Stock-Exchange-Losses-CodinGame/Stock-Exchange-Losses-img1.png){:class="img-responsive"}

```
Input
6
5 3 4 2 3 1

Output
-4
```

![Stock Exchange Losses]({{ site.url }}/assets/2021-03-17-Stock-Exchange-Losses-CodinGame/Stock-Exchange-Losses-img2.png){:class="img-responsive"}

```
Input
5
1 2 4 4 5

Output
0
```

![Stock Exchange Losses]({{ site.url }}/assets/2021-03-17-Stock-Exchange-Losses-CodinGame/Stock-Exchange-Losses-img3.png){:class="img-responsive"}

From the test cases, we can see that there is a loss only when the stock price at time `t2` is less than the stock price at time `t1` (where `t2` > `t1`). Otherwise, the loss is 0.

Therefore, we only need to process further when `values[i]` > `values[j]`, where `i` and `j` are iterable variables in a loop. We initialize `i` to `0` and `j` to `i+1`.

Before we proceed, we should initialize all the variables that we need.

```
int loss = std::numeric_limits<int>::max();
int maxima = std::numeric_limits<int>::min();
int minima = std::numeric_limits<int>::max();
int maxima_index = 0, minima_index = 0;
```

While `i` = 0, `j` will iterate from `i+1` to `values.size() - 1`.

And, when condition `values[i] > values[j]` is false, then the value of `i` becomes the last value of `j`.
 
In fig 1, `i = 0` and at `j = 2` (i.e V<sub>3</sub>), the condition `values[i] > values[j]` is false. So the new value of `i` is 2 and value of `j` is 3 (i.e `i+1`).
 
In fig 1, from V<sub>1</sub> to V<sub>2</sub>, the condition `values[i] > values[j]` is true. So this is one section of graph where we will calculate loss. This loss will be stored in an integer variable `local_loss`. There will be many sections in graph where we will calculate the `local_loss` and update variable `loss` with lowest values (i.e. highest loss). The values will be in negative so lowest value means highest loss.
 
Similarly, integer variable `maxima` and `minima` will store the maximum and minimum values in the graph, but the condition is that `maxima_index` < `minima_index`.

As the program proceeds, the value of these variables change.

The value of variables `maxima` and `minima` will change only when `values[i]` >= `maxima`. This means we have found the new maximum in the graph, so `maxima` is updated. And `minima` is updated because its index must be greater than the index of `maxima`.

So, we can only update value of `minima`, `maxima`, `minima_index` and `maxima_index` if conditions `values[i] > values[j]` and `local_loss < loss` and `values[i] >= maxima` are `true`.
 
In fig 1, from V<sub>1</sub>(`i = 0`) to V<sub>2</sub>(`j = 1`), `maxima` is 3 and `minima` is 2 and `loss` is -1. From V<sub>3</sub>(`i = 2`) to V<sub>4</sub>(`j = 3`), the `loss` becomes -2.The `values[i]`(=4) is also greater than `maxima`(=3). Thus all conditons `values[i] > values[j]` and `local_loss < loss` and `values[i] >= maxima` are true. So `maxima` and `minima` is updated.

At V<sub>3</sub>(`i = 2`) and V<sub>5</sub>(`j = 4`) again all conditions are true and new value of `loss` is calculated.
 
```
int i = 0;
int j = i+1;

while (j < values.size())
{
    if (values[i] > values[j])
    {
        int local_loss = values[j] - values[i];

        if (local_loss < loss)
        {
            loss = local_loss;
            if (values[i] >= maxima)
            {
                maxima = values[i];
                maxima_index = i;
                minima = values[j];
                minima_index = j;
            }
        }
    }
    else if (values[i] < values[j])
    {
        i = j;
        j = i;
    }
    j++;
}
```
 
For the third test case, the condition `values[i] > values[j]` is never true. As a result, the `loss` variable is never updated and retains its initial value, which is the highest value an integer can store.

Since `minima` and `maxima` are also never updated, their indices remain `0`. Therefore, the loss will be `0` when we calculate it using the formula `loss = values[minima_index] - values[maxima_index]`.
 
{% include ads.html %}<br/>

<h3>C++ Implementation</h3>

```cpp
#include <iostream>
#include <vector>
#include <limits>

int largest_loss(std::vector<int>& values)
{
    int loss = std::numeric_limits<int>::max();

    int maxima = std::numeric_limits<int>::min();
    int minima = std::numeric_limits<int>::max();
    int maxima_index = 0, minima_index = 0;

    int i = 0;
 
    int j = i+1;

    while (j < values.size())
    {
        if (values[i] > values[j])
        {
            int local_loss = values[j] - values[i];

            if (local_loss < loss)
            {
                loss = local_loss;
                if (values[i] >= maxima)
                {
                    maxima = values[i];
                    maxima_index = i;
                    minima = values[j];
                    minima_index = j;
                }
            }

        }
        else if (values[i] < values[j])
        {
            i = j;
            j = i;
        }
        j++;
    }
    loss = values[minima_index] - values[maxima_index];
    return loss;
}

int main()
{
    int n;
    std::cin >> n; std::cin.ignore();
    std::vector<int> values(n);
    for (int i = 0; i < n; i++) {
        std::cin >> values[i]; 
        std::cin.ignore();
    }

    int loss = largest_loss(values);

    std::cout << loss << std::endl;
}
```

Check out this on [Github](https://github.com/{{site.github_username}}/CodinGame-Solutions/blob/master/Stock_Exchange_Losses.cpp)

If you have a different solution for Stock Exchange Losses, please share it in the comments below.

<h3>Other Competitive Programming Problems and Solutions</h3>
[Dungeons and Maps - CodinGame]({{ site.url }}/blog/2021/03/01/Dungeons-and-Maps-CodinGame-C++-Implementation)<br/>
[Sudoku Validator - CodinGame ]({{ site.url }}/blog/2020/07/13/SUDOKU-VALIDATOR-CodinGame-C++-Implementation)<br/>

