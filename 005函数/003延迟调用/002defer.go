package main

import "fmt"

func f1()  {
	fmt.Println("start")
		defer func() {
			fmt.Println("aaa")
			fmt.Println("bbb")
		}()   // 用deter 执行匿名函数必须执行，加一个括弧调用
	fmt.Println("end")
}

// 返回匿名变量
func f2() int {
	var a int
	defer func() {
		a++
	}()
	return a
}
// 返回一个命名变量
func f3() (a int)  {
	defer func() {
		a++
	}()
	return a
}
// 返回时 x 赋值给 y,再返回出去
func f4() (y int)  {
	x:=5
	defer func() {
		x++
	}()
	return x
}


func main()  {
	/*
		在 Go 语言的函数中 return 语句在底层并不是原子操作，
		它分为给返回值 赋值和 RET 指令 两步。
		而 defer 语句执行的时机就在返回值赋值操作后，RET 指令执行前。
	*/


	f1()
	//start
	//end
	//aaa
	//bbb

 	fmt.Println(f2())	// 0

	fmt.Println(f3())   // 1

	fmt.Println(f4())   // 5

}
