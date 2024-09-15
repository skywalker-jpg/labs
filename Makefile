GO_BUILD_PATH ?= bin
GO_BUILD_LABS_PATH ?= $(GO_BUILD_PATH)/labs/

PATH := $(PATH):/home/linuxbrew/.linuxbrew/opt/go/libexec/bin
export PATH



.PHONY: build
build:
	CGO_ENABLED=$(CGO) GOOS=linux go build -o $(GO_BUILD_LABS_PATH) ./cmd/labs/

.PHONY: up
up:
	docker compose up --build -d

.PHONY: down
down:
	@sudo docker ps -q | xargs -r sudo docker stop
	@sudo docker ps -aq | xargs -r sudo docker rm

.PHONY: clean
clean:
	@rm -rf ./bin