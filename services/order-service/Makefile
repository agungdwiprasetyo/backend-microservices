.PHONY : build run

build:
	go build -o bin

run: build
	./bin

proto:
	$(foreach proto_file, $(shell find api/proto -name '*.proto'),\
	protoc --proto_path=api/proto --go_out=plugins=grpc:api/proto \
	--go_opt=paths=source_relative $(proto_file);)

docker:
	docker build -t order-service:latest .

run-container:
	docker run --name=order-service --network="host" -d order-service

# unit test & calculate code coverage
test:
	@if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@echo ">> running unit test and calculate coverage"
	@go test ./... -cover -coverprofile=coverage.txt -covermode=count -coverpkg=$(PACKAGES)
	@go tool cover -func=coverage.txt

clear:
	rm bin order-service
