package palindrome

func IsPalindrome(s string) bool {
	return s == reverse(s)
}

func reverse(s string) string {
	runes := []rune(s)
	size := len(runes)
	out := make([]rune, 0)

	for i := range runes {
		out = append(out, runes[size-i-1])
	}
	return string(out)
}
