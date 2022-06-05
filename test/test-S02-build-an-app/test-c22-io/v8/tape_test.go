package main

import (
	"io/ioutil"
	"testing"

	inputoutput "learn.go/S02-build-an-app/c22-io/v8"
)

func TestTape_Write(t *testing.T) {
	file, clean := CreateTempFile(t, "12345")
	defer clean()

	Tape := &inputoutput.Tape{file}

	Tape.Write([]byte("abc"))

	file.Seek(0, 0)
	newFileContents, _ := ioutil.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
