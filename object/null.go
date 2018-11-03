package object

const NULL_OBJ = "NULL"

type Null struct {}

func (n *Null) Inspect() string {
	return "null"
}

func(n *Null) Type() ObjectType {
	return NULL_OBJ
}