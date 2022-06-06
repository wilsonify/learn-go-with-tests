package marshal

import (
	"log"
	"net/http"
	marshal "learn.go/S02-build-an-app/c21-json/v2"
)

func main() {
	server := marshal.NewPlayerServer(marshal.NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}
