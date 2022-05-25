package main

import (
	"fmt"
	"sync"
	"time"
)
var wg = sync.WaitGroup{}

//  开启一个 协程     区分 并发，并行
//  如果主线程 执行得比较快的话，执行完成后 直接退出，不管协程是否执行完成;
//  使用 sync.WaitGroup 来解决这个问题;
func test() {
	for i := 0; i <10 ; i++ {
		fmt.Println("test(),Golang-",i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()  // 协程计数器 加 -1
}


func main()  {
	wg.Add(1)  // 协程计数器 加 1
	go test()  // 开启一个协程;

	for i := 0; i <10 ; i++ {
		fmt.Println("main(),Golang,",i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Wait()
	fmt.Println("主线程退出。。。")

}
