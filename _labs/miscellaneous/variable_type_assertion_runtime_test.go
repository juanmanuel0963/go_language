package miscellaneous

import (
	"testing"
)

func TestVariableTypeAssertionRuntime_v1(t *testing.T) {

	VariableTypeAssertionRuntime(21)
	VariableTypeAssertionRuntime("hello")
	VariableTypeAssertionRuntime(true)

}
