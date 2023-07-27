package main

import "fmt"

/*
在 go 语言中，没有类的概念但是可以给类型（结构体，自定义类型）定义方法。
所谓方法就是定义了接收者的函数。接收者的概念就类似于其他语言中的 this 或者 self。

func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
	函数体
}

五、给任意类型添加方法
在 Go 语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。
举个例子，我们基于内置的 int 类型使用 type 关键字可以定义新的自定义类型，然后为我们
的自定义类型添加方法。
package main
import "fmt"
type myInt int
func (m myInt) SayHello() {
	fmt.Println("Hello, 我是一个 int。")
}
func main() {
	var m1 myInt
	m1.SayHello() //Hello, 我是一个 int。
	m1 = 100
	fmt.Printf("%#v %T\n", m1, m1) //100 main.MyInt
}
*/


// 结构体也是值类型
type Person2 struct {
	Name	string
	Age 	int
	Sex 	string
	height  int
}

func (p Person2) PrintInfo()  { 
	fmt.Printf("姓名：%v  年龄：%v \n",p.Name,p.Age)
}

// 如果想修改结构体里面的值,这里必须是指针类型;
func (p *Person2) setInfo(name string,age int)  {
	p.Name = name
	p.Age = age
}



func main()  {
	var p1 = Person2{
		Name: "张三",
		Age: 24,
		Sex: "女",
	}
	p1.PrintInfo()  // 姓名：张三  年龄：24
	p1.setInfo("李斯特",23)  // 姓名：张三  年龄：24

	p1.PrintInfo() // 姓名：李斯特  年龄：23

}
