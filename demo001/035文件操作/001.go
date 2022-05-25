package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	// 只读方式打开文件
	file,err :=os.Open("D:/Go/GolangStudy/demo001/001_变量命名.go")
	defer file.Close()
	if err!=nil {
		fmt.Println(err)
		return
	}
	// 读取
	var tempSlice = make([]byte,256)
	var valeSlice []byte
	for  {
		n,err:=file.Read(tempSlice)
		if err == io.EOF{  // 表示读取完毕
			break
		}
		if err!=nil {
			return
		}
		//fmt.Printf("%v\n",n)
		valeSlice = append(valeSlice, tempSlice[:n]...)
	}

	fmt.Println(string(valeSlice))


}
