package evaluator

import (
	"bellamy/object"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvokingNonStaticFunc(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`foo()`,
			"string",
		},
		{
			`"string".length()`,
			6,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		fmt.Println(evaluated)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			// TODO note array oob indexing currently returns null, it should throw an error
			// testNullObject(t, evaluated)
			res, ok := evaluated.(*object.Error)
			assert.True(t, ok)
			assert.Contains(t, res.Message, "index out of bounds of array, ")
		}
	}
}
