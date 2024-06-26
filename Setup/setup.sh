#!/usr/bin/env bash

# Echo commands to stdout.
set -x

# Exit on first error.
set -e

# Treat undefined environment variables as errors.
set -u

#check if dwc2 exists in both the boot config and etc/modules, this is required for the USB gadget mode to be enabled
#enable deice tree overlay and add dtoverlay=dwc2 to /boot/config.txt
if ! grep 'dtoverlay=dwc2' /boot/config; then
  echo "dtoverlay=dwc2" >> /boot/config.txt
fi
# Add dwc2 to /etc/modules
if ! grep dwc2 /etc/modules; then
  echo "dwc2" >> /etc/modules
fi

#set path for secondary script that will enable the HID gadget mode 
ENABLE_RPI_HID_PATH=/opt/enable-rpi-hid
ENABLE_RPI_HID_DIR=$(dirname $ENABLE_RPI_HID_PATH)

mkdir -p "$ENABLE_RPI_HID_DIR"
cp enableHIDmode "$ENABLE_RPI_HID_PATH"
chmod +x "$ENABLE_RPI_HID_PATH"
# setup a service that can run on boot to preserve functionality after a reboot
sed -e "s@/usr/bin/hidninja@${ENABLE_RPI_HID_PATH}@g" \
  usb-gadget > /lib/systemd/system/usb-gadget.service

# start service and enable run on boot
systemctl daemon-reload
systemctl enable usb-gadget.service

# reboot system
reboot