package main

import "fmt"

func funcA() {
	fmt.Println("func A")
}
func funcB() {
	// 这样就可以抛出错误,不会中断执行;
	defer func() {
		err := recover()
		if(err != nil){
			fmt.Println("err:",err) //err: panic in B
		}
	}()
	panic("panic in B")

}
func funcC() {
	fmt.Println("func C")
}

/*
1、Golang 没有结构化异常，使用 panic 抛出错误，recover 捕获错误。
异常的使用场景简单描述：Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。

2、panic 可以在任何地方引发，但 recover 只有在 defer 调用的函数中有效

3、程序运行期间 funcB 中引发了 panic 导致程序崩溃，异常退出了。这个时候我们就可以通过 recover 将程序恢复回来，继续往后执行。


*/


func main()  {

	funcA()
	funcB()
	funcC()

//	func A
//	err: panic in B
//	func C

}
