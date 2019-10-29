package normalizer

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{"1234567890", "1234567890"},
		{"123 456 7891", "1234567891"},
		{"(123) 456 7892", "1234567892"},
		{"123-456-7894", "1234567894"},
	}

	for _, c := range testCases {
		normalized := Normalize(c.input)

		if c.want != normalized {
			t.Errorf("The normalized phone number for %v should be %v and not %v", c.input, c.want, normalized)
		}
	}
}
