package main

import "fmt"

type Aface interface {
	SetName(string)
}
type Bface interface {
	GetName() string
}
// 接口嵌套; 嵌套  Ainterface 和 Binterface;
type Animalers interface {
	Aface
	Bface
}

type Dogs struct {
	Name string
}
// 给 Dog类型 的结构体绑定方法;
func (d *Dogs) SetName(name string) {
	d.Name = name
}
func (d Dogs) GetName() string {
	return d.Name
}

func main()  {
	// Dog 实现Animal的接口
	d := &Dogs{
		Name:"big黑狗",
	}
	var d1s Animalers = d

	fmt.Printf(d1s.GetName()+"\n")
	d1s.SetName("small白狗")
	fmt.Printf(d1s.GetName())


}



