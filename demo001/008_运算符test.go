package main

import "fmt"

func main(){
	//1. 交换两个变量的值  不使用第三方变量
	var a = 10
	var b = 32
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a=%v b=%v \n",a,b)

	// 假如还有 100 天放假， n 个星期 零 n天
	var days = 100;
	var week = days/7
	var day = days % 7

	fmt.Printf("%v %v \n",week,day)

	// 位运算符


}
