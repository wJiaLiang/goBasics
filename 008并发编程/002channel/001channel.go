package main

import "fmt"

func main()  {
/*
	一、概念
		1、单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。

		2、虽然可以使用共享内存进行数据交换，但是共享内存在不同的goroutine中容易发生竞态问题。
	   	   为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。

		3、Go语言的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。

		4、如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。
	   		channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。

		5、Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，
	   	   总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。
	   	   每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。


	二、channel类型
		1、channel是一种类型，一种引用类型。声明通道类型的格式如下：
		    var 变量 chan 元素类型

			var ch1 chan int   // 声明一个传递整型的通道
       		var ch2 chan bool  // 声明一个传递布尔型的通道
       		var ch3 chan []int // 声明一个传递int切片的通道

		2、通道是引用类型，通道类型的空值是nil。

		3、声明的通道后需要使用make函数初始化之后才能使用。
		   make(chan 元素类型, [缓冲大小])

		4、通道有发送（send）、接收(receive）和关闭（close）三种操作。


*/
	var ch chan int
	fmt.Println(ch)  // <nil>

	ch1 := make(chan int,10)
	fmt.Println(ch1)   // 0xc000110000  返回地址

	//1、发送 将一个值发送到管道中
	ch1<-10  // 把10发送到ch中

	//2、接收  从一个通道中接收值。
	x:= <- ch1
	<-ch1      // 从ch1中接收值,忽略结果
	fmt.Println(x)

	//3、关闭
	close(ch1)

















}
