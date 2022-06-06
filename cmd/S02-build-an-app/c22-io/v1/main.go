package inputoutput

import (
	"log"
	"net/http"

	inputoutput "learn.go/S02-build-an-app/c22-io/v1"
)

func main() {
	server := inputoutput.NewPlayerServer(inputoutput.NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}
