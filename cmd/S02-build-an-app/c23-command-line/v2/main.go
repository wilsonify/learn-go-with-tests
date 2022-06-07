package poker

import (
	"fmt"

	poker "learn.go/S02-build-an-app/c23-command-line/v2"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	poker.Server_main()
}
