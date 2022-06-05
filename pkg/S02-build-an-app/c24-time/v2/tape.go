package poker

import (
	"os"
)

type Tape struct {
	FileSeeker *os.File
}

func (t *Tape) Write(p []byte) (n int, err error) {
	t.FileSeeker.Truncate(0)
	t.FileSeeker.Seek(0, 0)
	return t.FileSeeker.Write(p)
}
