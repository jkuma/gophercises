package stkhtml

import (
	"github.com/jkuma/gophercises/recover_chroma/fcnt"
	"testing"
)

func TestHtmlOutput(t *testing.T) {
	f, err := fcnt.GetFileContent("test/debugtrace")

	if err != nil {
		t.Error(err)
	}

	html, err := HtmlOutput(f.Content)

	if err != nil {
		t.Error(err)
	}

	if len(html) == 0 {
		t.Error("The html file should not be empty")
	}
}
