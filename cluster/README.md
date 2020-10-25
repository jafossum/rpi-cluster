# Raspberry PI Demos

Demos for Kubernetes training.
Most of the examplesd uses the [Blinkt LED strip from Pimoroni](https://shop.pimoroni.com/products/blinkt)
LED strip to visualize how Kubernetes LoadBalances and runs different types of jobs in the cluster

## Pixel Controller

Most examples uses the `pixel-controller` ReplicaSet as a driver for controlling the blinkt LED connected to the host GPIO header
Examples wanting to use the LEDs should start this replicaset first, and then be able to control
the LEDs without actually needing the Blinkt driver

## Rpi Blinkt

Illustartes the creation and destruction of replicas across the cluster.
Each replica will light up the next LED available on the LED stip.

The python Deployment will light up blue leds, while the GO app wil light up green ones.

## Rpi Cron

Illustartes the Kubernetes CronJob setup.
The pod will start every minute and light a random LED in the cluster with a random color,
and keep it on for 10 seconds.

## Rpi Web

Illustartes the LoadBalancer functionality with a webserver pod running in the background.
This is not a replicaset, so pay attentyion to where the pods are running.

This also contains a simple clioent that can be run to call the API witn N number of threads using a random sleep between each call.
