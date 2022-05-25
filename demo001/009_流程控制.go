package main

import "fmt"

func main()  {
	//if else 写法 表达式 可以用括号包括起来，也可以不用
	/*
		if 表达式 {
			分支1
		}else if 表达式2 {
			分支2
		}else{
			分支3
		}

	 */

	if age := 32;age>23 {
		fmt.Println("局部变量:",age)
	}

	/**
		for 循环
		for 初始语句;条件 表达式;结束语{
			循环体语句
		}
		for 条件 {
			循环体语句
		}
		for 循环初始语句和 结束语句都可以省略
	*/

	var sum = 0
	for i:=0; i<=100; i++ {
		sum += i
	}
	fmt.Println(sum)

	//	 计算 9 的倍数
	for j:=0;j<=100;j++ {
		if j%9 == 0 {
			fmt.Println(j)
		}
	}
/**********************************************************/
	/*
		for range (键值循环)
	*/
	var str = "你好测试golang"
	for key,v:= range str {
		fmt.Printf("key=%v v=%c \n",key,v)
	}

}
