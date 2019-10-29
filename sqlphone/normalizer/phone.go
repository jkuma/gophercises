package normalizer

import (
	"bytes"
)

func Normalize(number string) string {
	var buffer bytes.Buffer

	for _, ch := range number {
		if ch >= '0' && ch <= '9' {
			buffer.WriteRune(ch)
		}
	}

	return buffer.String()
}
