package poker

import (
	"fmt"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func Main() {
	store, close, err := FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	NewCLI(store, os.Stdin).PlayPoker()
}
