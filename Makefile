APP?=testapp-battleship
export GO111MODULE=on
VERSION?=$(shell git log --date=short --pretty=format:'%ad-%h' -n 1 | sed -E 's/([0-9]{4})\-([0-9]{2})\-([0-9]{2})/\1.\2.\3/g')
DOCKER_SOURCE_DIR=/go/src/battleship

.PHONY: generate
generate:
	rm -rf app/generated/*
	docker run -it \
	-v $(PWD)/app/generated:$(DOCKER_SOURCE_DIR)/app/generated \
	-v $(PWD)/swagger:$(DOCKER_SOURCE_DIR)/swagger \
	-v $(PWD)/gogi.yaml:$(DOCKER_SOURCE_DIR)/gogi.yaml \
	-w="$(DOCKER_SOURCE_DIR)" \
	gotools.docker.lamoda.ru/gogi generate

# *** DEV ***
.PHONY: dev-server
dev-server:
	docker-compose -f deployments/docker-compose.yaml up -d battleship-db

.PHONY: dev-goose-migrate
dev-goose-migrate:
	docker-compose -f deployments/docker-compose.yaml up battleship-goose-migrate