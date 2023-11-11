myLS:
	go build -o myLs main/main.go

clean:
	rm myLs

DOCKER_IMAGE_NAME := my-ls-1
DOCKER_TAG := latest

.PHONY: docker-build docker-run docker-test docker-clean

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) .

docker-run:
	docker run --rm -it $(DOCKER_IMAGE_NAME):$(DOCKER_TAG)

docker-clean:
	docker rmi $(DOCKER_IMAGE_NAME):$(DOCKER_TAG)
