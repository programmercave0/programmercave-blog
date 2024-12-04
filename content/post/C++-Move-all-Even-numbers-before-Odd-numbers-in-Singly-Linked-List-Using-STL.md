---
author: Programmercave
date: "2018-02-08T00:00:00Z"
description: Given a Singly Linked List, we have to modify it such that all Even numbers
  appear before Odd numbers.
header-img: /assets/images/moveevenbeforeodd.png
tags:
- Cpp
- Algorithm
- Linked-List
- Data-Structures
title: Move all Odd numbers after Even numbers in Singly Linked List | C++ Implementation
---



Given a Singly Linked List, we have to modify it such that all Even numbers appear before Odd numbers.

For eg. 
```
Given Linked List : 1, 2, 3, 4, 5, 6, 7
After Modification : 2, 4, 6, 1, 3, 5, 7
```

![Move Odd After Even](/assets/images/moveevenbeforeodd.png)

From the above fig. we can see that how the function will work.

While the head of the list is odd, the head is copied to an auxiliary node and element next to the head will become new head. The auxiliary node is added after tail and tail is updated. Similarly all odd value nodes are removed from their position and added after the tail of the list.

Here is the `advance`, `getLastNode`, `isOdd` and `exchangeEvenOdd` function.

```cpp
static void advance(Node*& node)
{
    assert (node != nullptr);
    node = node->next;
}
```

Using this function node advances to next node.

```cpp
Node* getLastNode()
{
    Node *node = head;
    while (node->next != nullptr)
            node = node->next;

    return node;
}
```

This function returns the last node of the linked list.

```cpp
bool isOdd(int num)
{
    if (num % 2 != 0)
        return true;
    else
        return false;
}
```

This function checks whether the entered value is odd or even.

<br/>

```cpp
void exchangeEvenOdd()
{
    Node *node = nullptr;
    Node *lastNodeToTest = getLastNode();
    Node *tail = lastNodeToTest;

    while (isOdd(head->data) == true)
    {
        node = head;
        advance(head);
        tail->next = node;
        advance(tail);
    }

    Node *tmp = head;
    Node *curr = head;

    while (tmp->next != lastNodeToTest)
    {
        if (isOdd(curr->next->data) == true)
        {
            node = curr->next;
            curr->next = node->next;
            tail->next = node;
            advance(tail);
        }
        else
        {
            //advance "curr" and "tmp" only when next node to it is even
            advance(curr);
            advance(tmp);
        }
    }

    if (isOdd(curr->next->data) == true && tmp->next == lastNodeToTest)
    {
        node = lastNodeToTest;
        curr->next = lastNodeToTest->next;
        tail->next = lastNodeToTest;
        advance(tail);
    }
    tail->next = nullptr;
    lastNodeToTest = nullptr;
    node = nullptr;
}
```


<h3>C++ Implementation</h3>

```cpp
#include <iostream>
#include <utility>
#include <cassert>

class LinkedList
{
    struct Node
    {
        int data;
        Node * next = nullptr;
        Node(int value)   : data(std::move(value)), next(nullptr) {}
    };
    Node *head;

  public:
    LinkedList() : head(nullptr) {}
    ~LinkedList()
    {
        Node *tmp = nullptr;
        while (head)
        {
            tmp = head;
            head = head->next;
            delete tmp;
        }
        head = nullptr;
    }

    void insert(int);
    void exchangeEvenOdd();
    void printList() const;

  private:
    static void advance(Node*& node)
    {
        assert (node != nullptr);
        node = node->next;
    }

    Node* getLastNode()
    {
        Node *node = head;
        while (node->next != nullptr)
              node = node->next;

        return node;
    }

    bool isOdd(int num)
    {
        if (num % 2 != 0)
            return true;
        else
            return false;
    }
};

void LinkedList::insert(int value)
{
    Node *node = new Node(std::move(value));
    Node *tmp = head;
    if (tmp == nullptr)
    {
        head = node;
    }
    else
    {
        tmp = getLastNode();
        tmp->next = node;
    }
}

void LinkedList::exchangeEvenOdd()
{
    Node *node = nullptr;
    Node *lastNodeToTest = getLastNode();
    Node *tail = lastNodeToTest;

    while (isOdd(head->data) == true)
    {
        node = head;
        advance(head);
        tail->next = node;
        advance(tail);
    }

    Node *tmp = head;
    Node *curr = head;

    while (tmp->next != lastNodeToTest)
    {
        if (isOdd(curr->next->data) == true)
        {
            node = curr->next;
            curr->next = node->next;
            tail->next = node;
            advance(tail);
        }
        else
        {
            //advance "curr" and "tmp" only when next node to it is even
            advance(curr);
            advance(tmp);
        }
    }

    if (isOdd(curr->next->data) == true && tmp->next == lastNodeToTest)
    {
        node = lastNodeToTest;
        curr->next = lastNodeToTest->next;
        tail->next = lastNodeToTest;
        advance(tail);
    }
    tail->next = nullptr;
    lastNodeToTest = nullptr;
    node = nullptr;
}

void LinkedList::printList() const
{
    if (head == nullptr)
    {
        std::cout << "Empty List \n";
        return;
    }

    Node *node = head;

    while (node != nullptr)
    {
        std::cout << node->data << " ";
        advance(node);
    }

    std::cout << "\n";
}

int main()
{
    LinkedList ll1;
    ll1.insert(1);
    ll1.insert(2);
    ll1.insert(3);
    ll1.insert(4);
    ll1.insert(5);
    ll1.insert(6);
    ll1.insert(7);
    std::cout << "Original List : ";
    ll1.printList();

    ll1.exchangeEvenOdd();
    std::cout << "New List : ";
    ll1.printList();
}
```

View this code on [Github](https://github.com/{{< param "github_username" >}}/Algo-Data-Structure/blob/master/Linked_List/moveevenbeforodd.cpp)

Reference:<br/>
[Introduction to Algorithms](https://amzn.to/2OarGBs)<br/>
[The Algorithm Design Manual](https://amzn.to/2CH9h9Z)<br/>
[Data Structures and Algorithms Made Easy](https://amzn.to/2NLM0dd)<br/>


 <input type="hidden" name="IL_IN_ARTICLE"> 
<h3>You may also like</h3>
[Merge two sorted Linked List (in-place)](/C-Merge-two-sorted-Linked-List-in-place)<br/>
[Split Singly Circular Linked List](/C-Split-Singly-Circular-Linked-List-program)<br/>
[Doubly Circular Linked List](/C-Doubly-Circular-Linked-List-program)<br/>
[Reverse the Linked List ](/C-Reverse-the-Linked-List-Iterative-Method-program)<br/>
[Finding Length of Loop in Linked List](/C-Linked-List-containing-Loop-Floyd-Cycle-finding-Algorithm-program)<br/>
[Doubly Linked List](/C-Doubly-Linked-List-using-Template-Data-Structure)<br/>
[Singly Linked List](/C-Singly-Linked-List-using-Template-Data-Structure)









