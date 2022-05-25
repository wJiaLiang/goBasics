package main

import "fmt"

func main()  {
/*
	我们可以在使用make函数初始化通道的时候为其指定通道的容量，

	只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。
	就像你小区的快递柜只有那么个多格子，格子满了就装不下了，就阻塞了，等到别人取走一个快递员就能往里面放一个。

	我们可以使用内置的len函数获取通道内元素的数量，使用cap函数获取通道的容量，虽然我们很少会这么做。
*/

	ch2:= make(chan int,10)
	ch2<-10
	ch2<-9
	ch2<-8
	ch2<-7
	ch2<-6
	close(ch2)
	fmt.Printf("容量：%d\n长度：%d\n",cap(ch2),len(ch2))  //10  5
	<-ch2
	<-ch2
	<-ch2
	<-ch2
	<-ch2

	if a,b := <-ch2;b {
		fmt.Printf("a:%#v,b:%#v",a,b)
	}else{
		fmt.Println(a,b)  // 0,false
	}




}
