package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main()  {
	//1 打开文件
	file,err:=os.Open("D:/Go/GolangStudy/demo001/001_变量命名.go")
	defer file.Close()

	//2.读取文件
	if err != nil{
		return
	}
	reader := bufio.NewReader(file)
	var strlong string
	for  {
		str,err:= reader.ReadString('\n')  // 表示一次读取一行
		if err == io.EOF{
			strlong+=str
			break
		}
		if err!=nil{
			return
		}
		strlong+=str
	}
	fmt.Println(strlong)

}
