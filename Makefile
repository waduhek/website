.PHONY: build
build:
	go build -o ./build/website ./cmd/website

.PHONY: run
run: build
	./build/website

.PHONY: test
test:
	go test ./...

.PHONY: test-cover
test-cover:
	go test -coverprofile=/tmp/cover.out ./...

.PHONY: browse-cover
browse-cover:
	go tool cover -html=/tmp/cover.out

.PHONY: lint
lint:
	go vet ./...

.PHONY: format
format:
	go fmt ./...

.PHONY: build-infra
build-infra:
	go build -o ./build/infra ./cmd/infra

.PHONY: preview-infra
preview-infra: build-infra
	pulumi preview

.PHONY: apply-infra
apply-infra: build-infra
	pulumi up

# NOTE: Intended to be used only for local development
.PHONY: build-docker
build-docker:
	docker build --tag website/server:latest .

# NOTE: Intended to be used only for local development
.PHONY: run-docker
run-docker:
	docker run --env-file .env --publish 8080:8080 --rm --network website --name website-server website/server:latest

.PHONY: build-nginx
build-nginx:
	docker build --build-arg NGINX_VERSION=1.29.5 --tag website/nginx:latest ./nginx

.PHONY: run-nginx
run-nginx:
	docker run --publish 80:80 --rm --network website website/nginx:latest

.PHONY: build-otelcol
build-otelcol:
	docker build --tag website/otelcol:latest ./otelcol
