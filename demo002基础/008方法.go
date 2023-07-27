package main

import "fmt"

func main() {
	fmt.Println("方法")
	/*
		一、方法定义：
		func (recevier type) methodName(参数列表)(返回值列表){}
		参数和返回值可以省略

		type Test struct{}

		// 无参数、无返回值
		func (t Test) method0() {

		}

		// 多参数、无返回值
		func (t Test) method2(x, y int) {

		}



	*/

	/*
		二、普通函数与方法的区别

		1.对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然。

		2.对于方法（如struct的方法），接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以。




	*/

}
