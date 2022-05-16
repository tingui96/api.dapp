## 📙 OpenAPI Specification
To generate the OpenAPI specification run the fallowing command:

```$ swag init```

The current OpenAPI version used in this project is __2.2__. The integration package has not being ported to the 3 version yet.
For visiting the documentation open a browser in ``` http://localhost:8080/swagger/index.html ```.

___

❗ If we have external dependencies type in endpoint special comments, problems may occurs during the swagger doc generation process. So 
we need to do something like: 

1. import the package in the **main.go** file Eg:
```go
_ "github.com/ic-matcom/model-traceability-go"
```
2. Use this command below to generate swagger doc rather then ```$ swag init```:
```shell
$ swag init --parseDependency --parseInternal --parseDepth 1
```
or
```shell
$ swag init --parseDependency --parseInternal
```
[read more...](https://github.com/swaggo/swag/issues/817)
