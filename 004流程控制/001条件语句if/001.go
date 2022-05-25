package main

import "fmt"

/*
条件语句需要开发者通过指定一个或多个条件，
并通过测试条件是否为 true 来决定是否执行指定语句，并在条件为 false 的情况在执行另外的语句。
注意:	go 中不支持三元表达式

一、if 语句 if 语句 由一个布尔表达式后紧跟一个或多个语句组成。
	• 可省略条件表达式括号。
    • 持初始化语句，可定义代码块局部变量。
    • 代码块左 括号必须在条件表达式尾部。
	 if 布尔表达式 {
    	在布尔表达式为 true 时执行
	}


*/

func main()  {
	// 初始化语句未必就是定义变量 n
	x:=0
	// 在这里初始化的变量 在整个条件语句里面多可以使用;
	if n:="abc";x>0{
		fmt.Println(n[2])
	}else if x < 0 {
		fmt.Println(n[1])
	}else{
		fmt.Println(string(n[0]))  //a
		fmt.Println(n[0])
	}

//	if 嵌套语句 你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。

	init1()

}

func init1()  {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	/* 判断条件 */
	if a == 100 {
		/* if 条件语句为 true 执行 */
		if b == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Printf("a 的值为 100 ， b 的值为 200\n" )
		}
	}
	fmt.Printf("a 值为 : %d\n", a )
	fmt.Printf("b 值为 : %d\n", b )

}