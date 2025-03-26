package intrange

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestIdentEqual(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a    ast.Expr
		b    ast.Expr
		name string
		want bool
	}{
		{
			name: "Both nil",
			a:    nil,
			b:    nil,
			want: false,
		},
		{
			name: "One nil, one non-nil",
			a:    &ast.Ident{Name: "x"},
			b:    nil,
			want: false,
		},
		{
			name: "Both Ident with same name",
			a:    &ast.Ident{Name: "x"},
			b:    &ast.Ident{Name: "x"},
			want: true,
		},
		{
			name: "Both Ident with different names",
			a:    &ast.Ident{Name: "x"},
			b:    &ast.Ident{Name: "y"},
			want: false,
		},
		{
			name: "SelectorExpr with matching structure",
			a: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "pkg"},
				Sel: &ast.Ident{Name: "Func"},
			},
			b: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "pkg"},
				Sel: &ast.Ident{Name: "Func"},
			},
			want: true,
		},
		{
			name: "SelectorExpr with different structure",
			a: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "pkg1"},
				Sel: &ast.Ident{Name: "Func"},
			},
			b: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "pkg2"},
				Sel: &ast.Ident{Name: "Func"},
			},
			want: false,
		},
		{
			name: "IndexExpr with matching structure",
			a: &ast.IndexExpr{
				X:     &ast.Ident{Name: "arr"},
				Index: &ast.BasicLit{Kind: token.INT, Value: "0"},
			},
			b: &ast.IndexExpr{
				X:     &ast.Ident{Name: "arr"},
				Index: &ast.BasicLit{Kind: token.INT, Value: "0"},
			},
			want: true,
		},
		{
			name: "IndexExpr with different structure",
			a: &ast.IndexExpr{
				X:     &ast.Ident{Name: "arr1"},
				Index: &ast.BasicLit{Kind: token.INT, Value: "0"},
			},
			b: &ast.IndexExpr{
				X:     &ast.Ident{Name: "arr2"},
				Index: &ast.BasicLit{Kind: token.INT, Value: "0"},
			},
			want: false,
		},
		{
			name: "BasicLit with same value",
			a:    &ast.BasicLit{Kind: token.INT, Value: "42"},
			b:    &ast.BasicLit{Kind: token.INT, Value: "42"},
			want: true,
		},
		{
			name: "BasicLit with different value",
			a:    &ast.BasicLit{Kind: token.INT, Value: "42"},
			b:    &ast.BasicLit{Kind: token.INT, Value: "43"},
			want: false,
		},
		{
			name: "Different types of expressions",
			a:    &ast.BasicLit{Kind: token.INT, Value: "42"},
			b:    &ast.Ident{Name: "x"},
			want: false,
		},
		{
			name: "Different types of expressions reversed",
			a:    &ast.Ident{Name: "x"},
			b:    &ast.BasicLit{Kind: token.INT, Value: "42"},
			want: false,
		},
		{
			name: "Incompatible types",
			a: &ast.CallExpr{
				Fun: &ast.Ident{Name: "int"},
				Args: []ast.Expr{
					&ast.BasicLit{Kind: token.INT, Value: "42"},
				},
			},
			b:    &ast.BasicLit{Kind: token.INT, Value: "42"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if result := identEqual(tt.a, tt.b); result != tt.want {
				t.Errorf(
					"identEqual(%v, %v) = %v; want %v",
					tt.a,
					tt.b,
					result,
					tt.want,
				)
			}
		})
	}
}

func TestIsNumberLit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		expr ast.Expr
		name string
		want bool
	}{
		{
			name: "BasicLit is an integer",
			expr: &ast.BasicLit{
				Kind:  token.INT,
				Value: "42",
			},
			want: true,
		},
		{
			name: "BasicLit is not an integer",
			expr: &ast.BasicLit{
				Kind:  token.STRING,
				Value: `"42"`,
			},
			want: false,
		},
		{
			name: "CallExpr with int cast and valid argument",
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
			want: true,
		},
		{
			name: "CallExpr with int cast and invalid argument",
			expr: &ast.CallExpr{
				Fun: &ast.Ident{
					Name: "int",
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: `"42"`,
					},
				},
			},
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
			want: false,
		},
		{
			name: "Unsupported expression type",
			expr: &ast.Ident{
				Name: "x",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if result := isNumberLit(tt.expr); result != tt.want {
				t.Errorf(
					"isNumberLit(%v) = %v; want %v",
					tt.expr,
					result,
					tt.want,
				)
			}
		})
	}
}

func TestCompareNumberLit(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			if result := compareNumberLit(tt.expr, tt.val); result != tt.want {
				t.Errorf(
					"compareNumberLit(%v, %d) = %v; want %v",
					tt.expr,
					tt.val,
					result,
					tt.want,
				)
			}
		})
	}
}
