---
layout: post
title: "Bandit Level 16 to Level 18 | OverTheWire"
description: "Learn linux command by playing Bandit wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Level 16 → Level 17, Level 17 → Level 18 and Level 18 → Level 19. In this post we will learn how to scan for open ports and how to private key to login in a remote machine. We will learn how to find difference in two text files and how to use psuedo terminals. The passwords are hidden, so you have to find the passwords for next level yourself."
author: "Programmercave"
header-img: "/assets/Bandit-Overthewire/overthewire_poster.jpg"
tags:  [Linux, OverTheWire-Bandit, CTF]
date: 2019-12-24
---

Learn linux command by playing [Bandit](https://overthewire.org/wargames/bandit/) wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Level 16 → Level 17, Level 17 → Level 18 and Level 18 → Level 19. 

In this post we will learn how to scan for open ports and how to private key to login in a remote machine. We will learn how to find difference in two text files and how to use psuedo terminals. 

The passwords are hidden, so you have to find the passwords for next level yourself.

![Bandit OverTheWire]({{ site.url }}/assets/Bandit-Overthewire/overthewire_poster.jpg){:class="img-responsive"}

### Previous Post

[Bandit Level 0 to Level 3]({{ site.url }}/blog/2019/12/21/Bandit-Level-0-to-Level-5-OverTheWire)<br/>
[Bandit Level 4 to Level 8]({{ site.url }}/blog/2019/12/22/Bandit-Level-4-to-Level-9-OverTheWire)<br/>
[Bandit Level 9 to Level 11]({{ site.url }}/blog/2019/12/24/Bandit-Level-9-to-Level-12-OverTheWire)<br/>
[Bandit Level 12 → Level 13]({{ site.url }}/blog/2019/12/24/Bandit-Level-12-Level-13-OverTheWire)<br/>
[Bandit Level 13 to Level 15]({{ site.url }}/blog/2019/12/24/Bandit-Level-13-to-Level-16-OverTheWire)

## [Bandit Level 16 → Level 17](https://overthewire.org/wargames/bandit/bandit17.html)

### Level Goal

The credentials for the next level can be retrieved by submitting the password of the current level to **a port on localhost in the range 31000 to 32000**. First find out which of these ports have a server listening on them. Then find out which of those speak SSL and which don’t. There is only 1 server that will give the next credentials, the others will simply send back to you whatever you send to it.

### Commands you may need to solve this level

ssh, telnet, nc, openssl, s_client, nmap

### Solution :

Command to connect remote host : `ssh bandit16@bandit.labs.overthewire.org -p 2220` password is `****` .

Nmap (“Network Mapper”) will help us to scan for ports. Option `-p` provides the port ranges to the tool. So our command is
```
nmap localhost -p31000-32000 
```

![Bandit Level 16 17]({{ site.url }}/assets/Bandit-Overthewire/bandit_l1617_terminal1.jpg){:class="img-responsive"}

We found two ports and port 31790 is open. Open port means server on that port is listening.

Now to connect to port 31790 we will use `openssl` with `s_client` tool, which we have done in previous challenge.
```
openssl s_client -connect localhost:31790
```

and the we enter the password for the current level. The output we received is a RSA private key.
```
-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAvmOkuifmMg6HL2YPIOjon6iWfbp7c3jx34YkYWqUH57SUdyJ
imZzeyGC0gtZPGujUSxiJSWI/oTqexh+cAMTSMlOJf7+BrJObArnxd9Y7YT2bRPQ
Ja6Lzb558YW3FZl87ORiO+rW4LCDCNd2lUvLE/GL2GWyuKN0K5iCd5TbtJzEkQTu
DSt2mcNn4rhAL+JFr56o4T6z8WWAW18BR6yGrMq7Q/kALHYW3OekePQAzL0VUYbW
JGTi65CxbCnzc/w4+mqQyvmzpWtMAzJTzAzQxNbkR2MBGySxDLrjg0LWN6sK7wNX
x0YVztz/zbIkPjfkU1jHS+9EbVNj+D1XFOJuaQIDAQABAoIBABagpxpM1aoLWfvD
KHcj10nqcoBc4oE11aFYQwik7xfW+24pRNuDE6SFthOar69jp5RlLwD1NhPx3iBl
J9nOM8OJ0VToum43UOS8YxF8WwhXriYGnc1sskbwpXOUDc9uX4+UESzH22P29ovd
d8WErY0gPxun8pbJLmxkAtWNhpMvfe0050vk9TL5wqbu9AlbssgTcCXkMQnPw9nC
YNN6DDP2lbcBrvgT9YCNL6C+ZKufD52yOQ9qOkwFTEQpjtF4uNtJom+asvlpmS8A
vLY9r60wYSvmZhNqBUrj7lyCtXMIu1kkd4w7F77k+DjHoAXyxcUp1DGL51sOmama
+TOWWgECgYEA8JtPxP0GRJ+IQkX262jM3dEIkza8ky5moIwUqYdsx0NxHgRRhORT
8c8hAuRBb2G82so8vUHk/fur85OEfc9TncnCY2crpoqsghifKLxrLgtT+qDpfZnx
SatLdt8GfQ85yA7hnWWJ2MxF3NaeSDm75Lsm+tBbAiyc9P2jGRNtMSkCgYEAypHd

...
```

Save this key locally on your computer with name bandit17.key.

![Bandit Level 16 17]({{ site.url }}/assets/Bandit-Overthewire/bandit_l1617_terminal2.jpg){:class="img-responsive"}


In level 13 we login using a ssh private key. We will try to do it here with command
```
ssh -i bandit17.key bandit17@bandit.labs.overthewire.org -p 2220
```

But we received *warning: unprotected private key file* because file was readable to group and to the world and writable by the owner. We change to read-only mode to the owner using command  
```
sudo chmod 400 bandit17.key
```

Now we can login into bandit17.

![Bandit Level 16 17]({{ site.url }}/assets/Bandit-Overthewire/bandit_l1617_terminal3.jpg){:class="img-responsive"}


Reference : [https://www.feistyduck.com/library/openssl-cookbook/online/ch-testing-with-openssl.html](https://www.feistyduck.com/library/openssl-cookbook/online/ch-testing-with-openssl.html)<br/>
[The Linux Command Line – A Complete Introduction](https://amzn.to/2PDVmZz)<br/>
[https://linux.die.net/man/1/nmap](https://linux.die.net/man/1/nmap)<br/>

{% include ads.html %}<br/>

## [Bandit Level 17 → Level 18](https://overthewire.org/wargames/bandit/bandit18.html)

### Level Goal

There are 2 files in the homedirectory: **passwords.old** and **passwords.new**. The password for the next level is in **passwords.new** and is the only line that has been changed between **passwords.old** and **passwords.new**

**NOTE: if you have solved this level and see ‘Byebye!’ when trying to log into bandit18, this is related to the next level, bandit19**

### Commands you may need to solve this level

cat, grep, ls, diff

### Solution :

Command to login `ssh -i bandit17.key bandit17@bandit.labs.overthewire.org -p 2220`

`diff` program compare files line by line and option `--normal` output a normal diff.

Command to find difference between passwords.new and passwords.old
```
diff --normal passwords.new passwords.old 
```

and the difference in passwords.new is `****` . 

![Bandit Level 17 18]({{ site.url }}/assets/Bandit-Overthewire/bandit_l1718_terminal.jpg){:class="img-responsive"}

Reference : [http://man7.org/linux/man-pages/man1/diff.1.html](http://man7.org/linux/man-pages/man1/diff.1.html)

{% include ads.html %}<br/>

## [Bandit Level 18 → Level 19](https://overthewire.org/wargames/bandit/bandit19.html)

### Level Goal

The password for the next level is stored in a file **readme** in the homedirectory. Unfortunately, someone has modified **.bashrc** to log you out when you log in with SSH.

### Commands you may need to solve this level

ssh, ls, cat

### Solution : 

After logging with command `ssh bandit18@bandit.labs.overthewire.org -p 2220` password `****` .

We receive Byebye ! as the output. In the question it is mentioned that someone has modified **.bashrc** to log you out when you log in with SSH.

With option `-t` in `ssh` command we can force psuedo-tty allocation. 

"Pseudo Terminals" emulates Terminal hardware, handling input and output in the same way a physical device would so that the software connected is not aware there's not a real device attached.

We have forced a psuedo terminal and we know that password is in the readme file, so we can view password using command.
```
ssh -t bandit18@bandit.labs.overthewire.org -p 2220 cat readme
```

and the password for the next level is `***` .

![Bandit Level 18 19]({{ site.url }}/assets/Bandit-Overthewire/bandit_l1819_terminal.jpg){:class="img-responsive"}

Reference : [https://linux.die.net/man/1/ssh](https://linux.die.net/man/1/ssh)<br/>
[https://unix.stackexchange.com/a/21150/244874](https://unix.stackexchange.com/a/21150/244874)<br/>
[https://serverfault.com/a/201158](https://serverfault.com/a/201158)

### Next Post

[Bandit Level 19 to Level 20]({{ site.url }}/blog/2019/12/25/Bandit-Level-19-to-Level-20-OverTheWire)<br/>
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
