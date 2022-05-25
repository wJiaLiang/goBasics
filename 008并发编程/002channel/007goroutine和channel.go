package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fn1(intChan chan int) {
	for i := 0; i < 100; i++ {
		intChan <- i + 1
		fmt.Println("writeData 写入数据-", i+1)
		time.Sleep(time.Millisecond * 100)
	}
	close(intChan)
	wg.Done()
}

func fn2(intChan chan int) {
	for v := range intChan {
		fmt.Printf("readData 读到数据=%v\n", v)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done()
}

func main() {
	/*
		需求 1：定义两个方法，一个方法给管道里面写数据，一个给管道里面读取数据。要求同步 进行。
		1、开启一个 fn1 的的协程给向管道 inChan 中写入 100 条数据
		2、开启一个 fn2 的协程读取 inChan 中写入的数据
		3、注意：fn1 和 fn2 同时操作一个管道
		4、主线程必须等待操作完成后才可以退出
	*/

	allChan := make(chan int, 100)
	wg.Add(1)
	go fn1(allChan)

	wg.Add(1)
	go fn2(allChan)

	wg.Wait()
	fmt.Println("读取完毕...")

}
