# HIDNinja

## Introduction

HIDNinja is Linux based wireless HID keystroke injector with a remote payload interface.

## Required Hardware

This project is based on a Linux SBC which will be mimicking an HID device and is capable of receiving payloads over an onboard [WNIC](https://https://en.wikipedia.org/wiki/Wireless_network_interface_controller). It would be also possible to adjust this project for a more compact hardware setup and alternative communication methods could be used to interface with.

For my purposes I used a Raspberry Pi 4 Model B since I had one laying around. Obviously something smaller and more stealthy would be recommended for the usage in penetration testing. You could use a Pi Zero or even better something like a Micro Beetle (if you don't mind switching the platform).
