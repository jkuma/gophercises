package dbghtml

import (
	"github.com/jkuma/gophercises/recover_chroma/dbghtml/fcnt"
	"testing"
)

func TestHtml(t *testing.T) {
	f, err := fcnt.GetFileContent("test/debugtrace")

	if err != nil {
		t.Error(err)
	}

	html := Html(f.Content)

	if len(html) == 0 {
		t.Error("The html file should not be empty")
	}
}
