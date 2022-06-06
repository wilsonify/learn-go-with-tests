package poker_test

import (
	"io/ioutil"
	"testing"
	poker "learn.go/S02-build-an-app/c25-websockets/v2"
)

func TestTape_Write(t *testing.T) {
	file, clean := CreateTempFile(t, "12345")
	defer clean()

	Tape := &poker.Tape{File: file}

	Tape.Write([]byte("abc"))

	file.Seek(0, 0)
	newFileContents, _ := ioutil.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
