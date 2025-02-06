---
author: Programmercave
date: "2019-10-09T00:00:00Z"
description: 'Here is the code to animate simple pendulum on Ubuntu Machine. '
tags:
- Cpp
- Ubuntu
- Graphics
- Project
title: 'C++: Simple Pendulum Animation on Ubuntu Machine'
---



Here is the code to animate simple pendulum on Ubuntu Machine.

We have to add `#include <graphics.h>` header file.

If you want to install C/C++ `graphic.h` header file on Ubuntu you can visit this post. [How to Install graphics.h C/C++ library on Ubuntu](/How-to-Install-graphics.h-C-C-library-on-Ubuntu)

We have to also write 
  int gd = DETECT, gm;
  initgraph(&gd, &gm, NULL);
  
`initgraph()` initializes the graphics system by loading a graphics driver from disk (or validating a registered driver), and putting the system into graphics mode. 

Equations used:-

Angular acceleration of the pendulum:

    θ'' = − g⁄R sin <br/>
        θ = angle of pendulum ( 0= vertical) <br/>
        R = length of rod <br/>
        g = gravitational constant <br/>
 
Angular displacement :

    θ(t) = θ0cos(2π/T * t + φ)<br/>
    φ is a constant value

Period of oscillation:

    T ≈ 2π√(R/g)

 <input type="hidden" name="IL_IN_ARTICLE"> 

C++ Implementation

```cpp
#include <iostream>
#include <cmath>
#include <graphics.h>

const float Pi = 3.1415926535897931f;

float angular_disp(float t, float angle, float T) //angular displacement
{
    //Simple Harmonic Motion equation
    float res = angle * cos((2*Pi / T) * t); //let additional const be 0
    return res;
}

void animate_pendulum(float x1, float y1,
                      float l, float g, float angle, float T)
{
   int gd = DETECT, gmAngular displacement :


    θ(t) = θ0cos(2π/T * t + φ)


    φ is a constant value


Period of oscillation:

    T ≈ 2π√(R/g);
   initgraph(&gd, &gm, NULL);

  float angular_acc;//angular acceleration

  for(float t = 1.0; t < 100.0;)
  {
     float new_angle = angular_disp(t, angle, T);
     angular_acc = -(g/l) * sin(new_angle);

     //To know other end coordinates
     float base = l * cos(new_angle);
     float height = l * sin(new_angle);

     //end coordinates
     float y2 = y1 + base;
     float x2 = x1 + height;

     setlinestyle(DASHED_LINE,0,NORM_WIDTH);
     setcolor(WHITE);
     line(x1, y1-20, x1, y1+250);

     setlinestyle(SOLID_LINE,0,THICK_WIDTH);
     setcolor(BLUE);
     line(x1, y1, x2, y2);

     setcolor(RED);
     circle(x2, y2, 10);
     floodfill(x2, y2, RED);

     delay(100);
     cleardevice();

     t = t + 1.0;
  }
  delay(100);
  closegraph();
}

int main()
{
   float x1 = 250, y1 = 100; //start coordinates of pendulum

   float l = 200; //length of pendulum
   float g = 10; //acceleration of gravity (assumed)

   float angle = Pi/2; // in radians

   float T = (2 * Pi) * sqrt(l/g); //period of swing

   animate_pendulum(x1, y1, l, g, angle, T);
}
```    

Compile it with the command `g++ simplependulum.cpp -o simplependulum -lgraph`

<iframe width="560" height="315" src="https://www.youtube.com/embed/0Loe_KassCk" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

Reference:

[https://en.wikipedia.org/wiki/Pendulum](https://en.wikipedia.org/wiki/Pendulum)<br/>
[https://www.cs.colorado.edu/~main/bgi/doc/initgraph.html](https://www.cs.colorado.edu/~main/bgi/doc/initgraph.html)<br/>
[https://www.myphysicslab.com/pendulum/pendulum-en.html](https://www.myphysicslab.com/pendulum/pendulum-en.html) 
  

