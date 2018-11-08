package object

const BUILTIN_OBJ = "BUILTIN"

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Inspect() string {
	return "builtin function"
}

func (_ *Builtin) Type() ObjectType {
	return BUILTIN_OBJ
}
