package main

import "fmt"

/*
在 defer 归属的函数即将返回 时，将延迟处理的语句按 defer 定义的逆序进行执行，
也就是说，先被 defer 的语句最后被 执行，最后被 defer 的语句，最先被执行

*/

func dede()  {
	// defer 延迟返回   在命名返回值 和匿名返回函数中表现不一样
	fmt.Println("开始")
		defer fmt.Println(1)
		defer fmt.Println(2)
		defer fmt.Println(3)
	fmt.Println("结束")
	//开始
	//结束
	//3
	//2
	//1
}
// 匿名返回值
func f2() int {
	var a int
	defer func() {
		a++
	}()
	return a
}
// 命名返回值
func f3() (a int) {

	defer func() {
		a++
	}()
	return a    // 赋值给 a 后在执行 ++ 后再返回就返回的是1 了;
}
//**************************************************************************
// 错误异常处理 监听异常   panic  recover 必须要写在 defer函数中
func A()  {
	fmt.Println("A")
}
func B()  {
	defer func() {
		err:=recover()
		if err !=nil{
			fmt.Println(err)
		}
	}()
	panic("error!")
	fmt.Println("B")
}
func C()  {
	fmt.Println("C")
}



func main()  {
	dede()
	fmt.Println( f2() ) // 0
	fmt.Println( f3() ) // 1
	A()
	B()
	fmt.Println("结束")







}
