package main

import (
	"fmt"
	"reflect"
)


// 反射 获取 任意变量的类型

func reflectFn(x interface{})  {

	b,_:=x.(int)  // 类型断言 来改变 数据类型
	var num = b+10
	fmt.Println(num)

	//反射来实现
	v:=reflect.ValueOf(x)
	//var n = v+12
	//fmt.Println(n)

	var m = v.Int() + 12
	fmt.Println(m)



}

func main()  {
	a:=10
	reflectFn(a)



}

