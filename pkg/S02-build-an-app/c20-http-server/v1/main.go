package httpserver

import (
	"log"
	"net/http"

	hs "github.com/wilsonify/S02-build-an-app/c20-http-server/v1"
)

func Main() {
	handler := http.HandlerFunc(hs.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
