package main

import "fmt"

// 2、指针接收者
type Mover1 interface {
	move()
}
type dog2 struct { }

func (d *dog2) move() {
	fmt.Println("狗会动2")
}

func main()  {
	/*
	此时实现Mover接口的是*dog2类型，所以不能给y传入dog类型的 wangwang，此时y只能存储*dog2类型的值。
			var y Mover1
			var wangwang = dog2{}
			y = wangwang  // 报错，无法付给指针类型的接收者
			y.move()

			fugui:= &dog2{}
			y = fugui
			y.move()  //狗会动2

	*/




}