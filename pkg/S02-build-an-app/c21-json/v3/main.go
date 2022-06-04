package marshal

import (
	"log"
	"net/http"
)

func Main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}
