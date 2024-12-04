---
author: Programmercave
date: "2018-02-04T00:00:00Z"
description: Given a Singly Circular Linked List, we have to split it into two equal
  halves. If the number of nodes in the given list is odd then first list will have
  one node more than the second list.
header-img: /assets/images/splitsinglycircularlinkedlist.png
tags:
- Cpp
- Algorithm
- Linked-List
- Data-Structure
title: Split Singly Circular Linked List | C++ Implementation
---



Given a Singly Circular Linked List, we have to split it into two equal halves. If the number of nodes in the given list is odd then first list will have one node more than the second list.

```
Input : { 2, 3, 18, 25, 5 }
Output List 1 : { 2, 3, 18 }
Output List 2 : { 25, 5 }
```

Here is a meme to understand Circular Linked List

![Split Linked List](/assets/images/cll.jpg)

![Split Linked List](/assets/images/splitsinglycircularlinkedlist.png)

From the above fig. first we have to find middle node of the given singly circular linked list.

Middle node can be found using `slow_ptr` and `fast_ptr`. `slow_ptr` increments by one node whereas `fast_ptr` increments by two nodes. When `fast_ptr->next == head` or `fast_pr->next->next == head` then no increment of `fast_ptr` or `slow_ptr` and `slow_ptr` is at the middle node of the list.

```cpp
Node *find_middle_node()
{
    //slow pointer advances by 1
    Node *slow_ptr = head;
    //fast pointer advances by 2
    Node *fast_ptr = head;

    //If odd number of nodes then fast_ptr->next = head
    //If even number of nodes then fast_ptr->next->next = head
    while(fast_ptr->next != head && fast_ptr->next->next != head)
    {
        fast_ptr = fast_ptr->next->next;
        slow_ptr = slow_ptr->next;
    }
    return slow_ptr;
}
```
<br/>

After getting middle node, the node next to middle node is the head of second ouput list. Then tail of the lists are updated so that circular list is maintained.

```cpp
void split_list(circular_linked_list<T>& second_half)
{
    Node *tail = find_tail_node();
    Node *mid = find_middle_node();

    second_half.head = mid->next;
    tail->next = second_half.head;

    mid->next = head;
}
```

<h3>C++ Implementation to Split Circular Linked List</h3>

```cpp
#include <iostream>
#include <utility>

template <class T>
class circular_linked_list
{
    struct Node
    {
        T data;
        Node * next;
        Node(T&& value) : data(std::move(value)), next(nullptr) {}
    };
    Node *head;

  public:
    circular_linked_list() : head(nullptr) {}
    //copy constructor
    circular_linked_list(const circular_linked_list& cll) = delete;
    //move constructor
    circular_linked_list(circular_linked_list&& cll) = delete;
    //copy assignment
    circular_linked_list& operator=(const circular_linked_list& cll) = delete;
    //move assignment
    circular_linked_list& operator=(circular_linked_list&& cll) = delete;
    ~circular_linked_list();

    void insert_node(T&&);
    void split_list(circular_linked_list<T>&);
    void print_list();

  private:

    Node *find_tail_node()
    {
        Node *tmp = head;
        while(tmp->next != head)
        {
            tmp = tmp->next;
        }
        return tmp;
    }

    Node *find_middle_node()
    {
        //slow pointer advances by 1
        Node *slow_ptr = head;
        //fast pointer advances by 2
        Node *fast_ptr = head;

        //If odd number of nodes then fast_ptr->next = head
        //If even number of nodes then fast_ptr->next->next = head
        while(fast_ptr->next != head && fast_ptr->next->next != head)
        {
            fast_ptr = fast_ptr->next->next;
            slow_ptr = slow_ptr->next;
        }
        return slow_ptr;
    }
};

template <class T>
void circular_linked_list<T>::insert_node(T&& value)
{
    Node *node = new Node(std::move(value));
    if(head == nullptr)
    {
        node->next = node;
        head = node;
        return;
    }
    Node *tmp = head;
    while(tmp->next != head)
    {
        tmp = tmp->next;
    }
    tmp->next = node;
    node->next = head;
}

template <class T>
void circular_linked_list<T>::split_list(circular_linked_list<T>& second_half)
{
    Node *tail = find_tail_node();
    Node *mid = find_middle_node();

    second_half.head = mid->next;
    tail->next = second_half.head;

    mid->next = head;
}

template <class T>
void circular_linked_list<T>::print_list()
{
    Node *tmp = head;
    while(tmp->next != head)
    {
        std::cout << tmp->data << ' ';
        tmp = tmp->next;
    }
    std::cout << tmp->data << '\n';
}

template <class T>
circular_linked_list<T>::~circular_linked_list()
{
    Node *tmp = nullptr;
    Node *tail = head;
    while(tail->next != head)
    {
        tail = tail->next;
    }
    tail->next = nullptr;

    while(head != nullptr)
    {
        tmp = head;
        head = head->next;
        delete tmp;
    }
}

int main()
{
    circular_linked_list<int> original_list;
    original_list.insert_node(2);
    original_list.insert_node(3);
    original_list.insert_node(4);
    original_list.insert_node(1);
    original_list.insert_node(0);

    std::cout << "The original list is : ";
    original_list.print_list();

    circular_linked_list<int> second_half;
    original_list.split_list(second_half);

    //Original list becomes first half
    std::cout << "The first half : ";
    original_list.print_list();
    std::cout << "The second half : ";
    second_half.print_list();
}
```

View this code on [Github](https://github.com/{{< param "github_username" >}}/Algo-Data-Structure/blob/master/Linked_List/split_circular_ll.cpp)



Reference:<br/>
[Introduction to Algorithms](https://amzn.to/2OarGBs)<br/>
[The Algorithm Design Manual](https://amzn.to/2CH9h9Z)<br/>
[Data Structures and Algorithms Made Easy](https://amzn.to/2NLM0dd)<br/>

 <input type="hidden" name="IL_IN_ARTICLE"> 
<h3>You may also like</h3>
[Move all Odd numbers after Even numbers in Singly Linked List](/C-Move-all-Even-numbers-before-Odd-numbers-in-Singly-Linked-List-Using-STL)<br/>
[Merge two sorted Linked List (in-place)](/C-Merge-two-sorted-Linked-List-in-place)<br/>
[Doubly Circular Linked List](/C-Doubly-Circular-Linked-List-program)<br/>
[Reverse the Linked List](/C-Reverse-the-Linked-List-Iterative-Method-program)<br/>
[Finding Length of Loop in Linked List](/C-Linked-List-containing-Loop-Floyd-Cycle-finding-Algorithm-program)<br/>
[Doubly Linked List](/C-Doubly-Linked-List-using-Template-Data-Structure)<br/>
[Singly Linked List](/C-Singly-Linked-List-using-Template-Data-Structure)






