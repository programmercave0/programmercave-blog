---
layout: post
title: "Singly Linked List | C++ Implementation"
description: "A node in a singly linked list contains a data item and a node pointer to the next node. In a singly linked list we can traverse only in one direction."
author: "Programmercave"
header-img: "/assets/singlylinkedlist.png"
tags:  [Cpp, Algorithm, Linked-List, Data-Structure]
date: 2017-07-27
---
* toc
{:toc}

A linked list is a linear data structure where each element, called a node, is connected to the next element through a pointer. In a singly linked list, each node contains a data item and a pointer to the next node in the list. The order of the list is determined by the pointers, and the first node is called the head while the last node is called the tail. If the head is NULL, then the list is empty. In C++, nodes can be created using the `struct` keyword, which allows you to define a data type that contains a data item and a pointer to the next node.

```cpp
struct Node
{
    T data;
    Node * next;
    Node(T val): data(val), next(nullptr){}
};
 ```   

The constructor `Node(T val): data(val), next(nullptr){}` is used to initialize the `data` and `next` members of the `struct Node`. The `T` parameter indicates that the `Node` struct is generic and can store values of any data type. To declare a head node in C++, you can use the syntax `Node *head;`. This creates a pointer to a `Node` struct and assigns it to the `head` variable. You can then use the `head` pointer to reference the first node in a linked list. 

![Singly Linked List]({{ site.url }}/assets/singlylinkedlist.png){:class="img-responsive"}

In the above fig. Node containing 5 is head, node containing 15 is tail and its next pointer points to nullptr.

<h1>Implementation</h1>
Linked lists are a common data structure that support several basic operations, including searching, insertion, and deletion.

1. Searching: Linked lists allow you to search for a specific element or node by traversing the list until you find the desired item. This can be done using a loop and a pointer to the current node.

2. Insertion: Linked lists also support the insertion of new elements or nodes. This can be done by creating a new node and adjusting the pointers of the surrounding nodes to include the new element in the list.

3. Deletion: Linked lists also allow you to delete elements or nodes. This can be done by adjusting the pointers of the surrounding nodes to remove the desired element from the list.

Overall, linked lists are useful for storing and organizing data, and the ability to perform these basic operations makes them a versatile data structure.

<h3>Searching</h3>

The `search` function is a useful tool for finding a specific element or node in a linked list. To use this function, you pass a value as an argument and the function will search through the list, starting at the head, to find a node with a matching `data` value. If the node is found, it is returned, otherwise a message is printed saying "No such element in the list" and a `nullptr` is returned.

Here is an example of code for an iterative search function in C++:

```cpp
struct Node *search(T n)
{                            //returns node of the given value
    Node *node = head;
    while(node != nullptr)
    {
        if(node->data == n)
            return node;
        node = node->next;
    }

    std::cerr << "No such element in the list \n";
    return nullptr;
}
```    
   
<h3>Insertion</h3>

The `insert` function is a useful tool for adding a new node to a linked list. In this function, a value is passed as an argument and a new node with that value is inserted at the end of the list. If the list is empty, the new node becomes the head of the list.

Here is an example of code for an insert function in C++:
```cpp
void insert(T data)
{
    Node *t = new Node(data);
    Node *tmp = head;
    if (tmp == nullptr)
    {
        head = t;
    }
    else
    {
        while (tmp->next != nullptr)
        {
            tmp = tmp->next;
        }
        tmp->next = t;
    }
}
```

In addition to inserting a new node at the end of a linked list, it is also possible to insert a node at the front of the list, making it the new head. This can be useful in certain situations, such as when you want to add new elements to the beginning of the list or when you want to maintain the order of the list.

Here is an example of code for an insert function in C++ that inserts a new node at the front of the list:

```
template <typename T>
void insertFront(Node<T>*& head, const T& val) {
  Node<T>* newNode = new Node<T>(val);
  newNode->next = head;
  head = newNode;
}
```

To use this function, you can call it with the head of your linked list and the value you want to insert as arguments. For example:

```
insertFront(head, 5);
```

This will insert a new node with a `data` value of 5 at the front of the list, making it the new head.

{% include ads.html %}<br/>

<h3>Deletion</h3>

The delete function is a useful tool for removing a specific node from a linked list. In this function, a value is passed as an argument and the function searches for a node with a matching `data` value using the search function. If the node is found, it is deleted from the list.

Here is an example of code for a delete function in C++:

```cpp
void deleteNode(T data)
{
    Node *node=search(data);
    Node *tmp = head;

    if(tmp == node)
    {
        head=tmp->next;
    }
    else if (node != nullptr)
    {
        while(node != nullptr)
        {
            if(tmp->next==node)
            {
                tmp->next=node->next;
                return ;
            }
            tmp=tmp->next;
        }
        delete tmp;
    }
}
```

<h3>C++ Implementation of Singly Linked List</h3>

```cpp
#include <iostream>

template <class T>
class LinkedList
{
    struct Node
    {
        T data;
        Node * next;
        Node(T val): data(val), next(nullptr){}
    };
    Node *head;

 public:
     LinkedList() : head(nullptr){}
     LinkedList(const LinkedList<T> & ll) = delete;
     LinkedList& operator=(LinkedList const&) = delete;
    ~LinkedList();
     void insert(T);
     void display(std::ostream& out = std::cout) const
     {
          Node *node = head;
          while(node != nullptr)
          {
              out << node->data << " ";
              node = node->next;
          }
      }
      
     void deleteNode(T);
     template <class U>
     friend std::ostream & operator<<(std::ostream & os, const LinkedList<U> & ll);

 private:
    struct Node *search(T n)
    {                            //returns node of the given value
        Node *node = head;
        while(node != nullptr)
        {
            if(node->data == n)
                 return node;
            node = node->next;
        }

        std::cerr << "No such element in the list \n";
        return nullptr;
    }

};

template <class U>
std::ostream & operator<<(std::ostream & os, const LinkedList<U>& ll)
{
    ll.display(os);
    return os;
}

template <class T>
void LinkedList<T>::insert(T data)
{
    Node *t = new Node(data);
    Node *tmp = head;
    if (tmp == nullptr)
    {
        head = t;
    }
    else
    {
        while (tmp->next != nullptr)
        {
            tmp = tmp->next;
        }
        tmp->next = t;
    }
}

template <class T>
void LinkedList<T>::deleteNode(T data)
{
    Node *node=search(data);
    Node *tmp = head;

    if(tmp == node)
    {
        head=tmp->next;
    }
    else if (node != nullptr)
    {
        while(node != nullptr)
        {
            if(tmp->next==node)
            {
                tmp->next=node->next;
                return ;
            }
            tmp=tmp->next;
        }
        delete tmp;
    }
}

template <class T>
LinkedList<T>::~LinkedList()
{
    Node *tmp = nullptr;
    while (head)
    {
        tmp = head;
        head = head->next;
        delete tmp;
    }
    head =nullptr;
}

int main()
{
    LinkedList<int> ll1;
    ll1.insert(5);
    ll1.insert(6);
    ll1.insert(7);
    ll1.insert(8);
    std::cout<<ll1<<std::endl;
    ll1.deleteNode(11);
    std::cout<<ll1<<std::endl;
    LinkedList<char> ll2;
    ll2.insert('a');
    ll2.insert('r');
    ll2.insert('d');
    ll2.insert('y');
    std::cout<<ll2<<std::endl;
    return 0;
}
```

View this code on [Github](https://github.com/{{site.github_username}}/Algo-Data-Structure/blob/master/Singly%20Linked%20List/C++/linkedlist.cpp)

In this code, the `LinkedList(const LinkedList<T> & ll) = delete;` and `LinkedList& operator=(LinkedList const&) = delete;` statements are used to delete the copy constructor and copy assignment operator for the `LinkedList` class. This means that the compiler will not generate default implementations for these functions and any attempt to use them will result in a compilation error.

The `void display(std::ostream& out = std::cout) const` function is a member function of the `LinkedList` class that allows you to display the contents of the list to an output stream, such as `cout`. The `friend std::ostream & operator<<(std::ostream & os, const LinkedList<U> & ll);` declaration is a non-member function that has been declared as a friend of the `LinkedList` class. This means that it can access the private and protected members of the `LinkedList` class. The function overloads the `operator<<` operator, allowing you to print a `LinkedList` object using the `cout` stream.

Reference:<br/>
[Introduction to Algorithms](https://amzn.to/2OarGBs)<br/>
[The Algorithm Design Manual](https://amzn.to/2CH9h9Z)<br/>
[Data Structures and Algorithms Made Easy](https://amzn.to/2NLM0dd)<br/>

<h3>You may also like</h3>
[Move all Odd numbers after Even numbers in Singly Linked List]({{ site.url }}/blog/2018/02/08/C++-Move-all-Even-numbers-before-Odd-numbers-in-Singly-Linked-List-(Using-STL))<br/>
[Merge two sorted Linked List (in-place)]({{ site.url }}/blog/2018/02/06/C++-Merge-two-sorted-Linked-List-(in-place))<br/>
[Split Singly Circular Linked List program]({{ site.url }}/blog/2018/02/04/C++-Split-Singly-Circular-Linked-List-program)<br/>
[Doubly Circular Linked List program]({{ site.url }}/blog/2018/02/02/C++-Doubly-Circular-Linked-List-program)<br/>
[Reverse the Linked List]({{ site.url }}/blog/2018/01/23/C++-Reverse-the-Linked-List-(Iterative-Method)-program)<br/>
[Finding Length of Loop in Linked List]({{ site.url }}/blog/2018/01/20/C++-Linked-List-containing-Loop-(Floyd-Cycle-finding-Algorithm)-program)<br/>
[Doubly Linked List using Template]({{ site.url }}/blog/2017/07/28/C++-Doubly-Linked-List-using-Template-(Data-Structure))




