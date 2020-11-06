# OpenFaaS with Minikube

![](../../docs/resources/OpenFaaS.png)

Training with [OpenFaaS](https://www.openfaas.com) on a local [Minikube](https://minikube.sigs.k8s.io/docs/).

## OpenFaaS Workshops

The OpenFaaS team has put together a OpenFaaS training workshop. This can be found [here](https://github.com/openfaas/workshop).
Theis will describe install, setup and usage of OpenFaas.

The rest of this setup will describe OpenFaaS running in minikube on Mac OSX.

## Install and Start Minikube (Mac OS)

    $ brew update
    $ brew install minikube

Start minikube

    $ minikube start
    $ minikube status

This will set up a local kubernetes cluster running on you local machine.
Minikube will also install `kubernetes-cli (kubectl)`.

Test that the installation is working with the following command:

    $ kubectl get nodes

## Install OpenFaaS

### Install faas-cli

For instructions on installing faas-cli, see [this link](https://docs.openfaas.com/deployment/kubernetes/#install-the-faas-cli).
Herer we will be using brew

    $ brew install faas-cli

### Install arkade

Install arkade using this [guide](https://docs.openfaas.com/deployment/kubernetes/#1-deploy-the-chart-with-arkade-fastest-option)

You can use arkade to install OpenFaaS to a regular cloud cluster, your laptop, a VM, a Raspberry Pi, or a 64-bit ARM machine.

Using the latest Zsh for Mac OS, this command will work if the install page command complains

    $ curl -SLs https://dl.get-arkade.dev/ | sh

Test the install with the following

    $ arkade --help

### Install OpenFaaS

Now we are ready to install OpenFaaS. There are a lot of options for installing OpenFaas,
but here we will use arkade as an easy option

Install OpenFaaS using arkade on minikube

    $ arkade install openfaas --load-balancer

Follow the output, and you will get information about your current deployment.

## Acces OpenFaaS in minikube

### 1. Using port forward (recomended)

Port-forward the OpenFaaS gateway service to local port 8080

    $ kubectl rollout status -n openfaas deploy/gateway
    $ kubectl port-forward -n openfaas svc/gateway 8080:8080

#### Set up faas-cli

Now run the command in the the output section `# If basic auth is enabled, you can now log into your gateway:`
to setup the faas-cli. This now defaults to http://127.0.0.1:8080

    $ PASSWORD=$(kubectl get secret -n openfaas basic-auth -o jsonpath="{.data.basic-auth-password}" | base64 --decode; echo)
    $ echo -n $PASSWORD | faas-cli login --username admin --password-stdin

faas-cli should now be connected to the minikube OpenFaaS deployment. Check with

    $ faas-cli list

### 2. Using minikube service

Minikube needs to export the LoadBalancer service, so that it can be accessed from the outside.
This will enable us to access the OpenFaaS web, and use `faas-cli`

    $ minikube service gateway-external -n openfaas

#### Set up fass-cli

When installing OpenFaaS using arkade, the output section `# If basic auth is enabled, you can now log into your gateway:`
gives a command for setting the password in the openfaas secret into a \$PASSWORD env var.

This is explained below

**Output from minikube service**

```
🏃  Starting tunnel for service gateway-external.
|-----------|------------------|-------------|------------------------|
| NAMESPACE |       NAME       | TARGET PORT |          URL           |
|-----------|------------------|-------------|------------------------|
| openfaas  | gateway-external |             | http://127.0.0.1:51869 |
|-----------|------------------|-------------|------------------------|
```

**faas-cli login**

Change the login to use the minikube extrernal URL

    $ export OPENFAAS_URL="http://127.0.0.1:51869"
    $ PASSWORD=$(kubectl get secret -n openfaas basic-auth -o jsonpath="{.data.basic-auth-password}" | base64 --decode; echo)
    $ echo -n $PASSWORD | faas-cli login --username admin --password-stdin

faas-cli should now be connected to the minikube OpenFaaS deployment. Check with

    $ faas-cli list

### Login OpenFaaS Web

Using the command from the arkade install, the admin password should now be stored in the \$PASSWORD env var

    $ echo -n $PASSWORD

The above xommand should give the password, but the last character might not be included.
To verify the password, run the following

    $ kubectl get secret -n openfaas basic-auth -o jsonpath="{.data.basic-auth-password}" | base64 --decode; echo

Use this passwrod in the web login fomr for user `admin`

## OpenFaaS Workshops

The OpenFaaS team has put together a OpenFaaS training workshop. This can be found [here](https://docs.openfaas.com/tutorials/workshop/)

## OpenFaas Function Examples

- [Python3-HTTP](test-function/README.md)
- [Golang-HTTP](test-function-go/README.md)

## Delete OpenFaas deployment

Clean up the resources created

    $ kubectl delete all --all -n openfaas
    $ kubectl delete --all secrets -n openfaas
    $ kubectl delete --all configmaps -n openfaas

**Stop minikube Cluster**

Stop the minikube cluster with the following command

    $ minikube stop

**Delete minikube Cluster**

Stop the entire minikube cluster with the following command

    $ minikube delete
