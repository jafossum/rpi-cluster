# Blinkt Pixel Controller Demo

Blinkt pixel controller demo inspired by [Sheldonwl](https://github.com/Sheldonwl/rpi-travel-case)

## single-pixel

Simple Python application that checks if a file between 0 and 7 exists on the attached volume and creates the next iterated number if the limit of 8 LED's has not been filled.

Variables:

- BRIGHTNESS_1 = Give this variable a number between 0 and 255 to set the brightness of LED color 1.
- BRIGHTNESS_2 = Give this variable a number between 0 and 255 to set the brightness of LED color 2.
- BRIGHTNESS_3 = Give this variable a number between 0 and 255 to set the brightness of LED color 3.
- SLEEP = Sets the sleep interval after running the script. This is set to 60 default.

**single-pixel.py**

Checks if a file named 0 through 7 exists, if not, it creates the next iterated number of any existing file.
BRIGHTNESS_X are written to the 0-7 files.

**single-pixel.yaml**

Contains a single-pxel deployment
