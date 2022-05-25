package main

import (
	"fmt"
	"reflect"
)


func main() {
	/*

		为什么要使用反射？
		有时我们需要写一个函数，这个函数有能力统一处理各种值类型，而这些类型可能无法共享 同一个接口，也可能布局未知，
		也有可能这个类型在我们设计函数时还不存在，这个时候我 们就可以用到反射。
		因为函数指定具体参数就无法，获取任意类型了，所以指定了参数为 空接口，空接口可以存储任意类型的值

		反射是指在程序运行期间对程序本身进行访问和修改的能力。
		正常情况程序在编译时，变量 被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取 自身的信息。
		支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、 结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，
		这样就可以在程序运 行期获取类型的反射信息，并且有能力修改它们。

		Golang 中反射可以实现以下功能：
			1、反射可以在程序运行期间动态的获取变量的各种信息，比如变量的类型 类别
			2、如果是结构体，通过反射还可以获取结构体本身的信息，比如结构体的字段、结构体的 方法、结构体的 tag
			3、通过反射，可以修改变量的值，可以调用关联的方法

		Go 语言中的变量是分为两部分的:
			• 类型信息：预先定义好的元信息。
			• 值信息：程序运行过程中可动态变化的

		在 GoLang 的反射机制中，任何接口值都由是一个具体类型 和 具体类型的值 两部分组成的。
		反射的相关功能由内置的 reflect 包提供，任意接口值在反射中都可以理解为 由 reflect.Type 和 reflect.Value 两部分组成 ，
		并且reflect包提供了reflect.TypeOf 和 reflect.ValueOf 两个重要函数来获取任意对象的Value和Type。


		reflect.TypeOf()函数可以接受任意 interface{}参数，可以获得任意值的 类型对象（reflect.Type）
		程序通过类型对象可以访问任意值的类型信息。

		reflect.TypeOf(3) 把3赋值给 interface{} 参数，把一个具体值赋值给一个接口类型时会发生一个隐式类型
		转换会生成一个包含两个部分内容的接口值：动态类型部分是操作数的类型，动态值部分是操作数的值;
		并且把接口中的动态类型以reflect.Type形式返回

		reflect.ValueOf(3)函数可以接受任意 interface{}参数，并将接口的动态值以reflect.Value的形式返回,
		mreflect.ValueOf的返回值也是具体值



	*/

	a:=10
	b:=23.5
	c:=true
	d:="你好go"

	reflectType(a)  	// int
	reflectType(b)  	// float64
	reflectType(c)  	// bool
	reflectType(d)  	// string

	reflectType(aa)  	// main.myint
	reflectType(e)      // main.user
	reflectType(&h)     // *int
}

// 定义类型
type myint int
var aa myint

// 结构体
type user struct {
	name string
	age int
}
var e = user{
	name:"呵呵阿",
	age:18,
}

var h = 23

// 通过反射获取任何变量的类型
func reflectType(x interface{}) {

	v := reflect.TypeOf(x)
	fmt.Printf("type:%v \n", v)
}
