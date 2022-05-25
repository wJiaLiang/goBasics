package main

import "fmt"

/*
在 go 语言中，没有类的概念但是可以给类型（结构体，自定义类型）定义方法。所谓方法 就是定义了接收者的函数。
接收者的概念就类似于其他语言中的 this 或者 self。
方法的定义格式如下：
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
	函数体
}
1.接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名的第一个小写字母，而不是self、this之类的命名。
	例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
2.接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
3.方法名、参数列表、返回参数：具体格式与函数定义相同。
注意: 方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。
	结构体可以作为函数的参数和返回值
*/


/*
	给结构体 Person 定义一个方法打印 Person 的信息
*/
type Persion struct {
	name string
	age int8
}
//构造函数
func newPerson(name string, age int8) *Persion {
	return &Persion{
		name: name,
		age:  age,
	}
}
// 定义一个 printlnfo 方法
func (p Persion) printlnfo()  {
	fmt.Printf("姓名:%#v  年龄:%#v \n", p.name,p.age)
}
/*
	指针类型的接收者:
   指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。
	这种方式就十分接近于其他语言中面向对象中的this或者self。 例如我们为Person添加一个SetAge方法，来修改实例变量的年龄。
*/
func (p *Persion) SetAge(newAge int8)  {
	p.age = newAge
}

/*
	值类型的接收者:
	当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。
	在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。
*/
func (p Persion) SetAge2(newAge int8)  {
	p.age = newAge
}



func main()  {
	p1 := Persion{
		name: "kg",
		age:  25,
	}
	p1.printlnfo()   		// 姓名:"kg"  年龄:25

	p1.SetAge(30)
	p1.printlnfo()   		// 姓名:"kg"  年龄:25

	p1.SetAge2(66)
	p1.printlnfo()  		//
	fmt.Println(p1.age)   	//30

	/*
		什么时候应该使用指针类型接收者
		 1.需要修改接收者中的值
	     2.接收者是拷贝代价比较大的大对象
	     3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
	*/


}

