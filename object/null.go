package object

const NULL_OBJ = "NULL"

// Only 1 way to be a null object
var NULL = &Null{}

type Null struct {}

func (n *Null) Inspect() string {
	return "null"
}

func(n *Null) Type() ObjectType {
	return NULL_OBJ
}