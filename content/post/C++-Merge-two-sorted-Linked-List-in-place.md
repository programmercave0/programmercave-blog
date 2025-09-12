---
author: Programmercave
date: "2018-02-06T00:00:00Z"
description: Given two sorted Linked List, we have to merge them without using another
  linked list.
header-img: /assets/images/mergesortedlinkedlist.png
tags:
- Cpp
- Algorithm
- Linked-List
- Data-Structure
title: Merge two sorted Linked List (in-place) | C++ Implementation
---



Given two sorted Linked List, we have to merge them without using another linked list.
  ```
  List 1 : { 5, 10, 18, 25 }
  List 2 : { 6, 8, 11, 20 }
  Merged List : { 5, 6, 8, 10, 11, 18, 20, 25 }
  ```
  
![Merge linked list](/assets/images/mergesortedlinkedlist.png)

From the above fig. we can see that merging two linked list is same as merging two sorted array in mergesort.

**Related:** [Merge Sort](/C-Implementation-of-Merge-Sort)

`node` stores the smallest element in both linked list and it will become the head of the merged linked list later.

`head1` and `head2` are the nodes whose data item is to be compared and node with smallest data item is added after tmp. `tmp` is always the last node in merged list.

Here is `mergeSortedList` functions.

```cpp
void mergeSortedList(LinkedList<T>& ll2)
{
    if(ll2.head == nullptr)
        return;

    if(head == nullptr)
    {
        head = ll2.head;
        ll2.head = nullptr;
        return;
    }

    Node *head1 = head;
    Node *head2 = ll2.head;

    Node *node;

    if(head1->data <= head2->data)
    {
        node = head1;
        advance(head1);
    }
    else if(head2->data <= head1->data)
    {
        node = head2;
        advance(head2);
    }

    Node *tmp = nullptr;
    tmp = node;

    while(head1 != nullptr && head2 != nullptr)
    {
        if(head1->data <= head2->data)
        {
            tmp->next = head1;
            advance(head1);
        }
        else
        {
            tmp->next = head2;
            advance(head2);
        }

        advance(tmp);
    }

    //if there are extra nodes in either list
    if(head1 != nullptr)
        tmp->next = head1;

    if(head2 != nullptr)
        tmp->next = head2;

    head = node;
        ll2.head = nullptr;
}
```

<br/>

<h3>C++ Implementation to Merge Sorted Linked List</h3>

```cpp
#include <iostream>
#include <utility>
#include <cassert>

template<class T>
class LinkedList
{
    struct Node
    {
        T data;
        Node * next = nullptr;
        Node()          : data(), next(nullptr) {}
        Node(T value)   : data(std::move(value)), next(nullptr) {}
    };

    Node *head;

  public:
    LinkedList() : head(nullptr) {}
    LinkedList(const LinkedList& ll) = delete; //copy constructor
    LinkedList& operator=(const LinkedList& ll) = delete; //copy assignment
    ~LinkedList()
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

    void insert(T);
    void mergeSortedList(LinkedList<T>&);
    void printList() const;

  private:
    static void advance(Node*& node)
    {
        assert(node != nullptr);
        node = node->next;
    }

    Node* getLastNode(Node*& node)
    {
        while(node->next != nullptr)
        node = node->next;

        return node;
    }
};

template <class T>
void LinkedList<T>::insert(T value)
{
    Node *node = new Node(std::move(value));
    Node *tmp = head;
    if(tmp == nullptr)
    {
        head = node;
    }
    else
    {
        tmp = getLastNode(tmp);
        tmp->next = node;
    }
}

template <class T>
void LinkedList<T>::mergeSortedList(LinkedList<T>& ll2)
{
    if(ll2.head == nullptr)
        return;

    if(head == nullptr)
    {
        head = ll2.head;
        ll2.head = nullptr;
        return;
    }

    Node *head1 = head;
    Node *head2 = ll2.head;

    Node *node;

    if(head1->data <= head2->data)
    {
        node = head1;
        advance(head1);
    }
    else if(head2->data <= head1->data)
    {
        node = head2;
        advance(head2);
    }

    Node *tmp = nullptr;
    tmp = node;

    while(head1 != nullptr && head2 != nullptr)
    {
        if(head1->data <= head2->data)
        {
            tmp->next = head1;
            advance(head1);
        }
        else
        {
            tmp->next = head2;
            advance(head2);
        }

        advance(tmp);
    }

    //if there are extra nodes in either list
    if(head1 != nullptr)
        tmp->next = head1;

    if(head2 != nullptr)
        tmp->next = head2;

    head = node;
        ll2.head = nullptr;
}

template <class T>
void LinkedList<T>::printList() const
{
    if(head == nullptr)
    {
        std::cout << "Empty List \n";
        return;
    }

    Node *node = head;

    while(node != nullptr)
    {
        std::cout << node->data << " ";
        advance(node);
    }

    std::cout << "\n";
}

int main()
{
    LinkedList<int> ll1;
    ll1.insert(5);
    ll1.insert(10);
    ll1.insert(18);
    ll1.insert(25);
    std::cout << "List 1 : ";
    ll1.printList();

    LinkedList<int> ll2;
    ll2.insert(6);
    ll2.insert(8);
    ll2.insert(11);
    ll2.insert(20);
    ll2.insert(23);
    ll2.insert(28);
    ll2.insert(40);
    std::cout << "List 2 : ";
    ll2.printList();

    ll1.mergeSortedList(ll2);
    std::cout << "Merged List : ";
    ll1.printList();
}
```

View this code on [Github](https://github.com/{{< param "github_username" >}}/Algo-Data-Structure/blob/master/Linked_List/mergesortedll.cpp)



Reference:<br/>
[Introduction to Algorithms](https://amzn.to/2OarGBs)<br/>
[The Algorithm Design Manual](https://amzn.to/2CH9h9Z)<br/>
[Data Structures and Algorithms Made Easy](https://amzn.to/2NLM0dd)<br/>

 <input type="hidden" name="IL_IN_ARTICLE"> 
<h3>You may also like</h3>
[Move all Odd numbers after Even numbers in Singly Linked List](/C-Move-all-Even-numbers-before-Odd-numbers-in-Singly-Linked-List-Using-STL)<br/>
[Split Singly Circular Linked List](/C-Split-Singly-Circular-Linked-List-program)<br/>
[Doubly Circular Linked List](/C-Doubly-Circular-Linked-List-program)<br/>
[Reverse the Linked List](/C-Reverse-the-Linked-List-Iterative-Method-program)
[Finding Length of Loop in Linked List](/C-Linked-List-containing-Loop-Floyd-Cycle-finding-Algorithm-program)<br/>
[Doubly Linked List](/C-Doubly-Linked-List-using-Template-Data-Structure)<br/>
[Singly Linked List](/C-Singly-Linked-List-using-Template-Data-Structure)




