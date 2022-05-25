package main

import (
	"fmt"
	"time"
)

// 在一些场景下，我们需要同时从多个通道接受数据

func main()  {

	// 定义一个管道10 个 数据 int类型
	intChan:=make(chan int,10)
	for i := 0; i <10; i++ {
		intChan<-i
	}
	// 定义一个管道 5 个数据 string
	stringChan:= make(chan string,5)
	for i := 0; i < 5; i++ {
		stringChan<-"hello"+ fmt.Sprintf("%d",i)
	}

	for  {
		// 使用 select 来获取 channel 里面的数据的时候 不需要 关闭 channel 关闭后有问题
		select {
		case v:= <- intChan:
			fmt.Printf("从intChan读取的数据%d\n",v)
			time.Sleep(time.Millisecond * 50)
		case v:= <- stringChan:
			fmt.Printf("从stringChan读取的数据%d\n",v)
			time.Sleep(time.Millisecond * 50)
		default:
			fmt.Printf("获取完毕")
			return   // 跳出 for 死 循环
		}

	}


}
