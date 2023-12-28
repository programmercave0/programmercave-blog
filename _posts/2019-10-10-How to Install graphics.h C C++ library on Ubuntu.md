---
layout: post
title: "How to Install graphics.h C/C++ library on Ubuntu"
description: "In this post we are going to install graphics.h library on Ubuntu and to use it with gcc or g++ compiler. graphics.h provides access to a simple graphics library that makes it possible to draw lines, rectangles, ovals, arcs, polygons, images, and strings on a graphical window."
author: "Programmercave"
tags:  [Graphics, Cpp, Ubuntu, Linux, HowTo]
date: 2019-10-10
---



In this post we are going to install `graphics.h` library on Ubuntu and to use it with `gcc` or `g++` compiler. `graphics.h` provides access to a simple graphics library that makes it possible to draw lines, rectangles, ovals, arcs, polygons, images, and strings on a graphical window. To know about functions in `graphic.h` library in C/C++ visit [here](https://web.stanford.edu/class/archive/cs/cs106b/cs106b.1126/materials/cppdoc/graphics.html).

To use `graphics.h`on Ubuntu platform we need to compile and install `libgraph`. [libgraph](https://savannah.nongnu.org/projects/libgraph/) is an implementation of the TurboC graphics API (`graphics.h`) on GNU/Linux using SDL.

1. First we need to add the *Universe* repository because some packages are not available in main repository.
    ```
    sudo add-apt-repository universe
    sudo apt-get update
    ``` 

2. Now we will install build-essential and some additional packages.<br/>
   **For versions prior to 18.4**
      ```
      sudo apt-get install libsdl-image1.2 libsdl-image1.2-dev guile-1.8 \
      guile-1.8-dev libsdl1.2debian libart-2.0-dev libaudiofile-dev \
      libesd0-dev libdirectfb-dev libdirectfb-extra libfreetype6-dev \
      libxext-dev x11proto-xext-dev libfreetype6 libaa1 libaa1-dev \
      libslang2-dev libasound2 libasound2-dev build-essential
      ``` 

   **For 18.04** we need to add repositories of `xenial` in `sources.list`.
     ```
     sudo nano /etc/apt/sources.list
     ``` 
   
   Now add these lines
      ```
      deb http://us.archive.ubuntu.com/ubuntu/ xenial main universe
      deb-src http://us.archive.ubuntu.com/ubuntu/ xenial main universe
      ``` 
  
   Run `sudo apt-get update` and then install packages using 
      ```
      sudo apt-get install libsdl-image1.2 libsdl-image1.2-dev guile-2.0 \
      guile-2.0-dev libsdl1.2debian libart-2.0-dev libaudiofile-dev \
      libesd0-dev libdirectfb-dev libdirectfb-extra libfreetype6-dev \
      libxext-dev x11proto-xext-dev libfreetype6 libaa1 libaa1-dev \
      libslang2-dev libasound2 libasound2-dev
      ```
    
3. Download `libgraph` from [here](download.savannah.gnu.org/releases/libgraph/libgraph-1.0.2.tar.gz) and extract `libgraph-1.0.2.tar.gz` file.    
 <input type="hidden" name="IL_IN_ARTICLE"> 
4. Go to the extracted folder and run following commands.
     ```
     ./configure
     make
     sudo make install
     sudo cp /usr/local/lib/libgraph.* /usr/lib
     ```
  
Now we can use `graphics.h` library after adding following lines to the code.

    int gd = DETECT, gm; 
    initgraph(&gd, &gm, NULL);
   
  
Here is a program of [Simple Pendulum Animation on Ubuntu]({{ site.url }}/blog/2019/10/09/C++-Simple-Pendulum-Animation-on-Ubuntu-Machine).  
    
