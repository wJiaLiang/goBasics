package main

import "fmt"

func main(){
/*
	 可以通过内置的close()函数关闭channel（如果你的管道不往里存值或者取值的时候一定记得关闭管道）
	 如果不关闭 channel 取到最后就会 报错 fatal error: all goroutines are asleep - deadlock!

	 关于关闭通道需要注意的事情是，只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。
	 通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。

	 1.对一个关闭的通道再发送值就会导致panic。
	 2.对一个关闭的通道进行接收会一直获取值直到通道为空。
	 3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
	 4.关闭一个已经关闭的通道会导致panic。

*/
	c:= make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c<-i
		}
		close(c)
	}()

	for {
		if data,ok :=<-c;ok {
			fmt.Println(data)

		}else{
			fmt.Println(data,ok) // 0,false
			break
		}
	}
	fmt.Println("main 结束")


}
