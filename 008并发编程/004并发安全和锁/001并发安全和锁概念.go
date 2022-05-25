package main

import (
	"fmt"
	"sync"
)

var x int = 0
var wg sync.WaitGroup

func add()  {
	for i := 0; i < 100; i++ {
		x = x+1
	}
	wg.Done()
}



func main()  {
/*
	有时候在Go代码中可能会存在多个goroutine同时操作一个资源（临界区），这种情况会发生竞态问题（数据竞态）。
	类比现实生活中的例子有十字路口被各个方向的的汽车竞争；还有火车上的卫生间被车厢里的人竞争。


*/
	wg.Add(1)

	go add()
	go add()

	wg.Wait()
	fmt.Println("main 结束")

}


/*
	上面的代码中我们开启了两个goroutine去累加变量x的值，
	这两个goroutine在访问和修改x变量的时候就会存在数据竞争，导致最后的结果与期待的不符。
*/






