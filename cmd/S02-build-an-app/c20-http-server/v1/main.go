package main

import (
	"log"
	"net/http"

	hs "S02-build-an-app/c20-http-server/v1"
)

func main() {
	handler := http.HandlerFunc(hs.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
