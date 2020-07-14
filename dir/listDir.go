package main

import (
"fmt"
"io/ioutil"
	"strings"
)

func GetAllFile(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			s, err = GetAllFile(fullDir, s)
			if err != nil {
				fmt.Println("read dir fail:", err)
				return s, err
			}
		} else {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}

func GetAllDir(pathname string) ([]string, error) {
	//fmt.Println(s,pathname)
	//s=append(s,pathname)
        var s []string
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			s1,_:=GetAllDir(fullDir)
			s=append(s,s1...)
			s=append(s,fullDir)
		}
	}
	return s, nil
}
func main() {
	//遍历打印所有的文件名
	var s []string
	dir:="/Users/didi/goLang/src/github.com/xiazemin/autotest"
	fmt.Println(rtrim("/Users/didi/goLang/src/github.com/xiazemin/autotest"))
	fmt.Println(rtrim("/Users/didi/goLang/src/github.com/xiazemin/autotest/"))
	s, _ = GetAllDir(dir)

	for _,ss:=range s{
		fmt.Println("slice: %v", ss)
	}

}

func rtrim(str string)string{
	if str==""{
		return str
	}
	if str[0]=='/'{
		return "/"+strings.Trim(str,"/")
	}
	return strings.Trim(str,"/")
}