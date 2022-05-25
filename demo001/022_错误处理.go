package main

import "fmt"


/*
	当b 为0 时候会报错，程序抛出异常；
*/
func fn2(a int,b int) int  {
	defer func() {
		err:= recover()
		if err !=nil {
			fmt.Println("错误信息",err)  // 错误信息 runtime error: integer divide by zero
		}
	}()
	return  a/b
}

func main()  {
 fmt.Println(	 fn2(10,0) ) // 0
 fmt.Println("继续执行")
	
	
	
	
	
	
}
