# Order-Service

## Build and run service
```
$ make run
```
If include GRPC handler, run `$ make proto` for generate rpc files from proto (must install `protoc` compiler min version `libprotoc 3.14.0`)

## Run unit test & calculate code coverage
```
$ make test
```

## Create docker image
```
$ make docker
```
