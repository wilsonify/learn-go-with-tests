package httpserver

import (
	"log"
	"net/http"

	httpserver "learn.go/S02-build-an-app/c20-http-server/v5"
)

func main() {
	server := &httpserver.PlayerServer{httpserver.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
