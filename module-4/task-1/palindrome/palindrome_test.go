package palindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		in  string
		exp bool
	}{
		{in: "aa", exp: true},
		{in: "ab", exp: false},
		{in: "aba", exp: true},
		{in: "mełłem", exp: true},
		{in: "Hello", exp: false},
	}

	for i, tt := range tests {
		if got := IsPalindrome(tt.in); tt.exp != got {
			t.Errorf("Test #%d, unexpected result, expected %t, got %t", i, tt.exp, got)
		}
	}
}
