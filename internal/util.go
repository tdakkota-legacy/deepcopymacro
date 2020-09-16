package internal

import (
	"go/ast"

	builders "github.com/tdakkota/astbuilders"
)

func createFunction(name string, typ ast.Expr) builders.FunctionBuilder   {
	selector := ast.NewIdent("m")
	return builders.NewFunctionBuilder(name).
		Recv(&ast.Field{
			Names: []*ast.Ident{selector},
			Type:  typ,
		})
}

func createCopyFunction(name string, typ ast.Expr, bodyFunc builders.BodyFunc) builders.FunctionBuilder {
	return createFunction(name, typ).
		AddResults([]*ast.Field{
			{
				Names: []*ast.Ident{ast.NewIdent("out")},
				Type:  typ,
			},
		}...).
		Body(bodyFunc)
}