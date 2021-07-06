package naming

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"testing_tasks/task_1/naming/testdata"
)

func Test_getNumber_0(t *testing.T) {
	assert.Equal(t, 0, getNumberRegexp("first.last"))
}

func Test_getNumber_01(t *testing.T) {
	assert.Equal(t, 0, getNumberRegexp("firstlast"))
}

func Test_getNumber_02(t *testing.T) {
	assert.Equal(t, 1, getNumberRegexp("firstlast.01"))
}

func Test_getNumber_03(t *testing.T) {
	assert.Equal(t, 0, getNumberRegexp("firstlast.001"))
}

func Test_getNumber_04(t *testing.T) {
	assert.Equal(t, 99, getNumberRegexp("first.last.99"))
}

func Test_nameWithNumber(t *testing.T) {

	for _, r := range testdata.SameNameData {
		assert.Equal(t, r.OutName, nameWithNumber(r.InName, r.InNum))
	}
}
