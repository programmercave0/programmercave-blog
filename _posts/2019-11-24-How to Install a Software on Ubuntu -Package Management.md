---
layout: post
title: "How to Install a Software on Ubuntu | Package Management"
subtitle: "A package file is a compressed collection of files necessary to install a software or an application. It also includes metadata about the package such as text description of package and name of dependencies the software needs."
author: "Programmercave"
header-img: "/assets/ubuntulogo.jpeg"
tags:  [Operating-System, Linux, Ubuntu]
date: 2019-11-24
---

![Range]({{ site.url }}/assets/ubuntulogo.jpeg){:class="img-responsive"}

**Package management** facilities helps us to install and maintain software on the system. The package file for Ubuntu and other Debian based OS ends with `.deb`. <br/>For eg. `draw.io-amd64-12.2.2.deb`

<h3>Package Files</h3>

A *package file* is a compressed collection of files necessary to install a software or an application. It also includes metadata about the package such as text description of package and name of dependencies the software needs.

<h3>Dependencies</h3>

To run a progam on a system, it rely on other programs. These other programs are called *dependencies*.

<h3>Repository</h3>

A *repository* is a collection of set of packages on a server. It can be accessed with package management tools like *apt-get*.

<h1>Package Management Tools</h1>

There are two types of package management tools in a package management system.
  1. Low-level tools for eg. `dpkg`
  2. High-level tools for eg. `apt-get`, `aptitude`
  
{% include ads.html %}<br/>

<h3>Difference between dpkg and apt-get</h3>

The **dpkg** tool is used to install and remove a package file. While installing if it encounters an unsatisfied dependency, it will unable to install the software or the application. It will simply list the missing dependencies. The Advanced Package Tool (**APT**), including **apt** and **apt-get** can automatically resolve these issues.

APT installs the latest package from an online source while dpkg installs from a package stored in our local system.

<h1>Commands</h1>
<h3>Finding a Package in a Repository</h3>

`apt-cache search search_string`<br/>
For eg. `apt-cache search draw.io`

<h3>Installing a Package from a Repository</h3>

`apt-get install package_name`<br/>
For eg. `apt-get install draw.io`

<h3>Installing a Package from a Package File</h3>

`dpkg --install package_file`<br/>
For eg. `dpkg --install draw.io-amd64-12.2.2.deb`

<h3>Removing a Package</h3>

`apt-get remove package_name`<br/>
For eg. `apt-get remove draw.io`

{% include ads.html %}<br/>

<h3>Updating Packages from a Repository</h3>

`apt-get update; apt-get upgrade`

<h3>Listing Installed Packages</h3>

`dpkg --list`

<h3>Determining if a Package is Installed</h3>

`dpkg --status package_name`<br/>
For eg. `dpkg --status draw.io`

<h3>Displaying Info About an Installed Package</h3>

`apt-cache show package_name`<br/>
For eg. `apt-cache show draw.io`

Get this post in pdf - [How to install a Software on Ubuntu](https://www.file-up.org/632tqcgzdvnv)

Reference:<br/>
[The Linux Command Line](https://amzn.to/2QIYQel)<br/>
[https://www.oreilly.com/openbook/debian/book/appc_01.html](https://www.oreilly.com/openbook/debi/book/appc_01.html)<br/>
[https://kali.training/topic/introduction-to-apt/](https://kali.training/topic/introduction-to-apt/)<br/>
[https://help.ubuntu.com/community/Repositories](https://help.ubuntu.com/community/Repositories)

<h3>You may also like</h3>
[How to create Password list using CUPP tool on Ubuntu](https://programmercave0.github.io//blog/2019/10/10/How-to-create-Password-list-using-CUPP-tool-on-ubuntu)

  
