package evaluator

import (
	"bellamy/ast"
	"bellamy/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatement(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return booleanObject(node.Value)
	default:
		return object.NULL
	}
}

func evalStatement(statements []ast.Statement) object.Object {
	var result object.Object
	for _, stmt := range statements {
		result = Eval(stmt)
	}
	return result
}

func booleanObject(b bool) *object.Boolean {
	if b == true {
		return object.TRUE
	}
	return object.FALSE
}