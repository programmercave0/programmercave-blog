---
author: Programmercave
date: "2018-03-05T00:00:00Z"
description: 'Depth first search explores on a single path in a graph as long as it
  find undiscovered vertices. When an edge leads to the discovered vertices it backtrack
  to the previous vertex and explores along the edge where it find undiscovered vertices.
  Finally it backtracks to the source vertex from where it started. '
header-img: /assets/images/dfs.png
tags:
- Cpp
- Algorithm
- Graph-Algorithms
title: Depth First Search using Adjacency List | Graph traversal
---



**Depth first search** explores on a single path in a graph as long as it find undiscovered vertices. When an edge leads to the discovered vertices it backtrack to the previous vertex and explores along the edge where it find undiscovered vertices. Finally it backtracks to the source vertex from where it started. 

Initially all the vertices are white and when a vertex is discovered it becomes gray and then black when it is finished or processed.

![Depth First Search](/assets/images/dfs.png)
![Depth First Search](/assets/images/dfs_.png)

First source vertex 1 is discovered so it becomes gray. Then vertex 4 is discovered. Since there are no vertex after vertex 4 on that path so vertex 4 is processed or finished and it becomes black. Then vertex 2 is discovered and so on. When vertex 3 is processed, control backtracks to the previous vertex which is 2.

<h1>Implementation</h1>

To store the color of the current vertex, we need a data structure.

Here is the implementation of the function `depth_first_search` and `depth_first_search_visit`.

```cpp
enum Color {WHITE, GRAY, BLACK};
   struct Vertex
   {
      int id;
      Color color;
      Vertex(int _id) : id(_id),
                        color(Color::WHITE)
                        {}
   };

void depth_first_search()
{
   for (const auto& u: vertices)
   {
      if (vertices[u.id].color == WHITE)
      {
         depth_first_search_visit(u.id);
      }
   }
}

void depth_first_search_visit(int u)
{
   vertices[u].color = GRAY;
   std::cout << vertices[u].id <<" ";
   const auto& v = adjacent[u].head;
      if (vertices[v->dest_id].color == WHITE)
      {
         depth_first_search_visit(v->dest_id);
      }
   vertices[u].color = BLACK;
}
```
Variable `color` in struct `Vertex` stores color for the given vertex.

The time complexity of Depth First Search is *O(n+m)* where *n* is the number of vertices and *m* is the number of edges.

<br/>

Here is the C++ Implementation for Depth First Search using Adjacency List

```cpp
#include <iostream>
#include <vector>

class Graph
{
   int vertex_count;
   enum Color {WHITE, GRAY, BLACK};

   struct Vertex
   {
      int id;
      Color color;
      Vertex(int _id) : id(_id),
                        color(Color::WHITE)
                        {}
   };

   struct Adj_list_node
   {
      int dest_id;
      Adj_list_node * next;
      Adj_list_node(int _dest_id) : dest_id(_dest_id),
                                    next (nullptr)
                                    {}
   };

   struct Adj_list
   {
      Adj_list_node *head;
   };

   std::vector<Vertex> vertices;
   std::vector<Adj_list> adjacent;

public:
   Graph(int);
   void add_edge(int, int);
   void depth_first_search();

 private:
   void depth_first_search_visit(int);
};

Graph::Graph(int v)
{
   vertex_count = v;
   adjacent.resize(vertex_count);

   for (int i = 0; i < vertex_count; i++)
   {
       vertices.push_back( Vertex(i) );
       adjacent[i].head = nullptr;
   }
}

void Graph::depth_first_search()
{
   for (const auto& u: vertices)
   {
      if (vertices[u.id].color == WHITE)
      {
         depth_first_search_visit(u.id);
      }
   }
}

void Graph::depth_first_search_visit(int u)
{
   vertices[u].color = GRAY;
   std::cout << vertices[u].id <<" ";
   const auto& v = adjacent[u].head;
      if (vertices[v->dest_id].color == WHITE)
      {
         depth_first_search_visit(v->dest_id);
      }
   vertices[u].color = BLACK;
}

void Graph::add_edge(int src, int dest)
{
   Adj_list_node *node = new Adj_list_node(dest);
   if (adjacent[src].head == nullptr)
   {
      adjacent[src].head = node;
   }
   else
   {
      Adj_list_node *tmp = adjacent[src].head;
      while (tmp->next != nullptr)
             tmp = tmp->next;

      tmp->next = node;
   }

   //undirected graph
   node = new Adj_list_node(src);
   if (adjacent[dest].head == nullptr)
   {
      adjacent[dest].head = node;
   }
   else
   {
     Adj_list_node *tmp = adjacent[dest].head;
     while (tmp->next != nullptr)
            tmp = tmp->next;

     tmp->next = node;
   }
}

int main()
{
   Graph grp(4);
   grp.add_edge(0, 1);
   grp.add_edge(1, 2);
   grp.add_edge(2, 3);
   grp.add_edge(2, 1);
   grp.add_edge(0, 3);
   grp.depth_first_search();
   std::cout << "\n";
}
```



Reference:<br/>
[Introduction to Algorithms](https://amzn.to/2OarGBs)<br/>
[The Algorithm Design Manual](https://amzn.to/2CH9h9Z)<br/>
[Data Structures and Algorithms Made Easy](https://amzn.to/2NLM0dd)<br/>
Competitive Programmer’s Handbook - Antti Laaksonen<br/>

 <input type="hidden" name="IL_IN_ARTICLE"> 
You may also like<br/>
[C++: Dijkstra's Algorithm using STL](/C-Dijkstra's-Algorithm-using-STL)<br/>
[C++: Bellman Ford Algorithm using STL](/C-Bellman-Ford-Algorithm-using-STL)<br/>
[Breadth First Search using Adjacency List](/C-Breadth-First-Search-using-Adjacency-List)<br/>


