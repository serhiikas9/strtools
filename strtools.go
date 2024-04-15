package strtools

import "unicode/utf8"

// Reverse returns input string with reversed order of characters
func Reverse(in string) string {
	runes := []rune(in)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Truncate(in string, maxLength int) string {
	if utf8.RuneCountInString(in) > maxLength {
		runes := []rune(in)
		return string(runes[:maxLength])
	}

	return in
}
