# Backend Microservices

## Made with
<p align="center">
  <img src="https://storage.googleapis.com/agungdp/static/logo/golang.png" width="80" alt="golang logo" />
  <img src="https://storage.googleapis.com/agungdp/static/logo/docker.png" width="80" hspace="10" alt="docker logo" />
  <img src="https://storage.googleapis.com/agungdp/static/logo/rest.png" width="80" hspace="10" alt="rest logo" />
  <img src="https://storage.googleapis.com/agungdp/static/logo/graphql.png" width="80" alt="graphql logo" />
  <img src="https://storage.googleapis.com/agungdp/static/logo/grpc.png" width="160" hspace="15" vspace="15" alt="grpc logo" />
  <img src="https://storage.googleapis.com/agungdp/static/logo/kafka.png" height="80" alt="kafka logo" />
</p>

This repository explain implementation of Go for building multiple microservices using a single codebase. Using [Standard Golang Project Layout](https://github.com/golang-standards/project-layout) and [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

## Create new service (for new project)
Please install [**candi**](https://github.com/agungdwiprasetyo/candi) CLI first (min version v1.3.3), and then:
```
$ make init
```
If include GRPC handler, run `$ make proto service={{service_name}}` after init service for generate rpc files from proto (must install `protoc` compiler min version `libprotoc 3.14.0`)

## Run specific service
```
$ make run service={{service_name}}
```

## Add module(s) in specific service (project)
```
$ make add-module service={{service_name}}
```

## Run unit test and calculate code coverage
* **Generate mocks first (using [mockery](https://github.com/vektra/mockery)):**
```
$ make mocks service={{service_name}}
```
* **Run test:**
```
$ make test service={{service_name}}
```

## Run sonar scanner
```
$ make sonar level={{level}} service={{service_name}}
```
`{{level}}` is service environment, example is one of `dev`, `staging`, or `prod`

## Create docker image a service
```
$ make docker service={{service_name}}
```

## Services

* [**Auth Service**](https://github.com/agungdwiprasetyo/backend-microservices/tree/master/services/auth-service)
* [**Line Chatbot**](https://github.com/agungdwiprasetyo/backend-microservices/tree/master/services/line-chatbot#line-chatbot-service)
* [**Notification Service**](https://github.com/agungdwiprasetyo/backend-microservices/tree/master/services/notification-service)
* [**Storage Service**](https://github.com/agungdwiprasetyo/backend-microservices/tree/master/services/storage-service)
* [**User Service**](https://github.com/agungdwiprasetyo/backend-microservices/tree/master/services/user-service)
