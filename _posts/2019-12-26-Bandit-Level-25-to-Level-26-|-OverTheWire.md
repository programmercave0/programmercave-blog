---
layout: post
title: "Bandit Level 25 to Level 26 | OverTheWire"
description: "Learn linux command by playing Bandit wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Bandit Level 25 → Level 26 and 26 → Level 27. In this level we will learn how to change shell and how size of the terminal window can also help us to crack the password. The passwords are hidden, so you have to find the passwords for next level yourself."
author: "Programmercave"
header-img: "/assets/Bandit-Overthewire/overthewire_poster.jpg"
tags:  [Linux, OverTheWire-Bandit, CTF]
date: 2019-12-26
---

Learn linux command by playing [Bandit](https://overthewire.org/wargames/bandit/) wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Bandit Level 25 → Level 26 and 26 → Level 27. 

In this level we will learn how to change shell and how size of the terminal window can also help us to crack the password. 

The passwords are hidden, so you have to find the passwords for next level yourself.

![Bandit OverTheWire]({{ site.url }}/assets/Bandit-Overthewire/overthewire_poster.jpg){:class="img-responsive"}

### Previous Post

[Bandit Level 0 to Level 3]({{ site.url }}/blog/2019/12/21/Bandit-Level-0-to-Level-5-OverTheWire)<br/>
[Bandit Level 4 to Level 8]({{ site.url }}/blog/2019/12/22/Bandit-Level-4-to-Level-9-OverTheWire)<br/>
[Bandit Level 9 to Level 11]({{ site.url }}/blog/2019/12/24/Bandit-Level-9-to-Level-12-OverTheWire)<br/>
[Bandit Level 12 → Level 13]({{ site.url }}/blog/2019/12/24/Bandit-Level-12-Level-13-OverTheWire)<br/>
[Bandit Level 13 to Level 15]({{ site.url }}/blog/2019/12/24/Bandit-Level-13-to-Level-16-OverTheWire)<br/>
[Bandit Level 16 to Level 18]({{ site.url }}/blog/2019/12/24/Bandit-Level-16-to-Level-19-OverTheWire)<br/>
[Bandit Level 19 to Level 20]({{ site.url }}/blog/2019/12/25/Bandit-Level-19-to-Level-20-OverTheWire)<br/>
[Bandit Level 21 to Level 22]({{ site.url }}/blog/2019/12/25/Bandit-Level-21-to-Level-23-OverTheWire)<br/>
[Bandit Level 23 → Level 24]({{ site.url }}/blog/2019/12/25/Bandit-Level-23-Level-24-OverTheWire)<br/>
[Bandit Level 24 → Level 25]({{ site.url }}/blog/2019/12/26/Bandit-Level-24-Level-25-OverTheWire)

## [Bandit Level 25 → Level 26](https://overthewire.org/wargames/bandit/bandit26.html)

### Level Goal

Logging in to bandit26 from bandit25 should be fairly easy… The shell for user bandit26 is not **/bin/bash**, but something else. Find out what it is, how it works and how to break out of it.

### Commands you may need to solve this level

ssh, cat, more, vi, ls, id, pwd

### Solution : 

Command to connect remote host : `ssh bandit25@bandit.labs.overthewire.org -p 2220` password is `****` .

There is a private key in the file *bandit26.sshkey* in */home/bandit25* directory. If we use this private key to login as user *bandit26* using command `ssh -i bandit26.sshkey bandit26@localhost` we are logged out. We receive a message “*Connection to localhost closed*”.

![Bandit Level 25 26]({{ site.url }}/assets/Bandit-Overthewire/bandit_l2526_terminal1.jpg){:class="img-responsive"}

Lets find some information of user bandit25 and bandit26. bandit25 uses */bin/sh* shell and bandit26 uses something */usr/bin/showtext*.

![Bandit Level 25 26]({{ site.url }}/assets/Bandit-Overthewire/bandit_l2526_terminal2.jpg){:class="img-responsive"}

Lets see what is inside the showtext file. The content is
``` 
#!/bin/sh
export TERM=linux
more ~/text.txt
exit 0
```

![Bandit Level 25 26]({{ site.url }}/assets/Bandit-Overthewire/bandit_l2526_terminal3.jpg){:class="img-responsive"}

Before exiting, command more `~/text.txt` is executed.

`more` is a filter for paging through text one screenful at a time. With command `v` we can start up an editor at current line.  The editor is taken from the environment variable VISUAL if defined, or EDITOR if VISUAL is not defined, or defaults to vi if neither VISUAL nor EDITOR is defined.

To enable `more` we have to decrease size of our terminal window.

![Bandit Level 25 26]({{ site.url }}/assets/Bandit-Overthewire/bandit_l2526_terminal4.jpg){:class="img-responsive"}

Enter command `ssh -i bandit26.sshkey bandit26@localhost`.

![Bandit Level 25 26]({{ site.url }}/assets/Bandit-Overthewire/bandit_l2526_terminal5.jpg){:class="img-responsive"}

After typing `v` we enter into vim editor then enter `:e /etc/bandit_pass/bandit26` to view password for next level. The password for next level is `****` .

![Bandit Level 25 26]({{ site.url }}/assets/Bandit-Overthewire/bandit_l2526_terminal6.jpg){:class="img-responsive"}

Reference : [How To View System Users in Linux on Ubuntu](https://www.digitalocean.com/community/tutorials/how-to-view-system-users-in-linux-on-ubuntu)<br/>
[http://man7.org/linux/man-pages/man1/more.1.html](http://man7.org/linux/man-pages/man1/more.1.html)<br/>
[https://www.billycody.com/otw-wargames/bandit/bandit-level-25](https://www.billycody.com/otw-wargames/bandit/bandit-level-25)

{% include ads.html %}<br/>

## [Bandit Level 26 → Level 27](https://overthewire.org/wargames/bandit/bandit27.html)

### Level Goal

Good job getting a shell! Now hurry and grab the password for bandit27!

### Commands you may need to solve this level

ls

### Solution : 

Command to connect remote host : `ssh bandit26@bandit.labs.overthewire.org -p 2220` password is `****` .

Before running the above command, shrink the terminal window like we have done in previous level.

Enter `v` to start up an editor mode. Command `:set shell ?` will tell the shell of the user and `:set shell=/bin/sh` will set it to */bin/sh*. Run `:set shell ?` again to confirm.

To execute a command in subshell, enter `:!command`. Lets execute `:!ls -la`.

There is a *bandit27-do* file with elevated privilege. It will help us to know the password.

![Bandit Level 26 27]({{ site.url }}/assets/Bandit-Overthewire/bandit_l2627_terminal1.jpg){:class="img-responsive"}

Run command 
```
:! ./bandit27-do cat /etc/bandit_pass/bandit27
```

and the password is `****` .

![Bandit Level 26 27]({{ site.url }}/assets/Bandit-Overthewire/bandit_l2627_terminal2.jpg){:class="img-responsive"}


Reference : Vim tips: [Working with external commands](https://www.linux.com/tutorials/vim-tips-working-external-commands/)<br/>
[http://man7.org/linux/man-pages/man1/more.1.html](http://man7.org/linux/man-pages/man1/more.1.html)

### Next Post

[Bandit Level 27 to Level 31]({{ site.url }}/blog/2019/12/26/Bandit-Level-27-to-Level-31-OverTheWire)<br/>
[Bandit Level 32 → Level 33]({{ site.url }}/blog/2019/12/26/Bandit-Level-32-Level-33-OverTheWire)<br/>

### Other Wargames
[Leviathan Wargame from OverTheWire All Level Solutions]({{ site.url }}/blog/#overthewire-leviathan)<br/> 
[Krypton Wargame from OverTheWire All Level Solutions]({{ site.url }}/blog/#overthewire-krypton)<br/>

