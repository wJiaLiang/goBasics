package main

import (
	"fmt"
	"reflect"
)

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
