package inputoutput

import (
	"log"
	"net/http"
	"os"
	inputoutput "learn.go/S02-build-an-app/c22-io/v9"

)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := inputoutput.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server := inputoutput.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5000", server))
}
