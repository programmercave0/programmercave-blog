---
layout: post
title: "Sudoku Validator - CodinGame | C++ Implementation"
description: "A sudoku grid consists of 9*9 = 81 cells split in 9 sub-grids of 3*3 = 9 cells. For the grid to be correct, each row must contain one occurrence of each digit(1 to 9), each column must contain one occurrence of each digit (1 to 9) and each sub-grid must contain one occurrence of each digit (1 to 9)"
author: "Programmercave"
header-img: "/assets/images/2020-07-13-SUDOKU-VALIDATOR-CodinGame-C++-Implementation/sudoku1.jpg"
tags:  [Cpp, Competitive-Programming, CodinGame]
date: 2020-07-13
---



The problem is from [CodinGame](https://www.codingame.com/home) with difficulty level Easy.

<h1>Problem:</h1>

A sudoku grid consists of 9\*9 = 81 cells split in 9 sub-grids of 3\*3 = 9 cells.

For the grid to be correct, each row must contain one occurrence of each digit(1 to 9), each column must contain one occurrence of each digit (1 to 9) and each sub-grid must contain one occurrence of each digit (1 to 9).

You shall answer `true` if the grid is correct or `false` if it is not.

Read full problem here : [Sudoku Validator](https://www.codingame.com/training/easy/sudoku-validator)

<h1>Solution:</h1>

![SUDOKU VALIDATOR CodinGame]({{ site.url }}/assets/images/2020-07-13-SUDOKU-VALIDATOR-CodinGame-C++-Implementation/sudoku1.jpg){:class="img-responsive"}


To verify that each row contains only one occurrence of each digit (1 to 9) , we will make an array which will verify that the count of each digit is one. If the count of any digit is more than 1 than the function returns false.

```cpp
bool status = true;

for (int i = 0; i < 9; ++i)
{
    std::vector<bool> row_values(9, false);
     
    for (int r = 0; r < 9; ++r)
    {
        if (row_values[grid[r][i] - 1] == false)
        {
            row_values[grid[r][i] - 1] = true;
        }
        else
        {
            status = false;
            return status;
        }
    }
}
```

Here, array `row_values` makes sure that count of each digit is one.

Similarly, we will verify for each column.

```cpp
bool status = true;
    
for (int i = 0; i < 9; ++i)
{
    std::vector<bool> col_values(9, false);

    for (int c = 0; c < 9; ++c)
    {
        if (col_values[grid[i][c] - 1] == false)
        {
            col_values[grid[i][c] - 1] = true;
        }
        else
        {
            status = false;
            return status;
        }
    }
}
```

Here, array `col_values` makes sure that count of each digit is one.

![SUDOKU VALIDATOR CodinGame]({{ site.url }}/assets/images/2020-07-13-SUDOKU-VALIDATOR-CodinGame-C++-Implementation/sudoku2.jpg){:class="img-responsive"}

To verify diagonal subgrids, this code is used. It happens when `i == 0` or `i % 3 == 0`.

```cpp
bool status = true;
    
for (int i = 0; i < 9; ++i)
{
        
    std::vector<bool> subgrid_values(9, false);
   
    for (int r_sub = i; r_sub < i+3; ++r_sub)
    {
        for (int c_sub = i; c_sub < i+3; ++c_sub)
        {
            if (subgrid_values[grid[r_sub][c_sub] - 1] == false)
            {
                subgrid_values[grid[r_sub][c_sub] - 1] = true;
            }
            else
            {
                status = false;
                return status;
            }
        }
    }
}
```

![SUDOKU VALIDATOR CodinGame]({{ site.url }}/assets/images/2020-07-13-SUDOKU-VALIDATOR-CodinGame-C++-Implementation/sudoku4.jpg){:class="img-responsive"}

To verify these subgrids, this code is used. It happens when `i % 3 == 0`.

```cpp
bool status = true;
    
for (int i = 0; i < 9; ++i)
{
        
    std::vector<bool> subgrid_values(9, false);
   
    for (int r_sub = i; r_sub < i+3; ++r_sub)
    {
        for(int c_sub = 0; c_sub < 3; ++c_sub)
        {
            if (subgrid_values[grid[r_sub][c_sub] - 1] == false)
            {
                subgrid_values[grid[r_sub][c_sub] - 1] = true;
            }
            else
            {
                status = false;
                return status;
            }
        }
    }

    std::fill(subgrid_values.begin(), subgrid_values.end(), false);

    for (int r_sub = 0; r_sub < 3; ++r_sub)
    {
        for(int c_sub = i; c_sub < i+3; ++c_sub)
        {
            if (subgrid_values[grid[r_sub][c_sub] - 1] == false)
            {
                subgrid_values[grid[r_sub][c_sub] - 1] = true;
            }
            else
            {
                status = false;
                return status;
            }
        }
    }

    std::fill(subgrid_values.begin(), subgrid_values.end(), false);
}
```

![SUDOKU VALIDATOR CodinGame]({{ site.url }}/assets/images/2020-07-13-SUDOKU-VALIDATOR-CodinGame-C++-Implementation/sudoku5.jpg){:class="img-responsive"}

To verify these subgrids, this code is used. It happens when `i == 6`.

```cpp
bool status = true;
    
for (int i = 0; i < 9; ++i)
{
        
    std::vector<bool> subgrid_values(9, false);
   
    for (int r_sub = i; r_sub < i+3; ++r_sub)
    {
        for(int c_sub = 3; c_sub < 6; ++c_sub)
        {
            if (subgrid_values[grid[r_sub][c_sub] - 1] == false)
            {
                subgrid_values[grid[r_sub][c_sub] - 1] = true;
            }
            else
            {
                status = false;
                return status;
            }
        }
    }

    std::fill(subgrid_values.begin(), subgrid_values.end(), false);

    for (int r_sub = 3; r_sub < 6; ++r_sub)
    {
        for(int c_sub = i; c_sub < i+3; ++c_sub)
        {
            if (subgrid_values[grid[r_sub][c_sub] - 1] == false)
            {
                subgrid_values[grid[r_sub][c_sub] - 1] = true;
            }
            else
            {
                status = false;
                return status;
            }
        }
    }
}        
```

Since in one iteration more than one subgrids are checked. So we have to reinitialize array to its initial state using `fill(subgrid_values.begin(), subgrid_values.end(), false)`

{% include ads.html %}<br/>

<h3>C++ Implementation</h3>

```cpp
#include <iostream>
#include <vector>
#include <algorithm>

bool sudoku_validator(std::vector<std::vector<int>>& grid)
{
    bool status = true;
    
    for (int i = 0; i < 9; ++i)
    {
        std::vector<bool> row_values(9, false);
        std::vector<bool> col_values(9, false);
        std::vector<bool> subgrid_values(9, false);

        for (int r = 0; r < 9; ++r)
        {
            if (row_values[grid[r][i] - 1] == false)
            {
                row_values[grid[r][i] - 1] = true;
            }
            else
            {
                status = false;
                return status;
            }
        }

        for (int c = 0; c < 9; ++c)
        {
            if (col_values[grid[i][c] - 1] == false)
            {
                col_values[grid[i][c] - 1] = true;
            }
            else
            {
                status = false;
                return status;
            }
        }

        if (i == 0)
        {
            for (int r_sub = i; r_sub < i+3; ++r_sub)
            {
                for (int c_sub = i; c_sub < i+3; ++c_sub)
                {
                    if (subgrid_values[grid[r_sub][c_sub] - 1] == false)
                    {
                        subgrid_values[grid[r_sub][c_sub] - 1] = true;
                    }
                    else
                    {
                        status = false;
                        return status;
                    }
                }
            }
        }
        else if (i % 3 == 0)
        {
            for (int r_sub = i; r_sub < i+3; ++r_sub)
            {
                for(int c_sub = 0; c_sub < 3; ++c_sub)
                {
                    if (subgrid_values[grid[r_sub][c_sub] - 1] == false)
                    {
                        subgrid_values[grid[r_sub][c_sub] - 1] = true;
                    }
                    else
                    {
                        status = false;
                        return status;
                    }
                }
            }

            std::fill(subgrid_values.begin(), subgrid_values.end(), false);

            for (int r_sub = 0; r_sub < 3; ++r_sub)
            {
                for(int c_sub = i; c_sub < i+3; ++c_sub)
                {
                    if (subgrid_values[grid[r_sub][c_sub] - 1] == false)
                    {
                        subgrid_values[grid[r_sub][c_sub] - 1] = true;
                    }
                    else
                    {
                        status = false;
                        return status;
                    }
                }
            }

            std::fill(subgrid_values.begin(), subgrid_values.end(), false);

            if (i == 6)
            {
                for (int r_sub = i; r_sub < i+3; ++r_sub)
                {
                    for(int c_sub = 3; c_sub < 6; ++c_sub)
                    {
                        if (subgrid_values[grid[r_sub][c_sub] - 1] == false)
                        {
                            subgrid_values[grid[r_sub][c_sub] - 1] = true;
                        }
                        else
                        {
                            status = false;
                            return status;
                        }
                    }
                }

                std::fill(subgrid_values.begin(), subgrid_values.end(), false);

                for (int r_sub = 3; r_sub < 6; ++r_sub)
                {
                    for(int c_sub = i; c_sub < i+3; ++c_sub)
                    {
                        if (subgrid_values[grid[r_sub][c_sub] - 1] == false)
                        {
                            subgrid_values[grid[r_sub][c_sub] - 1] = true;
                        }
                        else
                        {
                            status = false;
                            return status;
                        }
                    }
                }
            }
            std::fill(subgrid_values.begin(), subgrid_values.end(), false);

            for (int r_sub = i; r_sub < i+3; ++r_sub)
            {
                for (int c_sub = i; c_sub < i+3; ++c_sub)
                {
                    if (subgrid_values[grid[r_sub][c_sub] - 1] == false)
                    {
                        subgrid_values[grid[r_sub][c_sub] - 1] = true;
                    }
                    else
                    {
                        status = false;
                        return status;
                    }
                }
            }
        }
    }
    return status;  
}

int main()
{
    std::vector< std::vector<int> > sudoku_grid;
    for (int i = 0; i < 9; i++) 
    {
        std::vector<int> row;
        for (int j = 0; j < 9; j++) 
        {
            int n;
            std::cin >> n; std::cin.ignore();
            row.push_back(n);
        }
        sudoku_grid.push_back(row);
    }

    bool result = sudoku_validator(sudoku_grid);

    if (result)
    {
        std::cout << "true\n";
    }
    else
    {
        std::cout << "false\n";
    }
    
}
```

View this on [Github](https://github.com/{{site.github_username}}/CodinGame-Solutions/blob/master/Sudoku_validator.cpp)

If you have another solution then please share it in the comments below.

<h3>Other Competitive Programming Problems and Solutions</h3>
[Repeated String]({{ site.url }}/blog/2020/04/24/Repeated-String-HackerRank-Challenge-Cpp-Implementation)<br/>
[Bank Robbers - CodinGame]({{ site.url }}/blog/2020/04/26/Bank-Robbers-CodinGame-Challenge-Cpp-Implementation)<br/>
