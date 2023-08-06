---
layout: post
title: "Bandit Level 12 → Level 13 | OverTheWire"
description: "Learn linux command by playing Bandit wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Level 12 → Level 13. In this post we will learn about various compression techniques and how to decompress file. We will learn how to convert binary to hex file and vice-versa. The passwords are hidden, so you have to find the passwords for next level yourself."
author: "Programmercave"
header-img: "/assets/Bandit-Overthewire/overthewire_poster.jpg"
tags:  [Linux, OverTheWire-Bandit, CTF]
date: 2019-12-24
---
* toc
{:toc}

Learn linux command by playing [Bandit](https://overthewire.org/wargames/bandit/) wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Level 12 → Level 13. 

In this post we will learn about various compression techniques and how to decompress file. We will learn how to convert binary to hex file and vice-versa.

The passwords are hidden, so you have to find the passwords for next level yourself.

![Bandit OverTheWire]({{ site.url }}/assets/Bandit-Overthewire/overthewire_poster.jpg){:class="img-responsive"}

### Previous Post

[Bandit Level 0 to Level 3]({{ site.url }}/blog/2019/12/21/Bandit-Level-0-to-Level-5-OverTheWire)<br/>
[Bandit Level 4 to Level 8]({{ site.url }}/blog/2019/12/22/Bandit-Level-4-to-Level-9-OverTheWire)<br/>
[Bandit Level 9 to Level 11]({{ site.url }}/blog/2019/12/24/Bandit-Level-9-to-Level-12-OverTheWire)

## [Bandit Level 12 → Level 13](https://overthewire.org/wargames/bandit/bandit13.html)

### Level Goal

The password for the next level is stored in the file **data.txt**, which is a hexdump of a file that has been repeatedly compressed. For this level it may be useful to create a directory under /tmp in which you can work using mkdir. For example: mkdir /tmp/myname123. Then copy the datafile using cp, and rename it using mv (read the manpages!)

### Commands you may need to solve this level

grep, sort, uniq, strings, base64, tr, tar, gzip, bzip2, xxd, mkdir, cp, mv, file

### Solution : 

Command to connect remote host : `ssh bandit12@bandit.labs.overthewire.org -p 2220` password is `****` .

As mentioned in question make a new directory in /tmp and rename the file.

![Bandit Level 12 13]({{ site.url }}/assets/Bandit-Overthewire/bandit_l1213_terminal1.jpg){:class="img-responsive"}

`xxd` program is used to make a hexdump or to do the reverse. Option `-r` convert hexdump into the binary. File `myfile.txt` is a hexdump and convert it into a binary file `myfile1.bin` using command
```
xxd -r myfile.txt > myfile1.bin
```

![Bandit Level 12 13]({{ site.url }}/assets/Bandit-Overthewire/bandit_l1213_terminal2.jpg){:class="img-responsive"}

Using command `file myfile1.bin` , we found that `myfile1.bin` is a *gzip compressed data*.

`zcat` is a program supplied with `gzip` and is used to decompress *gzip compressed files*.
```
zcat myfile1.bin > myfile2
```

![Bandit Level 12 13]({{ site.url }}/assets/Bandit-Overthewire/bandit_l1213_terminal3.jpg){:class="img-responsive"}

Again using `file` command on `myfile2`, we found that it is *bzip2 compressed data*.

`bzcat` program is supplied with `bzip2` and is used to decompress *bzip2 compressed files*.
```
bzcat myfile2 > myfile3
```

![Bandit Level 12 13]({{ site.url }}/assets/Bandit-Overthewire/bandit_l1213_terminal4.jpg){:class="img-responsive"}


`myfile3` is *gzip compressed file* so use `zcat` program to decompress it in `myfile4`. `myfile4` is a POSIX tar archive.

`tar` program is used for archiving file and options `x` is used to extract an archive, `f` is used to specify name of the tar archive and `v` is used for more detailed listing.
```
tar -xvf myfile4
```

This command outputs file `data5.bin` which is again a *tar archive*. Again use `tar` program on `data5.bin` which outputs `data6.bin`. `data6.bin` is a *bzip2 compressed file* and use `bzcat` program to decompress it to `myfile7`.
 
`myfile7` is a *tar archive* and use tar program which outputs `data8.bin`. `data8.bin` is a *gzip compressed file* and use `zcat` to decompress it to file `myfile9`.

`myfile9` contains *ASCII text* and `cat myfile9` tells the password for the next level.

The password for the next level is `****` .

![Bandit Level 12 13]({{ site.url }}/assets/Bandit-Overthewire/bandit_l1213_terminal5.jpg){:class="img-responsive"}

Reference : [The Linux Command Line – A Complete Introduction](https://amzn.to/2PDVmZz)<br/>
[https://linux.die.net/man/1/xxd](https://linux.die.net/man/1/xxd)

### Next Post

[Bandit Level 13 to Level 15]({{ site.url }}/blog/2019/12/24/Bandit-Level-13-to-Level-16-OverTheWire)<br/>
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
