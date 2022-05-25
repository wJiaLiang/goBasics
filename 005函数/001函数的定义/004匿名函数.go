package main

import (
	"fmt"
	"math"
)

func main()  {
/*


	1、匿名函数是指不需要定义函数名的一种函数实现方式。

	2、拥有函数名的函数只能在包级语法块中被声明，通过函数字面量（function literal），
	我们可绕过这一限制，在任何表达式中表示一个函数值 函数字面量的语法和函数声明相似
	函数值字面量是一种表达式，它的值被称为匿名函数（anonymous function）。
	3、匿名函数由一个不带函数名的函数声明和函数体组成。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。

	4、Golang匿名函数可赋值给变量，做为结构字段，或者在 channel 里传送。
*/

	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)   //Sqrt 返回 x 的平方根。
	}
	fmt.Println(getSqrt(4))


	// 第一了一个函数类型的切片 返回值 int 类型;
	fns := [] func(x int) int {
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}
	println(fns[0](100))  //101;

	//****************************************************
	//squares3返回后，变量x仍然隐式的存在于f中。
	f:=squares3()          // 返回的匿名函数内部变量被缓存;
	fmt.Println(f())  //1
	fmt.Println(f())  //4

	fmt.Println(squares3()() ) //1
	fmt.Println(squares3()() ) //1

}

// squares返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方。
func squares3() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}