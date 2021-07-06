package naming_test

import (
	"testing"

	"testing_tasks/task_1/naming"
	"testing_tasks/task_1/naming/testdata"
)

func TestNameWithNumberIncrement(t *testing.T) {
	for _, r := range testdata.SameNameData {
		out, err := naming.NameWithNumberIncrement(r.BaseName, r.Names)
		if r.Err {
			if err == nil {
				t.Logf("must be an error for %s, %d ", r.InName, r.InNum)
				t.Fail()
			}
			continue
		}
		if err != nil {
			t.Errorf("enexpected error: %s for %s, %d ", err.Error(), r.InName, r.InNum)
		}
		if r.OutName != out {
			t.Errorf("expected %s, got %s", r.OutName, out)
		}
	}
}
