package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
//	ioutil 来读取文件

	byteStr,err:= ioutil.ReadFile("D:/Go/GolangStudy/demo001/001_变量命名.go")
	if err != nil{
		return
	}
	fmt.Println(string(byteStr))

}
