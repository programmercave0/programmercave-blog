---
author: Programmercave
date: "2022-09-27T00:00:00Z"
description: Are you having trouble getting natural scrolling to work on your touchpad
  after updating OpenSuse Tumbleweed? Don't worry, we've got you covered. We've gone
  through the process of searching various websites and experimenting with different
  techniques, and we've finally found a solution. We'll walk you through the steps
  we took to get natural scrolling working on our touchpad, including modifying the
  wrong configuration file, encountering the dreaded blank screen of death, and rolling
  back to a previous snapshot. Please note that this tutorial is specifically for
  enabling natural scrolling on a touchpad, not a mouse. There are plenty of other
  resources available for enabling natural scrolling on a mouse.
header-img: /assets/images/Enable-Natural-Scrolling/pic1.png
tags:
- Linux
- Ubuntu
- OpenSuse
- Tumbleweed
title: How to Enable Natural Scrolling for Touchpad on OpenSuse Tumbleweed
toc: true
---

## Introduction

Are you having trouble getting natural scrolling to work on your touchpad after updating OpenSuse Tumbleweed? Don't worry, we've got you covered. We've gone through the process of searching various websites and experimenting with different techniques, and we've finally found a solution. We'll walk you through the steps we took to get natural scrolling working on our touchpad, including modifying the wrong configuration file, encountering the dreaded blank screen of death, and rolling back to a previous snapshot. Please note that this tutorial is specifically for enabling natural scrolling on a touchpad, not a mouse. There are plenty of other resources available for enabling natural scrolling on a mouse.

<h2>How to Install and Use xinput on OpenSuse Tumbleweed </h2>

If you want to manage your input devices on OpenSuse Tumbleweed, xinput is the tool you need. To install xinput, simply open a terminal and run `sudo zypper install xinput`. Once xinput is installed, you can use it to list all the input devices currently connected to your system. To do this, run `xinput --list` in the terminal. With this information, you can easily configure and customize your input devices to suit your needs.

![Enable Natural Scrolling for Touchpad in Linux](/assets/images/Enable-Natural-Scrolling/pic1.png)

You should see entries for both your mouse and touchpad. To view the options available for a specific input device, use the `xinput --list-props` command followed by the device ID. For example, to view the options for your mouse, you would run `xinput --list-props <mouse device ID>`, and to view the options for your touchpad, you would run `xinput --list-props <touchpad device ID>`. This will give you a list of all the options you can set for each device, allowing you to customize them to your liking.

<h2>Option Natural Scrolling </h2>

To view the options available for your mouse using xinput on OpenSuse Tumbleweed, you can run the following command in the terminal:

`xinput --list-props 12`

Note that you will need to replace the number "12" with the actual device ID for your mouse, as identified by the `xinput --list` command. This command will display a list of all the options you can set for your mouse, such as the button mapping, pointer acceleration, and scroll speed. You can then use these options to customize your mouse to your liking.

For example, if you want to change the scroll speed of your mouse, you can use the `xinput --set-prop` command followed by the device ID and the option you want to set. For example:

`xinput --set-prop 12 "Device Accel Constant Deceleration" 2`

This would set the "Device Accel Constant Deceleration" option for your mouse (device ID 12) to a value of 2. You can experiment with different values to find the scroll speed that works best for you.

![Enable Natural Scrolling for Touchpad in Linux](/assets/images/Enable-Natural-Scrolling/pic2.png)

We can see Mouse has option for Natural Scrolling.

But for Touchpad there is no options such as Natural Scrolling.

If you run `xinput --list-props <touchpad device ID>` and do not see an option for natural scrolling, it may be because natural scrolling is not a feature that is directly supported by the touchpad itself. Instead, it is typically implemented by the operating system or desktop environment.

On OpenSuse Tumbleweed, you can try enabling natural scrolling for your touchpad by going to the settings for your desktop environment (such as Gnome or KDE) and looking for the option to enable natural scrolling. Alternatively, you can try using a tool such as "gpointing-device-settings" to configure natural scrolling for your touchpad.

If neither of these options work, it is possible that your touchpad simply does not support natural scrolling. In this case, you may need to look into using a different touchpad or mouse that does support natural scrolling.

<br/>

<h2>Editing synaptics configuration file </h2>

To make changes to your touchpad configuration permanent on OpenSuse Tumbleweed, you can edit the `70-synaptics.conf` file located in the `/usr/share/X11/xorg.conf.d/` directory. This file contains the configuration settings for the Synaptics touchpad driver, which is responsible for handling input from your touchpad.

To edit the file, you will need to have root privileges. You can do this by running the sudo command before the text editor of your choice. For example:

`sudo nano /usr/share/X11/xorg.conf.d/70-synaptics.conf`

This will open the `70-synaptics.conf` file in the Nano text editor. You can then make the desired changes to the file and save them by pressing `Ctrl+X` and then `Y`.

After making your changes, it is a good idea to copy the `70-synaptics.conf` file to the `/etc/X11/xorg.conf.d/` directory. This will ensure that your changes are applied every time you boot your system. To copy the file, you can use the `cp` command:

`sudo cp /usr/share/X11/xorg.conf.d/70-synaptics.conf /etc/X11/xorg.conf.d/`

Note that it is not necessary to create a new file in the `/etc/X11/xorg.conf.d/` directory. Instead, you should simply copy the existing `70-synaptics.conf` file to this location.

After copying the file and making your changes, you may need to restart your system or your X server for the changes to take effect. You can do this by using the `systemctl` command:

`sudo systemctl restart display-manager`

To enable natural scrolling for your touchpad on OpenSuse Tumbleweed, you can add the following options to the `70-synaptics.conf` file located in the `/usr/share/X11/xorg.conf.d/` directory:

```
Option "VertScrollDelta" "-27"
Option "HorizScrollDelta" "-27"
```

These options should be added under the `Identifier "touchpad catchall"` section of the file. The `VertScrollDelta` option controls the vertical scrolling behavior of your touchpad, while the `HorizScrollDelta` option controls the horizontal scrolling behavior. Setting these options to a negative value will cause your touchpad to scroll in the opposite direction of your finger movement, which is how natural scrolling works.

To make these changes permanent, you will need to copy the `70-synaptics.conf` file to the `/etc/X11/xorg.conf.d/` directory, as described above. After copying the file and making your changes, you may need to restart your system or your X server for the changes to take effect. You can do this by using the `systemctl` command:

`sudo systemctl restart display-manager`

Alternatively, you can try restarting your touchpad by running the following command:

`sudo modprobe -r psmouse; sudo modprobe psmouse`

This will unload and then reload the touchpad driver, which may cause your changes to take effect.

Keep in mind that these instructions are specific to the Synaptics touchpad driver. If you are using a different touchpad driver, the options and steps for enabling natural scrolling may be different.

![Enable Natural Scrolling for Touchpad in Linux](/assets/images/Enable-Natural-Scrolling/pic3.png)

 If you followed the steps outlined, you should now be able to scroll in the opposite direction of your finger movement on your touchpad, which can be a more intuitive and natural way of scrolling for some users.

<h2>References </h2>

[Change Mouse Settings Using Xinput](https://linuxhint.com/change_mouse_touchpad_settings_xinput_linux/) <br>
[Manually Configure a Synaptics Touchpad By Editing Xorg Configs](https://www.linkedin.com/pulse/manually-configure-synaptics-touchpad-editing-xorg-configs-basel-korj/) <br>
[Synaptics Natural Scrolling](https://bbs.archlinux.org/viewtopic.php?id=266547) <br>
[synaptics(4) - Linux man page](https://linux.die.net/man/4/synaptics) <br>
