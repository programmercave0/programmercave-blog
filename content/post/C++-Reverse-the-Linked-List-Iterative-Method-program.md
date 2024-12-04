---
author: Programmercave
date: "2018-01-23T00:00:00Z"
description: There are two ways to reverse a linked list, iterative method and recursive
  method.
header-img: /assets/images/reverselinkedlist.png
tags:
- Cpp
- Algorithm
- Linked-List
- Data-Structure
title: How to Reverse a Linked List | C++ Implementation
---



Given a singly linked list, we have to reverse it.

```
Input list: { a, b, c, d, e }
Output list: { e, d, c, b, a }
```
![Range](/assets/images/reverselinkedlist.png)

There are two ways to reverse a linked list, *iterative method* and *recursive method*.

<h3>Iterative Method</h3>

```cpp
Node* iterativeReverse(Node* head)
{
    Node *previous = nullptr;
    Node *nextNode = nullptr;
    while(head)
    {
        nextNode = head->next;
        head->next = previous;
        previous = head;
        head = nextNode;
    }
    return previous;
}
```

<br/>

<h3>Recursive Method</h3>

```cpp
Node* recursiveReverse(Node* head)
{
    if(head == nullptr)
        return nullptr;

    if(head->next == nullptr)
        return head;

    Node *firstElement = head;
    Node *secondElement = firstElement->next;
    head = firstElement->next;
    firstElement->next = nullptr; //unlink first node
    Node *remainingList = recursiveReverse(head);
    secondElement->next = firstElement;
    return remainingList;
}
```
<h3>C++ Implementation to reverse a Linked List</h3>

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
    LinkedList(const LinkedList&& ll) = delete; //move constructor
    LinkedList& operator=(const LinkedList& ll) = delete; //copy assignment
    LinkedList& operator=(const LinkedList&& ll) = delete; //move assignment
    ~LinkedList();
    void insert(T);
    void printList();
    void iterativeReverse()
    {
        head = iterativeReverse(head);
    }
  private:
    Node* iterativeReverse(Node* head)
    {
        Node *previous = nullptr;
        Node *nextNode = nullptr;
        while(head)
        {
            nextNode = head->next;
            head->next = previous;
            previous = head;
            head = nextNode;
        }
        return previous;
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
void LinkedList<T>::printList()
{
    Node *node = head;
    while(node)
    {
        std::cout << node->data << " ";
        node = node->next;
    }
    std::cout<<"\n";
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
    ll1.printList();
    ll1.iterativeReverse();
    ll1.printList();
}
```

Reference:<br/>
[Introduction to Algorithms](https://amzn.to/2OarGBs)<br/>
[The Algorithm Design Manual](https://amzn.to/2CH9h9Z)<br/>
[Data Structures and Algorithms Made Easy](https://amzn.to/2NLM0dd)<br/>

You may also like:
[Move all Odd numbers after Even numbers in Singly Linked List](/C-Move-all-Even-numbers-before-Odd-numbers-in-Singly-Linked-List-Using-STL)<br/>
[Merge two sorted Linked List (in-place)](/C-Merge-two-sorted-Linked-List-in-place)<br/>
[Split Singly Circular Linked List](/C-Split-Singly-Circular-Linked-List-program)<br/>
[Doubly Circular Linked List](/C-Doubly-Circular-Linked-List-program)<br/>
[Reverse the Linked List](/C-Reverse-the-Linked-List-Recursive-Method-program)<br/>
[Finding Length of Loop in Linked List](/C-Linked-List-containing-Loop-Floyd-Cycle-finding-Algorithm-program)<br/>
[Doubly Linked List](/C-Doubly-Linked-List-using-Template-Data-Structure)<br/>
[Singly Linked List](/C-Singly-Linked-List-using-Template-Data-Structure)<br/>




