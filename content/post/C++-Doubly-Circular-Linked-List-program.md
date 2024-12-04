---
author: Programmercave
date: "2018-02-02T00:00:00Z"
description: A node in a doubly circular linked list contains a data item and two
  node pointers, one to the previous node and one to the next node. In doubly linked
  list we can traverse in both direction.
header-img: /assets/images/circulardoublylinkedlist.png
tags:
- Cpp
- Algorithm
- Linked-List
- Data-Structure
title: Doubly Circular Linked List | C++ Implementation
---

The nodes in a linked list are connected through *pointers*. Pointers represent the address of a location in a memory. The order in a linked list is determined by a pointer in each node. A `node` in a **doubly circular linked list** contains a data item and two node pointers, one to the previous node and one to the next node. In doubly linked list we can traverse in both direction.

**Related:** [Doubly Linked List](/C-Doubly-Linked-List-using-Template-Data-Structure)

Here is a meme to understand Circular Linked List.

![Split Linked List](/assets/images/cll.jpg)

The first node of the linked list is the head and the last node is the tail. If head is NULL then the list is empty.

In C++, a node can be defined using `struct`, which contain a data item and node pointers.

```cpp
struct Node
{
    T data;
    Node * next;
    Node * prev;
    Node(T value) : data(std::move(value)),
                    next(nullptr),
                    prev(nullptr)
                    {}
};
``` 

`Node(T val): data(val), next(nullptr), prev(nullptr) {}` is the constructor for the `struct Node` which is used to initialise `data`, `next` and `prev`.  `T` means it is a generic struct and data can store values of all data types.
    
To declare head and tail: `Node *head, *tail;`


![Doubly Circular Linked List](/assets/images/circulardoublylinkedlist.png)

In the above fig. Node containing 5 is head and node containing 15 is tail. `prev` pointer in head points to the last node and `next` pointer in tail points to the head.

<h1>Implementation</h1>

The three basic operations supported by a linked list are searching, insertion and deletion.

<h3>Searching</h3>

In the `search` function a value is passed as an argument and its node is returned if found, otherwise a message says “No such element in the list” and `nullptr` is returned.

The function starts searching from the head to the last node and passed value is matched with every node’s data item.

Here is the code for iterative search.

```cpp
struct Node *search(T value)
{
    Node *node = head;
    while(node->next != head)
    {
        if(node->data == value)
        {
            return node;
         }
         node = node->next;
     }
     if(node->data == value)
     {
        return node;
     }
     std::cerr << "No such element in the list \n";
     return nullptr;
}
```

<h3>Insertion</h3>

`insert` function insert a node with the value at the end of the linked list. If the linked list does not contain any node then the new node becomes head and tail otherwise new node is added after tail and its `prev` pointer points to tail and `next` pointer points to head. Then new node becomes tail.

```cpp
void insert(T value)
{
    Node *node = new Node(value);

    if(head == nullptr)
    {
        node->next = node;
        node->prev = node;
        head = node;
        tail = node;
    }

    tail = head->prev;
    tail->next = node;
    node->prev = tail;
    node->next = head;
    head->prev = node;
    tail = node;
}
```

<br/>

<h3>Deletion</h3>

In `deleteNode` function the value is entered which is to be deleted. The function search the node containing the value using `search` function and then the node is deleted. 

If the searched node is `head` then node next to head is made head and then the searched node is deleted. The node is deleted only if the value exists means `if (node != nullptr)`.

<br/>

After deletion next and prev pointer of previous and next nodes are updated.

```cpp
void deleteNode(T value)
{
    Node *node = search(value);
    if(node == nullptr)
    {
        std::cerr << "No such value in the list\n";
        return;
    }
    else
    {
        Node *tmp = head;
        Node *tail = head->prev;
        if(tmp == node)
        {
            tail->next = tmp->next;
            tmp->prev->next->prev = tail;
            head = tail->prev;
            delete tmp;
            return;
        }
        else if(tail == node)
        {
            Node *curr = tail;
            tmp = tail->prev;
            tmp->next = curr->next;
            head->prev = tmp;
            tail = tmp;
            delete curr;
            return;
        }
        else
        {
            while(tmp->next != head)
            {
                if(tmp == node)
                {
                    tmp->prev->next = tmp->next;
                    tmp->prev->next->prev = tmp->prev;
                    delete tmp;
                    return;
                }
                tmp = tmp->next;
            }
        }
    }
}
```

<h3>C++ Implementation of Circular Doubly Linked List</h3>

In C++ implementation of this program we have copy constructor, move constructor, copy assignment operator, move assignment operator and destructor.

Because the presence of a user-defined destructor, copy-constructor, or copy-assignment operator prevents implicit definition of the move constructor and the move assignment operator, any class for which move semantics are desirable, has to declare all five special member functions. This is called [Rule of Five](https://en.cppreference.com/w/cpp/language/rule_of_three).

```cpp
#include <iostream>
#include <utility>

template <typename T>
class circular_doubly_linked_list
{
    struct Node
    {
        T data;
        Node * next;
        Node * prev;
        Node(T value) : data(std::move(value)),
                        next(nullptr),
                        prev(nullptr)
                        {}
    };
    Node *head, *tail;

  public:
    circular_doubly_linked_list() : head(nullptr),
                                    tail(nullptr)
                                    {}
    //copy constructor
    circular_doubly_linked_list(const circular_doubly_linked_list &);
    //copy assignment
    circular_doubly_linked_list& operator=(const circular_doubly_linked_list& cdll)
    {
        circular_doubly_linked_list  temp(cdll);
        temp.swap(*this);
        return *this;
    }
    //move constructor
    circular_doubly_linked_list(circular_doubly_linked_list&&) noexcept;
    //move assignment
    circular_doubly_linked_list& operator=(circular_doubly_linked_list&& cdll) noexcept
    {
        cdll.swap(*this);
        return *this;
    }
    ~circular_doubly_linked_list();

    void append_node(T);
    void delete_node(T);

    friend void swap(circular_doubly_linked_list& lhs, circular_doubly_linked_list& rhs)
    {
        std::swap(lhs.head, rhs.head);
    }

    template <typename U>
    friend std::ostream & operator<<(std::ostream & os, const circular_doubly_linked_list<U> & cdll)
    {
        cdll.print_list(os);
        return os;
    }

  private:

    struct Node *search(T value)
    {
        Node *node = head;
        while(node->next != head)
        {
            if(node->data == value)
            {
                return node;
            }
            node = node->next;
        }
        if(node->data == value)
        {
            return node;
        }
        std::cerr << "No such element in the list \n";
        return nullptr;
    }

  void print_list(std::ostream& os = std::cout) const
  {
      Node *tmp = head;
      while(tmp->next != head)
      {
          std::cout << tmp->data << ' ';
          tmp = tmp->next;
      }
      std::cout << tmp->data << '\n';
  }
};

template <typename T>
circular_doubly_linked_list<T>::circular_doubly_linked_list(const circular_doubly_linked_list & cdll)
{
    if(cdll.head == nullptr)
    {
        head = tail = nullptr;
    }
    else
    {
        head = new Node(cdll.head->data);
        Node *curr = head;
        Node *tmp = head;
        Node *obj_curr = cdll.head;

        while(obj_curr->next != cdll.head)
        {
            curr->next = new Node(obj_curr->next->data);
            obj_curr = obj_curr->next;
            curr = curr->next;
            curr->prev = tmp;
            tmp = tmp->next;
        }
        tail = curr;
        curr->next = head;
        head->prev = curr;
    }
}

template <typename T>
circular_doubly_linked_list<T>::circular_doubly_linked_list(circular_doubly_linked_list&& cdll) noexcept
{
    head = tail = nullptr;
    swap(*this, cdll);
}

template <typename T>
void circular_doubly_linked_list<T>::append_node(T value)
{
    Node *node = new Node(std::move(value));

    if(head == nullptr)
    {
        node->next = node;
        node->prev = node;
        head = node;
        tail = node;
    }

    tail = head->prev;
    tail->next = node;
    node->prev = tail;
    node->next = head;
    head->prev = node;
    tail = node;
}

template <typename T>
void circular_doubly_linked_list<T>::delete_node(T value)
{
    Node *node = search(value);
    if(node == nullptr)
    {
        std::cerr << "No such value in the list\n";
        return;
    }
    else
    {
        Node *tmp = head;
        Node *tail = head->prev;
        if(tmp == node)
        {
            tail->next = tmp->next;
            tmp->prev->next->prev = tail;
            head = tail->prev;
            delete tmp;
            return;
        }
        else if(tail == node)
        {
            Node *curr = tail;
            tmp = tail->prev;
            tmp->next = curr->next;
            head->prev = tmp;
            tail = tmp;
            delete curr;
            return;
        }
        else
        {
            while(tmp->next != head)
            {
                if(tmp == node)
                {
                    tmp->prev->next = tmp->next;
                    tmp->prev->next->prev = tmp->prev;
                    delete tmp;
                    return;
                }
                tmp = tmp->next;
            }
        }
    }
}

template <typename T>
circular_doubly_linked_list<T>::~circular_doubly_linked_list()
{
    if(head)
    {
        Node *tmp = head;
        while(tmp->next != head)
        {
            Node *t = tmp;
            tmp = tmp->next;
            delete t;
        }
        delete tmp;
        head = nullptr;
    }
}

int main()
{
    circular_doubly_linked_list<int> cdll1;
    cdll1.append_node(3);
    cdll1.append_node(4);
    cdll1.append_node(5);
    cdll1.append_node(6);
    cdll1.append_node(7);
    cdll1.append_node(8);
    std::cout << cdll1;
    cdll1.delete_node(6);
    std::cout << cdll1;
    circular_doubly_linked_list<int> cdll2(cdll1); // using copy constructor
    std::cout << "Linked List 2: " << cdll2;
    circular_doubly_linked_list<int> cdll3 = cdll1; //using copy assignment
    std::cout << "Linked List 3: " << cdll3;
    circular_doubly_linked_list<int> cdll4 = std::move(cdll2); //using move constructor
    std::cout << "Linked list 4: " << cdll4;
}
```

View this code on [Github](https://github.com/{{< param "github_username" >}}/Algo-Data-Structure/blob/master/Linked_List/circulardll.cpp).



<input type="hidden" name="IL_IN_ARTICLE"> 

Reference:<br/>
[Introduction to Algorithms](https://amzn.to/2OarGBs)<br/>
[The Algorithm Design Manual](https://amzn.to/2CH9h9Z)<br/>
[Data Structures and Algorithms Made Easy](https://amzn.to/2NLM0dd)<br/>

<h3>may also like</h3>
[Move all Odd numbers after Even numbers in Singly Linked List](/C-Move-all-Even-numbers-before-Odd-numbers-in-Singly-Linked-List-Using-STL)<br/>
[Merge two sorted Linked List (in-place)](/C-Merge-two-sorted-Linked-List-in-place)<br/>
[Split Singly Circular Linked List](/C-Split-Singly-Circular-Linked-List-program)<br/>
[Reverse the Linked List](/C-Reverse-the-Linked-List-Iterative-Method-program)<br/>
[Finding Length of Loop in Linked List](/C-Linked-List-containing-Loop-Floyd-Cycle-finding-Algorithm-program)<br/>
[Singly Linked List](/C-Singly-Linked-List-using-Template-Data-Structure)
