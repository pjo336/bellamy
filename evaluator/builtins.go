package evaluator

import (
	"bellamy/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: length,
	},
}

func length(args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. got %d, expected 1", len(args))
	}
	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	default:
		return object.NewError("argument to `len` not supported, got %s", args[0].Type())
	}
}
