package main

import (
	"fmt"
	"sync"
)

var count = 0

var wg sync.WaitGroup
var mutex sync.Mutex  // 互斥锁
						// 读写 互斥锁


func test()  {
	mutex.Lock()
	count++
	fmt.Println("the count is:",count)
	//time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()
}


func main()  {
	for r := 0; r <20 ; r++ {
		wg.Add(1)
		go test()
	}
	wg.Wait()


}
