#!/usr/bin/env python

import time
import blinkt
import os
import signal

dir = "/home/pi/pixels/"

class PixelController:

    def __init__(self):
        self.kill = False

        self.signals = {
            signal.SIGINT: 'SIGINT',
            signal.SIGTERM: 'SIGTERM'
        }

        self.BRIGHTNESS_1 = os.environ.get("BRIGHTNESS_1", 0)
        self.BRIGHTNESS_2 = os.environ.get("BRIGHTNESS_2", 0)
        self.BRIGHTNESS_3 = os.environ.get("BRIGHTNESS_3", 128)
        self.CLEAR_SLEEP = os.environ.get("CLEAR_SLEEP", 0.1)

        blinkt.set_clear_on_exit()
        
        signal.signal(signal.SIGINT, self.exit_gracefully)
        signal.signal(signal.SIGTERM, self.exit_gracefully)

    def run(self):
        while not self.kill:
            no_pixels = 8
            for i in range(0,8):
                if os.path.isfile(dir + str(i)):
                    with open(dir + str(i), 'r') as f:
                        data = f.read()
                        try:
                            br = data.split(':')
                            self.BRIGHTNESS_1 = int(br[0])
                            self.BRIGHTNESS_2 = int(br[1])
                            self.BRIGHTNESS_3 = int(br[2])
                        except:
                            pass
                    print('Set pixel {}: {} {} {}'.format(i, self.BRIGHTNESS_1, self.BRIGHTNESS_2, self.BRIGHTNESS_3))
                    blinkt.set_pixel(i, self.BRIGHTNESS_1, self.BRIGHTNESS_2, self.BRIGHTNESS_3)
                    no_pixels += 1
                else:
                    blinkt.set_pixel(i, 0, 0, 0)
                    no_pixels -= 1

            blinkt.show()
            time.sleep(float(self.CLEAR_SLEEP))

            if no_pixels == 0:
                for i in range(0,8):
                    blinkt.clear()
                    blinkt.set_pixel(i, 0, 0, 0)
                blinkt.show()

    def exit_gracefully(self, signum, frame):
        print("\nReceived {} signal".format(self.signals[signum]))
        self.kill = True

if __name__ == '__main__':
    pk = PixelController()
    pk.run()

