package main

import "fmt"

/*
Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。
	通过嵌套来实现继承
结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

*/
// Animal 动物结构体  父结构体
type Animal struct {
	name string
}
// 动物的公共方法
func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗  子结构体
type Dog struct {
	Feet    int8
	*Animal 	//通过嵌套匿名结构体实现继承
}
// dog 独有的方法;
func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}


func main()  {
	d1:= &Dog{
		Feet:  4,
		Animal: &Animal{
			name:"乐乐",
		},
	}
	fmt.Printf("%#v \n",d1)
	d1.wang()
	d1.move()
}
