package main

import "fmt"

// 结构体也是值类型

type Person struct {
	Name	string
	Age 	int
	Sex 	string
	height  int
}

func (p Person) PrintInfo()  {
	fmt.Printf("姓名：%v  年龄：%v \n",p.Name,p.Age)
}

// 如果想修改值,这里必须是指针类型;
func (p *Person) setInfo(name string,age int)  {
	p.Name = name
	p.Age = age
}



func main()  {
	var p1 = Person{
		Name: "张三",
		Age: 24,
		Sex: "女",
	}
	p1.PrintInfo()  // 姓名：张三  年龄：24
	p1.setInfo("李斯特",23)  // 姓名：张三  年龄：24

	p1.PrintInfo() // 姓名：李斯特  年龄：23

}
