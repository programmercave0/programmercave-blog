---
layout: post
title: "Leviathan Level 0 to Level 1 | Basic Exploitation Techniques"
description: "Learn linux command by playing Leviathan wargame from OverTheWire. This wargame doesn't require any knowledge about programming - just a bit of common sense and some knowledge about basic *nix commands. Below is the solution of Level 0 → Level 1 and Level 1 → Level 2. In this post we will learn how to use a debugging tool ltrace to exploit a program."
author: "Programmercave"
header-img: "/assets/Leviathan-OverTheWire/overthewire_poster.jpg"
tags:  [Linux, OverTheWire-Leviathan, CTF]
date: 2020-01-06
---
* toc
{:toc}

Learn linux command by playing [Leviathan](https://overthewire.org/wargames/leviathan/) wargame from OverTheWire. This wargame doesn't require any knowledge about programming - just a bit of common sense and some knowledge about basic *nix commands.

Below is the solution of Level 0 → Level 1 and Level 1 → Level 2. In this post we will learn how to use a debugging tool **ltrace** to exploit a program.

![Leviathan OverTheWire]({{ site.url }}/assets/Leviathan-OverTheWire/overthewire_poster.jpg){:class="img-responsive"}

## [Level 0](https://overthewire.org/wargames/leviathan/leviathan0.html)

To login to the first level use:

Username: leviathan0<br/>
Password: leviathan0

The command to login into level 0 is `ssh leviathan0@leviathan.labs.overthewire.org -p 2223`
and the password is `leviathan0` .

## [Leviathan Level 0 → Level 1](https://overthewire.org/wargames/leviathan/leviathan1.html)

It is given at the homepage of this wargame that data for the levels can be found in the **homedirectories**. You can look at **/etc/leviathan_pass** for the various level passwords.

In *leviathan0* directory there is a directory *.backup* whose user is *leviathan1*. This can be found by command `ls -la`.

![Leviathan OverTheWire]({{ site.url }}/assets/Leviathan-OverTheWire/levi_l01_terminal1.jpg){:class="img-responsive"}

In *.backup* directory there is a file *bookmarks.html*. This file may contain the password for next level. To find that, the command is `cat bookmarks.html | grep leviathan` and the password is `rioGegei8m`.

![Leviathan OverTheWire]({{ site.url }}/assets/Leviathan-OverTheWire/levi_l01_terminal2.jpg){:class="img-responsive"}


{% include ads.html %}<br/>

## [Leviathan Level 1 → Level 2](https://overthewire.org/wargames/leviathan/leviathan2.html)

Command to login is `ssh leviathan1@leviathan.labs.overthewire.org -p 2223` and password is  `rioGegei8m` .

We have an executable file check with setuid for user leviathan2 and it can be executed by leviathan1. 

![Leviathan OverTheWire]({{ site.url }}/assets/Leviathan-OverTheWire/levi_l12_terminal1.jpg){:class="img-responsive"}

But when we execute it, it asks for the password.

![Leviathan OverTheWire]({{ site.url }}/assets/Leviathan-OverTheWire/levi_l12_terminal2.jpg){:class="img-responsive"}

To find the password, which is asked after executing binary file, we will use `ltrace` command.

`ltrace` is a diagnostic and debugging tool for the command line. It intercepts and records the dynamic library calls which are called by the executed process and the signals which are received by that process.  It can also intercept and print the system calls executed by the program.

This command will tell about all the system call this binary file made when it is executed.

The command is 
```
ltrace ./check
```

 and we are prompted to enter password. We enter *test* and one of the function called is `strcmp`. It compares *test* with *sex*. This means *sex* is the password the binary file need.

![Leviathan OverTheWire]({{ site.url }}/assets/Leviathan-OverTheWire/levi_l12_terminal3.jpg){:class="img-responsive"}

After we enter *sex* as password, we are in root shell and then we can see the password for next level using command `cat /etc/leviathan_pass/leviathan2` and the password is `ougahZi8Ta` .

![Leviathan OverTheWire]({{ site.url }}/assets/Leviathan-OverTheWire/levi_l12_terminal4.jpg){:class="img-responsive"}

Reference : [Chapter 9. ltrace](https://access.redhat.com/documentation/en-us/red_hat_developer_toolset/7/html/user_guide/chap-ltrace)<br/>
[LTRACE(1)](http://man7.org/linux/man-pages/man1/ltrace.1.html)<br/>

### Next Posts
[Leviathan Level 2 → Level 3]({{ site.url }}/blog/2020/01/06/Leviathan-Level-2-3-Basic-Exploitation-Techniques)<br/>
[Leviathan Level 3 to Level 4]({{ site.url }}/blog/2020/01/06/Leviathan-Level34-Basic-Exploitation-Techniques)<br/>
[Leviathan Level 5 to Level 6]({{ site.url }}/blog/2020/01/06/Leviathan-Level-5-to-6-Basic-Exploitation-Techniques)<br/>

### Other Wargames
[Bandit Wargame from OverTheWire All Level Solutions]({{ site.url }}/blog/#overthewire-bandit) <br/>
[Krypton Wargame from OverTheWire All Level Solutions]({{ site.url }}/blog/#overthewire-krypton)<br/>
