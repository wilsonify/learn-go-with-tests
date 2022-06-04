package httpserver

import (
	"log"
	"net/http"
)

func Main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
