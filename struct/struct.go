package structFile

import "fmt"

type MyStruct struct {
	Name string
	Age *int
}

func SetStruct(s MyStruct,sp *MyStruct,a *int,c * string) MyStruct {
	b:=new(int)
	fmt.Println(*b)
	*a=5
	fmt.Println(*a)
	return s
}
