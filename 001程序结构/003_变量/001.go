package main

import "fmt"

/*
变量声明的一般语法如下：
var 变量名字 类型 = 表达式
"类型"  或  "=表达式"  两个部分可以省略其中的一个,
如果省略的是类型信息，那么将根据初始化表达式来推导变量的类型信息。
如果初始化表达式被省略，那么将用零值初始化该变量。

数值类型变量对应的零值是 0
布尔类型变量对应的零值是 false
字符串类型对应的零值是空 字符串
接口或引用类型（包括 slice、指针、map、chan 和函数）变量对应的零值是 nil
数组或结构体等聚合类型 对应的零值是每个 元素或字段 都是对应该类型的零值。


	var i, j, k int
	var b,f,c = true,5,"呵呵"

*/

func main()  {
	var s string
	fmt.Println(s)

	// 声明一组变量
	var i, j, k int  			  // int int int
	var b,f,c = true,5.6,"呵呵"   // bool,float64,string

	fmt.Print(b,f,c)
	fmt.Print(i,j,k)



}
