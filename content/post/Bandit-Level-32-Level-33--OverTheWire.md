---
author: Programmercave
date: "2019-12-26T00:00:00Z"
description: Learn linux command by playing Bandit wargame. The Bandit wargame is
  aimed at absolute beginners. It will teach the basics needed to be able to play
  other wargames. Below is the solution of Bandit Level 32 → Level 33. The passwords
  are hidden, so you have to find the passwords for next level yourself.
header-img: /assets/images/Bandit-Overthewire/overthewire_poster.jpg
tags:
- Linux
- OverTheWire-Bandit
- CTF
title: Bandit Level 32 → Level 33 | OverTheWire
---



Learn linux command by playing [Bandit](https://overthewire.org/wargames/bandit/) wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Bandit Level 32 → Level 33. 

The passwords are hidden, so you have to find the passwords for next level yourself.

![Bandit OverTheWire](/assets/images/Bandit-Overthewire/overthewire_poster.jpg)

### Previous Post

[Bandit Level 0 to Level 3](/Bandit-Level-0-to-Level-5-OverTheWire)<br/>
[Bandit Level 4 to Level 8](/Bandit-Level-4-to-Level-9-OverTheWire)<br/>
[Bandit Level 9 to Level 11](/Bandit-Level-9-to-Level-12-OverTheWire)<br/>
[Bandit Level 12 → Level 13](/Bandit-Level-12-Level-13-OverTheWire)<br/>
[Bandit Level 13 to Level 15](/Bandit-Level-13-to-Level-16-OverTheWire)<br/>
[Bandit Level 16 to Level 18](/Bandit-Level-16-to-Level-19-OverTheWire)<br/>
[Bandit Level 19 to Level 20](/Bandit-Level-19-to-Level-20-OverTheWire)<br/>
[Bandit Level 21 to Level 22](/Bandit-Level-21-to-Level-23-OverTheWire)<br/>
[Bandit Level 23 → Level 24](/Bandit-Level-23-Level-24-OverTheWire)<br/>
[Bandit Level 24 → Level 25](/Bandit-Level-24-Level-25-OverTheWire)<br/>
[Bandit Level 25 to Level 26](/Bandit-Level-25-to-Level-26-OverTheWire)<br/>
[Bandit Level 27 to Level 31](/Bandit-Level-27-to-Level-31-OverTheWire)

## [Bandit Level 32 → Level 33](https://overthewire.org/wargames/bandit/bandit33.html)

After all this git stuff its time for another escape. Good luck!

### Commands you may need to solve this level

sh, man

### Solution : 

Command to connect remote host : `ssh bandit32@bandit.labs.overthewire.org -p 2220` password is `****` .

![Bandit Level 32 33](/assets/images/Bandit-Overthewire/bandit_l3233_terminal1.jpg)

THE UPPERCASE SHELL is converting every command into uppercase.                  

`$0` expands to the name of the shell or shell script. This is set at shell initialization. If bash is invoked with a file of commands, `$0` is set to the name of that file.

`$0` helps us to come out of the interactive mode.

After running `ls -la` we found that uppershell has escalated privilege as user bandit33. `echo $SHELL` tells us the current shell which is uppershell. The command is
``` 
cat /etc/bandit_pass/bandit33 
```

and the password is `****` .

![Bandit Level 32 33](/assets/images/Bandit-Overthewire/bandit_l3233_terminal2.jpg)

Reference : [https://bash.cyberciti.biz/guide/$0](https://bash.cyberciti.biz/guide/$0)<br/>
[https://unix.stackexchange.com/a/280458/244874](https://unix.stackexchange.com/a/280458/244874)<br/>

<br/>

## Bandit Level 33 → Level 34

At this moment, level 34 does not exist yet.


