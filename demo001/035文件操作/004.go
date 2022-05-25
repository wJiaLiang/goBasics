package main

import "os"

func main()  {
   file,err:=os.OpenFile("C:/Users/wei/Desktop/go测试.txt",os.O_CREATE|os.O_RDONLY|os.O_TRUNC,0666)
   defer file.Close()
	if err!= nil{
		return
	}
	//file.WriteString("直接写入一个数据")
	var str = "值值值数值"
    file.Write([]byte(str))
}
