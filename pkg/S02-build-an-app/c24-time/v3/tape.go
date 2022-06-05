package poker

import (
	"os"
)

type Tape struct {
	FilerSeeker *os.File
}

func (t *Tape) Write(p []byte) (n int, err error) {
	t.FilerSeeker.Truncate(0)
	t.FilerSeeker.Seek(0, 0)
	return t.FilerSeeker.Write(p)
}
