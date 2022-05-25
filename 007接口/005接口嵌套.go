package main

import "fmt"

// Sayer 接口
type Sayer7 interface {
	say()
}

// Mover 接口
type Mover7 interface {
	move()
}

// 接口嵌套
type animal interface {
	Sayer7
	Mover7
}
type cat7 struct {
	name string
}
func (c cat7) say() {
	fmt.Println("喵喵喵")
}

func (c cat7) move() {
	fmt.Println("猫会动")
}


func main()  {
/*
	接口与接口间 可以通过嵌套创造出新的接口。
	嵌套得到的接口的使用与普通接口一样，这里我们让cat实现 animal接口：
*/
	var xx animal
	xx = cat7{name:"花猫"}
	xx.move()
	xx.say()


}
