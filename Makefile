.PHONY: all build test clean docker-build docker-push

# go params
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# binary names
RECOMMENDATION_BINARY=recommendation-engine
AGGREGATOR_BINARY=data-aggregator
GATEWAY_BINARY=api-gateway

# docker params
DOCKER_REGISTRY=your-registry
VERSION=latest

all: test build

build:
	$(GOBUILD) -o bin/$(RECOMMENDATION_BINARY) ./cmd/recommendation-engine
	$(GOBUILD) -o bin/$(AGGREGATOR_BINARY) ./cmd/data-aggregator
	$(GOBUILD) -o bin/$(GATEWAY_BINARY) ./cmd/api-gateway

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf bin/

docker-build:
	docker build -t $(DOCKER_REGISTRY)/$(RECOMMENDATION_BINARY):$(VERSION) -f deployments/docker/recommendation/Dockerfile .
	docker build -t $(DOCKER_REGISTRY)/$(AGGREGATOR_BINARY):$(VERSION) -f deployments/docker/aggregator/Dockerfile .
	docker build -t $(DOCKER_REGISTRY)/$(GATEWAY_BINARY):$(VERSION) -f deployments/docker/gateway/Dockerfile .

docker-push:
	docker push $(DOCKER_REGISTRY)/$(RECOMMENDATION_BINARY):$(VERSION)
	docker push $(DOCKER_REGISTRY)/$(AGGREGATOR_BINARY):$(VERSION)
	docker push $(DOCKER_REGISTRY)/$(GATEWAY_BINARY):$(VERSION)
