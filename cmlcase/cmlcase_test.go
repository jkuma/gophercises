package cmlcase

import "testing"

func TestCount(t *testing.T) {
	s := "oneTwoThree"

	if Count(s) != 3 {
		t.Errorf("The string %v has 3 words", s)
	}
}

