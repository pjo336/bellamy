package static

import (
	"bellamy/object"
	"fmt"
)

var StaticBuiltins = map[string]*object.Builtin{
	"print": &object.Builtin{Fn: print},
	"len":   &object.Builtin{Fn: length},
	"first": &object.Builtin{Fn: first},
	"last":  &object.Builtin{Fn: last},
	"tail":  &object.Builtin{Fn: tail},
	"push":  &object.Builtin{Fn: push},
}

func length(args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. got %d, expected 1", len(args))
	}
	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}
	default:
		return object.NewError("argument to `len` not supported, got %s", args[0].Type())
	}
}

func print(args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Println(arg.Inspect())
	}
	return object.NULL
}
