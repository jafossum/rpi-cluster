# Blinkt Web Server Demo

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
