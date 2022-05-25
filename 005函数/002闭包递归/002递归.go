package main

import "fmt"

// 递归计算 阶乘
func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

//
func fibonaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}


func main()  {
/*
	递归，就是在运行的过程中调用自己。 一个函数调用自己，就叫做递归函数。
    构成递归需具备的条件：
		1、子问题须与原始问题为同样的事，且更为简单。
		2、不能无限制地调用本身，须有个出口，化简为非递归状况处理。
*/

	var i int = 7
	fmt.Printf("Factorial of %d is %d\n", i, factorial(i))

/*
   斐波那契数列(Fibonacci)
   这个数列从第3项开始，每一项都等于前两项之和。
*/
	var i2 int
	for i2 = 0; i2 < 10; i2++ {
		fmt.Printf("%d\n", fibonaci(i2))
	}




}
