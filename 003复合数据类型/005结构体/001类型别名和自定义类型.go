package main

import "fmt"

func main()  {
/*
   结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。每个值称为结构体的成员。
	用结构体的经典案例是处理公司的员工信息,每个员工有唯一的员工编号,名字，住址,年龄等

	Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。

	一、自定义类型
		在Go语言中有一些基本的数据类型，如string、整型、浮点型、布尔等数据类型，Go语言中可以使用type关键字来定义自定义类型。
		自定义类型是定义了一个全新的类型。我们可以基于内置的基本类型定义，也可以通过struct定义。

		 //将MyInt定义为int类型
       	 type MyInt int
		 通过Type关键字的定义，MyInt就是一种新的类型，它具有int的特性。


	二、类型别名
		类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。
		就像一个孩子小时候有小名、乳名，上学后用学名，英语老师又会给他起英文名，但这些名字都指的是他本人。

	三、类型定义和类型别名的区别
		类型别名与类型定义表面上看只有一个等号的差异，


*/
	diy1()
	diy2()
}

// 自定义类型
func diy1()  {
	type mytypes int


}
// 类型别名
func diy2()  {
	type byte = uint8
	type rune = int32

	type newInt int
	type MyInt = int

	var a newInt
	var b MyInt

	fmt.Printf("%T,%v \n",a,a) // main.newInt,0
	fmt.Printf("%T,%v \n",b,b) // int,0
	/*
		结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型。
		b的类型是int。MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型。
	*/

}