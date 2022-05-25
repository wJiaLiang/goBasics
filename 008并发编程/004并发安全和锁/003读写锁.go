package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var mutex sync.RWMutex
var lock2 sync.Mutex
var wg7 sync.WaitGroup

// 写的方法
func write() {
	mutex.Lock()       // 加写锁
	//lock2.Lock()       // 加互斥锁

	fmt.Println("++++执行写操作")
	time.Sleep(time.Second * 2)   // 假设 写操作要的时间 2秒

	mutex.Unlock()     // 解写锁
	//lock2.Unlock()     // 解互斥锁

	wg7.Done()
}

// 读的方法
func read() {
	mutex.RLock()     // 加读锁
	//lock2.Lock()      // 加互斥锁

	fmt.Println("----执行读操作")
	time.Sleep(time.Second * 1)   // 假设 读操作需要的时间 1秒

	mutex.RUnlock()   // 解开读锁
	//lock2.Unlock()    // 解互斥锁

	wg7.Done()
}

func main() {
	/*
	      1、互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，
	   	这种场景下使用读写锁是更好的一种选择。读写锁在Go语言中使用sync包中的RWMutex类型。

	      2、读写锁可以让多个读操作并发，同时读取，但是对于写操作是完全互斥的。也就是说，
	   	当一 个 goroutine 进行写操作的时候，其他 goroutine 既不能进行读操作，也不能进行写操作。


	   	GO 中的读写锁由结构体类型 sync.RWMutex 表示。此类型的方法集合中包含两对方法：
	   		1、一组是对写操作的锁定和解锁，简称“写锁定”和“写解锁”：
	   			func (*RWMutex)lock2()
	   			func (*RWMutex)Unlock2()

	   		2、另一组表示对读操作的锁定和解锁，简称为“读锁定”与“读解锁”：
	   			func (*RWMutex)Rlock2()
	   			func (*RWMutex)RUnlock2()读写锁示例：

			互斥锁  串行执行
			读写锁  读并行执行，写串行执行
	*/

	// 开启 10 个协程执行写操作
	for i := 0; i < 10; i++ {
		wg7.Add(1)
		go read()
	}

	// 开启 10 个协程执行读操作
	for i := 0; i < 10; i++ {
		wg7.Add(1)
		go write()
	}
	wg7.Wait()
	fmt.Println("main 执行完成")

}
