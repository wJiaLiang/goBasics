package main

import "fmt"

func main()  {
/*
	一、结构体的定义
	类型名：标识自定义结构体的名称，在同一个包内不能重复。
	字段名：表示结构体字段名。结构体中的字段名必须唯一。
	字段类型：表示结构体字段的具体类型。
	  type 类型名 struct {
           字段名 字段类型
           字段名 字段类型
           …
       }
	同样类型的字段也可以写在一行，
	语言内置的基础数据类型是用来描述一个值的，而结构体是用来描述一组值的。比如一个人有名字、年龄和居住城市等，本质上是一种聚合型的数据类型
*/
type Person struct {
	name string
	city string
	age int8
}

/*
   二、结构体实例化
	只有当结构体实例化时，才会真正地分配内存，也就是说实例化后才能使用
	结构体本身也是一种类型,可以用 var 那样声明一个结构体
*/
	var p Person
	p.name = "kk"
	p.age = 12
	p.city ="北京"



/*
  三、匿名结构体
	在定义一些临时数据结构等场景下还可以使用匿名结构体。
*/
	var user struct{Name string; Age int}
	user.Name = "ss"
	user.Age = 22

/*
	四、创建指针类型结构体
	我们还可以通过使用new关键字对结构体进行实例化，得到的是结构体的地址
*/
	var p2 = new(Person)
	fmt.Printf("p2=%T \n",p2)  // *main.person  结构体指针类型
	p2.name = "呵呵"
	fmt.Printf("p2=%#v \n",p2)  // p2=&main.Person{name:"呵呵", city:"", age:0}

/*
	五、取结构体的地址 实例化
	使用&对结构体进行取地址操作相当于对该结构体类型进行了一次 new 实例化操作。
*/
	var p3 = &Person{}
	fmt.Printf("p3=%#v \n",p3)  // p3=&main.Person{name:"", city:"", age:0}

	/**********************************************************************************/

/*
	六、结构体的初始化
	使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值
	1、使用键值对初始化
	2、使用值的列表初始化
	3、
*/

 p4 := Person{
	name:"mmk",
	city:"广州",
	age:23,
}
fmt.Printf("p4=%#v \n",p4)





}
