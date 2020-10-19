
#!/usr/bin/env python

import time
import os
import signal

dir = "/home/pi/pixels/"

class SinglePixel:

    def __init__(self):
        self.kill = False

        self.signals = {
            signal.SIGINT: 'SIGINT',
            signal.SIGTERM: 'SIGTERM'
        }

        self.cf = ""
        self.BRIGHTNESS_1 = int(os.environ.get("BRIGHTNESS_1", 0))
        self.BRIGHTNESS_2 = int(os.environ.get("BRIGHTNESS_2", 0))
        self.BRIGHTNESS_3 = int(os.environ.get("BRIGHTNESS_3", 128))

        print('Brightness: {} {} {}'.format(self.BRIGHTNESS_1, self.BRIGHTNESS_2, self.BRIGHTNESS_3))

        signal.signal(signal.SIGINT, self.exit_gracefully)
        signal.signal(signal.SIGTERM, self.exit_gracefully)

    def run(self):
        for i in range(8):
            print(i)
            if(i == 8):                                     # If i == 8, the end has been reached and it should not try to enable anymore leds
                print("No more spots left")
                break
            elif not os.path.isfile(dir + str(i)):          # If the file is not there, create it, so the pixel controller will know to turn on the pixel
                print("Pixel ",i," available, activating!")
                self.cf = dir + str(i)
                with open(self.cf, 'w') as f:
                    f.write('{}:{}:{}'.format(self.BRIGHTNESS_1, self.BRIGHTNESS_2, self.BRIGHTNESS_3))
                break

    def exit_gracefully(self, signum, frame):
        print("\nReceived {} signal".format(self.signals[signum]))

        if self.cf != "":
            os.remove(self.cf)

        self.kill = True

if __name__ == '__main__':
    pk = SinglePixel()
    pk.run()

    while not pk.kill:
        time.sleep(1)

