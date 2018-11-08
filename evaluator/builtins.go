package evaluator

import (
	"bellamy/object"
)

var builtins = map[string]*object.Builtin{
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

func tail(args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. got %d, expected %d", len(args), 1)
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return object.NewError("argument to `tail` must be ARRAY, got %s", args[0].Type())
	}
	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		newEls := make([]object.Object, length-1, length-1)
		copy(newEls, arr.Elements[1:length])
		return &object.Array{Elements: newEls}
	}
	return object.NULL
}

func last(args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. got %d, expected %d", len(args), 1)
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return object.NewError("argument to `tail` must be ARRAY, got %s", args[0].Type())
	}
	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		return arr.Elements[length-1]
	}
	return object.NULL
}

func first(args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("wrong number of arguments. got %d, expected %d", len(args), 1)
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return object.NewError("argument to `first` must be ARRAY, got %s", args[0].Type())
	}
	arr := args[0].(*object.Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}
	return object.NULL
}

func push(args ...object.Object) object.Object {
	if len(args) != 2 {
		return object.NewError("wrong number of arguments. got %d, expected %d", len(args), 2)
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return object.NewError("argument to `first` must be ARRAY, got %s", args[0].Type())
	}
	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	newEl := make([]object.Object, length+1, length+1)
	copy(newEl, arr.Elements)
	newEl[length] = args[1]
	return &object.Array{Elements: newEl}
}
