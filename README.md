# HIDNinja

## Introduction

HIDNinja is Linux based wireless HID injector with a remote payload interface.

## Required Hardware

This project is based on a Linux based SBC which will be mimicking an HID device and is capable of receiving payloads over an onboard WNIC. It would be also possible to adjust this project for a more compact hardware setup and alternative communication methods could be used to interface with.

For my purposes I used a Raspberry Pi 4 Model B since I had one laying around. Obviously something smaller and more stealthy would be recommended for the usage in penetration testing. You could use a Pi Zero or even better something like a Micro Beetle (if you don't mind switching the platform).
