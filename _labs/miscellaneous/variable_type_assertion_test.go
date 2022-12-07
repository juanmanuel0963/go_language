package miscellaneous

import (
	"testing"
)

func TestVariableTypeAssertion_v1(t *testing.T) {

	VariableTypeAssertion(21)
	VariableTypeAssertion("hello")
	VariableTypeAssertion(true)

}
