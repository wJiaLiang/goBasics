package main

import "fmt"

/*
1、另一个创建变量的方法是调用内建的new函数。创建T类型的匿名变量，初始化为T类型的零值,返回值为指针类型 *T
	p := new(int)   // p, *int 类型, 指向匿名的 int 变量
2、new 创建的变量和普通变量声明语句没有什么区别;
3、每次调用new函数都是返回一个新的变量的地址;

*/


func main()  {
	p := new(int)   	// p, *int 类型, 指向匿名的 int 变量
	fmt.Println(p)  	// 返回这个变量的指针地址 0xc00000a0c0
	fmt.Println(*p) 	// 0


}
