package clockface

import (
	"os"
	"time"
)

func Mainly() {
	t := time.Now()
	SVGWriter(os.Stdout, t)
}
