// Writes an SVG clockface of the current time to Stdout.
package main

import (
	"os"
	"time"
)

func main() {
	t := time.Now()
	Write(os.Stdout, t)
}
