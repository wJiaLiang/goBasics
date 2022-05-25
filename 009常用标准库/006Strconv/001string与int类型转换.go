package main

import (
	"fmt"
	"strconv"
)

func main()  {
/*
	strconv包实现了基本数据类型与其字符串表示的转换，
	主要有以下常用函数： Atoi()、Itia()、parse系列、format系列、append系列。

*/

	test1()
	test2()

}

// Atoi()函数用于将字符串类型的整数转换为int类型，函数签名如下。func Atoi(s string) (i int, err error)
func test1()  {
	s1:="100"
	v,err := strconv.Atoi(s1)
	if err !=nil{
		fmt.Println("不能转换成 int")
	}else {
		fmt.Printf("值：%v--类型：%T \n",v,v) //值：100--类型：int
	}
}

//  Itoa()函数用于将int类型数据转换为对应的字符串表示  函数签名如下 func Itoa(i int) string
func test2()  {
	s2:=123
	str := strconv.Itoa(s2)
	fmt.Println(str)
}	