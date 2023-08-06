---
layout: post
title: "Selected Questions from Interviews with Bjarne Stroustrup Part 1| Creator of C++ "
description: "What tips would you give to a budding programmer? Programming can be fun, but it is not all fun and games. Our civilization depends on good software.Whatever you are interested in, there will be a use for programming: literature, automobile production, history, coffee making, wind and solar energy, movie making, rocket science, farm management, medicine, science, engineering, and so much more.You might eventually be able to make a significant contribution to whatever field you find important or interesting! To do that you’ll have to know the field and the tools and techniques of software development. Be sure you learn about fundamentals, such as data structures and how machines work."
author: "Programmercave"
header-img: "/assets/2019-12-12-Bjarne-Stroustrup/bjarne.jpg"
tags:  [Cpp, Articles]
date: 2019-12-12
---
* toc
{:toc}

[Bjarne Stroustrup](http://www.stroustrup.com/index.html) is a Technical Fellow and Managing Director at Morgan Stanley in New York City and a Visiting Professor at Columbia University. He's also the creator of C++.

In 1975, Bjarne got his MSc degree from Department of Mathematics at Aarhus University when the now Department of Computer Science was still embedded in the Math Department.

Bjarne Stroustrup created the C++ programming language in 1979, and it’s still just as important today as it was 38 years ago, especially with its expansion into mobile development.

![Bjarne Stroustrup]({{ site.url }}/assets/2019-12-12-Bjarne-Stroustrup/bjarne.jpg){:class="img-responsive"}

This is the first post of the series Selected Questions from Interviews with Bjarne Stroustrup.<br/>
Selected Questions from Interviews with Bjarne Stroustrup Part 1<br/>
[Selected Questions from Interviews with Bjarne Stroustrup Part 2]({{site.url}}/blog/2019/12/12/Selected-Questions-from-Interviews-with-Bjarne-Stroustrup-Part-2-Creator-of-C++).

{% include ads.html %}<br/>

### An interview by Sonny Li from [Codecademy](https://www.codecademy.com/): [Talking C++](https://news.codecademy.com/bjarne-stroustrup-interview/). September 2019.

**What tips would you give to a budding programmer?**

Programming can be fun, but it is not all fun and games. Our civilization depends on good software.

Whatever you are interested in, there will be a use for programming: literature, automobile production, history, coffee making, wind and solar energy, movie making, rocket science, farm management, medicine, science, engineering, and so much more.

You might eventually be able to make a significant contribution to whatever field you find important or interesting! To do that you’ll have to know the field and the tools and techniques of software development. Be sure you learn about fundamentals, such as data structures and how machines work.

Don’t get overwhelmed, and don’t think you will be an expert in just a few weeks. Think of how long it takes to learn to speak a natural language, how long it takes to become a good athlete, and how long it takes to learn to play an instrument well enough for someone who isn’t your mother to want to listen to.

Think about how much fun you can have along the way to reach such mastery and how many friends you might make along the way. Some of the nicest people work with software.

**[The Definitive C++ Book Guide and List](https://stackoverflow.com/questions/388242/the-definitive-c-book-guide-and-list)**

Unfortunately, there is no definite C++ book list. There can’t be one. Not everyone needs the same information, not everyone has the same background, and C++ best practices are evolving.

I did a search of the web and found a bewildering set of suggestions. Many were seriously outdated and some were bad from the start. A novice looking for a good book without guidance will be very confused!

You do need a book because the techniques that make C++ effective are not easily picked up from a few blogs on specific topics—and of course blogs also suffer from mistakes, being dated, and poor explanations. Often, they also focus on advanced new stuff and ignore the essential fundamentals.

I recommend my [Programming: Principles and Practice Using C++ (2nd Edition)](https://amzn.to/2PbUbQZ) for people just beginning to learn to program, and [A Tour of C++ (2nd Edition)](https://amzn.to/2Eabg7y) for people who are already programmers and need to know about modern C++. People with a strong mathematical background can start with Peter Gottschling‘s [Discovering Modern C++: An Intensive Course for Scientists, Engineers, and Programmers](https://amzn.to/2qMmNH8).

Once you start using C++ for real, you need a set of guidelines to distinguish what can be done and what is good practice. For that, I recommend the [C++ Core Guidelines on GitHub](https://github.com/isocpp/CppCoreGuidelines/blob/master/CppCoreGuidelines.md).

For good brief explanations of individual language features and standard-library functions, I recommend [www.cppreference.com](http://www.cppreference.com/).

{% include ads.html %}<br/>

### An interview by Jordi Pérez Colomé for the Spanish daily newspaper El Pais: [Nuestra civilización depende igual del software que del agua](https://elpais.com/tecnologia/2019/02/01/actualidad/1549023454_360652.html). In Spanish. [Translation back to English](http://www.stroustrup.com/stroustrup_la_pais.pdf). January 2019. 

**Were you surprised by how much and how fast your creation became popular (C++)?**

Definitely. It was just a personal project to solve specific needs. I thought of it as a tool, but then it turned out to be useful to my colleagues, and then to many more people. I think that the reason for that utility was partly that the need for a language that could cope with both lower-level needs and higher-level ones was a very wide-spread one and that that need was increasing as computers became more powerful (leading to larger programs) and because our expectations of what a computer would do for us increased dramatically. The fact that I was addressing higher-level needs through efficient abstraction mechanisms rather than specific solutions to specific problems was key to the unexpected broad appeal. People could build their own types to address their own specific needs.

**After the turn of the millennium, the popularity of C++ dropped and for almost a decade and a half, it saw its base eroding slowly but steadily. In the past few years however, according to several statistics (TIOBE, Stack Overflow Developer Survey), the language has dynamically grew again suggesting that more people are using it again. What do you think the reason is for this? Could there be a certain event or change of trends in the background of this development? Do you attach any importance to these statistics, numbers and how many people use the language?**

Unfortunately, these “surveys” measure noise, rather than usage. One enthusiastic student easily carries more weight in those than 2,000 busy developers. There are more professional surveys, such as [https://blog.jetbrains.com/clion/2015/07/infographics-cpp-facts-before-clion/](https://blog.jetbrains.com/clion/2015/07/infographics-cpp-facts-before-clion/) , but they are expensive to conduct, and much rarer.I’m pretty sure that the C++ community shrank a bit in the years before 2005 and grew steadily after that as single-processor performance stopped increasing and portability across platforms became a larger issue. My guess is that the C++ developer community now exceeds 4.5 million and is growing by about 100,000 a year, but I don’t think anyone have precise and reliable figures. The reason for this growth in C++ use is an increase in the importance of its fundamental design aims. There are simply more systems needing both performance and to manage complexity. In many areas, performance per watt (energy efficiency) and reliability are more important than developer cost.The fact that C++ is now a significantly better language than is was in year 2000, is probably also significant.

**What are the reasons programmers should (and do) consider using C++ / and in what cases are they better off using something else?**

If you need to use your hardware well, for reasons of run-time efficiency, latency, hardware costs, energy efficiency C++ becomes an obvious candidate. If you also need reliability, need to manage significant complexity, or expect your system to be maintained for decades the case for C++ becomes very strong.If you can afford to devote 90% or more of your hardware capacity to simplify the life for developers, other languages often gain an edge.The case for C++ is stronger for professional developers than for casual users who are less concerned about the engineering and maintainability of programs. 

The result is that C++ is most prominent in infrastructure and performance sensitive areas. For example, people who are not professional programmers might prefer JavaScript for its relative ease of use, but the JavaScript engine will be C++. Many of the languages that people see as competitors to C++ are actually C++ applications.

{% include ads.html %}<br/>

### [Alumni interview: Q&A med Bjarne Stroustrup, Managing Director for Morgan Stanley i New York](http://cs.au.dk/news-events/pages/alumni-interview-qa-med-bjarne-stroustrup-managing-director-for-morgan-stanley-i-new-york/). October 2017.

**When did you decide to specialize in programming languages?**

I did not have this objective from the beginning – at first, I was curious. Could you develop a set of instructions to support better computer systems? I was probably more interested in hardware and microprogramming than most of my fellow students. My curiosity steered me towards my PhD thesis, which focused on how to render hardware support and communication in distributed systems. Later, when I was working at Bell Labs, I changed direction into creating support by using software. It was too inflexible and slow to build support for abstractions into hardware.  C++ reflects the necessity of having a programming language that supports multiple levels in a computer system at the same time. In 1979, this was new – and today it is still not the norm outside C++.

**What is your best advice for a computer science student?**

Be curios – take in the fundamental knowledge during your studies. Even if it does not seem relevant to you now. You may find yourself in a potential and exciting situation, where you need to put the things you learned into use. In the US, people are very goal-oriented from an early age. They know exactly what kind of educations, jobs and courses they need to pursue in order to reach their goals. This is not necessarily a negative thing, but in my experience career paths can evolve and take you to unforeseen directions. If you do not stay curios, you might overlook opportunities on your way. You do not know what you are doing ten years from now, but you can improve your toolbox by learning the fundamental concepts and techniques within i.e. machine architecture, algorithms, data structures and operating systems. And math. It is also important to develop ”soft skills” – to learn to understand other people’s problems and to be able to put yourself in their place. I had a student job where I solved programming assignments for small companies in Aarhus – all sorts of companies from stone masonries and carpenters to banks that were building systems that could handle mortgages. The programming tasks funded my studies and taught me how to talk to customers. To listen and understand. I learned to take responsibility for my programming tasks, because the customer’s livelihood was dependent on me carrying out a solid piece of work.

**What was your best decision ever?**

That I allowed myself to be curios and take chances. I did not know what a PhD was, but I thought it was a good idea to try it out, and see if I could make it. I was interested in immersing myself into a subject area. Later on, it was my PhD degree that opened the doors for me to Bell Labs in New Jersey, which was the most amazing place to be if you wanted to engage in practical computer science and system building. I learned so much while I was there. I was lucky – but you can only gain from luck by being aware of your options and come prepared with a good education and a little confidence. 

{% include ads.html %}<br/>

### The Setup: [What do people use to get stuff done?](https://usesthis.com/interviews/bjarne.stroustrup/). I wonder why people care? July 2015. 

**What would be your dream setup?**

Two pounds, good keyboard, good pen for scribbling and drawing, dockable, running all the major operating systems. Good Internet connection, but able to perform perfectly when disconnected from the web. Able to act as an interface to other systems. Simple and predictable to use. Stable, never needs rebooting.

Obviously, such a marvel does not exist today.

<h3>Reference</h3>
[Interviews with Bjarne Stroustrup](http://www.stroustrup.com/interviews.html)

<h3>Books written by Bjarne Stroustrup:</h3>
[A Tour of C++ (2nd Edition)](https://amzn.to/2Eabg7y)<br/>
[Programming: Principles and Practice Using C++ (2nd Edition)](https://amzn.to/2PbUbQZ)<br/>
[The C++ Programming Language (4th Edition)](https://amzn.to/36B1DLt)<br/>
[The Design and Evolution of C++](https://amzn.to/349eIde)<br/>
