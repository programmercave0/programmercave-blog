---
layout: post
title: "Bandit Level 21 to Level 22 | OverTheWire"
description: "Learn linux command by playing Bandit wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Level 21 → Level 22 and Level 22 → Level 23. In this post we will learn about crontab files. The passwords are hidden, so you have to find the passwords for next level yourself."
author: "Programmercave"
header-img: "/assets/Bandit-Overthewire/overthewire_poster.jpg"
tags:  [Linux, OverTheWire-Bandit, CTF]
date: 2019-12-25
---

Learn linux command by playing [Bandit](https://overthewire.org/wargames/bandit/) wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Level 21 → Level 22, Level 22 → Level 23 and and Level 23 → Level 24. 

In this post we will learn about crontab files.

The passwords are hidden, so you have to find the passwords for next level yourself.

![Bandit OverTheWire]({{ site.url }}/assets/Bandit-Overthewire/overthewire_poster.jpg){:class="img-responsive"}

### Previous Post

[Bandit Level 0 to Level 3]({{ site.url }}/blog/2019/12/21/Bandit-Level-0-to-Level-5-OverTheWire)<br/>
[Bandit Level 4 to Level 8]({{ site.url }}/blog/2019/12/22/Bandit-Level-4-to-Level-9-OverTheWire)<br/>
[Bandit Level 9 to Level 11]({{ site.url }}/blog/2019/12/24/Bandit-Level-9-to-Level-12-OverTheWire)<br/>
[Bandit Level 12 → Level 13]({{ site.url }}/blog/2019/12/24/Bandit-Level-12-Level-13-OverTheWire)<br/>
[Bandit Level 13 to Level 15]({{ site.url }}/blog/2019/12/24/Bandit-Level-13-to-Level-16-OverTheWire)<br/>
[Bandit Level 16 to Level 18]({{ site.url }}/blog/2019/12/24/Bandit-Level-16-to-Level-19-OverTheWire)<br/>
[Bandit Level 19 to Level 20]({{ site.url }}/blog/2019/12/25/Bandit-Level-19-to-Level-20-OverTheWire)

## [Bandit Level 21 → Level 22](https://overthewire.org/wargames/bandit/bandit22.html)

### Level Goal

A program is running automatically at regular intervals from **cron**, the time-based job scheduler. Look in **/etc/cron.d/** for the configuration and see what command is being executed.

### Commands you may need to solve this level

cron, crontab, crontab(5) (use “man 5 crontab” to access this)

### Solution : 

Command to connect remote host : `ssh bandit21@bandit.labs.overthewire.org -p 2220` password is `****` .

Cron is a system daemon used to execute desired tasks (in the background) at designated times. 

A *crontab* file contains instructions for the cron(8) daemon in the following simplified manner: "run this command at this time on this date".  Each user can define their own crontab.  Commands defined in any given crontab are executed under the user who owns that particular crontab.

In directory */etc/crond.d/*, file important to us is *cronjob_bandit22*. It is a crontab file and contain commands. The content can be viewed using `cat cronjob_bandit22`.

![Bandit Level 21 22]({{ site.url }}/assets/Bandit-Overthewire/bandit_l2122_terminal1.jpg){:class="img-responsive"}

The command is executing the file *cronjob_bandit22.sh* which is in the directory */usr/bin/*. Lets see the content of file *cronjob_bandit22.sh* using command `cat /usr/bin/cronjob_bandit22.sh`.

![Bandit Level 21 22]({{ site.url }}/assets/Bandit-Overthewire/bandit_l2122_terminal2.jpg){:class="img-responsive"}

Now we know that file */tmp/t7O6lds9S0RqQh9aMcz6ShpAoZKF7fgv* contains the password for the next level. We can view the password using command `cat /tmp/t7O6lds9S0RqQh9aMcz6ShpAoZKF7fgv` and the password is `****` .

Reference : [http://man7.org/linux/man-pages/man5/crontab.5.html](http://man7.org/linux/man-pages/man5/crontab.5.html)<br/>
[https://help.ubuntu.com/community/CronHowto](https://help.ubuntu.com/community/CronHowto)<br/>

{% include ads.html %}<br/>

## [Bandit Level 22 → Level 23](https://overthewire.org/wargames/bandit/bandit23.html)

### Level Goal

A program is running automatically at regular intervals from **cron**, the time-based job scheduler. Look in **/etc/cron.d/** for the configuration and see what command is being executed.

**NOTE:** Looking at shell scripts written by other people is a very useful skill. The script for this level is intentionally made easy to read. If you are having problems understanding what it does, try executing it to see the debug information it prints.

## Commands you may need to solve this level

cron, crontab, crontab(5) (use “man 5 crontab” to access this)

## Solution : 

Command to connect remote host : `ssh bandit22@bandit.labs.overthewire.org -p 2220` password is `****` .

The initial part is same as the previous level. The commands executed are :
```
cd /etc/cron.d/
ls
cat cronjob_bandit23
```

![Bandit Level 22 23]({{ site.url }}/assets/Bandit-Overthewire/bandit_l2223_terminal1.jpg){:class="img-responsive"}

File *cronjob_bandit23.sh* copies password of bandit22 in another file and it is readable and executable by user bandit22. 
```
-rwxr-x--- 1 bandit23 bandit22 211 Oct 16  2018 cronjob_bandit23.sh
```

So we cannot modify it. But we want to copy password of bandit23 instead of bandit22.

![Bandit Level 22 23]({{ site.url }}/assets/Bandit-Overthewire/bandit_l2223_terminal2.jpg){:class="img-responsive"}

After executing *cronjob_bandit23.sh* using `./cronjob_bandit23.sh` it generates a file in *tmp* directory and the name of the file is md5 hash containing bandit22.
 
If we change bandit23 to bandit23 then new hash is generated which is the name of new file in tmp directory which may contain password for level 23.

To generate new hash run `echo I am user bandit23 | md5sum | cut -d ' ' -f 1` on terminal. The new hash is *8ca319486bfbbc3663ea0fbe81326349* and this is the name of file in *tmp* directory.

To see password run `cat /tmp/8ca319486bfbbc3663ea0fbe81326349` and the password is `****` .

![Bandit Level 22 23]({{ site.url }}/assets/Bandit-Overthewire/bandit_l2223_terminal3.jpg){:class="img-responsive"}

### Next Post

[Bandit Level 23 → Level 24]({{ site.url }}/blog/2019/12/25/Bandit-Level-23-Level-24-OverTheWire)<br/>
[Bandit Level 24 → Level 25]({{ site.url }}/blog/2019/12/26/Bandit-Level-24-Level-25-OverTheWire)<br/>
[Bandit Level 25 to Level 26]({{ site.url }}/blog/2019/12/26/Bandit-Level-25-to-Level-26-OverTheWire)<br/>
[Bandit Level 27 to Level 31]({{ site.url }}/blog/2019/12/26/Bandit-Level-27-to-Level-31-OverTheWire)<br/>
[Bandit Level 32 → Level 33]({{ site.url }}/blog/2019/12/26/Bandit-Level-32-Level-33-OverTheWire)<br/>

### Other Wargames
[Leviathan Wargame from OverTheWire All Level Solutions]({{ site.url }}/blog/#overthewire-leviathan)<br/> 
[Krypton Wargame from OverTheWire All Level Solutions]({{ site.url }}/blog/#overthewire-krypton)<br/>
