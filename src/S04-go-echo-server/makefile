
.openapi-generator/VERSION:
	docker run --rm \
	-v "${PWD}:/local" \
	openapitools/openapi-generator-cli:v5.4.0 generate \
	-i /local/openapi.yaml \
	-g go-echo-server \
	-o /local


run:
	go mod download
	go build -o app

docker-build:
	docker build --network=host -t openapi .

docker-run: docker-build
	docker run --rm -it openapi
