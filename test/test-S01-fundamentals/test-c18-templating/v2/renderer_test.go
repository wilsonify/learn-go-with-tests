package blogrendering_test

import (
	"bytes"
	"testing"

	blogrendering "learn.go/S01-fundamentals/c18-templating/v2"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrendering.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrendering.Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}
		got := buf.String()
		want := `<h1>hello world</h1>`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
