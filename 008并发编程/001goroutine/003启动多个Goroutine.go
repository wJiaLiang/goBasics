package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)
var wg2 sync.WaitGroup

func test(num int)  {
	for i := 1; i <=5 ; i++ {
		fmt.Printf("协程(%v)打印的第(%v)条数据\n",num,i)
		time.Sleep(time.Millisecond*50)
	}
	wg2.Done()
}

func main()  {
/*
	1、启动多个协程  循环启动 协程 调用打印函数
	2、获取当前计算机上面的 Cup 线程数
*/
	for i := 0; i < 5; i++ {
		wg2.Add(1)
		go test(i)
	}
	wg2.Wait()
	fmt.Println("主线程退出....")


	num:=runtime.NumCPU()
	fmt.Println("当前cup的核线程数:",num)  //12


}
