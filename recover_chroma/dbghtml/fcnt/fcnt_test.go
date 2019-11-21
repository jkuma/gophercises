package fcnt

import "testing"

func TestGetFileContent(t *testing.T) {
	fc, err := GetFileContent("test/test.html")

	if err != nil {
		t.Error(err)
	}

	if len(fc.Content) == 0 {
		t.Error("The file should not be empty.")
	}
}
