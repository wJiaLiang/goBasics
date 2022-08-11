package main

import (
	"fmt"
	"reflect"
)

/*
有时我们需要写一个函数，这个函数有能力统一处理各种值类型，而这些类型可能无法共享
同一个接口，也可能布局未知，也有可能这个类型在我们设计函数时还不存在，这个时候我
们就可以用到反射


反射是指在程序运行期间对程序本身进行访问和修改的能力。正常情况程序在编译时，变量
被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取
自身的信息。支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、
结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运
行期获取类型的反射信息，并且有能力修改它们。

1、反射可以在程序运行期间动态的获取变量的各种信息，比如变量的类型 类别
2、如果是结构体，通过反射还可以获取结构体本身的信息，比如结构体的字段、结构体的
方法、结构体的 tag。
3、通过反射，可以修改变量的值，可以调用关联的方法
*/


// 反射 获取 任意变量的类型

type  myInt int
type Person struct {
	Name string
	Age int
}
func reflectFn(x interface{})  {
	v:=reflect.TypeOf(x)
	//fmt.Println(v)
	fmt.Printf("类型:%v 类型名称:%v 类型种类:%v \n",v,v.Name(),v.Kind())
}

func main()  {
	a:=10
	b:=23.43
	c:=true
	d:="你好golang"
	reflectFn(a)
	reflectFn(b)
	reflectFn(c)
	reflectFn(d)

	var f myInt = 23
	g:=Person{
		Name: "sss",
		Age: 23,
	}
	reflectFn(f)  // main.myInt
	reflectFn(g)  // main.Person
	var e = 23
	reflectFn(&e) // *int  类型:*int 类型名称: 类型种类:ptr

	var i=[3]int{2,3,4}
	var j=[]int{11,33,55}
	reflectFn(i)  // 类型:[3]int 类型名称: 类型种类:array
	reflectFn(j)  // 类型:[]int 类型名称: 类型种类:slice

}
