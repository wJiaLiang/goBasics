package main

import (
	"fmt"
	"time"
)

func write(ch chan string) {
	//  死 for 循环
	i:=1
	for {
		i+=1
		select {
		// 写数据 满了会匹配到下个 case ,都执行完了就会执行  default
		case ch <- "hello":
			fmt.Printf("写入 hello: %d \n",i)
		case <-ch:
			fmt.Println("取出")
		default:
			fmt.Println("channel full 满了")  // 一个存一个取，这里default 永不执行;
			return  // 退出，注意退出for 循环
		}
		time.Sleep(time.Millisecond * 500)
	}
}


func main()  {
/*
	可以用于判断管道是否存满

	使用 select 可以解决从管道取数据的阻塞问题,传统的方法在遍历管道时，如果不关闭会阻 塞而导致 deadlock,
	在实际开发中，可能我们不好确定什么关闭该管道.

	这里满了，不会一直阻塞而 deadlock，会自动到 下一个 case 匹配
	default 中注意退出循环

*/
	// 创建管道
	output1 := make(chan string, 5)

	// 子协程写数据
	go write(output1)

	// 取数据
	for s := range output1 {
		fmt.Println("读取:", s)
		time.Sleep(time.Second)
	}

	fmt.Printf("main 线程结束")

}
