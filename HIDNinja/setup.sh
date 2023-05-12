#!/usr/bin/env bash

#enable deice tree overlay and add dtoverlay=dwc2 to /boot/config.txt
echo "dtoverlay=dwc2" | sudo tee -a /boot/config.txt
# Add dwc2 and libcomposite to /etc/modules
echo "dwc2" | sudo tee -a /etc/modules
echo "libcomposite" | sudo tee -a /etc/modules

ENABLE_RPI_HID_PATH=/opt/enable-rpi-hid
ENABLE_RPI_HID_DIR=$(dirname $ENABLE_RPI_HID_PATH)

mkdir -p "$ENABLE_RPI_HID_DIR"
cp enableHIDmode "$ENABLE_RPI_HID_PATH"
chmod +x "$ENABLE_RPI_HID_PATH"

cp usb-gadget.service usr/lib/systemd/system/usb-gadget.service

# start service and enable run on boot
systemctl daemon-reload
systemctl enable usb-gadget.service