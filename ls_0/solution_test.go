package solution

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetMessage(t *testing.T) {
	expected := string([]rune{72, 101, 108, 108, 111, 32, 102, 114, 111, 109, 32, 127477, 127473, 32, 33})
	msg := GetMessage()
	if !strings.EqualFold(msg, expected) {
		fmt.Println([]rune(msg))
		t.Errorf("Unexpected result:\n\tExpected: %q\n\tGot: %q", expected, msg)
	}

}
