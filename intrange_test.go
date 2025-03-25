package intrange

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestCompareNumberLit(t *testing.T) {
	tests := []struct {
		expr ast.Expr
		name string
		val  int
		want bool
	}{
		{
			name: "BasicLit matches integer",
			expr: &ast.BasicLit{
				Kind:  token.INT,
				Value: "42",
			},
			val:  42,
			want: true,
		},
		{
			name: "BasicLit does not match integer",
			expr: &ast.BasicLit{
				Kind:  token.INT,
				Value: "42",
			},
			val:  43,
			want: false,
		},
		{
			name: "BasicLit matches hexadecimal integer",
			expr: &ast.BasicLit{
				Kind:  token.INT,
				Value: "0x2A",
			},
			val:  42,
			want: true,
		},
		{
			name: "BasicLit is not an integer",
			expr: &ast.BasicLit{
				Kind:  token.STRING,
				Value: `"42"`,
			},
			val:  42,
			want: false,
		},
		{
			name: "CallExpr with int cast matches integer",
			expr: &ast.CallExpr{
				Fun: &ast.Ident{
					Name: "int",
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.INT,
						Value: "42",
					},
				},
			},
			val:  42,
			want: true,
		},
		{
			name: "CallExpr with int cast does not match integer",
			expr: &ast.CallExpr{
				Fun: &ast.Ident{
					Name: "int",
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.INT,
						Value: "43",
					},
				},
			},
			val:  42,
			want: false,
		},
		{
			name: "CallExpr with non-int cast",
			expr: &ast.CallExpr{
				Fun: &ast.Ident{
					Name: "float64",
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.INT,
						Value: "42",
					},
				},
			},
			val:  42,
			want: false,
		},
		{
			name: "CallExpr with multiple arguments",
			expr: &ast.CallExpr{
				Fun: &ast.Ident{
					Name: "int",
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.INT,
						Value: "42",
					},
					&ast.BasicLit{
						Kind:  token.INT,
						Value: "43",
					},
				},
			},
			val:  42,
			want: false,
		},
		{
			name: "CallExpr with unexpected function",
			expr: &ast.CallExpr{
				Fun: &ast.CallExpr{},
			},
			val:  42,
			want: false,
		},
		{
			name: "Unsupported expression type",
			expr: &ast.Ident{
				Name: "x",
			},
			val:  42,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := compareNumberLit(tt.expr, tt.val); result != tt.want {
				t.Errorf("compareNumberLit(%v, %d) = %v; want %v", tt.expr, tt.val, result, tt.want)
			}
		})
	}
}
