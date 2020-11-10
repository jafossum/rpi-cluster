# GO Bcrypt Validate OpenFaaS Function

GO has a very nice library for encrypting passwords ands strings using bcrypt.
The resulting hashed string can be stored in clear text, and will be sendt together
with the clear text password for validation.

This Validate functionality is now wrapped in a OpenFaaS function

## JSON Body

The Validate function only accepts a POST request with the following JSON body

```json
{
  "text": "some-passwd",
  "hash": "$2a$10$jmIre8ZqAYenwc5y8gTNe.DMTrbuVjCwoc/g6LPvmjcX3tcqtzbvK"
}
```

# Create an OpenFaas Function

Create a folder for your function

    $ mkdir bcrypt-validate && cd bcrypt-validate

Get the OpenFaas templates

    $ faas-cli template pull

Get a custom template. In this cae the golang-middleware template

    $ faas-cli template store pull golang-middleware

## Build the function

Create a new OpenFaaS function using the Python3-HTTP template

    $ faas-cli new bcrypt-validate --lang golang-middleware --prefix=<YOUR-DOCKER-ID>

Push your code to OpenFaaS

    $ faas-cli up -f bcrypt-validate.yml --build-arg GO111MODULE=on

## Invoke your function

Invoke your function through the OpenFaaS web console, curl, or with faas-cli
(Remember to put a `\` before each \$ in the resulting hash in the JSON request)

    $ curl -d "{\"text\":\"some-passwd\",\"hash\":\"\$2a\$10\$vftrptYx21rjd3v09Vhmn.Gw3dsDDxzY6RD.hA10KfgCuYltV9.wq\"}" http://127.0.0.1:8080/function/bcrypt-validate

### Test with Docker

If OpenFaaS is not running, you can test it using the resulting docker container

    $ docker run --rm -p 8080:8080 jafossum/bcrypt-validate:latest

Invoke with curl (Remember to put a `\` before each \$ in the resulting hash in the JSON request)

    $ curl -d "{\"text\":\"some-passwd\",\"hash\":\"\$2a\$10\$vftrptYx21rjd3v09Vhmn.Gw3dsDDxzY6RD.hA10KfgCuYltV9.wq\"}" localhost:8080

## Read the Logs

All OpenFaaS functions run in the `openfaas-fn` namespace

    $ kubectl get pods -n openfaas-fn
    $ kubectl logs bcrypt-validate-d6c8cdf6d-rnchn -n openfaas-fn

## Remove the function from OpenFaaS

List the functions on the cluster

    $ faas-cli list

Remove the selected function

    $ faas-cli remove bcrypt-validate
