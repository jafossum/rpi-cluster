# Blinkt Pixel Controller Demo

Blinkt pixel controller demo inspired by [Sheldonwl](https://github.com/Sheldonwl/rpi-travel-case)

## pixel-controller

Simple python application that checks if a file exists on the `/home/pi/pixel` directory in a shared volume on each node (Raspberry Pi). 
If the file 0 through 7 exists, the pixel-controller will enable the corresponding pixel on the Blinkt 8 LED RGB strip.
Using files will make the kubernetes cluster easier to demo. 

**pixel-controller.py**

Runs an endless loop, where it checks if any of the files names 0 through 7 are available on disk and updates the LED's accordingly.

**Variables:**

- BRIGHTNESS = Give this variable a number between 0 and 255 to set the brightness of the LED's.
- CLEAR_SLEEP = With this variable, you can adjust the sleep time in-between LED updates

**pixel-controller.yaml**

Kubernetes DaemonSet that runs the pixel-controller app and attaches it to a volume that is shared between all single-pixel replicas running on the node. The CLEAR_SLEEP and BRIGHTNESS variables are set in this manifest.

## single-pixel

Simple Python application that checks if a file between 0 and 7 exists on the attached volume and creates the next iterated number if the limit of 8 LED's has not been filled.

Variables:

SLEEP = Sets the sleep interval after running the script. This is set to 60 default.

**single-pixel.py**

Checks if a file named 0 through 7 exists, if not, it creates the next iterated number of any existing file.

**single-pixel.yaml**

Contains a preStop command, that deletes the last created file. Meaning, if the file named 5 was created last, deleting the replica will delete that file named 5, thus removing one LED from the strip.
