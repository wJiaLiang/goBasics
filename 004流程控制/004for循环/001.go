package main

import "fmt"

func main()  {
	init1()
/*
	Golang for支持三种循环方式，包括类似 while 的语法。
    for循环是一个循环控制结构，可以执行指定次数的循环。

*/

}

func init1()  {

	// 常见的 for 循环，支持初始化语句。
	s:= "abc"
	for i := 0; i < len(s); i++ {
		 fmt.Println(s[i])
	}

	//替代 while (n > 0) {}
	n:=len(s)
	for n>0{
		fmt.Println(s[n-1])
	}

	// 替代 while (true) {}
	for {
		fmt.Println(s)
	}

}