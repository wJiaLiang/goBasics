package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

func main() {
	fmt.Println("结构体")

	/*
		Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。
	*/

	/*
	   类型别名和自定义类型

	*/

	/*
		使用type和struct关键字来定义结构体，具体代码格式如下：
		type 类型名 struct {
			   字段名 字段类型
			   字段名 字段类型
			   …
		}
	*/

	var p1 person
	p1.name = "pprof.cn"
	p1.city = "北京"
	p1.age = 18

}
