package fibonacci

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	var nUint uint = 3
	rInterface := fibonacci(nUint)

	var expectedUint uint = 2
	if rUint, ok := rInterface.(uint); !ok || rUint != expectedUint {
		t.Errorf("unexpected result for uint case, expected type - %s, got type - %s, expected result - %d, got result - %d",
			reflect.TypeOf(expectedUint), reflect.TypeOf(rInterface), expectedUint, rUint)
	}

	var nInt int = 5
	rInterface = fibonacci(nInt)

	var expectedInt int = 5
	if rInt, ok := rInterface.(int); !ok || rInt != expectedInt {
		t.Errorf("unexpected result for int case, expected type - %s, got type - %s, expected result - %d, got result - %d",
			reflect.TypeOf(expectedInt), reflect.TypeOf(rInterface), expectedInt, rInt)
	}

	var nFloat64 float64 = 7.0
	rInterface = fibonacci(nFloat64)

	var expectedFloat64 float64 = 13.0
	if rFloat64, ok := rInterface.(float64); !ok || rFloat64 != expectedFloat64 {
		t.Errorf("unexpected result for float64 case, expected type - %s, got type - %s, expected result - %f, got result - %f",
			reflect.TypeOf(expectedFloat64), reflect.TypeOf(rInterface), expectedFloat64, rFloat64)
	}

}
