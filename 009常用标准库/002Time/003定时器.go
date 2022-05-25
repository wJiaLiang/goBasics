package main

import (
	"fmt"
	"time"
)

/*
	使用time.Tick(时间间隔)来设置定时器，定时器的本质上是一个通道（channel）。

*/

func tickDemo()  {
	ticker:= time.Tick(time.Second*1)
	for i:=range ticker {
		fmt.Printf("%T--",i)
		fmt.Println(i)
	}
}

func main()  {
	tickDemo()


}


