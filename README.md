### Description
Basic conversion from OpenAPI specification to Krakend config. This is extendable with custom OpenAPI 
attributes and more support for both krakend and openapi configurations.

x-timeout: In top level and in method level modifies timeout for the whole api or for a single endpoint 

### Arguments

-directory: folder where swagger files live. default is swagger folder in repository.
<br>
-encoding: backend encoding for whole endpoints. default is "json".
<br>
-global-timeout: sets global timeout, default is 3000ms 

### Usage

```shell 
go build -o openapi2krakend ./pkg
./openapi2krakend
``` 

### Dockerizing Application
```shell
docker build -t test/krakend:1.0.0 .
```

### Deployment to Kubernetes

````shell
kubectl apply -f ./deployment
````