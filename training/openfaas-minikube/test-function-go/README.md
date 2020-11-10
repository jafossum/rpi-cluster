# Create an OpenFaas Function

Create a folder for your function

    $ mkdir test-function-go && cd test-function-go

Get the OpenFaas templates

    $ faas-cli template pull

Get a custom template. In this cae the golang-middleware template

    $ faas-cli template store pull golang-http
    $ faas-cli template store pull golang-middleware

## Build a Python3-HTTP function

Create a new OpenFaaS function using the Python3-HTTP template

    $ faas-cli new test-function-go --lang golang-middleware --prefix=<YOUR-DOCKER-ID>

Push your code to OpenFaaS

    $ faas-cli up -f test-function-go.yml --build-arg GO111MODULE=on

## Invoke your function

Invoke your function through the OpenFaaS web console, curl, or with faas-cli

    $ curl http://127.0.0.1:50569/function/test-function-go
    $ faas-cli invoke test-function-go

## LoadTest the function 

### Curl

Run the following command, replacing the adress with the function path

    $ while true; do curl http://127.0.0.1:50569/function/test-function-go ; sleep 0.1; done

Load test

    $ for run in {1..10000}; do curl http://127.0.0.1:50569/function/test-function-go ; done

On the last run you should see the OpenFaaS console start to start more replicas of the function to handle the load

### LoadTester

There is provided a GO script for load-testing the REST endpoint

    $ cd load-function
    $ go run load-function.go -p <EXPOSED_OPENFAAS_PORT>

To list the different input arguments run:

    $ go run load-function.go -h

## Read the Logs

All OpenFaaS functions run in the `openfaas-fn` namespace

    $ kubectl get pods -n openfaas-fn-go
    $ kubectl logs test-function-go-58899b5bfc-qgpwz -n openfaas-fn

## Remove the function from OpenFaaS

List the functions on the cluster

    $ faas-cli list

Remove the selected function

    $ faas-cli remove test-function-go
