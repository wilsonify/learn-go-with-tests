package poker

import (
	"log"
	"net/http"
)

func main_webserver() {
	store, close, err := FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5000", server))
}
