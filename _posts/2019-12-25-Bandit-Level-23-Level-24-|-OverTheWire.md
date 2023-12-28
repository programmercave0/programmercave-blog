---
layout: post
title: "Bandit Level 23 → Level 24 | OverTheWire"
description: "Learn linux command by playing Bandit wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Bandit Level 23 → Level 24.In this level we will learn about corntab files and write our first shell-script in this series. The passwords are hidden, so you have to find the passwords for next level yourself."
author: "Programmercave"
header-img: "/assets/images/Bandit-Overthewire/overthewire_poster.jpg"
tags:  [Linux, OverTheWire-Bandit, CTF]
date: 2019-12-25
---



Learn linux command by playing [Bandit](https://overthewire.org/wargames/bandit/) wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Bandit Level 23 → Level 24. 

In this level we will learn about corntab files and write our first shell-script in this series. 

The passwords are hidden, so you have to find the passwords for next level yourself.

![Bandit OverTheWire]({{ site.url }}/assets/images/Bandit-Overthewire/overthewire_poster.jpg){:class="img-responsive"}

### Previous Post

[Bandit Level 0 to Level 3]({{ site.url }}/blog/2019/12/21/Bandit-Level-0-to-Level-5-OverTheWire)<br/>
[Bandit Level 4 to Level 8]({{ site.url }}/blog/2019/12/22/Bandit-Level-4-to-Level-9-OverTheWire)<br/>
[Bandit Level 9 to Level 11]({{ site.url }}/blog/2019/12/24/Bandit-Level-9-to-Level-12-OverTheWire)<br/>
[Bandit Level 12 → Level 13]({{ site.url }}/blog/2019/12/24/Bandit-Level-12-Level-13-OverTheWire)<br/>
[Bandit Level 13 to Level 15]({{ site.url }}/blog/2019/12/24/Bandit-Level-13-to-Level-16-OverTheWire)<br/>
[Bandit Level 16 to Level 18]({{ site.url }}/blog/2019/12/24/Bandit-Level-16-to-Level-19-OverTheWire)<br/>
[Bandit Level 19 to Level 20]({{ site.url }}/blog/2019/12/25/Bandit-Level-19-to-Level-20-OverTheWire)<br/>
[Bandit Level 21 to Level 22]({{ site.url }}/blog/2019/12/25/Bandit-Level-21-to-Level-23-OverTheWire)

## [Bandit Level 23 → Level 24](https://overthewire.org/wargames/bandit/bandit24.html)

### Level Goal

A program is running automatically at regular intervals from **cron**, the time-based job scheduler. Look in **/etc/cron.d/** for the configuration and see what command is being executed.

**NOTE:** This level requires you to create your own first shell-script. This is a very big step and you should be proud of yourself when you beat this level!

**NOTE 2:** Keep in mind that your shell script is removed once executed, so you may want to keep a copy around…

### Commands you may need to solve this level

cron, crontab, crontab(5) (use “man 5 crontab” to access this)

### Solution : 

Command to connect remote host : `ssh bandit23@bandit.labs.overthewire.org -p 2220` password is `****` .

Execute these commands to know the content of crontab file and then content of the script file.
```
cd /etc/cron.d/
ls
cat cronjob_bandit24
cat /usr/bin/cronjob_bandit24.sh
```

The script in *cronjob_bandit24.sh* is :
```
#!/bin/bash
myname=$(whoami)
cd /var/spool/$myname
echo "Executing and deleting all scripts in /var/spool/$myname:"
for i in * .*;
do
    if [ "$i" != "." -a "$i" != ".." ];
    then
	echo "Handling $i"
	timeout -s 9 60 ./$i
	rm -f ./$i
    fi
done
```

The command `cd /vars/pool/$myname` will change to *bandit24* directory and then all scripts in directory *bandit24* will be executed and then deleted. This is done in every 60 seconds as `timeout` is provided.

![Bandit Level 23 24]({{ site.url }}/assets/images/Bandit-Overthewire/bandit_l2324_terminal1.jpg){:class="img-responsive"}

So if we put our script in */var/spool/bandit24* then it will be executed and then deleted. So we will write our script such that it will output password for next level.

We will create a directory *pcdir123* in */tmp* directory and change the permissions to readable, writeable and executable for *pcdir123*. (This is very important). The command are:
```
mkdir /tmp/pcdir123
chmod 777 pcdir123
```

![Bandit Level 23 24]({{ site.url }}/assets/images/Bandit-Overthewire/bandit_l2324_terminal2.jpg){:class="img-responsive"}

Now create a script file *pass24script.sh* in */tmp/pcdir123* which will contain our script. The script is 
```
#!/bin/sh
cat /etc/bandit_pass/bandit24 > /tmp/pcdir123/pass24file
```

*pass24script.sh* is not executable, so make it executable for all users using command `chmod 777 pass24script`.

![Bandit Level 23 24]({{ site.url }}/assets/images/Bandit-Overthewire/bandit_l2324_terminal3.jpg){:class="img-responsive"}

Copy this script in */var/spool/bandit/24* using command `cp pass24script.sh /var/spool/bandit24`

When new minute starts, file pass24file is produced with password for the next level if the directory pcdir123 has write permission enabled. The password is `****` .

![Bandit Level 23 24]({{ site.url }}/assets/images/Bandit-Overthewire/bandit_l2324_terminal4.jpg){:class="img-responsive"}
 
Reference: [https://www.reddit.com/r/HowToHack/comments/cg3zu9/is_overthewires_bandit_level_2324_broken_spoilers/](https://www.reddit.com/r/HowToHack/comments/cg3zu9/is_overthewires_bandit_level_2324_broken_spoilers/)<br/>
[https://www.billycody.com/otw-wargames/bandit/bandit-level-23](https://www.billycody.com/otw-wargames/bandit/bandit-level-23)

### Next Post

[Bandit Level 24 → Level 25]({{ site.url }}/blog/2019/12/26/Bandit-Level-24-Level-25-OverTheWire)<br/>
[Bandit Level 25 to Level 26]({{ site.url }}/blog/2019/12/26/Bandit-Level-25-to-Level-26-OverTheWire)<br/>
[Bandit Level 27 to Level 31]({{ site.url }}/blog/2019/12/26/Bandit-Level-27-to-Level-31-OverTheWire)<br/>
[Bandit Level 32 → Level 33]({{ site.url }}/blog/2019/12/26/Bandit-Level-32-Level-33-OverTheWire)<br/>

### Other Wargames
[Leviathan Wargame from OverTheWire All Level Solutions]({{ site.url }}/blog/#overthewire-leviathan)<br/> 
[Krypton Wargame from OverTheWire All Level Solutions]({{ site.url }}/blog/#overthewire-krypton)<br/>


