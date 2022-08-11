package main

import "fmt"

func main()  {
	fmt.Println("main")

	/*
		指针地址、指针类型和指针取值。
		每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go语言中使用&字符放在变量前面对变量进行“取地址”操作
		Go语言中的值类型（int、float、bool、string、array、struct）都有对应的指针类型，如：*int、*int64、*string
	取变量指针的语法如下：
	    ptr := &v    // v的类型为T
	    v:代表被取地址的变量，类型为T
	    ptr:用于接收地址的变量，ptr的类型就为*T，称做T的指针类型。*代表指针。


		在对普通变量使用&操作符取地址后会获得这个变量的指针，然后可以对指针使用*操作，也就是指针取值，代码如下。

	*/
	//指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)   // type of b:*int          (b是指针类型 )
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)   // type of c:int
	fmt.Printf("value of c:%v\n", c)  // value of c:10



	//  空指针
	//	当一个指针被定义后没有分配到任何变量时，它的值为 nil
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%v\n", p)  //<nil>
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值") //空值
	}
/*
	   在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
		而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。要分配内存，就引出来今天的new和make。
		Go语言中new和make是内建的两个函数，主要用来分配内存

		new是一个内置的函数，它的函数签名如下：
			func new(Type) *Type
			1.Type表示类型，new函数只接受一个参数，这个参数是一个类型
			2.*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。

		   a := new(int)
		   b := new(bool)
		   fmt.Printf("%T\n", a) // *int
		   fmt.Printf("%T\n", b) // *bool
		   fmt.Println(*a)       // 0
		   fmt.Println(*b)       // false


		var a *int只是声明了一个指针变量a但是没有初始化，指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值。
		应该按照如下方式使用内置的new函数对a进行初始化之后就可以正常对其赋值了：

		 var a *int        // 声明
		 a = new(int)      // 分配空间
		 *a = 10           // 赋值
		 fmt.Println(*a)


		//	make
		make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，
		因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。

		func make(t Type, size ...IntegerType) Type
		make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作

		new与make的区别
     	1.二者都是用来做内存分配的。
       2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
       3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
*/



}
