package structFile

import (
	"reflect"
	"testing"
)

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
