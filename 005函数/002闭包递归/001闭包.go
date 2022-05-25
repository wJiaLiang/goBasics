package main

import "fmt"

// 外部引用函数参数局部变量  闭包函数
func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

// 返回2个函数类型的 返回值
func test01(base int) (func(int) int, func(int) int) {
	// 定义2个函数，并返回
	// 相加
	add := func(i int) int {
		base += i
		return base
	}
	// 相减
	sub := func(i int) int {
		base -= i
		return base
	}
	// 返回
	return add, sub
}


func main()  {
/*
	闭包是由函数及其相关引用环境组合而成的实体(即：闭包=函数+引用环境)
*/
	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))   //11  13

	// 此时tmp1 和 tmp2不是一个实体了   add执行一次,里面的环境每次都不一样;
	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))  //101 103

/*
	*****************************************************************************
   返回2个闭包
*/

	f1, f2 := test01(10)

	// base一直是没有被回收,依然保存在内存中;
	fmt.Println(f1(1), f2(2))   // 11   9

	fmt.Println(f1(3), f2(4))  // 12    8



}
