---
layout: post
title: "How to find the Length of Loop in Linked List | C++ Implementation"
subtitle: "Given a Linked List, we have to find does loop exist in Linked List and if yes, find the length of loop."
author: "Programmercave"
header-img: "/assets/linkedlistwithloop.png"
tags:  [Cpp, Algorithm, Linked-List, Data-Structure]
date: 2018-01-20
---

Given a Linked List, we have to find does loop exist in Linked List and if yes, find the length of loop.

![Floyd Cycle Finding Algorithm]({{ site.url }}/assets/linkedlistwithloop.png){:class="img-responsive"}

To find loop in the linked list, we need two node pointers `slowPtr` and `fastPtr` which starts from the head. `slowPtr` increments by one node while `fastPtr` increments by two nodes. If these pointers point at the same node after starting from head then loop exists. This algorithm is known as *Floyd Cycle Finding Algorithm*.

```cpp
Node* doesLoopExist()
{
    Node *slowPtr = head;
    Node *fastPtr = head;

    while(slowPtr && fastPtr && fastPtr->next)
    {
        slowPtr = slowPtr->next;
        fastPtr = fastPtr->next->next;

        if(slowPtr == fastPtr)
        {
            return slowPtr;
        }
    }
    return nullptr;
}
```

If loop exists then this function returns node which is pointed by `slowPtr` and `fastPtr` and if node doesn’t exist then it returns `nullptr`.

To calculate length of loop, `fastPtr` start from its current node and count number of nodes until `slowPtr` is equal to `fastPtr`.

```cpp
int lengthOfLoop()
{
    int count = 0;
    bool loopExist = false;
    Node *slowPtr = nullptr, *fastPtr = nullptr;

    slowPtr = doesLoopExist();
    fastPtr = slowPtr;
    if(slowPtr != nullptr)
    {
        loopExist = true;
    }
    else
    {
        return 0;
    }
    if(loopExist)
    {
        fastPtr = fastPtr->next;
        count++;
        while(slowPtr != fastPtr)
        {
            fastPtr = fastPtr->next;
            count++;
        }
        return count;
    }
    return 0;
}
```

{% include ads.html %}<br/>

<h3>C++ Implementation to find Length of Loop in a Linked List</h3>

```cpp
#include <iostream>
#include <utility>

template <class T>
class LinkedList
{
    struct Node
    {
        T data;
        Node * next;
        Node(T value) : data(std::move(value)), next(nullptr) {}
    };
    Node *head;

  public:
    LinkedList() : head(nullptr) {}
    LinkedList(const LinkedList& ll) = delete; //copy constructor
    LinkedList& operator=(const LinkedList& ll) = delete; //copy assignment
    ~LinkedList();
    void insert(T);
    void createLoop(int);
    int lengthOfLoop();
    Node* doesLoopExist()
    {
        Node *slowPtr = head;
        Node *fastPtr = head;

        while(slowPtr && fastPtr && fastPtr->next)
        {
            slowPtr = slowPtr->next;
            fastPtr = fastPtr->next->next;

            if(slowPtr == fastPtr)
            {
                return slowPtr;
            }
        }
        return nullptr;
    }
};

template <class T>
void LinkedList<T>::insert(T data)
{
    Node *node = new Node(std::move(data));
    Node *tmp = head;
    if(tmp == nullptr)
    {
        head = node;
    }
    else
    {
        while(tmp->next != nullptr)
        {
            tmp = tmp->next;
        }
        tmp->next = node;
    }
}

template <class T>
void LinkedList<T>::createLoop(int n)
{
    Node *tmp = head;
    Node *tail = head;
    for(int i = 0; i < n-1; i++)
    {
        tmp = tmp->next;
    }

    while(tail->next != nullptr)
    {
        tail = tail->next;
    }
    tail->next = tmp;
}

template <class T>
int LinkedList<T>::lengthOfLoop()
{
    int count = 0;
    bool loopExist = false;
    Node *slowPtr = nullptr, *fastPtr = nullptr;

    slowPtr = doesLoopExist();
    fastPtr = slowPtr;
    if(slowPtr != nullptr)
    {
        loopExist = true;
    }
    else
    {
        return 0;
    }
    if(loopExist)
    {
        fastPtr = fastPtr->next;
        count++;
        while(slowPtr != fastPtr)
        {
            fastPtr = fastPtr->next;
            count++;
        }
        return count;
    }
    return 0;
}

template <class T>
LinkedList<T>::~LinkedList()
{
    Node *tmp = nullptr;
    while(head)
    {
        tmp = head;
        head = head->next;
        delete tmp;
    }
    head = nullptr;
}

int main()
{
    LinkedList<char> ll1;
    ll1.insert('p');
    ll1.insert('r');
    ll1.insert('o');
    ll1.insert('g');
    ll1.insert('r');
    ll1.insert('a');
    ll1.insert('m');
    ll1.insert('m');
    ll1.insert('e');
    ll1.insert('r');

    int nodeNumber = 5;
    //Node number starts from 1

    ll1.createLoop(nodeNumber);
    bool result = ll1.doesLoopExist();
    if(result == true)
    {
        std::cout <<"Loop exist in the Linked List\n";
    }
    else
    {
        std::cout <<"Loop does not exist in the Linked List\n";
    }

    int len = ll1.lengthOfLoop();
    std::cout << "Length of Loop is " << len <<"\n";

    exit(0);
}
```

View this code on [Github](https://github.com/{{site.github_username}}/Algo-Data-Structure/blob/master/Linked%20List%20containing%20Loop/C++/linkedListwithLoop.cpp)

`std::move` is used to indicate that an object t may be "moved from", i.e. allowing the efficient transfer of resources from t to another object. 
It is defined in header `<utility>`.

Get this post in pdf - [Find Length og Loop in a Linked List](https://www.file-up.org/it0eldmtml0l)

Reference:<br/>
[Introduction to Algorithms](https://amzn.to/2OarGBs)<br/>
[The Algorithm Design Manual](https://amzn.to/2CH9h9Z)<br/>
[Data Structures and Algorithms Made Easy](https://amzn.to/2NLM0dd)<br/>

<h3>You may also like:</h3>
[Move all Odd numbers after Even numbers in Singly Linked List]({{ site.url }}/blog/2018/02/08/C++-Move-all-Even-numbers-before-Odd-numbers-in-Singly-Linked-List-(Using-STL))<br/>
[Merge two sorted Linked List (in-place)]({{ site.url }}/blog/2018/02/06/C++-Merge-two-sorted-Linked-List-(in-place))<br/>
[Split Singly Circular Linked List]({{ site.url }}/blog/2018/02/04/C++-Split-Singly-Circular-Linked-List-program)<br/>
[Doubly Circular Linked List]({{ site.url }}/blog/2018/02/02/C++-Doubly-Circular-Linked-List-program)<br/>
[Reverse the Linked List]({{ site.url }}/blog/2018/01/23/C++-Reverse-the-Linked-List-(Iterative-Method)-program)<br/>
[Doubly Linked List]({{ site.url }}/blog/2017/07/28/C++-Doubly-Linked-List-using-Template-(Data-Structure))<br/>
[Singly Linked List]({{ site.url }}/blog/2017/07/27/C++-Singly-Linked-List-using-Template-(Data-Structure))









