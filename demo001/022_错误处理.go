package main

import "fmt"

/*
当b 为0 时候会报错，程序抛出异常；
*/
func fn22(a int, b int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("错误信息", err) // 错误信息 runtime error: integer divide by zero
		}
	}()
	return a / b
}

func main() {
	fmt.Println(fn22(10, 0)) // 0
	fmt.Println("继续执行")

	/*
		nil是一个预声明的标识符
		指针、切片、映射、通道、函数和接口的零值则是 nil。
		预声明标识符nil没有默认类型
			预声明标识符true和false的默认类型均为内置类型bool。
			预声明标识符iota的默认类型为内置类型int。
		但是，预声明标识符nil没有一个默认类型，尽管它有很多潜在的可能类型
		nil不是一个关键字
		预声明标识符nil可以被更内层的同名标识符所遮挡。


	*/
	var m map[int]string
	var ptr *int
	var c chan int
	var sl []int
	var f func()
	var i interface{}
	fmt.Printf("%#v\n", m)
	fmt.Printf("%#v\n", ptr)
	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", sl)
	fmt.Printf("%#v\n", f)
	fmt.Printf("%#v\n", i)

	// map[int]string(nil)
	// (*int)(nil)
	// (chan int)(nil)
	// []int(nil)
	// (func())(nil)
	// <nil>
}
