# Raspberry Pi 4b Hardware Setup and Configuration Guide

## Components Needed
1. Raspberry Pi Camera v2
2. Raspberry Pi 4b
3. Raspberry Pi TFT screen 7-inch touch display
4. SSD card

## Steps to Connect the Hardware Components

1. **Connect the Raspberry Pi Camera v2** to the Raspberry Pi camera port.
2. **Connect the Raspberry Pi 4b GPIO 2 pin** to the SDA pin on the Raspberry Pi TFT screen display.
3. **Connect the Raspberry Pi 4b power (5V) pin** to the 5V pin on the Raspberry Pi TFT display.
4. **Connect the Raspberry Pi 4b GPIO 5 pin** to the SCL pin on the Raspberry Pi TFT display.
5. **Connect the Raspberry Pi 4b ground pin** to the GND pin on the Raspberry Pi TFT display.

## Preparation Steps

1. **Download the Operating System to SSD card**:
   - Download the appropriate operating system for your task (e.g., Raspberry Pi OS Legacy Lite, Debian Bullseye with desktop environment).

### Tips
- **Avoid Ubuntu**: It lacks pre-installed packages, requiring significant installation time. Without proper cooling, this can overheat the processor.
- **Camera Compatibility**: The camera may only work with operating systems that include Debian Bullseye.

## How to Set Up the Raspberry Pi 4b

1. **Ensure all parts are connected properly**:
   - Tip: Do not disconnect and reconnect hardware components during setup to avoid damage.
2. **Insert the Micro SD Card**.
3. **Connect Mouse & Keyboard**.
4. **Connect the Power Supply**.
5. **Follow On-Screen Instructions**.

## Steps to Enable the Camera

1. **Enable the Camera through the Raspberry Pi Configuration Tool**:
   - Go to Preferences > Raspberry Pi Configuration > Interfaces and enable the camera.
2. If the camera is not detected:
   - Open Terminal and run: 
     ```bash
     sudo raspi-config
     ```
   - Enable the camera through the Interface Options.
   - Tip: Check if the camera is working with:
     ```bash
     libcamera-hello
     ```
3. If the camera still isn't working:
   - Edit the configuration file:
     ```bash
     sudo nano /boot/config.txt
     ```
   - Change the camera section from `start_x=1` to `camera_auto_detect=1`.
   - Tip: Verify the camera functionality again.
4. If the camera still doesn't respond:
   - Check the connection.
   - Expand the file system:
     ```bash
     sudo raspi-config
     ```
     Go to Advanced Options > Expand File System.
   - Tip: Ensure you have the correct operating system and a stable Wi-Fi connection.

## Implementation of the Code

The initial goal was to use OpenCV for video processing and facial recognition. However, due to hardware limitations, we opted to implement a neural network on a server to process images and compare them to a database.

### Errors Encountered

1. **Web Server Hosting Issues**:
   - Attempted using Apache and Nginx but faced difficulties connecting Python scripts to the web server.
   
2. **Steps Taken**:
   1. Created a basic HTML page with a button.
   2. Created a Python script to capture a picture when the button is pressed.
   3. Developed HTML success and fail pages.
   4. Created a client page to handle basic logic.

### Watchdog Library Issues

The `watchdog` library monitors folder changes to detect when a picture is taken.

### Steps to Resolve Watchdog Issues

1. **Install Watchdog**:
   ```bash
   pip install watchdog
   sudo apt-get install python3-pip python3-dev
   sudo apt-get install build-essential
   ```

2. **Check WDT Module in Kernel**:
   ```bash
   sudo cat /lib/modules/$(uname -r)/modules.builtin | grep wdt
   sudo cat /var/log/kern.log* | grep watchdog
   ```

3. **Enable Watchdog**:
   - Add `dtparam=watchdog=on` in `/boot/config.txt`.
   - Install and enable watchdog:
     ```bash
     sudo apt install watchdog
     sudo systemctl enable watchdog
     ```

4. **Configure WDT Service**:
   - Edit `/etc/watchdog.conf`:
     ```bash
     max-load-1 = 24
     watchdog-device = /dev/watchdog
     realtime = yes
     priority = 1
     ```

5. **Test WDT Service**:
   - Simulate heavy load:
     ```bash
     # :(){ :|: & };:
     ```
   - **Warning**: This command will reboot your device.

## References

- [Raspberry Pi Watchdog Setup](https://gist.github.com/PSJoshi/803a0419e568cc95c6bec24ebb0d44dc)