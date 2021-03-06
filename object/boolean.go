package object

import "fmt"

const BOOLEAN_OBJ = "BOOLEAN"

// Since only 2 types of booleans exist, lets hard code them to save memory
var (
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (b *Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 2
	}
	return HashKey{Type: b.Type(), Value: value}
}
