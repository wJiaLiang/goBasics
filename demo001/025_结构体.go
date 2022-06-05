package main

import (
	"fmt"
)

// 自定义类型
type myInt int
type myFn func(int,int) int

// 类型别名 1.9版本后增加
type myFloat = float64

//go 中没有类的概念
/*
结构体的定义
使用 type 和 struct 关键字来定义结构体，
type 类型名 struct {
	字段名  字段类型
	字段名  字段类型
	。。。。
}
*/
type Person struct {
	name string
	age int
	sex string
}




func main()  {
	var a myInt = 10
	fmt.Printf("%v---%T \n",a,a) //10---main.myInt

	var b myFloat = 12.6
	fmt.Printf("%v---%T \n",b,b) //12.6---float64

	//	 实例化结构体 方法一：
	var p1 Person
	p1.name = "张三"
	p1.sex = "女"
	p1.age = 30
	fmt.Printf("值：%v 类型：%T  \n",p1,p1)  // 值：{张三 30 女} 类型：main.Person
	fmt.Printf("值：%#v 类型：%T \n",p1,p1) // 值：main.Person{name:"张三", age:30, sex:"女"} 类型：main.Person

	//	 实例化结构体 方法二  通过new 关键字对结构体进行实例化  返回的是指针类型
	var p2 = new(Person)
	p2.name="李四"
	fmt.Printf("值：%#v 类型：%T \n",p2,p2) //值：&main.Person{name:"李四", age:0, sex:""} 类型：*main.Person


	//	 实例化结构体 方法三  通过 &符号 关键字对结构体进行实例化   返回的是指针类型
	var p3 = &Person{}
	p3.name = "王五"
	fmt.Printf("值：%#v 类型：%T \n",p3,p3) //值：&main.Person{name:"王五", age:0, sex:""} 类型：*main.Person

	//	 实例化结构体 方法四
	var p4 = Person{
		name:"呵呵呵",
		sex:"1",
		age:33,
	}
	fmt.Printf("值4：%#v 类型4：%T \n",p4,p4) // 值4：main.Person{name:"呵呵呵", age:33, sex:"1"} 类型4：main.Person


	//	 实例化结构体 方法五
	var p5 = &Person{
		name:"呵呵呵",
		sex:"1",
		age:33,
	}
	fmt.Printf("值5：%#v 类型5：%T \n",p5,p5) // 值5：&main.Person{name:"呵呵呵", age:33, sex:"1"} 类型5：*main.Person

	//	 实例化结构体 方法六 部分赋值，其他不赋值的是该数据类型的默认值;
	var p6 = &Person{
		name:"呵呵呵",
		sex:"1",
		age:33,
	}
	fmt.Printf("值5：%#v 类型5：%T \n",p6,p6)

	//	 实例化结构体 方法七 顺序要一致
	var p7 = &Person{
		"呵呵呵",
		38,
		"男",
	}
	fmt.Printf("值5：%#v 类型5：%T \n",p7,p7)



}
