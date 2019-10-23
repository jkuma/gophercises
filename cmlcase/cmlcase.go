package cmlcase

import "unicode"

func Count(s string) int16 {
	var wc int16

	runes := []rune(s)

	if len(runes) > 0 {
		wc++
	}

	for _, r := range runes {
		if unicode.IsUpper(r) {
			wc++
		}
	}

	return wc
}
