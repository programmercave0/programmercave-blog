---
author: Programmercave
date: "2019-12-26T00:00:00Z"
description: Learn linux command by playing Bandit wargame. The Bandit wargame is
  aimed at absolute beginners. It will teach the basics needed to be able to play
  other wargames. Below is the solution of Bandit Level 27 → Level 28, Level 28 →
  Level 29, Level 29 → Level 30, Level 30 → Level 31 and 31 → Level 32. In this post
  we will learn about git and its terminal command. The passwords are hidden, so you
  have to find the passwords for next level yourself.
header-img: /assets/images/Bandit-Overthewire/overthewire_poster.jpg
tags:
- Linux
- OverTheWire-Bandit
- CTF
title: Bandit Level 27 to Level 31 | OverTheWire
---



Learn linux command by playing [Bandit](https://overthewire.org/wargames/bandit/) wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Bandit Level 27 → Level 28, Level 28 → Level 29, Level 29 → Level 30, Level 30 → Level 31 and 31 → Level 32. 

In this post we will learn about git and its terminal command.

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
[Bandit Level 25 to Level 26](/Bandit-Level-25-to-Level-26-OverTheWire)

## [Bandit Level 27 → Level 28](https://overthewire.org/wargames/bandit/bandit28.html)

### Level Goal

There is a git repository at `ssh://bandit27-git@localhost/home/bandit27-git/repo`. The password for the user `bandit27-git` is the same as for the user `bandit27`.

Clone the repository and find the password for the next level.

### Commands you may need to solve this level

git

### Solution : 

Command to connect remote host : `ssh bandit27@bandit.labs.overthewire.org -p 2220` password is `****` .

We will use command `git clone` to clone repository. But we cannot clone it in a home directory, so make a new directory *programmercave* in *tmp* directory.

Command to clone 
```
git clone ssh://bandit27-git@localhost/home/bandit27-git/repo
```

and enter this level’s password.

The password is in the README in *repo* directory and the password is `****` .

![Bandit Level 27 28](/assets/images/Bandit-Overthewire/bandit_l2728_terminal.jpg)

<br/>

## [Bandit Level 28 → Level 29](https://overthewire.org/wargames/bandit/bandit29.html)

### Level Goal

There is a git repository at `ssh://bandit28-git@localhost/home/bandit28-git/repo`. The password for the user `bandit28-git` is the same as for the user `bandit28`.

Clone the repository and find the password for the next level.

### Commands you may need to solve this level

git

### Solution : 

Command to connect remote host : `ssh bandit28@bandit.labs.overthewire.org -p 2220` password is `****` .

Initial part is same as the previous level. Clone the repository in your directory in *tmp* directory.
``` 
git clone  ssh://bandit28-git@localhost/home/bandit28-git/repo
```

But README.md file does not contain password.

![Bandit Level 28 29](/assets/images/Bandit-Overthewire/bandit_l2829_terminal1.jpg)


The file README.md has been updated. The `git show` command shows the most recent commit on the current branch. This command shows the changes made in the README.md file. The password for the next level is `****` .

![Bandit Level 28 29](/assets/images/Bandit-Overthewire/bandit_l2829_terminal2.jpg)

Reference : [https://git-scm.com/docs/user-manual](https://git-scm.com/docs/user-manual)

<br/>

## [Bandit Level 29 → Level 30](https://overthewire.org/wargames/bandit/bandit30.html)

### Level Goal

There is a git repository at `ssh://bandit29-git@localhost/home/bandit29-git/repo`. The password for the user `bandit29-git` is the same as for the user `bandit29.`

Clone the repository and find the password for the next level.

### Commands you may need to solve this level

git

### Solution : 

Command to connect remote host : `ssh bandit29@bandit.labs.overthewire.org -p 2220` password is `****` .

Initial part is same as the previous level. Clone the repository under *tmp* directory using `git clone  ssh://bandit29-git@localhost/home/bandit29-git/repo`

The README file does not contain password. Infact password is not in the production. Lets check other branches.

![Bandit Level 29 30](/assets/images/Bandit-Overthewire/bandit_l2930_terminal1.jpg)

`git branch` command shows you the list of branch heads.

The repository may also have had other branches, though, and your local repository keeps branches which track each of those remote branches, called remote-tracking branches, which you can view using the `-r` option to `git branch`.
 
`git checkout` command lets you navigate between the branches created by `git branch`. 

The README.md file in dev branch has the password for the next level and the password is `****` .

![Bandit Level 29 30](/assets/images/Bandit-Overthewire/bandit_l2930_terminal2.jpg)

Reference : [https://git-scm.com/docs/user-manual](https://git-scm.com/docs/user-manual)<br/>
[Git Checkout](https://www.atlassian.com/git/tutorials/using-branches/git-checkout)

<br/>

## [Bandit Level 30 → Level 31](https://overthewire.org/wargames/bandit/bandit31.html)

### Level Goal

There is a git repository at `ssh://bandit30-git@localhost/home/bandit30-git/repo`. The password for the user `bandit30-git` is the same as for the user `bandit30`.

Clone the repository and find the password for the next level.

### Commands you may need to solve this level

git

### Solution : 

Command to connect remote host : `ssh bandit30@bandit.labs.overthewire.org -p 2220` password is `****` .

Clone the repository in tmp directory using command `git clone ssh://bandit30-git@localhost/home/bandit30-git/repo`

The README.md file does not contain the password. 

![Bandit Level 30 31](/assets/images/Bandit-Overthewire/bandit_l3031_terminal1.jpg)

`git show`, `git log` and `git branch -r` does not help us in this level.

![Bandit Level 30 31](/assets/images/Bandit-Overthewire/bandit_l3031_terminal2.jpg)

`git tag` create, list, delete or verify a tag object signed with GPG. This command tells us about the `secret` tag. We can view this tag using `git show secret` and the password is `****` .

![Bandit Level 30 31](/assets/images/Bandit-Overthewire/bandit_l3031_terminal3.jpg)

Reference : [https://git-scm.com/docs/user-manual](https://git-scm.com/docs/user-manual)<br/>
[https://git-scm.com/docs/git-tag](https://git-scm.com/docs/git-tag)

<br/>

## [Bandit Level 31 → Level 32](https://overthewire.org/wargames/bandit/bandit32.html)

### Level Goal

There is a git repository at `ssh://bandit31-git@localhost/home/bandit31-git/repo`. The password for the user `bandit31-git` is the same as for the user `bandit31`.

Clone the repository and find the password for the next level.

### Commands you may need to solve this level

git

### Solution : 

Command to connect remote host : `ssh bandit31@bandit.labs.overthewire.org -p 2220` password is `****` .

First clone the repository in *tmp* directory using `git clone ssh://bandit31-git@localhost/home/bandit31-git/repo`

The README.md file tells us to push file *key.txt* to the remote directory. The commands to push a file on a git repository is
```
git add key.txt
git commit -m “commit message”
git push origin master
```

![Bandit Level 31 32](/assets/images/Bandit-Overthewire/bandit_l3132_terminal1.jpg)

But when we try to add file it tell us that *.gitignore* file is ignoring our file.

![Bandit Level 31 32](/assets/images/Bandit-Overthewire/bandit_l3132_terminal2.jpg)

We can delete the *.gitignore* file.

![Bandit Level 31 32](/assets/images/Bandit-Overthewire/bandit_l3132_terminal3.jpg)

Now if add and push our file it will show us the password for next level and the password is `****` .

![Bandit Level 31 32](/assets/images/Bandit-Overthewire/bandit_l3132_terminal4.jpg)

### Next Post

[Bandit Level 32 → Level 33](/Bandit-Level-32-Level-33-OverTheWire)<br/>


