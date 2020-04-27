FROM golang:1.12.7-alpine3.10

ARG SERVICE_NAME
ARG LOG_DIR=/${SERVICE_NAME}/logs
ARG BUILD_PACKAGES="git curl make g++ tzdata build-base autoconf automake libtool"

RUN mkdir -p ${LOG_DIR}

WORKDIR /usr/app
ENV SRC_DIR=/usr/app/
ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log

COPY . $SRC_DIR

RUN apk update && apk add --no-cache $BUILD_PACKAGES

# Install protoc
ENV PROTOBUF_URL https://github.com/google/protobuf/releases/download/v3.7.1/protobuf-cpp-3.7.1.tar.gz
RUN curl -L -o /tmp/protobuf.tar.gz $PROTOBUF_URL
WORKDIR /tmp/
RUN tar xvzf protobuf.tar.gz
WORKDIR /tmp/protobuf-3.7.1
RUN mkdir /export
RUN ./autogen.sh && \
    ./configure --prefix=/export && \
    make -j 3 && \
    make check && \
    make install

# Install protoc-gen-go
RUN go get github.com/golang/protobuf/protoc-gen-go
RUN cp /go/bin/protoc-gen-go /export/bin/

# Export dependencies
RUN cp /usr/lib/libstdc++* /export/lib/
RUN cp /usr/lib/libgcc_s* /export/lib/

RUN go mod download \
  && make prepare ${SERVICE_NAME} \
  && CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -o bin .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
RUN mkdir -p /root/api
RUN mkdir -p /root/config/key
RUN mkdir -p /root/cmd/${SERVICE_NAME}
RUN mkdir -p /root/web
COPY --from=0 /usr/app/bin .
COPY --from=0 /usr/app/.env .
COPY --from=0 /usr/app/api /root/api
COPY --from=0 /usr/app/cmd/${SERVICE_NAME} /root/cmd/${SERVICE_NAME}
COPY --from=0 /usr/app/config/key /root/config/key
COPY --from=0 /usr/app/web /root/web

ENTRYPOINT ["sh", "-c", "./bin"]