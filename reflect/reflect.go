package main

import (
	"fmt"
	"reflect"
	"bytes"
	"encoding/json"
)

type MyStruct struct {
	Name string
	Age *int
}

type MyInterface interface{
	Get()string
	Set(string)
}


func main()  {
	var a MyStruct
	t:=reflect.TypeOf(a)
	fmt.Println(t,t.Kind())
	for i:=0;i<t.NumField();i++{
		fmt.Println(i,t.Field(i).Type.Kind())
	}
	var b *MyStruct
	bt:=reflect.TypeOf(b)
	fmt.Println(bt,bt.Kind())
	if bt.Kind()==reflect.Ptr{
		fmt.Println(bt.Elem())
	}
	for i:=0;i<bt.Elem().NumField();i++{
		fmt.Println(i,bt.Elem().Field(i).Type.Kind())
	}
	var c *MyInterface
	fmt.Println(c)
	ct:=reflect.TypeOf(c)
	fmt.Println(ct.Elem().NumMethod())
	for i:=0;i<ct.Elem().NumMethod();i++{
		fmt.Println(i,ct.Elem().Method(i))
	}

	d:=func (s MyStruct,y *int)(*int){return y}
	fmt.Println(reflect.TypeOf(d),reflect.ValueOf(d))
	cd:=reflect.TypeOf(d)
	fmt.Println(cd.NumIn(),cd.NumOut())
	var f struct{}
	fmt.Println(f)
	structOf()

	fmt.Printf("%#v",reflect.FuncOf([]reflect.Type{bt},[]reflect.Type{bt,bt},false))
	fmt.Println(reflect.FuncOf([]reflect.Type{bt},[]reflect.Type{bt,bt},false))
	var s  []MyStruct=make([]MyStruct,3)
	st:=reflect.TypeOf(s)
	fmt.Println(reflect.FuncOf([]reflect.Type{st,st},[]reflect.Type{bt,bt},true))
}

func structOf (){
	typ := reflect.StructOf([]reflect.StructField{
		{
			Name: "Height",
			Type: reflect.TypeOf(float64(0)),
			Tag:  `json:"height"`,
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(int(0)),
			Tag:  `json:"age"`,
		},
	})

	v := reflect.New(typ).Elem()
	v.Field(0).SetFloat(0.4)
	v.Field(1).SetInt(2)
	s := v.Addr().Interface()

	w := new(bytes.Buffer)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}

	fmt.Printf("value: %#v\n", s)
	fmt.Printf("json:  %s", w.Bytes())

	r := bytes.NewReader([]byte(`{"height":1.5,"age":10}`))
	if err := json.NewDecoder(r).Decode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)
}
