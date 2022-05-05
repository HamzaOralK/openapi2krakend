.PHONY: build dockerize test run

build:
	rm -rf ./build && \
	env GOOS=linux GOARCH=amd64 go build -o ./build/openapi2krakend ./pkg
	upx -9 ./build/openapi2krakend

dockerize: build
	docker buildx build --platform=linux/amd64 -f docker/Dockerfile -t okhuz/openapi2krakend:0.0.6 .

test:
	go test ./... -v

run:
	./scripts/run.sh