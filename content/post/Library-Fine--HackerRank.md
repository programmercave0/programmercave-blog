---
author: Programmercave
date: "2019-12-20T00:00:00Z"
description: 'Your local library needs your help! Given the expected and actual return
  dates for a library book, create a program that calculates the fine (if any). The
  fee structure is as follows: • If the book is returned on or before the expected
  return date, no fine will be charged (i.e.: fine = 0). • If the book is returned
  after the expected return day but still within the same calendar month and year
  as the expected return date, fine = 15 Hackos * (the number of days late). • If
  the book is returned after the expected return month but still within the same calendar
  year as the expected return date, the fine = 500 Hackos * (the number of months
  late). • If the book is returned after the calendar year in which it was expected,
  there is a fixed fine of 10000 Hackos.'
header-img: /assets/images/2019-12-20-Library-Fine/Hackerrank_poster.jpg
tags:
- Cpp
- Algorithm
- Data-Structure
- Competitive-Programming
- Hackerrank
title: Library Fine | HackerRank
---



This is an easy hackerrank challenge which will help you to become good at competitive programming. There are various competitive programming websites like [CodeChef](https://www.codechef.com/), [HackerEarth](https://www.hackerearth.com/challenges/), [Codeforces](https://codeforces.com/) where you can practice coding.

<h1>Problem : </h1>

Your local library needs your help! Given the expected and actual return dates for a library book, create a program that calculates the fine (if any). The fee structure is as follows: 

 • If the book is returned on or before the expected return date, no fine will be charged (i.e.: *fine = 0*).<br/>
 • If the book is returned after the expected return day but still within the same calendar month and year as the expected return date, *fine = 15 Hackos * (the number of days late)*.<br/>
 • If the book is returned after the expected return month but still within the same calendar year as the expected return date, the *fine = 500 Hackos * (the number of months late)*. <br/>
 • If the book is returned after the calendar year in which it was expected, there is a fixed fine of *10000 Hackos*.

Charges are based only on the least precise measure of lateness. For example, whether a book is due January 1, 2017 or December 31, 2017, if it is returned January 1, 2018, that is a year late and the fine would be *10,000 Hackos*.
 
Read full problem : [Library Fine](https://www.hackerrank.com/challenges/library-fine/problem)

![Library Fine HackerRank](/assets/images/2019-12-20-Library-Fine/Hackerrank_poster.jpg)

<h2>Solution : </h2>

Let integer variables `return_day, return_month, return_year` represents returned date and `due_day, due_month and due_year` represents expected return date of a book.

If `return_year < due_year`, then no fine will be charged.

If `return_year == due_year`, then we have to check for `return_month` and `return_day` since book is returned within the same calendar year. 

If book is returned within the same calendar year but after expected return month (`due_month`) i. e. `(return_year == due_year && return_month > due_month)` then fine is `500 Hackos * (return_month – due_month)`.

If the book is returned within the same calendar year and within the same calendar month but after the expected return day i. e. `(return_year == due_year && return_month == due_month && return_day > due_day)` then fine is `15 Hackos * (return_day – due_day)`.

If the book is returned after the expected return year i. e. `(return_year > due_year)` then fine is `10000 Hackos`.

```cpp
int library_fine(int return_day, int return_month, int return_year, int due_day, int due_month, int due_year) 
{
    int fine = 0;
    if (return_year == due_year)
    {
        if (return_month == due_month)
        {
            if (return_day > due_day)
            {
                fine = 15 * (return_day - due_day);
            }
        }
        else if (return_month > due_month)
        {
            fine = 500 * (return_month - due_month);
        }
    }
    else if (return_year > due_year)
    {
        fine = 10000;
    }
    return fine;
}
```

<br/>

<h2>C++ Implementation</h2>

```cpp
#include <iostream>

int library_fine(int return_day, int return_month, int return_year, int due_day, int due_month, int due_year) 
{
    int fine = 0;
    if (return_year == due_year)
    {
        if (return_month == due_month)
        {
            if (return_day > due_day)
            {
                fine = 15 * (return_day - due_day);
            }
        }
        else if (return_month > due_month)
        {
            fine = 500 * (return_month - due_month);
        }
    }
    else if (return_year > due_year)
    {
        fine = 10000;
    }
    return fine;
}

int main()
{
    int return_day, return_month, return_year;
    int due_day, due_month, due_year;

    std::cin >> return_day >> return_month >> return_year;
    std::cin >> due_day >> due_month >> due_year;

    std::cout << library_fine(return_day, return_month, return_year, due_day, due_month, due_year) << "\n";
}    
```

View this code on [Github](https://github.com/{{< param "github_username" >}}/Competitive-Programming/blob/master/Hackerrank/Library_Fine.cpp).

If you have another solution then please share it in the comments below.

<h3>Other Competitive Programming Problems and Solutions</h3>
[Sherlock and Squares](/Sherlock-and-Squares-HackerRank)<br/>
[Circular Array Rotation](/Circular-Array-Rotation-HackerRank)<br/>


