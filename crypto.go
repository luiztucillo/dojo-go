package main

import "strings"

func Encrypt(value string) string {
	result := strings.Map(func(r rune) rune {
		return caesar(r, 3)
	}, value)

	return result
}

func Decrypt(value string) string {
	result := strings.Map(func(r rune) rune {
		return caesar(r, -3)
	}, value)

	return result
}

func caesar(r rune, shift int) rune {
	if int(r) < 'a' || int(r) > 'z' {
		return r
	}

	s := int(r) + shift

	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}

	return rune(s)
}
