package main

import "fmt"

/*
	接口 （interface）是一种类型，一种抽象的类型，是一组函数的集合 ,不能包含任何变量;
	type 接口名 interface {
		方法名1（参数列表1） 返回值列表1
		方法名2（参数列表2） 返回值列表2
	}
*/

type Usber interface {
	start()
	stop()
}

// 如果接口里面有方法的话，必须要通过结构体或者通过自定义类型实现这个接口
type Phone struct {
	Name string
}

// 手机要实现usb 接口的话必须得实现 Usber 接口中的所有方法
func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}
func (p Phone) stop() {
	fmt.Println(p.Name, "关闭")
}
func main() {
	p := Phone{
		Name: "苹果手机",
	}
	var p1 Usber = p // 实现  Usber 接口 ,通过 结构体;实现;
	//p.start()
	p1.start()
	p1.start()

}
