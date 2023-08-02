---
layout: post
title: "How to create Password list using CUPP tool on Ubuntu"
subtitle: "People do not create complex passwords. Their password consists of their name or surname, date of birth, favourite football team etc, so they can remember it easily. There are many tools which create password list using these parameters. Here we are going to discuss about one such tool CUPP and how to create password lists using it."
author: "Programmercave"
header-img: "/assets/cupp1.png"
tags:  [Python, Ubuntu, Linux, HowTo, Hacking]
date: 2019-10-10
---

People do not create complex passwords. Their password consists of their name or surname, date of birth, favourite football team etc, so they can remember it easily. There are many tools which create password list using these parameters. Here we are going to discuss about one such tool CUPP and how to create password lists using it.

CUPP(Common User Passwords Profiler) is python based free and open source tool which generates wordlist for a specific user based on the information provided. It provides an interactive interface to build its custom password lists.

**Requirements**<br/>
To run CUPP install Python3 on your Ubuntu.

**Quickstart**<br/>
Clone [cupp](https://github.com/Mebus/cupp) repository from Github using this command:
  ```
  git clone https://github.com/Mebus/cupp.git
  cd cupp
  ```
  
**Options**<br/>
You can view the available options using this command:
  ```
  python3 cupp.py -h
  ```
![Options]({{ site.url }}/assets/cupp.png)

Before starting you should gather enough information about the victim so that chance of password in the list increases.
<input type="hidden" name="IL_IN_ARTICLE"> 
To start tool in interactive mode write this command and enter details.
  ```
  python3 cupp.py -i
  ```
  
![Output]({{ site.url }}/assets/cupp1.png)

The file is saved in the same folder with the name provided above.


