---
layout: post
title: "Selected Questions from Interviews with Bjarne Stroustrup Part 2| Creator of C++"
description: "Why did you choose C as the language to add object oriented features of Simula? There were a few Procedural languages like COBOL, PASCAL to add OO Features. Also Your favorite language was Algol68.Bjarne: I knew quite a few languages at the time, maybe 25 or thereabouts. It was easier to learn languages at that time. They were not as big or complicated, and their implementations were not so complicated. Algol68 was quite a beautiful language, very elegant.The reason I chose C was that I needed to manipulate hardware. I needed to write things like schedulers, memory allocators, device drivers."
author: "Programmercave"
header-img: "/assets/2019-12-12-Bjarne-Stroustrup/bjarne1.jpg"
tags:  [Cpp, Articles]
date: 2019-12-12
---

![Bjarne Stroustrup]({{ site.url }}/assets/2019-12-12-Bjarne-Stroustrup/bjarne1.jpg){:class="img-responsive"}

This is the second post of the series Selected Questions from Interviews with Bjarne Stroustrup.<br/>
[Selected Questions from Interviews with Bjarne Stroustrup Part 1]({{site.url}}/blog/2019/12/12/Selected-Questions-from-Interviews-with-Bjarne-Stroustrup-Inventor-of-Cpp)<br/>
Selected Questions from Interviews with Bjarne Stroustrup Part 2.

{% include ads.html %}<br/>

### Interview by Pramod Shadishidhaa: [Mapping the Journey](https://www.mappingthejourney.com/single-post/Interview-with-Bjarne-Stroustrup). A podcast with transcript. July 2017. 

**Pramod: BELL labs was the go to place for top computer scientists. How did you get this opportunity? What was the experience working with some great engineers at Bell Labs?**

Bjarne: It was the place to be to if you we’re a practical computer science person or even some of the theoreticians. I got there because there were some connections between the Computer Lab in Cambridge and Bell Labs. Some people had moved over and put some work together. Things like the language BCPL had been invented at MIT by somebody from Cambridge. From there I got into to Bell Labs.

One of the things that were great about Bell Labs is that it was a really good environment with some really great people there. There was Dennis Ritchie, Brian Kernighan, Al Aho, and in particular a guy called Douglas McIlroy, who was everybody’s favorite listener and one of the best catalyst I’ve ever met. You could explain things to him, and that was hard because he asks really good questions and then after a while, you learned something from that exercise. So there were some really great people, and there were a load of challenges.

Bell Labs was doing just about everything the time; hardware, software, communication networks, fibers, and CCDs invented for what later became digital cameras. So there’s a lot of stuff there, real enormous diversity of taste and tasks, and people that you could collaborate with. It was not a place where people were hiding their ideas from each other or hoarding them for their next paper. Really good exchange of ideas and enlightened management that allows us to try actually to develop and develop good things.

**Pramod: Why did you choose C as the language to add object oriented features of Simula? There were a few Procedural languages like COBOL, PASCAL to add OO Features. Also Your favorite language was Algol68.**

Bjarne: I knew quite a few languages at the time, maybe 25 or thereabouts. It was easier to learn languages at that time. They were not as big or complicated, and their implementations were not so complicated. Algol68 was quite a beautiful language, very elegant.

The reason I chose C was that I needed to manipulate hardware. I needed to write things like schedulers, memory allocators, device drivers. That was not a strength of Algol68, and in general, it’s not a strength of the more elegant languages because the ability to manipulate hardware is missing. I had to choose one of the close-to-the-machine languages. There were a few. C seems a pretty good one of them and there was local support at Bell Labs of course. When I was confused about something I could go and ask Dennis or Brian and that’s important. Local support is always important. But the real need was to manipulate hardware directly and efficiently.

**Pramod: Any newer languages like Rust or Go interests you?**

Bjarne: Yeah, I look at new languages roughly all the time. You learn something every time. It is interesting but as I said I’m more of a systems guy than the languages guys. I probably look less at languages than you’d think. I look at what I think is interesting, I try them out but what I really do is to look at what problems do my colleagues and students have and how do I solve them.

{% include ads.html %}<br/>

### [An interview by Aditya Bhushan Dwivedi](http://yourstory.com/2013/12/bjarne-stroustrup-interview/) for [Your Story](http://yourstory.com/). 

**YS: Very little is known about your early days. What sparked your interest in computers?**

Bjarne: Like many high-school students, I didn’t have a clue what to do after high-school. I read up on various possible careers, and dreamt of becoming an architect, a historian, a sociologist, an engineer, a psychologist, and more. My family was encouraging me to choose something with a clear job and career path after my studies, such as a business major. 

In the end, I chose mathematics with computer science because mathematics was one of my favorite topics in high school and I was under the mistaken impression that computer science was a form of applied mathematics. 

I had never seen a computer – in those days computers were rare and hidden away in dedicated machine rooms. I chose mathematics because it would make a good career and a lousy hobby over my second choice, history, which I thought would make a lousy career and a great hobby. To this day, I have more history books than computing books, and I’m part-way through a book on the history of England in the 14th century. 

I was right that mathematics would make a great career, but only because I had been completely mistaken about computer science. It turned out to be far more interesting, practical, and useful than I had imagined. Also, real applied mathematics would probably have been beyond my limited talents in that direction. I ended up doing pure math and programming. So, I was signed up for computer science from the start, but what really sparked my interest was the introductory computer science course.

**YS: As an academic, what steps should be taken to improve the way programming is taught in schools and colleges around the world to create better and employable programmers.**

Bjarne: We need a better balance between ‘theory’ and practice. There is a lot of talk about teaching methods, but I think the real problem is that the content is lacking in appropriateness; not that the delivery methods are inadequate (though of course that is also a problem in many places).

Good software and software development require a mixture practical skills and book learning. I find that teachers and students often fail to emphasize one or the other. For more details and examples, see my CACM paper: ‘What should we teach software developers? Why?’ Basically, I argue for a conventional foundation (including mathematics, machine architecture, data structures, and algorithms) complemented by some form of specialization (e.g., graphics, networking, or real-time systems).

Unsurprisingly, I think that greater emphasis should be placed on the practical aspects of constructing quality software (correctness, reliability, maintainability, tools, and performance). Finally, I think that a good developer needs a minor in some non-computer field to learn to appreciate other fiends and to communicate with non-programmers. That’s obviously a lot, so I consider a master’s degree (rather than a bachelor degree) the ideal for developers of demanding applications.

I think we need to place much greater emphasis on professionalism. Our civilization depends on software, so we need professionals to build and maintain our foundational software, not just semi-educated craftsmen willing to ship any program their managers tell them to ship. I see software development as a potentially noble profession, like medicine or some of the classical engineering disciplines, but we still have a long way to go to get there.

**YS: What is the biggest mistake you think people commit while learning to program in any language?**

Bjarne: I’m not sure what the biggest mistake is, but the first mistake is typically writing code using the style and techniques from some previous language. Some people never get over that. Some people even try to make a virtue out of their lack of understanding, insisting that they know how best to program independently of programming languages. To use a language well, you have to learn its strengths and weaknesses, and to use it idiomatically, the way expert programmers in that language use it to its greatest effect. You can write Fortran in any language, and C, but you don’t have to, and for most languages trying to avoid facilities that go beyond what is available in most languages lead to suboptimal designs, hard-to maintain, low-level, complicated application code that consumes excessive resources (time, space, bandwidth, power).

There are many candidates for the ‘biggest mistake’ title, but maybe not adopting a colloquial style of a language and using a mostly language-independent, low-level, style of code (emulating ‘foreign’ programming styles) could be it. To use any programming language well, you need an understanding of its fundamentals and its ‘native’ techniques.

**YS: What are your thoughts about competitive programming? is it really important to learn one language in depth or learn multiple languages at a basic level to be able to write error free applications?**

Bjarne: To be a professional, you need to learn to use several languages idiomatically. Knowing just one language is limiting and a shallow understanding of several languages make you write unnecessarily complicated code in every one of those.

**YS: What are the three key lessons you would like to share with programmers in their early twenties who are just starting up?**

Bjarne: Why three? Well, OK: 

Learn several programming language well. 

Learn to communicate well, verbally and in writing. 

Develop and maintain interests outside software development. Programming is a way of expressing idea, but unless you have an understanding of an application area you won’t have the insights to make contributions to any application area. 

In addition to your computer science basics (such as machine architecture, data structures, and algorithms), you should spend significant time on some non-computer topic (mathematics, physics, biology, history, architecture or literature), and learn to speak and write effectively.

<h3>Reference</h3>
[Interviews with Bjarne Stroustrup](http://www.stroustrup.com/interviews.html)

<h3>Books written by Bjarne Stroustrup:</h3>
[A Tour of C++ (2nd Edition)](https://amzn.to/2Eabg7y)<br/>
[Programming: Principles and Practice Using C++ (2nd Edition)](https://amzn.to/2PbUbQZ)<br/>
[The C++ Programming Language (4th Edition)](https://amzn.to/36B1DLt)<br/>
[The Design and Evolution of C++](https://amzn.to/349eIde)<br/>
