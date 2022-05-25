package main

import (
	"fmt"
	"sync"
)

var (
	myMap = make(map[int]int)
	wg2 sync.WaitGroup
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	lock.Lock()    //加锁
	myMap[n] = res

	lock.Unlock()  //解锁
	wg2.Done()
}

func main() {
	/*
		互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。
		Go语言中使用sync包的Mutex类型来实现互斥锁

		使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；
		当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的。

		标准库 sync 中的 Mutex 结构体类型表示。sync.Mutex 类型只有两个公开的指针方法，Lock 和 Unlock。
		Lock 锁定当 前的共享资源，Unlock 进行解锁

	    需求：现在要计算 1-60 的各个数的阶乘，并且把各个数的阶乘放入到 map 中。最后显示出 来。要求使用 goroutine 完成。

	   	1. 编写一个函数，来计算各个数的阶乘，并放入到 map 中
	   	2. 启动多个协程，将统计的将结果放入到 map 中

	   	只使用 Goroutine 实现,运行的时候可能会出现资源争夺问题 concurrent map writes



	*/

	for i := 1; i <= 60; i++ {
		wg2.Add(1)
		go test(i)
	}

	wg2.Wait()

	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}
