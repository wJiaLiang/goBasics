package main

import "fmt"

// 默认的管道是可以读可写的
func main()  {

	ch1 := make(chan int,2)
	ch1<-10
	ch1<-12
	var m = <-ch1
	var m1 = <-ch1
	fmt.Println(m,m1)


	//	声明为只写的管道
	ch2:= make(chan<- int ,2)
	ch2 <-32
	ch2 <-21

//	 声明一个只读的管道
	ch3 := make(<-chan int,2)
	ch3<-30


}
