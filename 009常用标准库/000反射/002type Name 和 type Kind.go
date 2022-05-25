package main

import (
	"fmt"
	"reflect"
)

//定义一个函数 通过反射获取任何变量的类型
func reflectFn(x interface{}) {
	v := reflect.TypeOf(x)
	v.Name()  // 类型名称
	v.Kind()  // 类型种类  种类（Kind）就是指底层的类型

	fmt.Printf("类型:%v 类型名称:%v 类型种类%v\n", v,v.Name(),v.Kind())
}
// 定义类型
type myInt int
// 定义结构体
type Person struct {
	Name string
	Age  int
}

func main()  {
/*
	在反射中关于类型还划分为两种：类型（Type）和种类（Kind）。
	因为在 Go 语言中我们可 以使用 type 关键字构造很多自定义类型，而种类（Kind）就是指底层的类型，
	但在反射中， 当需要区分指针、结构体等大品种的类型时，就会用到种类（Kind）。

	Go 语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。
*/
	a := 10
	b := 23.4
	c := true
	d := "你好"
	reflectFn(a)	// int
	reflectFn(b)	// float64
	reflectFn(c)	// bool
	reflectFn(d) 	// string

	var e myInt = 666
	var f = Person{
		Name:"李白",
		Age:77,
	}
	reflectFn(e)  // 类型:main.myInt 类型名称:myInt 类型种类int
	reflectFn(f)  // 类型:main.Person 类型名称:Person 类型种类struct

	var h = 33

	var i = [6]int{1,2,3,4,5,6}

	var j = []int{123}

	reflectFn(&h)  // 类型:*int 类型名称: 类型种类ptr
	reflectFn(i)   // 类型:[6]int 类型名称: 类型种类array
	reflectFn(j)   // 类型:[]int 类型名称: 类型种类slice
}

