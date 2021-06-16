package pointers

import (
	"testing"
)

func Test_doubleValues(t *testing.T) {
	tests := []struct {
		in       [5]int
		expected [5]int
	}{
		{in: [5]int{2, 4, 6, 8, 10}, expected: [5]int{4, 8, 12, 16, 20}},
		{in: [5]int{1, 3, 5, 7, 9}, expected: [5]int{2, 6, 10, 14, 18}},
	}

	for i, tt := range tests {
		out, expected := &tt.in, &tt.expected
		if DoubleValues(out); expected == out {
			t.Errorf("Unexpected result in Test #%d, expected %+v, got %+v", i, tt.expected, tt.in)
		}
	}
}
