### Description
Basic conversion from OpenAPI specification to Krakend config. This is extendable with custom OpenAPI 
attributes and more support for both krakend and openapi configurations.

### Supported custom OpenAPI extensions

Example of custom extensions can be found in the `./swagger/pet-store.json`

x-timeout: At top level or at method level modifies timeout for the whole api or for a single endpoint 

### Usage

openapi2krakend can be run before the krakend container to generate krakend.json for krakend to consume.
Services can serve their swagger definitions or even from their documentation pages one can download those swagger
files and convert.

### Arguments

-directory: folder where swagger files live. default is swagger folder in repository.
<br>
-encoding: backend encoding for whole endpoints. default is "json".

### Usage

In make file image creation and build has been declared 

To build:
```shell 
make build
``` 

To dockerize
````shell
make dockerize
````

To run with sample environment variables
```shell
make run
```

### Deployment to Kubernetes

In deployment/deployment.yaml file you can set environment variables to "https://service-1/api-specification,https://service-2/api-specification"
then image will download all the specifications supplied and merge and convert all swagger files into a single krakend
configuration file.

#### Environment Variables for openApi2krakend

| Name             | Description                                                               | Type     | Default                                                | Required |
|------------------|---------------------------------------------------------------------------|----------|--------------------------------------------------------|:--------:|
| ENABLE_LOGGING   | Enable logging plugin for KrakenD                                         | `bool`   | `false`                                                |    no    |
| LOG_LEVEL        | Log level                                                                 | `string` | `WARNING`                                              |    no    |
| LOG_PREFIX       | Log prefix for filtering                                                  | `string` | `[KRAKEND]`                                            |    no    |
| LOG_SYSLOG       | Enable syslog                                                             | `bool`   | `true`                                                 |    no    |
| LOG_STDOUT       | Enable stdout                                                             | `bool`   | `true`                                                 |    no    |
| ENABLE_CORS      | Enable CORS plugin for KrakenD                                            | `bool`   | `false`                                                |    no    |
| ALLOWED_ORIGINS  | Comma seperated allowed origins, it will be used when ENABLE_CORS is true | `string` | `*`                                                    |    no    |
| ALLOWED_METHODS  | Comma seperated allowed methods, it will be used when ENABLE_CORS is true | `string` | `GET,HEAD,POST,PUT,DELETE,CONNECT,OPTIONS,TRACE,PATCH` |    no    |
 | GLOBAL_TIMEOUT   | Sets global timeout across all endpoints                                  | `string` | `3000ms`                                               |    no    |
| ENCODING         | Sets default encoding. Values are json, safejson, xml, rss, string, no-op | `string` | `json`                                                 |    no    |
| ADDITIONAL_PATHS | Comma seperated set of URLs to add every swagger for additional paths     | `string` | ``                                                     |    no    |

````shell
kubectl apply -f ./deployment
````
#### Contribution
Project itself quite easy to understand and maintain, and I am changing as I need it in my stack. Please feel free to contribute and open a PR if you see fit.