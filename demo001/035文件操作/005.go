package main

import (
	"bufio"
	"os"
)

func main()  {
	// bufio 来 写入
	file,err:=os.OpenFile("C:/Users/wei/Desktop/go测试.txt",os.O_CREATE|os.O_RDONLY|os.O_TRUNC,0666)
	defer file.Close()
	if err!= nil{
		return
	}
	w:=bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		w.WriteString("bufilo写入数据")  // 先写入到缓存中
	}
	w.Flush()  // 写入到文件中

}
