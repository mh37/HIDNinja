#!/usr/bin/env bash

# Echo commands to stdout.
set -x

# Exit on first error.
set -e

# Treat undefined environment variables as errors.
set -u


#enable deice tree overlay and add dtoverlay=dwc2 to /boot/config.txt
if ! grep 'dtoverlay=dwc2' /boot/config; then
  echo "dtoverlay=dwc2" >> /boot/config.txt
fi
# Add dwc2 and libcomposite to /etc/modules
if ! grep dwc2 /etc/modules; then
  echo "dwc2" >> /etc/modules
fi

ENABLE_RPI_HID_PATH=/opt/enable-rpi-hid
ENABLE_RPI_HID_DIR=$(dirname $ENABLE_RPI_HID_PATH)

mkdir -p "$ENABLE_RPI_HID_DIR"
cp enableHIDmode "$ENABLE_RPI_HID_PATH"
chmod +x "$ENABLE_RPI_HID_PATH"

sed -e "s@/usr/bin/hidninja@${ENABLE_RPI_HID_PATH}@g" \
  usb-gadget.service > /lib/systemd/system/usb-gadget.service

# start service and enable run on boot
systemctl daemon-reload
systemctl enable usb-gadget.service

# reboot system
reboot