# Create an OpenFaas Function

Create a folder for your function

    $ mkdir test-function && cd test-function

Get the OpenFaas templates

    $ faas-cli template pull

Get a custom template. In this cae the python3-http template

    $ faas-cli template store pull python3-http

## Build a Python3-HTTP function

Create a new OpenFaaS function using the Python3-HTTP template

    $ faas-cli new test-function --lang python3-http --prefix=<YOUR-DOCKER-ID>

Push your code to OpenFaaS

    $ faas-cli up -f test-function.yml

## Invoke your function

Invoke your function through the OpenFaaS web console, curl, or with faas-cli

    $ curl http://127.0.0.1:50569/function/test-function
    $ faas-cli invoke test-function

## LoadTest the function with Curl

Run the following command, replacing the adress with the function path

    $ while true; do curl http://127.0.0.1:50569/function/test-function ; sleep 0.1; done

Load test

    $ for run in {1..10000}; do curl http://127.0.0.1:50569/function/test-function ; done

On the last run you should see the OpenFaaS console start to start more replicas of the function to handle the load

## Read the Logs

All OpenFaaS functions run in the `openfaas-fn` namespace

    $ kubectl get pods -n openfaas-fn
    $ kubectl logs test-function-58899b5bfc-qgpwz -n openfaas-fn

## Remove the function from OpenFaaS

List the functions on the cluster

    $ faas-cli list

Remove the selected function

    $ faas-cli remove test-function
