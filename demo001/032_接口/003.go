package main

import "fmt"

// 定义一个 Animal 的接口，Animal中定义两个方法,分别是 SetName 和 GetName, dog结构体 和 cat结构体去实现

type Animaler interface {
	SetName(string)
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


type Cat struct {
	Name string
}
// 给 Dog类型 的结构体绑定方法;
func (c *Cat) SetName(name string) {
	c.Name = name
}
func (c Cat) GetName() string {
	return c.Name
}


func main()  {
	// Dog 实现Animal的接口
	d := &Dog{
		Name:"黑狗",
	}
	var d1 Animaler = d

	fmt.Printf(d1.GetName())
	d1.SetName("白狗")
	fmt.Printf(d1.GetName())

	//用 cat 实现 Animaler 的接口


}
