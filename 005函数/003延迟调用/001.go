package main

import "fmt"

func main()  {
/*
	一、延迟调用（defer）
 		1. 关键字 defer 用于注册延迟调用。
		2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
		3. 多个defer语句，按先进后出的方式执行。
		4. defer语句中的变量，在defer声明时就决定了。

	二、defer 用途：
		1. 关闭文件句柄
		2. 锁资源释放
	    3. 数据库连接释放

	defer 语句会将其后面跟随的语句进行延迟处理。在 defer 归属的函数即将返回 时，将延迟处理的语句按 defer 定义的逆序进行执行，
	先被 defer 的语句最后被 执行，最后被 defer 的语句，最先被执行

	由于 defer 语句延迟调用的特性，所以 defer 语句能非常方便的处理资源释放问题。比如： 资源清理、文件关闭、解锁及记录时间等。

*/

	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")

	//输出先后顺序 start end 3 2 1



}
