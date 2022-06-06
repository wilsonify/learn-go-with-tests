package inputoutput

import (
	"log"
	"net/http"
	"os"
	inputoutput "learn.go/S02-build-an-app/c22-io/v6"

)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store := inputoutput.NewFileSystemPlayerStore(db)

	server := inputoutput.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5000", server))
}
