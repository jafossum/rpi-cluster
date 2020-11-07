# GO Bcrypt Encrypt OpenFaaS Function

GO has a very nice library for encrypting passwords ands strings using bcrypt.
The resulting hashed string can be stored in clear text, and will be sendt together
with the clear text password for validation.

This Encrypt functionality is now wrapped in a OpenFaaS function

## JSON Body

The Encrypt function only accepts a POST request with the following JSON body

```json
{
  "encrypt": "some-passwd",
  "cost": 10
}
```

The cost is optional, and will default to the bcrypt default value of 10. Must be between 4 and 31.

# Create an OpenFaas Function

Create a folder for your function

    $ mkdir bcrypt-encrypt && cd bcrypt-encrypt

Get the OpenFaas templates

    $ faas-cli template pull

Get a custom template. In this cae the golang-middleware template

    $ faas-cli template store pull golang-middleware

## Build the function

Create a new OpenFaaS function using the Python3-HTTP template

    $ faas-cli new bcrypt-encrypt --lang golang-middleware --prefix=<YOUR-DOCKER-ID>

Push your code to OpenFaaS

    $ faas-cli up -f bcrypt-encrypt.yml --build-arg GO111MODULE=on

## Invoke your function

Invoke your function through the OpenFaaS web console or with curl

    $ curl -d "{\"encrypt\":\"some-passwd\",\"cost\":10}" http://127.0.0.1:8080/function/bcrypt-encrypt

### Test with Docker

If OpenFaaS is not running, you can test it using the resulting docker container

    $ docker run --rm -p 8080:8080 jafossum/bcrypt-encrypt:latest

Invoke with curl

    $ curl -d "{\"encrypt\":\"some-passwd\",\"cost\":10}" localhost:8080

## Read the Logs

All OpenFaaS functions run in the `openfaas-fn` namespace

    $ kubectl get pods -n openfaas-fn
    $ kubectl logs bcrypt-encrypt-b8566d6bd-8hzqj -n openfaas-fn

## Remove the function from OpenFaaS

List the functions on the cluster

    $ faas-cli list

Remove the selected function

    $ faas-cli remove bcrypt-encrypt
