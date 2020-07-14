package main

import (
	"go/ast"
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test case example",
		},
	}
	for range tests {
		main()
	}
}

func Test_visitor_Visit(t *testing.T) {
	type args struct {
		n ast.Node
	}
	tests := []struct {
		name string
		v    visitor
		args args
		want ast.Visitor
	}{
		// TODO: Add test cases.
		{
			name: "test case example",
			args: args{
				//n: ast.Node{},
			},
		},
	}
	for _, tt := range tests {
		if got := tt.v.Visit(tt.args.n); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. visitor.Visit() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
		c float32
		d string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test case example",
			args: args{
				a: -1,
				b: -1,
				c: 0,
				d: "",
			},
		},
	}
	for _, tt := range tests {
		if got := add(tt.args.a, tt.args.b, tt.args.c, tt.args.d); got != tt.want {
			t.Errorf("%q. add() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_inArray(t *testing.T) {
	type args struct {
		a int
		b []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test case example",
			args: args{
				a: -1,
				b: []int{-1, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		if got := inArray(tt.args.a, tt.args.b); got != tt.want {
			t.Errorf("%q. inArray() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSetStruct(t *testing.T) {
	type args struct {
		s  MyStruct
		sp *MyStruct
		a  *int
		c  *string
	}
	tests := []struct {
		name string
		args args
		want MyStruct
	}{
		// TODO: Add test cases.
		{
			name: "test case example",
			args: args{
				s:  MyStruct{},
				sp: &MyStruct{},
				a:  new(int),
				c:  new(string),
			},
		},
	}
	for _, tt := range tests {
		if got := SetStruct(tt.args.s, tt.args.sp, tt.args.a, tt.args.c); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. SetStruct() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
