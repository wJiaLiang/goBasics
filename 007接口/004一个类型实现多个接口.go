package main

import "fmt"

// Sayer6 接口
type Sayer6 interface {
	say()
}
// Mover6 接口
type Mover6 interface {
	move()
}
type dog6 struct {
	name string
}
// 实现Sayer接口
func (d dog6) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}
// 实现Mover接口
func (d dog6) move() {
	fmt.Printf("%s会动\n", d.name)
}


func main()  {
/*
	一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。
	例如，狗可以叫，也可以动。我们就分别定义Sayer接口和Mover接口
	dog 既可以实现Sayer接口，也可以实现Mover接口。
	实现结构体的实例赋值到 给多个 接口
*/
	var x Sayer6
	var y Mover6

	a:= dog6{name:"旺财"}
	x = a
	y = a
	x.say()
	y.move()

}
