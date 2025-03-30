package intrange

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestRecursiveOperandToString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		expr         ast.Expr
		name         string
		want         string
		incrementInt bool
	}{
		{
			name: "BasicLit integer without increment",
			expr: &ast.BasicLit{
				Kind:  token.INT,
				Value: "42",
			},
			incrementInt: false,
			want:         "42",
		},
		{
			name: "BasicLit integer with increment",
			expr: &ast.BasicLit{
				Kind:  token.INT,
				Value: "42",
			},
			incrementInt: true,
			want:         "43",
		},
		{
			name: "BasicLit invalid integer",
			expr: &ast.BasicLit{
				Kind:  token.INT,
				Value: "foo",
			},
			incrementInt: true,
			want:         "foo",
		},
		{
			name: "BasicLit hex integer",
			expr: &ast.BasicLit{
				Kind:  token.INT,
				Value: "0x2A",
			},
			incrementInt: true,
			want:         "43",
		},
		{
			name: "BasicLit non-integer",
			expr: &ast.BasicLit{
				Kind:  token.STRING,
				Value: "hello",
			},
			incrementInt: true,
			want:         "hello",
		},
		{
			name: "Ident",
			expr: &ast.Ident{
				Name: "x",
			},
			incrementInt: false,
			want:         "x",
		},
		{
			name: "SelectorExpr",
			expr: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "pkg"},
				Sel: &ast.Ident{Name: "Func"},
			},
			incrementInt: false,
			want:         "pkg.Func",
		},
		{
			name: "IndexExpr",
			expr: &ast.IndexExpr{
				X:     &ast.Ident{Name: "arr"},
				Index: &ast.BasicLit{Kind: token.INT, Value: "0"},
			},
			incrementInt: false,
			want:         "arr[0]",
		},
		{
			name: "BinaryExpr",
			expr: &ast.BinaryExpr{
				X:  &ast.Ident{Name: "x"},
				Op: token.ADD,
				Y:  &ast.BasicLit{Kind: token.INT, Value: "1"},
			},
			incrementInt: false,
			want:         "x + 1",
		},
		{
			name: "StarExpr",
			expr: &ast.StarExpr{
				X: &ast.Ident{Name: "ptr"},
			},
			incrementInt: false,
			want:         "*ptr",
		},
		{
			name: "CallExpr with single argument",
			expr: &ast.CallExpr{
				Fun: &ast.Ident{Name: "int"},
				Args: []ast.Expr{
					&ast.BasicLit{Kind: token.INT, Value: "42"},
				},
			},
			incrementInt: true,
			want:         "int(43)",
		},
		{
			name: "CallExpr with multiple arguments",
			expr: &ast.CallExpr{
				Fun: &ast.Ident{Name: "max"},
				Args: []ast.Expr{
					&ast.BasicLit{Kind: token.INT, Value: "1"},
					&ast.BasicLit{Kind: token.INT, Value: "2"},
				},
			},
			incrementInt: false,
			want:         "max(1, 2)",
		},
		{
			name: "Nested CallExpr",
			expr: &ast.CallExpr{
				Fun: &ast.CallExpr{
					Fun: &ast.Ident{Name: "outer"},
					Args: []ast.Expr{
						&ast.Ident{Name: "x"},
					},
				},
				Args: []ast.Expr{
					&ast.BasicLit{Kind: token.INT, Value: "42"},
				},
			},
			incrementInt: false,
			want:         "outer(x)(42)",
		},
		{
			name: "Unexpected expression type",
			expr: &ast.UnaryExpr{
				Op: token.ADD,
				X:  &ast.Ident{Name: "x"},
			},
			incrementInt: false,
			want:         "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if result := recursiveOperandToString(
				tt.expr,
				tt.incrementInt,
			); result != tt.want {
				t.Errorf(
					"recursiveOperandToString(%v, %v) = %v; want %v",
					tt.expr,
					tt.incrementInt,
					result,
					tt.want,
				)
			}
		})
	}
}

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
		{
			name: "BasicLit cannot be parsed",
			expr: &ast.BasicLit{
				Kind:  token.INT,
				Value: "foo",
			},
			val:  0,
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
