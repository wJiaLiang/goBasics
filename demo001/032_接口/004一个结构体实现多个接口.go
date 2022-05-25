package main

import "fmt"

// 定义一个 Animal 的接口，Animal中定义两个方法,分别是 SetName 和 GetName, dog结构体 和 cat结构体去实现

type Animaler1 interface {
	SetName(string)
}
type Animaler2 interface {
	GetName() string
}

type Dog struct {
	Name string
}
// 给 Dog类型 的结构体绑定方法;
func (d *Dog) SetName(name string)  {
	d.Name = name
}
func (d Dog) GetName() string  {
	return d.Name
}


func main()  {
	// Dog 实现Animal的接口
	d := &Dog{
		Name:"big黑狗",
	}
	var d1 Animaler1 = d
	var d2 Animaler2 = d

	fmt.Printf(d2.GetName()+"\n")
	d1.SetName("small白狗")
	fmt.Printf(d2.GetName())



}

