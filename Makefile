.PHONY : prepare build run

$(eval $(service):;@:)

check:
	@[ "${service}" ] || ( echo "\x1b[31;1mERROR: 'service' is not set\x1b[0m"; exit 1 )
	@if [ ! -d "services/$(service)" ]; then  echo "\x1b[31;1mERROR: service '$(service)' undefined\x1b[0m"; exit 1; fi

prepare: check
	@if [ ! -f services/$(service)/.env ]; then cp services/$(service)/.env.sample services/$(service)/.env; fi;

init:
	@candi -packageprefix=monorepo/services -withgomod=false -protooutputpkg=monorepo/sdk -output=services/ -scope=1

add-module: check
	@candi -packageprefix=monorepo/services -withgomod=false -protooutputpkg=monorepo/sdk -output=services/ -servicename=$(service) -scope=2

build: check
	@go build -o services/$(service)/bin services/$(service)/*.go

run: build
	@WORKDIR="services/$(service)/" ./services/$(service)/bin

proto: check
	@if [ ! -d "sdk/$(service)" ]; then echo "creating new proto files..." &&  mkdir sdk/$(service) && mkdir sdk/$(service)/proto; fi
	$(foreach proto_file, $(shell find services/$(service)/api/proto -name '*.proto'),\
	protoc --proto_path=services/$(service)/api/proto --go_out=plugins=grpc:sdk/$(service)/proto \
	--go_opt=paths=source_relative $(proto_file);)

docker: check
	docker build --build-arg SERVICE_NAME=$(service) -t $(service):latest .

run-container:
	docker run --name=$(service) --network="host" -d $(service)

# mocks all interfaces in sdk for unit test
mocks:
	@mockery --all --keeptree --dir=sdk --output=./sdk/mocks
	@if [ -f sdk/mocks/Option.go ]; then rm sdk/mocks/Option.go; fi;

# unit test & calculate code coverage from selected service (please run mocks before run this rule)
test: check
	@echo "\x1b[32;1m>>> running unit test and calculate coverage for service $(service)\x1b[0m"
	@if [ -f services/$(service)/coverage.txt ]; then rm services/$(service)/coverage.txt; fi;
	@go test ./services/$(service)/... -cover -coverprofile=services/$(service)/coverage.txt -covermode=count -coverpkg=$(PACKAGES)
	@go tool cover -func=services/$(service)/coverage.txt

sonar: check
	@echo "\x1b[32;1m>>> running sonar scanner for service $(service)\x1b[0m"
	@[ "${level}" ] || ( echo "\x1b[31;1mERROR: 'level' is not set\x1b[0m" ; exit 1 )
	@sonar-scanner -Dsonar.projectKey=$(level)-$(service) \
	-Dsonar.projectName=$(level)-$(service) \
	-Dsonar.sources=./services/$(service) \
	-Dsonar.exclusions=**/mocks/**,**/module.go \
	-Dsonar.test.inclusions=**/*_test.go \
	-Dsonar.test.exclusions=**/mocks/** \
	-Dsonar.coverage.exclusions=**/mocks/**,**/*_test.go,**/main.go \
	-Dsonar.go.coverage.reportPaths=./services/$(service)/coverage.txt

generate-rsa-key:
	sh scripts/generate_rsa_key.sh

clear:
	rm -rf sdk/mocks services/$(service)/mocks services/$(service)/bin bin
