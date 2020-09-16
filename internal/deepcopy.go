package internal

import (
	"go/ast"
	"go/token"
	"go/types"

	builders "github.com/tdakkota/astbuilders"
	"github.com/tdakkota/gomacro/derive"
)

type DeepCopy struct{}

func (m *DeepCopy) CallFor(d *derive.Derive, field derive.Field, kind types.BasicKind) (*ast.BlockStmt, error) {
	s := builders.NewStatementBuilder()
	sel := field.Selector

	var out ast.Expr
	switch v := sel.(type) {
	case *ast.SelectorExpr:
		out = &ast.SelectorExpr{
			X:   v,
			Sel: ast.NewIdent("out"),
		}
	case *ast.Ident:
		out = ast.NewIdent("out")
	}

	s = s.Assign(out)(token.ASSIGN)(sel)
	return s.CompleteAsBlock(), nil
}

func (m *DeepCopy) Array(d *derive.Derive, field derive.Field, arr derive.Array) (*ast.BlockStmt, error) {
	return nil, nil
}

func (m *DeepCopy) Impl(d *derive.Derive, field derive.Field) (*ast.BlockStmt, error) {
	return nil, nil
}

func (m *DeepCopy) Target() *types.Interface {
	return nil
}

func (m *DeepCopy) Callback(d *derive.Derive, typeSpec *ast.TypeSpec) error {
	if _, ok := typeSpec.Type.(*ast.InterfaceType); ok {
		return nil
	}

	var err error
	builder := createCopyFunction("DeepCopy", typeSpec.Name, func(s builders.StatementBuilder) builders.StatementBuilder {
		s, err = d.Derive(typeSpec, s)
		return s.Return(builders.Nil())
	})

	d.AddDecls(builder.CompleteAsDecl())
	return err
}
