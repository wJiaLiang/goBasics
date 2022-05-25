package main

import "fmt"

func test1() (int, int) {
	return 1, 2
}
func add1(x, y int) int {
	return x + y
}
func test2(x,y int) (cc int) {
	cc =x+y
	return
}

// 命名返回参数可被同名局部变量遮蔽，此时需要显式返回。
func add2(x, y int) (z int) {
	{ // 不能在一个级别，引发 "z redeclared in this block" 错误。
		var z = x + y
		// return   // Error: z is shadowed during return
		return z // 必须显式返回。
	}
}

// 命名返回参数允许 defer 延迟调用通过闭包读取和修改。
// defer 延迟执行;
func add3(x, y int) (z int) {
	defer func() {
		z += 100
	}()
	z = x + y
	return
}

func main()  {
	/*
		1、"_"标识符，用来忽略函数的某个返回值
		2、Golang返回值不能用容器对象接收多返回值。只能用多个变量，或 "_" 忽略。
		3、没有参数的 return 语句返回各个返回变量的当前值。这种用法被称作“裸”返回。
		4、多返回值可直接作为其他函数调用实参。
		5、命名返回参数可看做与形参类似的局部变量，最后由 return 隐式返回。
		6、命名返回参数可被同名局部变量遮蔽，此时需要显式返回。

		7、命名返回参数允许 defer 延迟调用通过闭包读取和修改。
	*/

	add1(test1())
	var aa = test2(3,6)
	fmt.Println(aa)

	fmt.Println(add3(1,2))  //103

}
