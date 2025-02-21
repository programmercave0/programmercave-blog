---
author: Programmercave
date: "2019-12-24T00:00:00Z"
description: Learn linux command by playing Bandit wargame. The Bandit wargame is
  aimed at absolute beginners. It will teach the basics needed to be able to play
  other wargames. Below is the solution of Level 9 → Level 10, Level 10 → Level 11
  and Level 11 → Level 12. In this post we will learn how to find a sting of certain
  pattern in a text file. How to decode a file from certain encoding and how to transform
  string in command line. The passwords are hidden, so you have to find the passwords
  for next level yourself.
header-img: /assets/images/Bandit-Overthewire/overthewire_poster.jpg
tags:
- Linux
- OverTheWire-Bandit
- CTF
title: Bandit Level 9 to Level 11 | OverTheWire
---



Learn linux command by playing [Bandit](https://overthewire.org/wargames/bandit/) wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Level 9 → Level 10, Level 10 → Level 11 and Level 11 → Level 12. 

In this post we will learn how to find a sting of certain pattern in a text file. How to decode a file from certain encoding and how to transform string in command line.

The passwords are hidden, so you have to find the passwords for next level yourself.

![Bandit OverTheWire](/assets/images/Bandit-Overthewire/overthewire_poster.jpg)

### Previous Post

[Bandit Level 0 to Level 3](/Bandit-Level-0-to-Level-5-OverTheWire)<br/>
[Bandit Level 4 to Level 8](/Bandit-Level-4-to-Level-9-OverTheWire)

## [Bandit Level 9 → Level 10](https://overthewire.org/wargames/bandit/bandit10.html)

### Level Goal

The password for the next level is stored in the file **data.txt** in one of the few human-readable strings, beginning with several ‘=’ characters.

### Commands you may need to solve this level
grep, sort, uniq, strings, base64, tr, tar, gzip, bzip2, xxd

### Solution : 
Command to connect remote host : `ssh bandit9@bandit.labs.overthewire.org -p 2220` password is `****` .

Command `strings` print the sequences of printable characters in files and `grep =` will print lines with ‘=’ in it. The command is :
```
strings data.txt | grep =
```

and the password for next level is `****` .

![Bandit Level 9 10](/assets/images/Bandit-Overthewire/bandit_l910_terminal.jpg)

Reference : [http://man7.org/linux/man-pages/man1/strings.1.html](http://man7.org/linux/man-pages/man1/strings.1.html)

<br/>

## [Bandit Level 10 → Level 11](https://overthewire.org/wargames/bandit/bandit11.html)

### Level Goal

The password for the next level is stored in the file **data.txt**, which contains base64 encoded data

### Commands you may need to solve this level

grep, sort, uniq, strings, base64, tr, tar, gzip, bzip2, xxd

### Solution :
 
Command to connect remote host : `ssh bandit10@bandit.labs.overthewire.org -p 2220` password is `****` .

Command `base64` with option `-d` decodes the file. So the command is
``` 
base64 -d data.txt
```

and the password is `****` .

![Bandit Level 10 11](/assets/images/Bandit-Overthewire/bandit_l1011_terminal.jpg)

Reference : [https://linux.die.net/man/1/base64](https://linux.die.net/man/1/base64)

<br/>

## [Bandit Level 11 → Level 12](https://overthewire.org/wargames/bandit/bandit12.html)

### Level Goal

The password for the next level is stored in the file **data.txt**, where all lowercase (a-z) and uppercase (A-Z) letters have been rotated by 13 positions

### Commands you may need to solve this level

grep, sort, uniq, strings, base64, tr, tar, gzip, bzip2, xxd

### Solution : 

Command to connect remote host : `ssh bandit11@bandit.labs.overthewire.org -p 2220` password is `****` .

Command `tr` is used to translate character. Since characters are rotated by 13 positions means *a* becomes *n* and *n* becomes *a*, *b* becomes *o* and *o* becomes *b* and so on. The command is
``` 
cat data.txt | tr "[a-zA-Z]" "[n-za-mN-ZA-M]"
```

and the password for the next level is `****` .

![Bandit Level 11 12](/assets/images/Bandit-Overthewire/bandit_l1112_terminal.jpg)

Reference : [http://man7.org/linux/man-pages/man1/tr.1p.html](http://man7.org/linux/man-pages/man1/tr.1p.html)<br/>
[https://unix.stackexchange.com/questions/19772/how-does-tr-a-z-n-za-m-work](https://unix.stackexchange.com/questions/19772/how-does-tr-a-z-n-za-m-work)

### Next Post

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







