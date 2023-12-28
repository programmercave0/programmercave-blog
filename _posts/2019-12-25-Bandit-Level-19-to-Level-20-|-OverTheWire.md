---
layout: post
title: "Bandit Level 19 to Level 20 | OverTheWire"
description: "Learn linux command by playing Bandit wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Level 19 → Level 20 and Level 20 → Level 21. In this post we will learn about setuid and how to use it to access files of other users. The passwords are hidden, so you have to find the passwords for next level yourself."
author: "Programmercave"
header-img: "/assets/images/Bandit-Overthewire/overthewire_poster.jpg"
tags:  [Linux, OverTheWire-Bandit, CTF]
date: 2019-12-25
---



Learn linux command by playing [Bandit](https://overthewire.org/wargames/bandit/) wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Level 19 → Level 20 and Level 20 → Level 21. 

In this post we will learn about setuid and how to use it to access files of other users.

The passwords are hidden, so you have to find the passwords for next level yourself.

![Bandit OverTheWire]({{ site.url }}/assets/images/Bandit-Overthewire/overthewire_poster.jpg){:class="img-responsive"}

### Previous Post

[Bandit Level 0 to Level 3]({{ site.url }}/blog/2019/12/21/Bandit-Level-0-to-Level-5-OverTheWire)<br/>
[Bandit Level 4 to Level 8]({{ site.url }}/blog/2019/12/22/Bandit-Level-4-to-Level-9-OverTheWire)<br/>
[Bandit Level 9 to Level 11]({{ site.url }}/blog/2019/12/24/Bandit-Level-9-to-Level-12-OverTheWire)<br/>
[Bandit Level 12 → Level 13]({{ site.url }}/blog/2019/12/24/Bandit-Level-12-Level-13-OverTheWire)<br/>
[Bandit Level 13 to Level 15]({{ site.url }}/blog/2019/12/24/Bandit-Level-13-to-Level-16-OverTheWire)<br/>
[Bandit Level 16 to Level 18]({{ site.url }}/blog/2019/12/24/Bandit-Level-16-to-Level-19-OverTheWire)<br/>

## [Bandit Level 19 → Level 20](https://overthewire.org/wargames/bandit/bandit20.html)

### Level Goal

To gain access to the next level, you should use the setuid binary in the homedirectory. Execute it without arguments to find out how to use it. The password for this level can be found in the usual place (/etc/bandit_pass), after you have used the setuid binary.

### Solution : 

Command to connect remote host : `ssh bandit19@bandit.labs.overthewire.org -p 2220` password is `****` .

**Setuid**, which stands for set user ID on execution, is a special type of file permission in Unix and Unix-like operating systems such as Linux and BSD. It is a security tool that permits users to run certain programs with escalated privileges.

When an executable file's setuid permission is set, users may execute that program with a level of access that matches the user who owns the file.

When viewing a file's permissions with the `ls -l` command, the setuid permission is displayed as an "**s**" in the "user execute" bit position.

![Bandit Level 19 20]({{ site.url }}/assets/images/Bandit-Overthewire/bandit_l1920_terminal.jpg){:class="img-responsive"}

Here setuid is set for the binary file (`-rwsr-x--`) and the user is *bandit20*. So if execute `./bandit20-do` it will run as user *bandit20* and password for *bandit20* is stored in */etc/bandit_paas/bandit20* which is only accessible through user *bandit20*.

The command is 
```
./bandit20-do cat /etc/bandit_pass/bandit20
```

and the password is `****` .

Reference : [https://www.computerhope.com/jargon/s/setuid.htm](https://www.computerhope.com/jargon/s/setuid.htm)

{% include ads.html %}<br/>

## [Bandit Level 20 → Level 21](https://overthewire.org/wargames/bandit/bandit21.html)

### Level Goal

There is a setuid binary in the homedirectory that does the following: it makes a connection to localhost on the port you specify as a commandline argument. It then reads a line of text from the connection and compares it to the password in the previous level (bandit20). If the password is correct, it will transmit the password for the next level (bandit21).

**NOTE:** Try connecting to your own network daemon to see if it works as you think

### Commands you may need to solve this level

ssh, nc, cat, bash, screen, tmux, Unix ‘job control’ (bg, fg, jobs, &, CTRL-Z, …)

### Solution : 

Command to connect remote host : `ssh bandit20@bandit.labs.overthewire.org -p 2220` password is `****` .

First lets see how to create a TCP listener connection on our local machine. For this purpose we will use netcat program and option `-l` specify that `nc` should listen for an incoming connection rather than initiate a connection to a remote host. We will open two terminal and in first terminal the command is `nc -l localhost -p 5000`. The netcat server is listening to port 5000. `-p` specifies the source port `nc` should use.

Now in second terminal netcat client establishes connection to the server on port 5000. The command to do that is `nc -v localhost 5000`. Option `-v` give more verbose output. 

![Bandit Level 20 21]({{ site.url }}/assets/images/Bandit-Overthewire/bandit2021.gif){:class="img-responsive"}

Now login to a remote machine as user bandit20.

Binary file *suconnect* has escalated privilege as user *bandit21*. After running command `./suconnect` it shows how to use command i.e. *Usage: ./suconnect* <portnumber>

It also shows “*This program will connect to the given port on localhost using TCP. If it receives the correct password from the other side, the next password is transmitted back*”.

![Bandit Level 20 21]({{ site.url }}/assets/images/Bandit-Overthewire/bandit_l2021_terminal2.jpg){:class="img-responsive"}

First we have to establish a TCP connection like we did above. We can connect on any port for eg. port 12345. Then we pass the password of current level (bandit20). And if the password is correct, it will display password for level 21.

In terminal 1 server will listen on port 12345 and password of the current level is passed. Command is
```
nc -l localhost -p 12345 < /etc/bandit_pass/bandit20
```

In terminal 2 binary file suconncet will connect to server on port 12345 and if password matches, password for next level is displayed on the terminal 1. Command is 
```
./suconnect 12345
```

and the password is `****` .

Here connection is also established on port 12346.

![Bandit Level 20 21]({{ site.url }}/assets/images/Bandit-Overthewire/bandit_l2021_terminal3.jpg){:class="img-responsive"}

Reference : [https://unix.stackexchange.com/a/214480/244874](https://unix.stackexchange.com/a/214480/244874)<br/>
[https://linux.die.net/man/1/nc](https://linux.die.net/man/1/nc)<br/>
[https://www.billycody.com/otw-wargames/bandit/bandit-level-20](https://www.billycody.com/otw-wargames/bandit/bandit-level-20)<br/>

### Next Post

[Bandit Level 21 to Level 22]({{ site.url }}/blog/2019/12/25/Bandit-Level-21-to-Level-23-OverTheWire)
<br/>
[Bandit Level 23 → Level 24]({{ site.url }}/blog/2019/12/25/Bandit-Level-23-Level-24-OverTheWire)<br/>
[Bandit Level 24 → Level 25]({{ site.url }}/blog/2019/12/26/Bandit-Level-24-Level-25-OverTheWire)<br/>
[Bandit Level 25 to Level 26]({{ site.url }}/blog/2019/12/26/Bandit-Level-25-to-Level-26-OverTheWire)<br/>
[Bandit Level 27 to Level 31]({{ site.url }}/blog/2019/12/26/Bandit-Level-27-to-Level-31-OverTheWire)<br/>
[Bandit Level 32 → Level 33]({{ site.url }}/blog/2019/12/26/Bandit-Level-32-Level-33-OverTheWire)<br/>

### Other Wargames
[Leviathan Wargame from OverTheWire All Level Solutions]({{ site.url }}/blog/#overthewire-leviathan)<br/> 
[Krypton Wargame from OverTheWire All Level Solutions]({{ site.url }}/blog/#overthewire-krypton)<br/>
