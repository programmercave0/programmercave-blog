---
layout: post
title: "Low Level Design - What and Why?"
description: "Notes on Low Level Design from Scaler"
author: "Programmercave"
header-img: "/assets/images/LLD/lld.png"
tags:  [System-Design, Software-Engineering]
date: 2024-06-21
toc: true
---
![Low Level Design]({{ site.url }}/assets/images/LLD/lld.png){:class="img-responsive"}

### What is LLD:
- component-level design process that follows a step-by-step refinement process. 
- used for designing data structures, required software architecture, source code and ultimately, performance algorithms. 
- Overall, the data organization may be defined during requirement analysis and then refined during data design work. 
- Post-build, each component is specified in detail

#### Why LLD?
- goal of LLD or a low-level design (LLD) is to give the internal logical design of the actual program code.
- describes the class diagrams with the methods and relations between classes and program specs. 
- It describes the modules so that the programmer can directly code the program from the document.
- Ultimately, LLD has the following goals:
  - Low level implementation details of a system
  - organization of code
  - write good software
- A good software is a software that is
  - easy to maintain
  - easy to scale.
  - easy to extend

##### Maintainability:
- is a long-term aspect that describes how easily software can evolve and change, which is especially important in todayʼs agile environment.
- ISO 25010 states that a highly maintainable software system must possess the following qualities:
  - **Modularity** - The product is composed of discrete components such that a change to one component has minimal impact on other components.
  - **Reusability** - The product makes use of assets that can be re-used in building other assets or in other systems.
  - **Analyzability** - The impact of an intended change on the product, diagnosis of deficiencies, causes of failures or identification of the components that need to be changed can be analyzed effectively and efficiently.
  - **Modifiability** - The product can be effectively and efficiently modified without introducing defects or degrading existing product quality.
  - **Testability** - The test criteria can be established effectively and efficiently, and the product can be tested to determine whether those criteria have been met.

##### Scalability:
- is an attribute of a tool or a system to increase its capacity and functionalities based on its usersʼ demand. 
- Scalable software can remain stable while adapting to changes, upgrades, overhauls, and resource reduction.

##### Extensibility:
- is a measure of the ability to extend a system and the level of effort required to implement the extension. 
- can be through the addition of new functionality or through modification of existing functionality. 
- The principle provides for enhancements without impairing existing system functions

> These are my notes from the Low-Level Design (LLD) course I took at Scaler.