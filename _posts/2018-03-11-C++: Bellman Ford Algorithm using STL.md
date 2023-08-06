---
layout: post
title: "Bellman Ford Algorithm | Single-Source Shortest Path"
description: "Bellman–Ford algorithm finds shortest path from the source vertex to all vertices in the graph. The graph can contain negative-weight edges, but it should not contain a negative-weight cycle that is reachable from the source vertex."
author: "Programmercave"
header-img: "/assets/bellman.png"
tags:  [Cpp, Algorithm, Graph-Algorithms]
date: 2018-03-11
---
* toc
{:toc}

**Bellman–Ford** algorithm finds shortest path from the source vertex to all vertices in the graph. The graph can contain negative-weight edges, but it should not contain a negative-weight cycle that is reachable from the source vertex.

The algorithm returns TRUE if there is no negative-weight cycle and FALSE if there is a negative-weight cycle reachable from the source vertex. If there is a negative-weight cycle that is reachable from source vertex, then no solution exists.

![Bellman-Ford]({{ site.url }}/assets/bellman.png){:class="img-responsive"}

In fig. (a) there is no negative-weight cycle, so Bellman Ford algorithm finds the shortest path from source if fig. (a) is in a graph whereas fig. (b) contains a negative-weight cycle therefore no solution exists.

![Bellman-Ford]({{ site.url }}/assets/bellman1.png){:class="img-responsive"}
![Bellman-Ford]({{ site.url }}/assets/bellman2.png){:class="img-responsive"}

In fig. (a) the graph is in inital configuration. All vertices were at distance infinity from source vertex *a*. Then vertices *b* and *c* are reached and their distance from vertex *a* is updated. 

In fig. (d) the sum of distance of path *a -> b -> d* is 2 + 5 = 7 whereas sum of distance of path *a -> c -> d* is 7 + (-3) = 4. Hence path *a -> c -> d* is preferred which is shown in red line. In fig. (e) red line shows shortest path from *a* to *e*.

<h1>Implementation</h1>

A data structure is needed to store the distance of a vertex from the source vertex.

Here is the implementation of `bellman_ford` and `relax` function.

```cpp
struct Vertex
    {
        std::size_t id;
        int distance = std::numeric_limits<int>::max();
        Vertex(std::size_t id) : id(id) {}
    };

void relax(std::size_t src, std::size_t dest, int weight)
{
    if (vertices[dest].distance > (vertices[src].distance + weight))
    {
        vertices[dest].distance = (vertices[src].distance + weight);
    }
}

bool bellman_ford(std::size_t src)
{
    //initialize distance of source
    vertices[src].distance = 0;

    for (int i = 0; i < vertices.size() - 1; i++)
    {
        for (auto it = edge_weight.begin(); it != edge_weight.end(); it++)
        {
            relax(it->first.first, it->first.second, it->second);
        }
    }

    for (auto it = edge_weight.begin(); it != edge_weight.end(); it++)
    {
        if (vertices[it->first.second].distance
            > (vertices[it->first.first].distance + it->second))
           return false;
    }
    return true;
}
```

{% include ads.html %}<br/>

`relax` function updates the distance of the vertex from the source vertex if new calculated distance is smaller than the stored distance.

The time complexity of Bellman Ford algorithm is *O(nm)* where *n* is the number of vertices and *m* is the number of edges.

**Related:** [Dijkstra's Algorithm]({{ site.url }}/blog/2018/03/14/C++-Dijkstra's-Algorithm-using-STL)

Here is the C++ Implementation of Bellman-Ford Algorithm

```cpp
#include <iostream>
#include <vector>
#include <limits>
#include <map>

class Graph
{

    struct Vertex
    {
        std::size_t id;
        int distance = std::numeric_limits<int>::max();
        Vertex(std::size_t id) : id(id) {}
    };
    std::vector<Vertex> vertices;
    std::map<std::pair<std::size_t, std::size_t>, int> edge_weight;

  public:
    Graph(std::size_t);
    void add_edge(std::size_t, std::size_t, int); //source, destination, weight
    bool bellman_ford(std::size_t); //source
    std::ostream& distance_from_source(std::ostream&) const;

  private:
    void relax(std::size_t, std::size_t, int); //source, destination, weight
};

Graph::Graph(std::size_t size)
{
    vertices.reserve(size);
    for (int i = 0; i < size; i++)
    {
        vertices.push_back(Vertex(i));
    }
}

void Graph::add_edge(std::size_t src, std::size_t dest, int weight)
{
    if (src == dest)
        throw std::logic_error("Source and destination vertices are same");

    if (src < 0 || vertices.size() <= src)
        throw std::out_of_range("Enter correct source vertex");

    if (dest < 0 || vertices.size() <= dest)
        throw std::out_of_range("Enter correct destination vertex");

    const auto inserted = edge_weight.insert(std::make_pair(
                          std::make_pair(src, dest), weight ));
    if (!inserted.second)
        throw std::logic_error("Existing edge");
}

void Graph::relax(std::size_t src, std::size_t dest, int weight)
{
    if (vertices[dest].distance > (vertices[src].distance + weight))
    {
        vertices[dest].distance = (vertices[src].distance + weight);
    }
}

bool Graph::bellman_ford(std::size_t src)
{
    //initialize distance of source
    vertices[src].distance = 0;

    for (int i = 0; i < vertices.size() - 1; i++)
    {
        for (auto it = edge_weight.begin(); it != edge_weight.end(); it++)
        {
            relax(it->first.first, it->first.second, it->second);
        }
    }

    for (auto it = edge_weight.begin(); it != edge_weight.end(); it++)
    {
        if (vertices[it->first.second].distance
            > (vertices[it->first.first].distance + it->second))
           return false;
    }
    return true;
}

std::ostream& Graph::distance_from_source(std::ostream& os) const
{
    os << "Vertex\t\tDistance from Source\n";
    for (int i = 0; i < vertices.size(); i++)
    {
        os << vertices[i].id <<"\t\t" << vertices[i].distance <<"\n";
    }
}


int main()
{
    Graph grp(5);
    grp.add_edge(0, 1, 6);
    grp.add_edge(0, 2, 7);
    grp.add_edge(1, 3, 5);
    grp.add_edge(1, 4, -4);
    grp.add_edge(1, 2, 8);
    grp.add_edge(2, 3, -3);
    grp.add_edge(2, 4, 9);
    grp.add_edge(3, 1, -2);
    grp.add_edge(4, 0, 2);
    grp.add_edge(4, 3, 7);

    bool res = grp.bellman_ford(0);
    if (res == true)
       std::cout << "Graph contains no negative cycle \n";
    else
       std::cout << "Graph conatins negative cycle \n";

    grp.distance_from_source(std::cout);
}
```

Output

![Output]({{ site.url }}/assets/BellmanOut.png){:class="img-responsive"}

Get this post in pdf - [Bellman Ford Algorithm](https://www.file-up.org/1i0k5ezjgq66)

Reference:<br/>
[Introduction to Algorithms](https://amzn.to/2OarGBs)<br/>
[The Algorithm Design Manual](https://amzn.to/2CH9h9Z)<br/>
[Data Structures and Algorithms Made Easy](https://amzn.to/2NLM0dd)<br/>
Competitive Programmer’s Handbook - Antti Laaksonen<br/>


 <input type="hidden" name="IL_IN_ARTICLE"> 
You may also like<br/>
[Breadth First Search using Adjacency List | Graph traversal]({{ site.url }}/blog/2018/03/06/C++-Breadth-First-Search-using-Adjacency-List)<br/>
[Depth First Search using Adjacency List | Graph traversal]({{ site.url }}/blog/2018/03/05/C++-Depth-First-Search-using-Adjacency-List)<br/>

