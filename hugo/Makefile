# https://gist.github.com/mpneuried/0594963ad38e68917ef189b4e6a269db

cnf ?= config.env
include $(cnf)
export $(shell sed 's/=.*//' $(cnf))

dpl ?= deploy.env
include $(dpl)
export $(shell sed 's/=.*//' $(dpl))

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

build: ## Build the container
	docker build -t $(DOCKER_REPO_USER)/$(APP_NAME):v$(VERSION) .
	docker build -t $(DOCKER_REPO_USER)/$(APP_NAME):latest .

run: ## Run container on port configured in `config.env`
	docker run -i -t --rm --env-file=./config.env -p=$(PORT):$(PORT) --name="$(APP_NAME)" $(APP_NAME)

up: build run

stop: ## Stop and remove a running container
	docker stop $(APP_NAME); docker rm $(APP_NAME)

release: build-nc publish

publish: repo-login publish-latest publish-version

publish-latest: tag-latest ## Publish the `latest` taged container to ECR
	@echo 'publish latest to $(DOCKER_REPO)'
	docker push $(DOCKER_REPO_USER)/$(APP_NAME):latest

publish-version: tag-version ## Publish the `{version}` taged container to ECR
	@echo 'publish $(VERSION) to $(DOCKER_REPO)'    
	docker push $(DOCKER_REPO_USER)/$(APP_NAME):v$(VERSION) 

repo-login: ## Logs in to dockerhub
	@echo 'logging into dockerhub'
	docker login

tag-latest: ## Generate container `{version}` tag
	@echo 'create tag latest'
	docker tag $(DOCKER_REPO_USER)/$(APP_NAME) $(DOCKER_REPO_USER)/$(APP_NAME):latest

tag-version: ## Generate container `latest` tag
	@echo 'create tag $(VERSION)'
	docker tag $(DOCKER_REPO_USER)/$(APP_NAME) $(DOCKER_REPO_USER)/$(APP_NAME):v$(VERSION)

version: ## Output the current version
	@echo $(VERSION)
