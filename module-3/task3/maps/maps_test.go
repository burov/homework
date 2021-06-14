package maps

import (
	"testing"
)

func TestSortedValues(t *testing.T) {
	tests := []struct {
		in       map[int]string
		expected string
	}{
		{in: map[int]string{1: "A", 5: "D", 8: "G", 3: "C", 2: "B", 19: "H", 6: "E", 7: "F"}, expected: "ABCDEFGH"},
		{in: map[int]string{3: "X", 1: "Z", 5: "W", 18: "S", 9: "U", 2: "Y", 11: "T", 6: "V"}, expected: "ZYXWVUTS"},
	}

	for i, tt := range tests {
		if got := SortedValues(tt.in); tt.expected != got {
			t.Errorf("Unexpected result in Test #%d, expected %q, got %q", i, tt.expected, got)
		}
	}

}
