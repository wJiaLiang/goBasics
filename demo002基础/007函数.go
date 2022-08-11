package main

import "fmt"

/*
1.函数声明：
函数声明包含一个函数名，参数列表， 返回值列表和函数体。如果函数没有返回值，则返回列表可以省略。
函数从第一条语句开始执行，直到执行return语句或者执行函数的最后一条语句。

函数可以没有参数或接受多个参数。

当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。

注意类型在变量名之后 。

函数可以返回任意数量的返回值

使用关键字 func 定义函数，左大括号依旧不能另起一行。


*/

func test(x, y int, s string) (int, string) {
	// 类型相同的相邻参数，参数类型可合并。 多返回值必须用括号。
	n := x + y
	return n, fmt.Sprintf(s, n)
}
func test2(x,y int,s string) string  {
	return "yes"
}
func main()  {
  fmt.Println(test(1,2,"a") )
  fmt.Println(test2(1,2,"b") )
}




















