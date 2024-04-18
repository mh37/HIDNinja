# HIDNinja

## Introduction

HIDNinja is Linux based wireless HID keystroke injector with a remote payload interface.

USB is a flexible protocol that provides a lot of functions and HID functionality which is trust by default on most host systems. This makes it not only versatile but also an attractive point of attack. Among the wide range of USB based attacks, even simple methods like keystroke injections represent a significant security risk.  

## Required Hardware

This project is based on a Linux SBC which will be mimicking an HID device and is capable of receiving payloads over an onboard [WNIC](https://https://en.wikipedia.org/wiki/Wireless_network_interface_controller). It would be also possible to adjust this project for a more compact hardware setup and alternative communication methods could be used to interface with.


## TODOs & Planned future enhancements

The current implementation is very barebones, therefore, we still have an extensive list of missing features:

 • Payload repository and payload management (CRUD)
 • Communication support with the SBC for Bluetooth
 • Communication support with the SBC for radiofrequency (RF)
 • Dynamic USB peripheral emulation (on the fly change of device types) for the purpose of local exfiltration
 • Data receival of encoded exfiltration transfers from to Host to SBC through the USB control transfer pipeline to evade detection. 
 • Dynamic keystroke timing which mimics human input patterns as closely as possible to evade detection algorithms.
 • Customizable delays as part of the payload
 • Support for modifier keys
 • Make the scan code translation case sensitive.
 • Support for the handling and managing of multiple SBCs (swarm) from one user interface, and deployment of payloads to multiple targets at the same time.  
 • Implement relay/mesh networking over multiple SBCs to extend communication range. 
