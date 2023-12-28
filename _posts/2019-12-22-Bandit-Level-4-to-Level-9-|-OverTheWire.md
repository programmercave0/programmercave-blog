---
layout: post
title: "Bandit Level 4 to Level 8 | OverTheWire"
description: "Learn linux command by playing Bandit wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Level 4 → Level 5, Level 5 → Level 6, Level 6 → Level 7, Level 7 → Level 8, and Level 8 → Level 9. In this post we will learn how to find a human readable file with certain size in bytes and with certain user. The passwords are hidden, so you have to find the passwords for next level yourself."
author: "Programmercave"
header-img: "/assets/images/Bandit-Overthewire/overthewire_poster.jpg"
tags:  [Linux, OverTheWire-Bandit, CTF]
date: 2019-12-22
---



Learn linux command by playing [Bandit](https://overthewire.org/wargames/bandit/) wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Level 4 → Level 5, Level 5 → Level 6, Level 6 → Level 7, Level 7 → Level 8, and Level 8 → Level 9. 

In this post we will learn how to find a human readable file with certain size in bytes and with certain user.

The passwords are hidden, so you have to find the passwords for next level yourself.

![Bandit OverTheWire]({{ site.url }}/assets/images/Bandit-Overthewire/overthewire_poster.jpg){:class="img-responsive"}

### Previous Post

[Bandit Level 0 to Level 3]({{ site.url }}/blog/2019/12/21/Bandit-Level-0-to-Level-5-OverTheWire)

## [Bandit Level 4 → Level 5](https://overthewire.org/wargames/bandit/bandit5.html)

### Level Goal

The password for the next level is stored in the only human-readable file in the **inhere** directory.<br/> Tip: if your terminal is messed up, try the “reset” command.

### Commands you may need to solve this level

ls, cd, cat, file, du, find

### Solution : 

Command to connect remote host : `ssh bandit4@bandit.labs.overthewire.org -p 2220` password is  `****` .

`file` command is used to determine a file’s type or what file contains.

In *inhere* directory, there are 10 files *-file00, -file01, …, -file09*. The human-readable file means the content of that file is ASCII and we can find the type of content of a file by running command 
```
file ./-file00 
```
```
file ./-file01
```
...
```
file ./file07
```

We found that *-file07* contains ASCII text.

Instead of checking each file we can use `find` and `xargs` command.

`find` program searches a given directory (and its subdirectories) for files based on a variety of attributes.
 
Command `find . -type f` searches all regular files in the current directory. Current directory is specified by . (dot).

The `xargs` command performs an interesting function. It accepts input from standard input and converts it into an argument list for a specified command. 

Command `find . -type f | xargs file` finds all the regular files in the current directory and `xargs` constructs an argument list for `file` command and then executes it.

The password for the next level is `****` .

![Bandit Level 4 5]({{ site.url }}/assets/images/Bandit-Overthewire/bandit_l45_terminal.jpg){:class="img-responsive"}


Reference : [The Linux Command Line – A Complete Introduction](https://amzn.to/2PDVmZz)<br/>
[https://stackoverflow.com/questions/12654026/how-to-count-all-the-human-readable-files-in-bash](https://stackoverflow.com/questions/12654026/how-to-count-all-the-human-readable-files-in-bash)

{% include ads.html %}<br/>

## [Bandit Level 5 → Level 6](https://overthewire.org/wargames/bandit/bandit6.html)

### Level Goal

The password for the next level is stored in a file somewhere under the **inhere** directory and has all of the following properties:<br/>
 * human-readable 
 * 1033 bytes in size 
 * not executable 

### Commands you may need to solve this level

ls, cd, cat, file, du, find

### Solution : 

Command to connect remote host : `ssh bandit5@bandit.labs.overthewire.org -p 2220` password is  `****` .

In manual pages of `find` command it is mentioned that option `-size` is used to specify size of the file and `c` is used for bytes. `-executable` matches the executable files, so `! -executables` matches the non executable files. We then pipe this output to `xargs` command which tells the content of the file found. Command is 

```
find . -type f -size 1033c ! - executable | xargs file
```

The password for the next level is `****` .

![Bandit Level 5 6]({{ site.url }}/assets/images/Bandit-Overthewire/bandit_l56_terminal.jpg){:class="img-responsive"}


Reference : [The Linux Command Line – A Complete Introduction](https://amzn.to/2PDVmZz)<br/>
[http://man7.org/linux/man-pages/man1/find.1.html](http://man7.org/linux/man-pages/man1/find.1.html)

{% include ads.html %}<br/>

## [Bandit Level 6 → Level 7](https://overthewire.org/wargames/bandit/bandit7.html)

### Level Goal

The password for the next level is **stored somewhere on the server** and has all of the following properties:<br/>
 * owned by user bandit7 
 * owned by group bandit6 
 * 33 bytes in size 

### Commands you may need to solve this level

ls, cd, cat, file, du, find, grep

### Solution :
 
Command to connect remote host : `ssh bandit6@bandit.labs.overthewire.org -p 2220` password is `****` .

Since the password is stored *somewhere on the server*. Lets go to the root directory by running command `cd ..` two times.

From manual page of `find` command :

 `-user` *uname*<br/>
    File is owned by user uname (numeric user ID allowed).<br/>
 `-size` *n[cwbkMG]*<br/>
    File uses n units of space, rounding up.  The following suffixes can be used:<br/>
    'c'    for bytes<br/>
 `-group` *gname*
    File belongs to group gname (numeric group ID allowed).<br/>

From the given, `uname` is *bandit7*, `gname` is *bandit6* and `n` is *33c*.

We want file with ASCII content and `xargs file` will tell us about that.

So the command is :
```
find -user bandit7 -group bandit6 -size 33c | xargs file
```

The password for the next level is `****` .

![Bandit Level 6 7]({{ site.url }}/assets/images/Bandit-Overthewire/bandit_l67_terminal.jpg){:class="img-responsive"}

Reference : [The Linux Command Line – A Complete Introduction](https://amzn.to/2PDVmZz)<br/>
[http://man7.org/linux/man-pages/man1/find.1.html](http://man7.org/linux/man-pages/man1/find.1.html)

{% include ads.html %}<br/>

## [Bandit Level 7 → Level 8](https://overthewire.org/wargames/bandit/bandit8.html)

### Level Goal

The password for the next level is stored in the file **data.txt** next to the word **millionth**

### Commands you may need to solve this level

grep, sort, uniq, strings, base64, tr, tar, gzip, bzip2, xxd

### Solution :
 
Command to connect remote host : `ssh bandit7@bandit.labs.overthewire.org -p 2220` password is `****` .

Here we can use `grep` program. `grep` is used to find text patterns within file. The text we have to find is *millionth* and the password for next level is next to it.

The command is :
```
cat data.txt | grep millionth
```
The password for next level is `****` .

![Bandit Level 7 8]({{ site.url }}/assets/images/Bandit-Overthewire/bandit_l78_terminal.jpg){:class="img-responsive"}


Reference : [The Linux Command Line – A Complete Introduction](https://amzn.to/2PDVmZz)<br/>

{% include ads.html %}<br/>

## [Bandit Level 8 → Level 9](https://overthewire.org/wargames/bandit/bandit9.html)

### Level Goal

The password for the next level is stored in the file **data.txt** and is the only line of text that occurs only once

### Commands you may need to solve this level

grep, sort, uniq, strings, base64, tr, tar, gzip, bzip2, xxd

### Solution :
 
Command to connect remote host : `ssh bandit8@bandit.labs.overthewire.org -p 2220` password is `****` .

We will use command `sort` to sort all texts in the file. Command `uniq` with option `u` i.e. `uniq -u` only prints unique lines. So the command is :
```
sort data.txt | uniq -u
```

and password is `****` .

![Bandit Level 8 9]({{ site.url }}/assets/images/Bandit-Overthewire/bandit_l89_terminal.jpg){:class="img-responsive"}


Reference : [http://man7.org/linux/man-pages/man1/uniq.1.html](http://man7.org/linux/man-pages/man1/uniq.1.html)<br/>
[https://askubuntu.com/questions/915570/how-do-i-find-a-single-unique-line-in-a-file](https://askubuntu.com/questions/915570/how-do-i-find-a-single-unique-line-in-a-file)

### Next Post

[Bandit Level 9 to Level 11]({{ site.url }}/blog/2019/12/24/Bandit-Level-9-to-Level-12-OverTheWire)<br/>
[Bandit Level 12 → Level 13]({{ site.url }}/blog/2019/12/24/Bandit-Level-12-Level-13-OverTheWire)<br/>
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



