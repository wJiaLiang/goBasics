package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)
var wg = sync.WaitGroup{}

func test(num int)  {
	for i := 1; i <=5 ; i++ {
		fmt.Printf("协程(%v)打印的第(%v)条数据\n",num,i)
		time.Sleep(time.Millisecond*50)
	}
	wg.Done()
}

func main()  {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=",cpuNum)

	for i := 0; i <=6 ; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
}
