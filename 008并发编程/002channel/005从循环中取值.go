package main

import "fmt"

func main()  {
/*
	当通过通道发送有限的数据时，我们可以通过close函数关闭通道 来告知 从该通道接收值的 goroutine 停止等待。
	当通道被关闭时，往该通道发送值会引发panic，从该通道里接收的值一直都是类型零值。

*/
	ch5 := make(chan int)
	ch6 := make(chan int)

	// 开启goroutine将0~100的数发送到 ch5 中
	go func() {
		for i := 0; i < 100; i++ {
			ch5 <- i
			//fmt.Println("写入数据",i)
		}
		close(ch5)
	}()

	// 开启goroutine从 ch5 中接收值，并将该值的平方发送到 ch6 中
	go func() {
		for {
			i, ok := <-ch5 		// 通道关闭后再取值 ok= false
			if !ok {
				break
			}
			ch6 <- i * i
			//fmt.Println("读取数据",i)
		}
		close(ch6)
	}()


	// 在主 goroutine中从 ch6 中接收 值打印
	for i := range ch6 { 		// 通道关闭后会退出for range循环
		fmt.Println(i)
	}

/*
    1、从上面的例子中我们看到有两种方式在接收值的时候判断通道是否被关闭，我们通常使用的是for range的方式。

    2、使用 for range 遍历管道，当管道被关闭的时候就会退出 for range,
	   如果没有关闭管道 就会报个错误 fatal error: all goroutines are asleep - deadlock!

	3、通过 for range 来遍历管道数据 管道返回值没有 key


*/



}
