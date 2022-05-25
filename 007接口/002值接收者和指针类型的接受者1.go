package main

import "fmt"

// 1、值接收者
type Mover interface {
	move()
}
type dog1 struct {}

func (d dog1) move() {
	fmt.Println("狗会动")
}




func main()  {
	/*
		使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。
		因为Go语言中有对指针类型变量求值的语法糖，dog指针fugui内部会自动求值*fugui。
	*/
	var x Mover
	var wanagwang = dog1{}
	x = wanagwang
	x.move()

	var fugui = &dog1{}
	x = fugui
	x.move()




}




