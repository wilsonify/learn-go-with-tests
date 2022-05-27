package main

import (
	"os"
	"time"

	"github.com/wilsonify/learn-go-with-tests/math/v10/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
