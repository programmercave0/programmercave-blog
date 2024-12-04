---
author: Programmercave
date: "2019-12-21T00:00:00Z"
description: Learn linux command by playing Bandit wargame. The Bandit wargame is
  aimed at absolute beginners. It will teach the basics needed to be able to play
  other wargames. Below is the solution of Level0, Level 0 → Level 1, Level 1 → Level
  2, Level 2 → Level 3, and Level 3 → Level 4. In this post we will learn how to connect
  to a remote machine using ssh and how to find a file with certain attributes in
  the machine. The passwords are hidden, so you have to find the passwords for next
  level yourself.
header-img: /assets/images/Bandit-Overthewire/overthewire_poster.jpg
tags:
- Linux
- OverTheWire-Bandit
- CTF
title: Bandit Level 0 to Level 3 | OverTheWire
---



Learn linux command by playing [Bandit](https://overthewire.org/wargames/bandit/) wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Level0, Level 0 → Level 1, Level 1 → Level 2, Level 2 → Level 3, and Level 3 → Level 4. 

In this post we will learn how to connect to a remote machine using ssh and how to find a file with certain attributes in the machine.

The passwords are hidden, so you have to find the passwords for next level yourself.

![Bandit OverTheWire](/assets/images/Bandit-Overthewire/overthewire_poster.jpg)

## [Bandit Level 0](https://overthewire.org/wargames/bandit/bandit0.html)

### Level Goal 

The goal of this level is for you to log into the game using SSH. The host to which you need to connect is **bandit.labs.overthewire.org**, on port 2220. The username is **bandit0** and the password is **bandit0**. Once logged in, go to the [Level 1](https://overthewire.org/wargames/bandit/bandit1.html) page to find out how to beat Level 1.

### Commands you may need to solve this level

ssh

### Solution :
 
**SSH** (Secure Shell) provides secure connection with a remote host. It prevents “man in the middle” attack by authenticating that the remote host is who it says it is. It encrypts all of the communications between the local and remote hosts.

Command `man ssh` tells us more about ssh.

We have given an address - *bandit.labs.overthewire.org*, port - *2220*, username - *bandit0* and password – *bandit0*.
The option `-p` is tell the port to connect and the general command to connect is `ssh username@address -p port`

So the command to connect to bandit server is :
```
ssh bandit0@bandit.labs.overthewire.org -p 2220
```

![Bandit Level 0](/assets/images/Bandit-Overthewire/bandit_l0_terminal.jpg)

A message “The authenticity of host … can’t be established” is displayed when connection is established for first time. To connect enter “yes” and once the connection is established, the user is asked to enter the password which is *bandit0* for this level.

Reference : [The Linux Command Line – A Complete Introduction](https://amzn.to/2PDVmZz)

<br/>

## [Bandit Level 0 → Level 1](https://overthewire.org/wargames/bandit/bandit1.html)

### Level Goal

The password for the next level is stored in a file called **readme** located in the home directory. Use this password to log into bandit1 using SSH. Whenever you find a password for a level, use SSH (on port 2220) to log into that level and continue the game.

### Commands you may need to solve this level

ls, cd, cat, file, du, find

### Solution :

`ls` command is used to see list of files and subdirectories contained in the current working directory and determine variety of important files and directory attributes.

Current working directory can be found using `pwd` command.

`cat` command is used to view the content of a file, concatenate file and redirect output in terminal or a file. It can accept more than one file as an argument, so it is used to join files together. Here we are going to use `cat` to view the content of a file.

Enter command `ls` to know the files and directories. There is a file *readme* in the current working directory which is */home/bandit0*. The *readme* file stores the password for level 1. The password is displayed on the terminal using command `cat readme` and the password is `****` .

Exit the remote session using command `exit`.

![Bandit Level 0 1](/assets/images/Bandit-Overthewire/bandit_l01_terminal.jpg)

Reference : [The Linux Command Line – A Complete Introduction](https://amzn.to/2PDVmZz)

<br/>

## [Bandit Level 1 → Level 2](https://overthewire.org/wargames/bandit/bandit2.html)

### Level Goal

The password for the next level is stored in a file called **-** located in the home directory.

### Commands you may need to solve this level

ls, cd, cat, file, du, find

### Solution : 

Command to connect remote host : `ssh bandit1@bandit.labs.overthewire.org -p 2220` password is  `****` .

In UNIX and Linux, a filename can start with  – (dash) or can be just – (dash). Above it is given that  the file is called – (dash).

But content of the file – can not be displayed using command `cat –` because it reads from standard input and it is waiting for us to type something.

So to view the content of the file - , the path to the file is prefixed with the filename. 
```
cat ./-
```

The password for the next level is `****` .

![Bandit Level 1 2](/assets/images/Bandit-Overthewire/bandit_l12_terminal.jpg)


Reference: [https://www.cs.ait.ac.th/~on/O/oreilly/unix/upt/ch23_14.htm](https://www.cs.ait.ac.th/~on/O/oreilly/unix/upt/ch23_14.htm)<br/>
[https://unix.stackexchange.com/questions/16357/usage-of-dash-in-place-of-a-filename](https://unix.stackexchange.com/questions/16357/usage-of-dash-in-place-of-a-filename)<br/>

<br/>

## [Bandit Level 2 → Level 3](https://overthewire.org/wargames/bandit/bandit3.html)

### Level Goal

The password for the next level is stored in a file called **spaces in this filename** located in the home directory

### Commands you may need to solve this level

ls, cd, cat, file, du, find

### Solution :
 
Command to connect remote host : `ssh bandit2@bandit.labs.overthewire.org -p 2220` password is  `****` .

When we run the `ls` command we find that the name of the file is *spaces in this filename* means there are spaces in the filename.

When there are spaces in a filename use **\\** after every word. 

“A non-quoted backslash (\\) is the escape character. It preserves the literal value of the next character that follows, with the exception of `<newline>`.”

The command is 
```
cat spaces\ in\ this\ filename
```

Since in that directory there is only file we can also use *tab button*, after typing s, which writes the full name of file which starts with s.

The password for the next level is `****` .

![Bandit Level 2 3](/assets/images/Bandit-Overthewire/bandit_l23_terminal.jpg)


Reference: [https://askubuntu.com/questions/101587/how-do-i-enter-a-file-or-directory-with-special-characters-in-its-name](https://askubuntu.com/questions/101587/how-do-i-enter-a-file-or-directory-with-special-characters-in-its-name)

<br/>

## [Bandit Level 3 → Level 4](https://overthewire.org/wargames/bandit/bandit4.html)

### Level Goal

The password for the next level is stored in a hidden file in the inhere directory.

### Commands you may need to solve this level

ls, cd, cat, file, du, find

### Solution : 

Command to connect remote host : `ssh bandit3@bandit.labs.overthewire.org -p 2220` password is  `****` .

`cd` command is used to change our current working directory. `cd` is followed by the pathname of the desired working directory.

Our current working directory is */home/bandit3* and our desired working directory is */home/bandit3/inhere* . So we can either use command `cd inhere/` or `cd /home/bandit3/inhere/`

It is given that the password is stored in the hidden file and after running command `ls` we do not find any file in the directory. Files whose name starts with a period (**.**) are hidden file and command `ls -a` list all files, even those with names that begin with a period, which are normally not listed (i. e., hidden). So the name of the file is *.hidden* and command `cat .hidden` is used to see the content of the file. The password to the next level is `****` .

![Bandit Level 3 4](/assets/images/Bandit-Overthewire/bandit_l34_terminal.jpg)


Reference : [The Linux Command Line – A Complete Introduction](https://amzn.to/2PDVmZz)

### Next Post

[Bandit Level 4 to Level 8](/Bandit-Level-4-to-Level-9-OverTheWire)<br/>
[Bandit Level 9 to Level 11](/Bandit-Level-9-to-Level-12-OverTheWire)<br/>
[Bandit Level 12 → Level 13](/Bandit-Level-12-Level-13-OverTheWire)<br/>
[Bandit Level 13 to Level 15](/Bandit-Level-13-to-Level-16-OverTheWire)<br/>
[Bandit Level 16 to Level 18](/Bandit-Level-16-to-Level-19-OverTheWire)<br/>
[Bandit Level 19 to Level 20](/Bandit-Level-19-to-Level-20-OverTheWire)<br/>
[Bandit Level 21 to Level 22](/Bandit-Level-21-to-Level-23-OverTheWire)
<br/>
[Bandit Level 23 → Level 24](/Bandit-Level-23-Level-24-OverTheWire)<br/>
[Bandit Level 24 → Level 25](/Bandit-Level-24-Level-25-OverTheWire)<br/>
[Bandit Level 25 to Level 26](/Bandit-Level-25-to-Level-26-OverTheWire)<br/>
[Bandit Level 27 to Level 31](/Bandit-Level-27-to-Level-31-OverTheWire)<br/>
[Bandit Level 32 → Level 33](/Bandit-Level-32-Level-33-OverTheWire)<br/>

