package main

import (
	"fmt"
	"math"
)

func main()  {

	/*

	一、Go 语言支持两种浮点型数：float32 和 float64。
		常量math.MaxFloat32表示float32能表示的最大数值
		常量math.MaxFloat64表示float64能表示的最大数值

	二、Golang 中 float 精度丢失问题
		几乎所有的编程语言都有精度丢失这个问题
		在一定条件下，二进制小数和十进制小数互转可能有精度丢失。
		 %f 输出float 类型   %.2f 输出数据的时候保留2位小数;

		用第三方包解决https://github.com/shopspring/decimal

	三、Golang 科学计数法表示浮点类型



	*/
	a1()


}

func a1()  {
	fmt.Println(math.MaxFloat32) // 3.4028234663852886e+38
	fmt.Println(math.MaxFloat64) // 1.7976931348623157e+308

	n:=2.345E-12   		// 2.345 除以 10的12次方
	n1:=2.345E12   		// 2.345 乘以 10的12次方
	fmt.Println(n,n1)   // 2.345e-12    2.345e+12
	fmt.Printf("%.2f \n",math.Pi) //3.14  保留两位小数

	//默认 是 float64
	num1:=1.1
	fmt.Printf("值:%v----类型:%[1]T",num1)   // 值:1.1----类型:float64




}