
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
	docker build --network=host -t openapi .

docker-run:
	docker run --rm -it openapi

