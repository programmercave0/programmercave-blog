---
author: Programmercave
date: "2020-01-07T00:00:00Z"
description: Learn basic cryptographic techniques by playing Krypton wargame from
  OverTheWire. Below is the solution of Level 3 → Level 4, Level 4 → Level 5 and Level
  5 → Level 6.
header-img: /assets/images/Krypton-OverTheWire/overthewire_poster.jpg
tags:
- Linux
- OverTheWire-Krypton
- CTF
- Cryptography
title: Krypton Level 3 to Level 5 | Basic Cryptographic Techniques
---



Learn basic cryptographic techniques by playing [Krypton](https://overthewire.org/wargames/krypton/) wargame from OverTheWire. Below is the solution of Level 3 → Level 4, Level 4 → Level 5 and Level 5 → Level 6.

![Krypton OverTheWire](/assets/images/Krypton-OverTheWire/overthewire_poster.jpg)

### Previous Post
[Krypton Level 0 to Level 2](/Krypton-Level-0-to-Level-2-Basic-Cryptographic-Techniques)

## [Krypton Level 3 → Level 4](https://overthewire.org/wargames/krypton/krypton3.html)

### Level Info

In this example, the cipher mechanism is not available to you, the attacker.

However, you have been lucky. You have intercepted more than one message. The password to the next level is found in the file ‘krypton4’. You have also found 3 other files. (found1, found2, found3)

You know the following important details:<br/>
 * The message plaintexts are in English (*** very important) - They were produced from the same key (*** even better!) 

### Solution

Command to login `ssh krypton3@krypton.labs.overthewire.org -p 2222` and password is `CAESARISEASY` .

First lets change to directory where we will find *krypton4* . `cd /krypton/krypton3`

In the directory there are *found1*, *found2* and *found3* files which are encrypted using same key. These files contains the encrypted text like *krypton4*.

![Krypton OverTheWire](/assets/images/Krypton-OverTheWire/kryp_l34_terminal1.jpg)

Using these files and an automated cryptogram solver [quipquip](https://quipqiup.com/), we can crack the password.

First copy all the content of files *found1*, *found2* and *found3* in the puzzle box and click solve. The output we get make sense, so it has decrypted the text.

![Krypton OverTheWire](/assets/images/Krypton-OverTheWire/kryp_l34_terminal2.jpg)

Now in another tab enter the content of file *krypton4*, but there are nine possible output.

![Krypton OverTheWire](/assets/images/Krypton-OverTheWire/kryp_l34_terminal3.jpg)

Since the quipquip produced only one output when we ran it with the content of *found1*, *found2* and *found3*. So we can add the content of *krypton4* with the content of these files. We have added the content of *krypton4* file at the end, so our password will be at the end. The decrypted text at the end is *DONE THE LEVEL FOUR PASSWORD IS BRUTE*. So password for next level is `BRUTE` . 

![Krypton OverTheWire](/assets/images/Krypton-OverTheWire/kryp_l34_terminal4.jpg)

<br/>

## [Krypton Level 4 → Level 5](https://overthewire.org/wargames/krypton/krypton4.html)

### Level Info

An example of a polyalphabetic cipher is called a Vigenère Cipher. It works like this:

If we use the key(K) ‘GOLD’, and P = PROCEED MEETING AS AGREED, then “add” P to K, we get C. When adding, if we exceed 25, then we roll to 0 (modulo 26).
```
P P R O C E E D M E E T I N G A S A G R E E D\
K G O L D G O L D G O L D G O L D G O L D G O\
```

becomes:
```
P 15 17 14 2 4 4 3 12 4 4 19 8 13 6 0 18 0 6 17 4 4 3\
K 6 14 11 3 6 14 11 3 6 14 11 3 6 14 11 3 6 14 11 3 6 14\
C 21 5 25 5 10 18 14 15 10 18 4 11 19 20 11 21 6 20 2 8 10 17\
```

So, we get a ciphertext of:

`VFZFK SOPKS ELTUL VGUCH KR`

This level is a Vigenère Cipher. You have intercepted two longer, english language messages. You also have a key piece of information. You know the key length!

For this exercise, the key length is 6. The password to level five is in the usual place, encrypted with the 6 letter key.

### Solution

Command to login `ssh krypton4@krypton.labs.overthewire.org -p 2222` and password is `BRUTE` .

Read [The Vigenère Cipher Encryption and Decryption](https://pages.mtu.edu/~shene/NSF-4/Tutorial/VIG/Vig-Base.html)

In directory */krypton/krypton4* , file *kryton5* contains the encrypted password of the next level. We know the key length is 6. We will use [Vigenère Cipher – Decoder](https://www.dcode.fr/vigenere-cipher) to get the key and then the password.

In content in file *found1* and *found2* will help to find us the key. Copy the content of *found1* and *found2* in the decoder on the website and enter key length 6.

The first decrypted text make sense and the key we get from here is `FREKEY`. 

![Krypton OverTheWire](/assets/images/Krypton-OverTheWire/kryp_l45_terminal1.jpg)

Now enter the encrypted password from the file *krypton5* in the decoder box and the key `FREKEY` in the key section. The output we get is `CLEAR TEXT`, and the password for the next level is `CLEARTEXT`.

![Krypton OverTheWire](/assets/images/Krypton-OverTheWire/kryp_l45_terminal2.jpg)

<br/>

## [Krypton Level 5 → Level 6](https://overthewire.org/wargames/krypton/krypton5.html)

### Level Info

FA can break a known key length as well. Lets try one last polyalphabetic cipher, but this time the key length is unknown.

Enjoy.

### Solution

Command to login `ssh krypton5@krypton.labs.overthewire.org -p 2222` and password is `CLEARTEXT` .

Change into directory */krypton/krypton5*. We have *found1*, *found2* and *found3* files which will help us. We do not know key length and key.

![Krypton OverTheWire](/assets/images/Krypton-OverTheWire/kryp_l56_terminal1.jpg)

We can use [Vigenère Cipher – Decoder](https://www.dcode.fr/vigenere-cipher) to get the key. Copy the content of files *found1*, *found2* and *found3* in the decoder box and click Automatic Decryption. The first output make sense and the key is `KEYLENGTH`.

![Krypton OverTheWire](/assets/images/Krypton-OverTheWire/kryp_l56_terminal2.jpg)

Now copy the encrypted text from file *krypton6* and enter the key. The ouput we get is `RANDO M` and the password is `RANDOM` .

![Krypton OverTheWire](/assets/images/Krypton-OverTheWire/kryp_l56_terminal3.jpg)

