# Blinkt Web Server Demo

## pixel-controller

Simple python application that checks if a file exists on the `/home/pi/pixel` directory in a shared volume on each node (Raspberry Pi). 
If the file 0 through 7 exists, the pixel-controller will enable the corresponding pixel on the Blinkt 8 LED RGB strip.
Using files will make the kubernetes cluster easier to demo.
The RGB LED information is stored in the file.

**pixel-controller.py**

Runs an endless loop, where it checks if any of the files names 0 through 7 are available on disk and updates the LED's accordingly.

**Variables:**

- BRIGHTNESS_1 = Give this variable a number between 0 and 255 to set the brightness of LED color 1.
- BRIGHTNESS_2 = Give this variable a number between 0 and 255 to set the brightness of LED color 2.
- BRIGHTNESS_3 = Give this variable a number between 0 and 255 to set the brightness of LED color 3.
- CLEAR_SLEEP = With this variable, you can adjust the sleep time in-between LED updates

**pixel-controller.yaml**

Kubernetes DaemonSet that runs the pixel-controller app and attaches it to a volume that is shared between all web-app replicas running on the node. 
The CLEAR_SLEEP and BRIGHTNESS_X variables are set in this manifest.
BRIGHTNESS_X are atempted read from the 0-7 files, so the manifest holds the default / fallback values

## web-app

Simple GO application that runs a webserver at `:8081/blinkt`, and creates a random file at `/home/pi/pixel`, with random RGB colors. 
After a second the file will be deleted, and the respons returned.

**web-app.yaml**

Contains a Deployment and a LoadBalancer Service, that exposes thw webserver at port `:30000`.

## client-app

Standalone GO application that calls the LoadBalancer endpoint at port `:30000`.
It takes an optional argument `-n` that specifies the number of threds to run in paralell.
If the application i for example called like this:

    $ go run main.go -n 8

Then 10 seoarate threads will run and call the endpoint repeatedly at random times. The LEDs will now illustrate how the LoadBalancer distributes teh load accross the nodes in the cluster.