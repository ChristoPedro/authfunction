# Auth Function for OCI API Gateway

Two samples of how create an Auth Function for OCI API Gateway.

Bouth codes do the same thing, they receive a body like the example bellow. Compare the "token" field with the secreat storaged on OCI Secrets and return true or false.

```json
{"type": "token", "token":"Samples_Token"}
```

- [Python](/authpy/): A python sample that uses the default docker image from fn project.
- [Go](/authfn/): A Go Samples that uses a custom docker image.

> **obs:** The Golang Function use a [Custom Dockerfile](https://docs.oracle.com/en-us/iaas/Content/Functions/Tasks/functionsusingcustomdockerfiles.htm) it allows to use diffent SO and lang versions to build your functions.

# How to use

## Create an OCI Function Application

https://docs.oracle.com/en-us/iaas/Content/Functions/Tasks/functionscreatingapps.htm

## Configure you fn enviroment

https://docs.oracle.com/en-us/iaas/Content/Functions/Tasks/functionsuploading.htm

## Create an OCI Secret

https://docs.oracle.com/en-us/iaas/Content/KeyManagement/Tasks/managingsecrets_topic-To_create_a_new_secret.htm

## Update Function with Secret OCI

For each function, update the func.yaml file with the ocid from the secret create earlier.

## Deploy the function

On the function folder execute:

```bash
fn deploy --app [You Application Name] -v
```

## Update API Gateway to use Functions as Auth

https://docs.oracle.com/en-us/iaas/Content/APIGateway/Tasks/apigatewayusingauthorizerfunction.htm