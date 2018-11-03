package object

import "fmt"

const BOOLEAN_OBJ = "BOOLEAN"

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func(b *Boolean) Boolean() ObjectType {
	return BOOLEAN_OBJ
}