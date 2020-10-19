#!/usr/bin/env python

import time
import blinkt
import os

dir = "/home/pi/pixels/"
BRIGHTNESS_1 = os.environ.get("BRIGHTNESS_1", 0)
BRIGHTNESS_2 = os.environ.get("BRIGHTNESS_2", 0)
BRIGHTNESS_3 = os.environ.get("BRIGHTNESS_3", 128)
CLEAR_SLEEP = os.environ.get("CLEAR_SLEEP", 0.1)

blinkt.set_clear_on_exit()

while True:
    no_pixels = 8
    for i in range(0,8):
        if os.path.isfile(dir + str(i)):
            with open(dir + str(i), 'r') as f:
                data = f.read()
                try:
                    br = data.split(':')
                    BRIGHTNESS_1 = int(br[0])
                    BRIGHTNESS_2 = int(br[1])
                    BRIGHTNESS_3 = int(br[2])
                except:
                    pass
            print('Set pixel {}: {} {} {}'.format(i, BRIGHTNESS_1, BRIGHTNESS_2, BRIGHTNESS_3))
            blinkt.set_pixel(i, BRIGHTNESS_1, BRIGHTNESS_2, BRIGHTNESS_3)
            no_pixels += 1
        else:
            blinkt.set_pixel(i, 0, 0, 0)
            no_pixels -= 1

    blinkt.show()
    time.sleep(float(CLEAR_SLEEP))

    if no_pixels == 0:
        for i in range(0,8):
            blinkt.clear()
            blinkt.set_pixel(i, 0, 0, 0)
        blinkt.show()
