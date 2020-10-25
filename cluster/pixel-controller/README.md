# Blinkt Pixel Controller Demo

Blinkt pixel controller demo inspired by [Sheldonwl](https://github.com/Sheldonwl/rpi-travel-case)

## pixel-controller

Simple python application that checks if a file exists on the `/home/pi/pixel` directory in a shared volume on each node (Raspberry Pi).
If the file 0 through 7 exists, the pixel-controller will enable the corresponding pixel on the Blinkt 8 LED RGB strip.
Using files will make the kubernetes cluster easier to demo.

**pixel-controller.py**

Runs an endless loop, where it checks if any of the files names 0 through 7 are available on disk and updates the LED's accordingly.

**Variables:**

- BRIGHTNESS_1 = Give this variable a number between 0 and 255 to set the brightness of LED color 1.
- BRIGHTNESS_2 = Give this variable a number between 0 and 255 to set the brightness of LED color 2.
- BRIGHTNESS_3 = Give this variable a number between 0 and 255 to set the brightness of LED color 3.
- CLEAR_SLEEP = With this variable, you can adjust the sleep time in-between LED updates

**pixel-controller.yaml**

Kubernetes DaemonSet that runs the pixel-controller app and attaches it to a volume that is shared between all single-pixel replicas running on the node. The CLEAR_SLEEP and BRIGHTNESS-X variables are set in this manifest.
BRIGHTNESS_X are atempted read from the 0-7 files, so the manifest holds the default / fallback values
