.PHONY: test
ROOT_DIR = $(shell pwd)

#################################################################################
# RUN COMMANDS
#################################################################################
run:
	docker compose --file ./docker-compose.yml --project-directory . up --build; \
	docker compose --file ./docker-compose.yml --project-directory . down --volumes;

#################################################################################
# LINT COMMANDS
#################################################################################
tidy:
	goimports -w .
	gofmt -s -w .
	go mod tidy

#################################################################################
# BUILD COMMANDS
#################################################################################
protoc:
	rm -rf pkg/*
	protoc -I=proto --go_out=. proto/*.proto

#################################################################################
# TEST COMMANDS
#################################################################################
generate-messages:
	CGO_ENABLED=1 go run cmd/generate-kafka-messages/main.go
