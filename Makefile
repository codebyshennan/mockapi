.PHONY: static
SHELL := /bin/bash

APP_VERSION := $$(cat .version)
IMG_TAG := gcr.io/$(PROJECT_ID)/mockserver-:$(APP_VERSION)

dev-be:
	go run ./cmd/server/server.go

dev-air:
	air

dev-fe:
	cd web; yarn; yarn dev

test:
	echo TODO

ci-test: test

ci-build:
	cd web; npm i -g yarn; yarn; yarn build; cp -r dist/* ../cmd/server/; 
	docker build \
	--build-arg MONGO_URI=mongodb://localhost:27017 \
	--build-arg DB_NAME=api-03 \
	--build-arg JWT_SIGNING_KEY=123 \
	--build-arg API_MODE=mock \
	--build-arg GOOGLE_CLIENT_ID=257544260761-mjmgllu9gvvbt4vut09urq4ridltl80m.apps.googleusercontent.com \
	-t $(IMG_TAG) .

ci-push:
	docker push $(IMG_TAG)
