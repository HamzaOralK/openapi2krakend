.PHONY: build dockerize

build:
	rm -rf ./build && \
	env GOOS=linux GOARCH=amd64 go build -o ./build/openapi2krakend ./pkg

dockerize: build
	docker buildx build --platform=linux/amd64 -f docker/Dockerfile -t okhuz/openapi2krakend:latest .

