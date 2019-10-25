package caesarcipher

func Marshal(s string, k int) string {
	var ret []rune

	for _, r := range []rune(s) {
		ret = append(ret, cipher(r, k))
	}

	return string(ret)
}

func cipher(r rune, delta int) rune {
	if r >= 'A' && r <= 'Z' {
		return rotate(r, 'A', delta)
	}
	if r >= 'a' && r <= 'z' {
		return rotate(r, 'a', delta)
	}
	return r
}

func rotate(r rune, base, delta int) rune {
	tmp := int(r) - base
	tmp = (tmp + delta) % 26
	return rune(tmp + base)
}
