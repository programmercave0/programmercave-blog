---
layout: post
title: "Rearranging an Array: Transforming arr[i] into arr[arr[i]]"
description: "Aspiring software developers and computer science engineering enthusiasts often encounter captivating challenges that put their problem-solving abilities to the test. In this beginner's guide to data structures and algorithms, we delve into a fascinating problem involving arrays. We will explore a step-by-step solution to rearrange the elements in the input array, leading to a transformation where each element at index `i` becomes the value at index `arr[i]`."
author: "Programmercave"
header-img: "/assets/rearrange-array/rearrange-array-1.png"
tags:  [Cpp, Competitive-Programming, Algorithm, Data-Structure, Python, Array-Manipulation]
date: 2023-07-22
---
* toc
{:toc}

# Introduction

Aspiring software developers and computer science engineering enthusiasts often encounter captivating challenges that put their problem-solving abilities to the test. In this beginner's guide to data structures and algorithms, we delve into a fascinating problem involving arrays. We will explore a step-by-step solution to rearrange the elements in the input array, leading to a transformation where each element at index `i` becomes the value at index `arr[i]`.

![Rearranging an Array: Transforming arr[i] into arr[arr[i]]]({{ site.url }}/assets/rearrange-array/rearrange-array-1.png){:class="img-responsive"}

# The Problem

Let's consider the problem statement:

Given an array containing numbers in the range from 0 to N-1, the task is to rearrange the array such that each number at the index `i` in the array becomes the number at the index `a[i]` in the same array.

# The Solution

## Step 1: Create a New Array and Multiply Every Element by N

We begin by traversing the input array and create a new array by multiplying each element by `N`. This multiplication ensures that all elements in the new array become greater than or equal to `N`. Since the original array's elements are in the range `[0, N-1]`, this step guarantees that no element will be less than `N`.

```python
def create_new_array(arr, N):
    new_array = [element * N for element in arr]
    return new_array
```

For example, let's take an initial array `[3, 2, 0, 1]` with N=4:

```plaintext
Initial array: [3, 2, 0, 1]
N = 4

After multiplying by N:
new_array = [12, 8, 0, 4]
```

## Step 2: Adding the Value of `a[new_array[i]/N]` to Each Element

If we take the mod of every element of the new\_array by `N`, it will give 0. We don't want that. We want `new_array[i]` at index `i` to be equal to `a[a[i]]`. So, we add `new_array[i]` with `a[new_array[i]/N]`. By doing so, when we take the mod by `N`, it will return `a[new_array[i]/N]`, and `a[new_array[i]/N]` is nothing but `a[a[i]]`.

Now, we iterate through each element of the `new_array`, and for each element at the index `i`, we need to find the value to be added such that the condition `new_array[i]` = `a[a[i]]` is satisfied.

```python
def rearrange_array(new_array, arr, N):
    for i in range(len(new_array)):
        new_index = new_array[i] // N
        new_array[i] += arr[new_index]
    return new_array
```

Let's walk through the process of adding the values for each element:

1. **At index 0 (**`i=0`): We take the value at `new_array[0]`, which is `12`. To find the value to be added, we calculate the `new_index` by dividing `new_array[0]` by N, resulting in `new_index = 12/4 = 3`. Now, `new_index` represents the value of `a[0]` in the original array. We need to find the value `a[new_index]` in the original array and add it to `new_array[0]`. In this case, `a[0]` is `3`, so `a[new_index]` is `a[3]`, which is `1`. Thus, we add `1` to `new_array[0]`, resulting in `new_array[0] += 1 = 12 + 1 = 13`.
    

```plaintext
new_array = [12, 8, 0, 4]

At index 0 (i=0):
new_array[0] = 12
new_index = 12/4 = 3
a[new_index] = a[3] = 1
new_array[0] += 1 = 12 + 1 = 13
Resulting array: [13, 8, 0, 4]
```

![Rearranging an Array: Transforming arr[i] into arr[arr[i]]]({{ site.url }}/assets/rearrange-array/arr-1.png){:class="img-responsive"}


1. **At index 1 (**`i=1`): We take the value at `new_array[1]`, which is `8`. We calculate the `new_index` by dividing `new_array[1]` by N, resulting in `new_index = 8/4 = 2`. Now, `new_index` represents the value of `a[1]` in the original array. We need to find the value `a[new_index]` in the original array and add it to `new_array[1]`. In this case, `a[1]` is `2`, so `a[new_index]` is `a[2]`, which is `0`. Thus, we add `0` to `new_array[1]`, resulting in `new_array[1] += 0 = 8`.
    

```plaintext
new_array = [13, 8, 0, 4]

At index 1 (i=1):
new_array[1] = 8
new_index = 8/4 = 2
a[new_index] = a[2] = 0
new_array[1] += 0 = 8
Resulting array: [13, 8, 0, 4]
```

![Rearranging an Array: Transforming arr[i] into arr[arr[i]]]({{ site.url }}/assets/rearrange-array/arr-2.png){:class="img-responsive"}

1. **At index 2 (**`i=2`): We take the value at `new_array[2]`, which is `0`. We calculate the `new_index` by dividing `new_array[2]` by N, resulting in `new_index = 0/4 = 0`. Now, `new_index` represents the value of `a[2]` in the original array. We need to find the value `a[new_index]` in the original array and add it to `new_array[2]`. In this case, `a[2]` is `0`, so `a[new_index]` is `a[0]`, which is `3`. Thus, we add `3` to `new_array[2]`, resulting in `new_array[2] += 3 = 0 + 3 = 3`.
    

```plaintext
new_array = [13, 8, 0, 4]

At index 2 (i=2):
new_array[2] = 0
new_index = 0/4 = 0
a[new_index] = a[0] = 3
new_array[2] += 3 = 0 + 3 = 3
Resulting array: [13, 8, 3, 4]
```

![Rearranging an Array: Transforming arr[i] into arr[arr[i]]]({{ site.url }}/assets/rearrange-array/arr-3.png){:class="img-responsive"}

1. **At index 3 (**`i=3`): We take the value at `new_array[3]`, which is `4`. We calculate the `new_index` by dividing `new_array[3]` by N, resulting in `new_index = 4/4 = 1`. Now, `new_index` represents the value of `a[3]` in the original array. We need to find the value `a[new_index]` in the original array and add it to `new_array[3]`. In this case, `a[3]` is `1`, so `a[new_index]` is `a[1]`, which is `2`. Thus, we add `2` to `new_array[3]`, resulting in `new_array[3] += 2 = 4 + 2 = 6`.
    

```plaintext
new_array = [13, 8, 3, 4]

At index 3 (i=3):
new_array[3] = 4
new_index = 4/4 = 1
a[new_index] = a[1] = 2
new_array[3] += 2 = 4 + 2 = 6
Resulting array: [13, 8, 3, 6]
```

![Rearranging an Array: Transforming arr[i] into arr[arr[i]]]({{ site.url }}/assets/rearrange-array/arr-4.png){:class="img-responsive"}

## Step 3: Take the Mod of Each Element by N

In this final step, we complete the rearrangement by taking the modulo of each element in the `new_array` by N (4). This step is necessary because we previously multiplied each element by N, and now we need to revert them to their original values.

Let's take the modulo of each element in `new_array` by N (4):

```python
def take_mod(arr, N):
    return [element % N for element in arr]
```

```plaintext
After taking the modulo by N:
new_array = [1, 0, 3, 2]
```

Now the array is rearranged as required, and each element `a[i]` is equal to `a[a[i]]`.

By mastering this array rearrangement challenge, you will sharpen your problem-solving skills and be better prepared for software developer interviews and similar technical assessments.

## Space and Time Complexity Analysis:

Let's analyze the space and time complexities of the first approach, where we created a new array.

### Using a New Array

**Space Complexity**: In this approach, we created a new array of the same size as the original array. Therefore, the space complexity is `O(N)`, where `N` is the size of the array. 

**Time Complexity**: The time complexity involves iterating over the array multiple times, which is done in linear time. Thus, the time complexity is `O(N)`, where `N` is the size of the array. Now, let's move to the second approach, where we avoid creating a new array by modifying the original array directly.

### Modifying the Original Array

**Space Complexity**: In this approach, we directly modify the original array without creating a new array. Hence, there is no additional space used, and the space complexity is `O(1)`, constant space complexity. 

**Time Complexity**: Similar to the first approach, time complexity involves iterating over the array multiple times, which is done in linear time. Therefore, the time complexity is `O(N)`, where `N` is the size of the array.

**For C++ Implementation using Modifying the Original Array:**

If you're interested in exploring the C++ implementation of this magical array rearrangement algorithm, you can find the complete code on our [Github repository](https://github.com/{{site.github_username}}/Algo-Data-Structure/blob/master/Maths/rearrange.cpp){:target="_blank"}.

# Conclusion

Congratulations on successfully navigating the array rearrangement challenge! Through this beginner's guide to data structures and algorithms, you've gained valuable insights into problem-solving techniques for arrays. The step-by-step solution presented here empowers you to transform an array in a way that each element `arr[i]` becomes `arr[arr[i]]`.

As you continue your journey in computer science engineering, keep exploring new problems and algorithms. Practicing such challenges will enhance your proficiency as a software developer and strengthen your ability to tackle a wide array of interview questions with confidence.

Remember that continuous learning and practice are the cornerstones of excellence in the world of algorithms and data structures. Happy coding!