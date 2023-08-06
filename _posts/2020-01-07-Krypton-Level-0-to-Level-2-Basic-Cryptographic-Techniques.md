---
layout: post
title: "Krypton Level 0 to Level 2 | Basic Cryptographic Techniques"
description: "Learn basic cryptographic techniques by playing Krypton wargame from OverTheWire. Below is the solution of Level 0 → Level 1, Level 1 → Level 2 and Level 2 → Level 3."
author: "Programmercave"
header-img: "/assets/Krypton-OverTheWire/overthewire_poster.jpg"
tags:  [Linux, OverTheWire-Krypton, CTF, Cryptography]
date: 2020-01-07
---
* toc
{:toc}

Learn basic cryptographic techniques by playing [Krypton](https://overthewire.org/wargames/krypton/) wargame from OverTheWire. Below is the solution of Level 0 → Level 1, Level 1 → Level 2 and Level 2 → Level 3.

![Krypton OverTheWire]({{ site.url }}/assets/Krypton-OverTheWire/overthewire_poster.jpg){:class="img-responsive"}

## [Krypton Level 0 → Level 1](https://overthewire.org/wargames/krypton/krypton0.html)

### Level Info

Welcome to Krypton! The first level is easy. The following string encodes the password using Base64:
`S1JZUFRPTklTR1JFQVQ=`

Use this password to log in to krypton.labs.overthewire.org with username krypton1 using SSH on port 2222. You can find the files for other levels in /krypton/

### Solution 

To get the password of the next level decode the given string using command
```
echo S1JZUFRPTklTR1JFQVQ= | base64 -d
```

and the password is `KRYPTONISGREAT` .

![Krypton OverTheWire]({{ site.url }}/assets/Krypton-OverTheWire/kryp_l01_terminal1.jpg){:class="img-responsive"}


Reference : [How can I decode a base64 string from the command line?](https://askubuntu.com/questions/178521/how-can-i-decode-a-base64-string-from-the-command-line)


## [Krypton Level 1 → Level 2](https://overthewire.org/wargames/krypton/krypton1.html)

### Level Info

The password for level 2 is in the file ‘krypton2’. It is ‘encrypted’ using a simple rotation. It is also in non-standard ciphertext format. When using alpha characters for cipher text it is normal to group the letters into 5 letter clusters, regardless of word boundaries. This helps obfuscate any patterns. This file has kept the plain text word boundaries and carried them to the cipher text. Enjoy!

### Solution

Command to login `ssh krypton1@krypton.labs.overthewire.org -p 2222` and password is `KRYPTONISGREAT` .

The file *krypton2* is in directory */krypton/krypton1*. First change into that directory `cd /krypton/krypton1` .

![Krypton OverTheWire]({{ site.url }}/assets/Krypton-OverTheWire/kryp_l12_terminal1.jpg){:class="img-responsive"}


The password in *krypton2* file is encrypted using simple rotation. We can decrypt it like we did in [Bandit Level 11 → Level 12]({{ site.url }}/blog/2019/12/24/Bandit-Level-9-to-Level-12-OverTheWire). The command is
```
cat krypton2 | tr "[a-zA-Z]" "[n-za-mN-ZA-M]
```

and the password is `ROTTEN` .

![Krypton OverTheWire]({{ site.url }}/assets/Krypton-OverTheWire/kryp_l12_terminal2.jpg){:class="img-responsive"}


{% include ads.html %}<br/>

## [Krypton Level 2 → Level 3](https://overthewire.org/wargames/krypton/krypton2.html)

### Level Info

This level contains an old form of cipher called a ‘Caesar Cipher’. A Caesar cipher shifts the alphabet by a set number. For example:
```
plain:  a b c d e f g h i j k ...
cipher: G H I J K L M N O P Q ...
```

In this example, the letter ‘a’ in plaintext is replaced by a ‘G’ in the ciphertext so, for example, the plaintext ‘bad’ becomes ‘HGJ’ in ciphertext.

The password for level 3 is in the file krypton3. It is in 5 letter group ciphertext. It is encrypted with a Caesar Cipher. Without any further information, this cipher text may be difficult to break. You do not have direct access to the key, however you do have access to a program that will encrypt anything you wish to give it using the key. If you think logically, this is completely easy.

### Solution

Command to login `ssh krypton2@krypton.labs.overthewire.org -p 2222` and password is `ROTTEN` .

The encrypted password is in the *krypton3* which is in the */krypton/krypton2* directory. So first change into that directory. `cd /krypton/krypton2`.

In the directory there is encrypt binary and *keyfile.dat* which contains the key but we cannot open it.  When we execute the binary the output says that a file containing plaintext should be executed along with the binary.

![Krypton OverTheWire]({{ site.url }}/assets/Krypton-OverTheWire/kryp_l23_terminal1.jpg){:class="img-responsive"}

We can create a directory */tmp/programmercave* . This directory mush have executable permission set because we will execute the *encrypt* binary. This directory will contain a *plaintext* file with text ABCD. 
```
mkdir /tmp/programmercave
chmod 777 /tmp/programmercave
cd /tmp/programmercave
cat > plaintext
ABCD 
^C
```

![Krypton OverTheWire]({{ site.url }}/assets/Krypton-OverTheWire/kryp_l23_terminal2.jpg){:class="img-responsive"}

We need to create symbolic link to file */krypton/krypton2/keyfile.dat* because when encrypt is executed, key should there in that directory. This can be done using 
```
ln -s /krypton/krypton2/keyfile.dat  
```

![Krypton OverTheWire]({{ site.url }}/assets/Krypton-OverTheWire/kryp_l23_terminal3.jpg){:class="img-responsive"}

The command `/krypton/krypton2/encrypt plaintext` will encrypt the text ABCD in *plaintext* using key from *keyfile.dat* to new file *ciphertext*. The file *ciphertext* contains MNOP. This means the key is converting ABCD to MNOP.

Using this we can decrypt the password in *krypton3* file using `tr` program. The command is 
```
cat /krypton/krypton2/krypton3 | tr "[m-za-lM-ZA-L]" "[a-zA-Z]"
```
and the password is `CAESARISEASY` .

![Krypton OverTheWire]({{ site.url }}/assets/Krypton-OverTheWire/kryp_l23_terminal4.jpg){:class="img-responsive"}

### Next Post
[Krypton Level 3 to Level 5]({{ site.url }}/blog/2020/01/07/Krypton-Level-3-to-Level-5-Basic-Cryptographic-Techniques)

### Other Wargames
[Bandit Wargame from OverTheWire All Level Solutions]({{ site.url }}/blog/#overthewire-bandit)<br/>
[Leviathan Wargame from OverTheWire All Level Solutions]({{ site.url }}/blog/#overthewire-leviathan)<br/> 


