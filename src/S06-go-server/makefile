
.openapi-generator/VERSION:
	docker run --rm \
	-v "${PWD}:/local" \
	openapitools/openapi-generator-cli:v5.4.0 generate \
	-i /local/openapi.yaml \
	-g go-server \
	-o /local

run:
	go run main.go

docker-build:
	docker build -t go-server .

docker-run:
	docker run --rm -d -p 49154:8080 go-server

go.mod:
	go mod init go-server

go.sum: go.mod
	go list && go mod tidy