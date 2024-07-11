# Auth Function for OCI API Gateway

Two samples of how create an Auth Function for OCI API Gateway.

Bouth codes do the same thing, they receive a body like the example bellow. Compare the "token" field with the secreat storaged on OCI Secrets and return true or false.

```json
{"type": "token", "token":"Samples_Token"}
```

- Python: A python sample that uses the default docker image from fn project.
- Go: A Go Samples that uses a custom docker image.