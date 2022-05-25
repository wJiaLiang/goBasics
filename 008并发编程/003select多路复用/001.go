package main

import "fmt"

//
func main()  {
/*
	1、在某些场景下我们需要同时从多个通道接收数据。通道在接收数据时，如果没有数据可以接收将会发生阻塞。
    for  {
   		data,ok := <-ch1
   		data,ok := <-ch2
   		data,ok := <-ch3
   		fmt.Println(data,ok)
   	}
    2、这种方式虽然可以实现从多个通道接收值的需求，但是运行性能会差很多。
	为了应对这种场景，Go内置了select关键字，可以同时响应多个通道的操作。

	3、select的使用类似于switch语句，它有一系列case分支和一个默认的分支。每个case会对应一个通道的通信（接收或发送）过程。
	   select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句。

	4、select可以同时监听一个或多个channel，直到其中一个channel ready
	5、如果多个channel同时ready，则随机选择一个执行
	6、可以用于判断管道是否存满

*/

	ch1:= make(chan int,10)
	ch2:= make(chan int,10)
	ch3:= make(chan int,10)

	select {
	case <-ch1:
		// 如果 ch1 成功读到数据，则进行该case处理语句
		fmt.Println("case1")
	case ch2<-1:
		// 如果成功向 ch2 写入数据，则进行该case处理语句
		fmt.Println("case2")   //case2
	case ch3<-2:
		// 如果成功向 ch3 写入数据，则进行该case处理语句
		fmt.Println("case3")
	default:
		// 如果上面都没有成功，则进入default处理流程
		fmt.Println("default...")

	}




}
