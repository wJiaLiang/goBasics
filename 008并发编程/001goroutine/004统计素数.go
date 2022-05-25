package main

import (
	"fmt"
	"sync"
	"time"
)
var wg6 sync.WaitGroup

func test5()  {
	countNum:=0
	start:= time.Now().Unix()
	for i := 2; i < 100000; i++ {
		flag:=true
		for n := 2; n < i; n++ {
			if i%n == 0 {
				flag = false
				break

			}
		}
		if flag == true{
			//fmt.Println(i,":是素数")
			countNum++
		}
	}
	end:= time.Now().Unix()
	fmt.Println("总个数",countNum)
	fmt.Println("循环开始到结束的时间",end-start)
}
/*
	通过goroutine来实现,开启4个协程去实现
	i == 1 统计 1--30000
	i == 2 统计 30001--60000
	i == 3 统计 60001--90000
	i == 4 统计 90001--120000
   每次循环的范围	(num-1) * 30000+1   num*30000
*/

func test6(num int)  {
	countNum:=0
	start:= time.Now().Unix()
	for i := (num-1) * 30000+1; i < num*30000; i++ {
		if i>1 {
			flag:=true
			for n := 2; n < i; n++ {
				if i%n == 0 {
					flag = false
					break
				}
			}
			if flag == true{
				//fmt.Println(i,":是素数")
				countNum++
			}
		}
	}
	end:= time.Now().Unix()
	fmt.Println("总个数",countNum)
	fmt.Println("循环开始到结束的时间",end-start)
	wg6.Done()
}


func main()  {
	// test5()
	for i := 1; i <=4; i++ {
		wg6.Add(1)
		go test6(i)
	}
	wg6.Wait()  // 等待协程执行完成;
}