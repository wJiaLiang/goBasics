package main

import (
	"fmt"
)

/*
	switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上直下逐一测试，直到匹配为止。
	Golang switch 分支表达式可以是任意类型，不限于常量。可省略 break，默认自动终止。
	语法：
		switch var1 {
		case val1:
			...
		case val2:
			...
		default:
			...
		}
*/
func main()  {
	init1()



}

func init1()  {
/*
	变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。
	类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。
	 您可以同时测试多个可能符合条件的值，使用逗号分割它们，例如：case val1, val2, val3。
*/
	expr:=23
	switch expr {
		case 1:
			fmt.Println(1)
		case 2:
			fmt.Println(2)
		case 4,5:
			fmt.Println(4,5)
		case 23:
			fmt.Println(23)
		default:
			fmt.Println("default")
	}

	var age int8 = 11
	age = 12
	switch  {
	case age == 1:
		fmt.Println(age)
	case age ==3:
		fmt.Println(age)
	case age ==7:
		fmt.Println(age)
	case age ==8||age==12 :
		fmt.Println(age)
	default:
		fmt.Println("没有匹配到")
	}



}
