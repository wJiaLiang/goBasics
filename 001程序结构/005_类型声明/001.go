package main

import "fmt"

/*
1、变量或表达式的类型定义了对应存储值的属性特征，例如数值在内存的存储大小（或者是元素的bit个数），是否支持一些操作符，以及它们自己关联的方法等。

2、一个类型声明语句创建了一个新的类型名称，和现有类型具有相同的底层结构。
	新命名的类型提供了一个方法，用来分隔不同概念的类型，这样即使它们底层类型相同也是不兼容的

	type 类型名字 底层类型

3、类型声明语句一般出现在包一级，如果类型名字的首字符大写，则在包外部也可以使用。
4、对于每一个类型T，都有一个对应的类型转换操作T(x)，用于将x转为T类型

	type Celsius float64
	type Fahrenheit float64

5、 Celsius(f)  并不是函数调用，而是类型转换;

*/

func main()  {
	type Celsius float64    	// 声明了一个数据类型 Celsius 底层类型是float64
	type Fahrenheit float64

	const (
		AbsoluteZeroC Celsius = -273.15 // 绝对零度
		FreezingC     Celsius = 0       // 结冰点温度
		BoilingC      Celsius = 100     // 沸水温度
	)

	var c Celsius = 23.5
	var f Fahrenheit
	fmt.Println(c == 0)				// true
	fmt.Println(f==0)			    // true
	fmt.Println(c== Celsius(f))	    // true


}



