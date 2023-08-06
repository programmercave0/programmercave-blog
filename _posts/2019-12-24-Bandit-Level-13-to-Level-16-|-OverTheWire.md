---
layout: post
title: "Bandit Level 13 to Level 15 | OverTheWire"
description: "Learn linux command by playing Bandit wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Level 13 → Level 14, Level 14 → Level 15 and Level 15 → Level 16. In this post we will learn how to use ssh key instead of password to login in a remote machine. We will learn about Secure Scoket Layer and how to establish a connection to a remote machine on a port. The passwords are hidden, so you have to find the passwords for next level yourself."
author: "Programmercave"
header-img: "/assets/Bandit-Overthewire/overthewire_poster.jpg"
tags:  [Linux, OverTheWire-Bandit, CTF]
date: 2019-12-24
---

Learn linux command by playing [Bandit](https://overthewire.org/wargames/bandit/) wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Level 13 → Level 14, Level 14 → Level 15 and Level 15 → Level 16. 

In this post we will learn how to use ssh key instead of password to login in a remote machine. We will learn about Secure Scoket Layer and how to establish a connection to a remote machine on a port.

The passwords are hidden, so you have to find the passwords for next level yourself.

![Bandit OverTheWire]({{ site.url }}/assets/Bandit-Overthewire/overthewire_poster.jpg){:class="img-responsive"}

### Previous Post

[Bandit Level 0 to Level 3]({{ site.url }}/blog/2019/12/21/Bandit-Level-0-to-Level-5-OverTheWire)<br/>
[Bandit Level 4 to Level 8]({{ site.url }}/blog/2019/12/22/Bandit-Level-4-to-Level-9-OverTheWire)<br/>
[Bandit Level 9 to Level 11]({{ site.url }}/blog/2019/12/24/Bandit-Level-9-to-Level-12-OverTheWire)<br/>
[Bandit Level 12 → Level 12]({{ site.url }}/blog/2019/12/24/Bandit-Level-12-Level-13-OverTheWire)

## [Bandit Level 13 → Level 14](https://overthewire.org/wargames/bandit/bandit14.html)

### Level Goal

The password for the next level is stored in **/etc/bandit_pass/bandit14** and **can only be read by user bandit14**. For this level, you don’t get the next password, but you get a private SSH key that can be used to log into the next level. Note: **localhost** is a hostname that refers to the machine you are working on

## Commands you may need to solve this level

ssh, telnet, nc, openssl, s_client, nmap

## Solution : 

Command to connect remote host : `ssh bandit13@bandit.labs.overthewire.org -p 2220` password is `****` .

To access password we need to login as bandit14 and for that the host is localhost. In the directory */home/bandit13* file sshkey.private contains SSH key to login into bandit14.

`ssh` command is used to login and execute commands on a remote machine. Option `-i` selects a file from which the identity (private key) for RSA or DSA authentication is read and the file `sshkey.private`. The command is
```
ssh -i sshkey.private bandit14@localhost
```
The password is stored in */etc/bandit_pass/bandit14*.

The command is `cat /etc/bandit_pass/bandit14` and the password is `****` .

![Bandit Level 13 14]({{ site.url }}/assets/Bandit-Overthewire/bandit_l1314_terminal.jpg){:class="img-responsive"}

Reference : [https://linux.die.net/man/1/ssh](https://linux.die.net/man/1/ssh)<br/>
[https://support.rackspace.com/how-to/logging-in-with-an-ssh-private-key-on-linuxmac/](https://support.rackspace.com/how-to/logging-in-with-an-ssh-private-key-on-linuxmac/)

{% include ads.html %}<br/>

## [Bandit Level 14 → Level 15](https://overthewire.org/wargames/bandit/bandit15.html)

### Level Goal

The password for the next level can be retrieved by submitting the password of the current level to port 30000 on localhost.

### Commands you may need to solve this level

ssh, telnet, nc, openssl, s_client, nmap

### Solution : 

Command to connect remote host : `ssh bandit14@bandit.labs.overthewire.org -p 2220 password` is `****` .

`netcat` is a simple unix utility which reads and writes data across network connections, using TCP or UDP protocol. 

`nc` host port creates a TCP connection to the given port on the given target host. Your standard input is then sent to the host, and anything that comes back across the connection is sent to your standard output. The command is 
```
nc localhost 30000
```
and then enter password of this level. The password for the next level is `****` .

![Bandit Level 14 15]({{ site.url }}/assets/Bandit-Overthewire/bandit_l1415_terminal.jpg){:class="img-responsive"}

Reference : [https://www.commandlinux.com/man-page/man1/nc.1.html](https://www.commandlinux.com/man-page/man1/nc.1.html)

{% include ads.html %}<br/>

## [Bandit Level 15 → Level 16](https://overthewire.org/wargames/bandit/bandit16.html)

### Level Goal

The password for the next level can be retrieved by submitting the password of the current level to **port 30001 on localhost** using SSL encryption.

**Helpful note: Getting “HEARTBEATING” and “Read R BLOCK”? Use -ign_eof and read the “CONNECTED COMMANDS” section in the manpage. Next to ‘R’ and ‘Q’, the ‘B’ command also works in this version of that command…**

### Commands you may need to solve this level

ssh, telnet, nc, openssl, s_client, nmap

### Solution :

Command to connect remote host : `ssh bandit15@bandit.labs.overthewire.org -p 2220` password is `****` .

OpenSSL comes with a client tool that you can use to connect to a secure server. The tool is similar to telnet or nc, in the sense that it handles the SSL/TLS layer but allows you to fully control the layer that comes next.

To connect to a server, you need to supply a hostname and a port. For example: `$ openssl s_client -connect www.feistyduck.com:443`
 
So our command is 
```
openssl s_client -connect localhost:30001
```
and then enter password for the current level. The password for the next level is `****` .

![Bandit Level 15 16]({{ site.url }}/assets/Bandit-Overthewire/bandit_l1516_terminal.jpg){:class="img-responsive"}

Reference : [https://www.feistyduck.com/library/openssl-cookbook/online/ch-testing-with-openssl.html](https://www.feistyduck.com/library/openssl-cookbook/online/ch-testing-with-openssl.html)

### Next Post

[Bandit Level 16 to Level 18]({{ site.url }}/blog/2019/12/24/Bandit-Level-16-to-Level-19-OverTheWire)<br/>
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
