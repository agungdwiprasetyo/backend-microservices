# Stage 1
FROM golang:1.16.4-alpine3.13 AS dependency_builder

WORKDIR /go/src
ENV GO111MODULE=on

RUN apk update
RUN apk add --no-cache bash ca-certificates git

COPY go.mod .
COPY go.sum .

RUN go mod download

# Stage 2
FROM dependency_builder AS service_builder

ARG SERVICE_NAME
WORKDIR /usr/app

COPY sdk sdk
COPY services/$SERVICE_NAME services/$SERVICE_NAME
COPY go.mod .
COPY go.sum .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin services/$SERVICE_NAME/*.go

# Stage 3
FROM alpine:latest  

ARG SERVICE_NAME
ARG BUILD_NUMBER
RUN apk --no-cache add ca-certificates tzdata
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime
WORKDIR /root/
ENV WORKDIR=services/$SERVICE_NAME/
ENV BUILD_NUMBER=$BUILD_NUMBER

RUN mkdir -p /root/services/$SERVICE_NAME
RUN mkdir -p /root/services/$SERVICE_NAME/api
RUN mkdir -p /root/services/$SERVICE_NAME/api/configs
COPY --from=service_builder /usr/app/bin bin
COPY --from=service_builder /usr/app/services/$SERVICE_NAME/.env /root/services/$SERVICE_NAME/.env
COPY --from=service_builder /usr/app/services/$SERVICE_NAME/api /root/services/$SERVICE_NAME/api

ENTRYPOINT ["./bin"]
