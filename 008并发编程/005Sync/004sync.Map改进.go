package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m2 = sync.Map{}


func main()  {
/*
	像这种场景下就需要为map加锁来保证并发的安全性了，Go语言的sync包中提供了一个开箱即用的并发安全版map–sync.Map。
       开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用

   	同时 sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法。

*/

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key:=strconv.Itoa(n)

			m2.Store(key, n)
			value, _ := m2.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)

			wg.Done()
		}(i)
	}
	wg.Wait()

}
