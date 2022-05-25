package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 每100毫秒打印一次
func test1() {
	for i := 0; i <10 ; i++ {
		fmt.Println("test1(),Golang-",i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()  // 协程计数器 加 -1
}

// 每50毫秒打印一次
func test2() {
	for i := 0; i <10 ; i++ {
		fmt.Println("test2(),Golang-",i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done()  // 协程计数器 加 -1
}

func main()  {
/*
	如果主线程 执行得比较快的话，执行完成后 直接退出，不管协程是否执行完成;
   	使用 sync.WaitGroup 来解决这个问题;
	sync.WaitGroup 可以实现主线程等待协程执行完毕。
		首先全局定义一个变量 wg
		后面每开启一个协程的时候 wg.Add(1) 加一
		每个协程执行完成后 调用  wg.Done() 方法
*/
	wg.Add(1)
 	go test1()
	wg.Add(1)
 	go test2()

 	wg.Wait()
 	fmt.Println("主线程退出。。。")



}
