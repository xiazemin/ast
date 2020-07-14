package main

import (
	"reflect"
	"testing"
)

func TestGetAllFile(t *testing.T) {
	type args struct {
		pathname string
		s        []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test case example",
			args: args{
				pathname: "",
				s:        nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		got, err := GetAllFile(tt.args.pathname, tt.args.s)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. GetAllFile() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. GetAllFile() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGetAllDir(t *testing.T) {
	type args struct {
		pathname string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test case example",
			args: args{
				pathname: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		got, err := GetAllDir(tt.args.pathname)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. GetAllDir() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. GetAllDir() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

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

func Test_rtrim(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test case example",
			args: args{
				str: "",
			},
		},
	}
	for _, tt := range tests {
		if got := rtrim(tt.args.str); got != tt.want {
			t.Errorf("%q. rtrim() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
