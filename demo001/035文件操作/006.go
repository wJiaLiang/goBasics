package main

import "io/ioutil"

func main()  {
	// ioutil 写入
	str:="hello,ioutil"
	err:= ioutil.WriteFile("C:/Users/wei/Desktop/go测试.txt",[]byte(str),0666)
	if err!=nil{
		return
	}

}
