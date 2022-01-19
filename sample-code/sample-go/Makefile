# Image URL to use all building/pushing image targets
REPOSITORY ?= ethernoy/sample-go
IMG := $(REPOSITORY):latest
VERSION := v1.1.0

# Build sample app image
docker-build: 
	docker build . -t ${IMG}

# Tag for Release
docker-tag-release:
	@docker tag $(IMG) $(REPOSITORY):$(VERSION)

# Push for Release
docker-push-release: docker-tag-release
	@docker push $(REPOSITORY):$(VERSION)

.PHONY: vendor
vendor:
	go mod vendor
	go mod tidy