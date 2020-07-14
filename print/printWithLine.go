package print
import ("fmt";"runtime")

/*
golang 的runtime库，提供Caller函数，可以返回运行时正在执行的文件名和行号：
函数定义：
func Caller(skip int) (pc uintptr, file string, line int, ok bool) {}

函数用法:
_, file, line, ok := runtime.Caller(0)
*/

func PrintWithLine(v ...interface{}) {
	funcName,file,line,ok := runtime.Caller(1)
	if(ok){
		fmt.Println("func name: " + runtime.FuncForPC(funcName).Name()+fmt.Sprintf("file: %s, line: %d  %v",file,line,v))

	}else {
		fmt.Println(v)
	}

}
