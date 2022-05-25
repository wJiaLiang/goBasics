package main

import "fmt"

func main()  {
/*
	1、常量表达式的值在编译期计算，而不是在运行期
	2、个常量的声明语句定义了常量的名字，和变量的声明语法类似，常量的值不可修改

	3、当操作数是常量时，一些运行时的错误也可以在编译时被发现，例如整数除零、字符串索引越界、任何导致无效浮点数的操作等。

	4、常量间的所有算术运算、逻辑运算和比较运算的结果也是常量，对常量的类型转换操作或以下函数调用都是返回常量结
		len、cap、real、imag、complex和unsafe.Sizeof
	5、一个常量的声明也可以包含一个类型和一个值，但是如果没有显式指明类型，那么将从右边的表达式推断类型。

	6、批量声明，除了第一个外其它的常量右边的初始化表达式都可以省略，如果省略初始化表达式则表示使用前面常量的初始化表达式

	7、简单地复制右边的常量表达式，并没有太实用的价值。但是它可以带来其它的特性，那就是iota常量生成器语法


	一、iota 常量生成器
	常量声明可以使用iota常量生成器初始化，它用于生成一组以相似规则初始化的常量，
	在一个const声明语句中，在第一个声明的常量所在的行，iota将会被置为0，然后在每一个有常量声明的行加一。

*/

}
func init()  {
	const pi = 3.14159

	//或者
	const(
		e = 2.718
		pi2 = 3.14
	)

	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a,b,c,d)  // 1  1  2  2

	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday,Monday,Tuesday,Wednesday,Thursday,Friday,Saturday) //0 1 2 3 4 5 6



/*
   无类型常量

*/



}
