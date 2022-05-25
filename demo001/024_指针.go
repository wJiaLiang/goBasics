package main

import "fmt"

/**
指针也是一个变量，单他是一种特殊的变量，他存的数据不是一个普通值，而是一个变量的内存地址
所有的变量都有内存地址
*/

func fn1(x int)  {
	x =20
}
func fn2(x *int)  {
	*x = 50
}
//**************************************************************
/*
	make  和 new
	都是用来分配内存的，make 只用于 slice map, channel 的初始化 ，返回的还是这三个引用类型本身
	new 用于 类型的内存分配，并且内存对应的值为类型的零值，返回的是指向类型的指针
	应用类型 没有办法直接声明后 就使用， 必须要先分配内存
*/

func test()  {
	var userinfo = make(map[string]string)
	userinfo["username"] = "张3"
	fmt.Println(userinfo)
}
func newH()  {
	// 使用new 函数得到是一个指针类型，并且该指针类型对应的值 为该类型的零值
	// a 是一个指针变量，类型是 *int 的 指针类型  指针变量对应的值是0
	var a = new(int)
	fmt.Printf("%v--%T",a,a,)
}


func main()  {
	var a int =10
	fmt.Printf("值：%v  类型：%T 内存地址：%p \n",a,a,&a) // 值：10  类型：int 地址：0xc00000a0f0


	 p := &a  // p 是指针变量 p的类型 *int (指针类型)
	 fmt.Printf("值：%v  类型：%T \n",p,p) // 值：0xc00000a0f0  类型：*int
	// *p 表示取出p这个变量对应的内存地址的值;
	 fmt.Println(p) // a的地址   0xc000106018
	 fmt.Println(*p) // 打印变量 a 内存地址对应的值    10

	 *p = 30   // 改变p 这个变量对应的内存地址的值
	 fmt.Println(a) //30 

	var s = 5
	fn1(s)
	fmt.Println(s)  //5
	fn2(&s)
	fmt.Println(s)  //50

	test()
	newH()

}

