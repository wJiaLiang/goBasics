package main

import "fmt"

func fn1()  {
	fmt.Println("方法1")
}
func fn2()  {
	fn1()
	fmt.Println("方法2")
}

// 递归
func sum(n int) int  {
	if n>1{
		return n + sum(n-1)
	}else{
		return 1
	}

}
//******************************************************
// 闭包 函数里面嵌套一个函数，之后返回里面的函数
// 定义一个函数 返回一个函数
func adder() func() int {
	var i=10
	return func() int {
		return i+1
	}
}
func adder1() func(int) int {
	var i=10
	return func(y int) int {
		i += y
		return i
	}
}


func main()  {
	/*
		函数的递归调用
	*/
	fn2()

	//实现 1到100 数值的和  使用递归实现
	fmt.Println( sum(100) ) // 5050

/*
	闭包  让一个变量 常驻内存 单不污染全局变量
	闭包可以理解为 定义在函数内部的函数，本质上闭包是将函数内部和函数外部链接起来的桥梁，
	或者说是函数 和 其引用环境的组合体

*/

	var v = adder()
	fmt.Println(v()) // 11
	fmt.Println(v()) // 11
	fmt.Println(v()) // 11

	var v1 = adder1()
	fmt.Println(v1(10)) // 20
	fmt.Println(v1(10)) // 30
	fmt.Println(v1(10)) // 40







}
