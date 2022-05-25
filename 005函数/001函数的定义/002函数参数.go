package main

import (
	"fmt"
)


func swap(x, y *int) {
	/* 定义相互交换值的函数 */
	var temp int

	temp = *x 		/* 保存 x 的值 */
	*x = *y   		/* 将 y 值赋给 x */
	*y = temp 		/* 将 temp 值赋给 y*/
}

func myfn(argssss...int){ fmt.Println(argssss) } //[1,2]
func myfn1(a int,args...int)  {  }
func myfn2(a int,b int,args...int)  {  }

func main()  {
/*
	一、函数参数
	1、函数定义时有参数，该变量可称为函数的形参。形参就像定义在函数体内的局部变量。
	2、值传递：指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
	3、引用传递：是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。


	注意1：无论是值传递，还是引用传递，传递给函数的都是变量的副本，
	注意2：map、slice、chan、指针、interface默认以引用的方式传递。


*/
	var a,b int = 1,2
	swap(&a,&b)
	fmt.Println(a,b)  // 2,1  同传递指针交换两个变量的值;

/*
	二、
	可变参数：
		不定参数传值 就是函数的参数不是固定的，后面的类型是固定的。（可变参数）
   		Golang 可变参数本质上就是 slice。只能有一个，且必须是最后一个。

   	任意类型的不定参数：
		就是函数的参数和每个参数的类型都不是固定的。
		用interface{}传递任意类型数据是Go语言的惯例用法，而且interface{}是类型安全的。
		func myfunc(args ...interface{}) {  }
*/
	var a1,b1 int = 1,2
	myfn(a1,b1)

/*
	三、使用 slice 对象做变参时，必须展开。（slice...）
*/
	a2 := []int{1, 2, 3}
	myfn(a2...)

}


