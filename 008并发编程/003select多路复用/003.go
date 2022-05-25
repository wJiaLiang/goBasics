package main

import "fmt"

func main()  {
/*
	如果多个channel同时ready，则随机选择一个执行

*/

	// 创建2个管道
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)

	go func() {
		//time.Sleep(2 * time.Second)
		int_chan <- 1
	}()

	go func() {
		string_chan <- "hello"
	}()

	select {
	case value:=<-int_chan:
		fmt.Println("int:",value)
	case value:=<-string_chan:
		fmt.Println("string",value)
		
	}

	fmt.Println("main 线程结束")

}
