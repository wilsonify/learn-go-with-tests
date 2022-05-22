
.openapi-generator/VERSION:
	docker run --rm \
	-v "${PWD}:/local" \
	openapitools/openapi-generator-cli:v5.4.0 generate \
	-i /local/openapi.yaml \
	-g go-gin-server \
	-o /local

run:
	go run main.go

docker-build:
	docker build --network=host -t go-bin-server .

docker-run: docker-build
	docker run --rm -it openapi

