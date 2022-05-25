package main

import (
	"fmt"
	"time"
)

// 需求 统计1--100 中的素数  除了1 和他本身 能够整除的数，不能被其他数整除的数;

func main() {
	start:= time.Now().Unix()
	for i := 2; i < 120000; i++ {
		var flag = true
		for n := 2; n < i; n++ {
			if i%n == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i, "是素数")
		}
	}
	end:= time.Now().Unix()
	fmt.Println(end,start)
	fmt.Println(end - start)  // 秒
}
