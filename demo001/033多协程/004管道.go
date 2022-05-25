package main

import "fmt"

// 通过管道实现 多协程 并行 和 并发
/*
	channel 类型;
	channel 是一种类型，一种引用类型
	var ch1 chan int     //
	var ch2 chan bool    //
	var ch3 chan []int   // int 切片的管道

	声明 管道 后需要使用 make 函数初始话之后才能使用
	make(chan 元素类型 容量)

	发送 (将数据放在管道内)
	ch<-10   // 把10 发送到ch中

	接收 (从管道内取值)
	x:= <- ch  从 ch中接受 值 并赋值给变量x
	<- ch 从 ch中接受 值 忽略结果

	一存一取 后 管道就空了
*/

func main()  {
	ch:=make(chan int,3)
	ch<-19
	a:=<-ch
	fmt.Println(a)




//	for range 循环遍历通道 当通过关闭的时候 会退出 for range ;如果没有就会报错
	var chn = make(chan int,10)
	for i := 1; i <= 10; i++ {
		chn<-i
	}
	close(chn)  //关闭管道

	for val := range chn {
		fmt.Println(val)
	}

}
