package main

import "fmt"

func main()  {

	/*
		fatal error: all goroutines are asleep - deadlock!
		报错是因为我们使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候才能发送值。
	*/

	ch:= make(chan int)

	ch <- 10
	fmt.Println("发送成功")


}