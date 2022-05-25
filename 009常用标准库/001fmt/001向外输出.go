package main

import (
	"fmt"
	"os"
)

func main()  {
	/*
		标准库fmt提供了以下几种输出相关函数。

	*/
	test1()
	//test2()
	test3()
	test4()
}

/*
	Print
	Print系列函数会将内容输出到系统的标准输出，区别在于Print函数直接输出内容
	Printf函数支持格式化输出字符串，
	Println函数会在输出内容的结尾添加一个换行符。

*/
func test1()  {
	fmt.Print("在终端打印该信息。")
	name := "枯藤"
	fmt.Printf("我是：%s\n", name)
	fmt.Println("在终端打印单独一行显示")

}

/*
	Fprint
	Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。
	注意，只要满足io.Writer接口的类型都支持写入。


*/
func test2()  {
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "枯藤"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)

}

/*
	Sprint
	Sprint系列函数会把传入的数据生成并返回一个字符串。
*/
func test3()  {
	s1 := fmt.Sprint("枯藤\n")
	name := "枯藤"
	age := 18
	s2 := fmt.Sprintf("name:%sage:%d \n", name, age)
	s3 := fmt.Sprintln("枯藤")
	fmt.Println(s1, s2, s3)
}

func test4()  {
	// Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。
	err := fmt.Errorf("这是一个错误")
	fmt.Println(err)

}




