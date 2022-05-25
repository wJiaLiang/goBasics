package main

import "fmt"

/*
在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。
举个例子，我们基于内置的int类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。

注意事项： 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。
*/
type  myInt int

func (m myInt) SayHello()  {
	fmt.Println("Hello,int类型的类型")
}

func main()  {
	var m1 myInt
	m1.SayHello()    	// Hello,int类型的类型
	m1 = 100
	fmt.Printf("%#v  %T\n",m1,m1) //100  main.myInt


}
