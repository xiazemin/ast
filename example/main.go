package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"fmt"
 "github.com/davecgh/go-spew/spew"
	//"bytes"
	//"encoding/json"
	//"github.com/kubernetes/kubernetes/staging/src/k8s.io/apimachinery/pkg/util/validation/field"
	"strings"
	//"go/types"
	"reflect"
)

func main() {
	//var aliases = [...]*types.Basic{
	//	{types.Byte, types.IsInteger | types.IsUnsigned, "byte"},
	//	{types.Rune, types.IsInteger, "rune"},
	//}
	//fmt.Println(aliases)
	var a=[]MyStruct{{"1",new(int)},{"1",new(int)}}
	fmt.Println(a,reflect.TypeOf(a))
	var b=[...]MyStruct{{"1",new(int)},{"1",new(int)}}
	fmt.Println(b,reflect.TypeOf(b))
	// 这就是上一章的代码.
	src := `
	package main
	type MyStruct struct {
	Name string
	Age *int
}

func (this *MyStruct)SetStruct(s MyStruct,sp *MyStruct,a *int,c * string) MyStruct {
	b:=new(int)
	fmt.Println(*b)
	*a=5
	fmt.Println(*a)
	return s
}
	var a int
	func main() {
	b:="test"
	    println("Hello, World!",b)
	}
	func test(a ,b int)int{
	  return a+b
	}
	`
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	fmt.Println(fset)
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(*f)
	//var prettyJSON bytes.Buffer
	//error := json.Indent(&prettyJSON, f, "", "\t")
	//fmt.Println(error)
	//fmt.Println(string(prettyJSON.Bytes()))

	// Print the AST.
	ast.Print(fset, f)
	spew.Dump(f)
	ast.Inspect(f, func(x ast.Node) bool {
		s, ok := x.(*ast.FuncType)
		if !ok {
			fmt.Println("not function")
			return true
		}

		fmt.Printf("param: %v\n",s.Params)
		for _,p:=range s.Params.List{
			fmt.Println(p)
			for _,v:=range p.Names {
				fmt.Println(v)
			}
		}
		fmt.Printf("result:   %v\n", s.Results)
		return false
	})
	var v visitor
	ast.Walk(v,f)
}

type visitor int

func (v visitor) Visit(n ast.Node) ast.Visitor {
	fmt.Println(fmt.Sprintf("%#v,-----> %T",n,n))
	if n == nil {
		return nil
	}
	if decl,ok:=n.(ast.Decl);ok{
		fmt.Println(fmt.Sprintf("%#v,-----> %T--->***%#v",n,n,decl))
	}
	fmt.Println("测试:%v", n.Pos().IsValid())
	fmt.Printf("%s%T\n", strings.Repeat("\t", int(v)), n)
	return v + 1
}

func add(a int,b int,c float32,d string)int{
	fmt.Println(c,d)
	return a+b
}

func inArray(a int ,b []int)bool{
	for _,i:=range b{
if a==i{
	return true
}
	}
	return false
}

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
//0  *ast.File {
//1  .  Package: 2:2
//2  .  Name: *ast.Ident {
//3  .  .  NamePos: 2:10
//4  .  .  Name: "main"
//5  .  }
//6  .  Decls: []ast.Decl (len = 1) {
//7  .  .  0: *ast.FuncDecl {
//8  .  .  .  Name: *ast.Ident {
//9  .  .  .  .  NamePos: 3:7
//10  .  .  .  .  Name: "main"
//11  .  .  .  .  Obj: *ast.Object {
//12  .  .  .  .  .  Kind: func
//13  .  .  .  .  .  Name: "main"
//14  .  .  .  .  .  Decl: *(obj @ 7)
//15  .  .  .  .  }
//16  .  .  .  }
//17  .  .  .  Type: *ast.FuncType {
//18  .  .  .  .  Func: 3:2
//19  .  .  .  .  Params: *ast.FieldList {
//20  .  .  .  .  .  Opening: 3:11
//21  .  .  .  .  .  Closing: 3:12
//22  .  .  .  .  }
//23  .  .  .  }
//24  .  .  .  Body: *ast.BlockStmt {
//25  .  .  .  .  Lbrace: 3:14
//26  .  .  .  .  List: []ast.Stmt (len = 1) {
//27  .  .  .  .  .  0: *ast.ExprStmt {
//28  .  .  .  .  .  .  X: *ast.CallExpr {
//29  .  .  .  .  .  .  .  Fun: *ast.Ident {
//30  .  .  .  .  .  .  .  .  NamePos: 4:6
//31  .  .  .  .  .  .  .  .  Name: "println"
//32  .  .  .  .  .  .  .  }
//33  .  .  .  .  .  .  .  Lparen: 4:13
//34  .  .  .  .  .  .  .  Args: []ast.Expr (len = 1) {
//35  .  .  .  .  .  .  .  .  0: *ast.BasicLit {
//36  .  .  .  .  .  .  .  .  .  ValuePos: 4:14
//37  .  .  .  .  .  .  .  .  .  Kind: STRING
//38  .  .  .  .  .  .  .  .  .  Value: "\"Hello, World!\""
//39  .  .  .  .  .  .  .  .  }
//40  .  .  .  .  .  .  .  }
//41  .  .  .  .  .  .  .  Ellipsis: -
//42  .  .  .  .  .  .  .  Rparen: 4:29
//43  .  .  .  .  .  .  }
//44  .  .  .  .  .  }
//45  .  .  .  .  }
//46  .  .  .  .  Rbrace: 5:2
//47  .  .  .  }
//48  .  .  }
//49  .  }
//50  .  Scope: *ast.Scope {
//51  .  .  Objects: map[string]*ast.Object (len = 1) {
//52  .  .  .  "main": *(obj @ 11)
//53  .  .  }
//54  .  }
//55  .  Unresolved: []*ast.Ident (len = 1) {
//56  .  .  0: *(obj @ 29)
//57  .  }
//58  }
