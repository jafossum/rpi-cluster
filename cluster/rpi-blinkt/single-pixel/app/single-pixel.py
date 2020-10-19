
#!/usr/bin/env python

import time
import os

dir = "/home/pi/pixels/"
cf = ""
SLEEP = int(os.environ.get("SLEEP",60))
BRIGHTNESS_1 = int(os.environ.get("BRIGHTNESS_1", 0))
BRIGHTNESS_2 = int(os.environ.get("BRIGHTNESS_2", 0))
BRIGHTNESS_3 = int(os.environ.get("BRIGHTNESS_3", 128))

print('Brightness: {} {} {}, Sleep: {}'.format(BRIGHTNESS_1, BRIGHTNESS_2, BRIGHTNESS_3, SLEEP))

for i in range(8):
    print(i)
    if(i == 8):                                     # If i == 8, the end has been reached and it should not try to enable anymore leds
        print("No more spots left")
        break
    elif not os.path.isfile(dir + str(i)):          # If the file is not there, create it, so the pixel controller will know to turn on the pixel
        print("Pixel ",i," available, activating!")
        cf = dir + str(i)
        with open(cf, 'w') as f:
            f.write('{}:{}:{}'.format(BRIGHTNESS_1, BRIGHTNESS_2, BRIGHTNESS_3))
        break

try:
    time.sleep(SLEEP)
finally:
    if cf != "":
        os.remove(cf)
