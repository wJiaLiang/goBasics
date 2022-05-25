package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	// go 中浮点类型有两种 float32  	float64
//	 定义浮点类型
	var a float32 = 3.123
	fmt.Printf("值:%v---%f,类型%T \n",a,a,a)  // 值:3.123---3.123000,类型float32
	fmt.Println(unsafe.Sizeof(a))  // 4 占4个字节;


	var b float64 = 3.123
	fmt.Printf("值:%v---%f,类型%T \n",b,b,b)  // 值:3.123---3.123000,类型float64

	fmt.Println(unsafe.Sizeof(b))  // 8 占8个字节;

	//2. %f 输出float 类型   %.2f 输出数据的时候保留2位小数;
	var c float64 = 3.141592552
	fmt.Printf("%v---%.2f \n",c,c)
	fmt.Printf("%v---%.4f \n",c,c)  // 保留4位小数

	//3. 64位系统中，默认的浮点数是 float64

	//4. 科学计数法表示浮点类型
	var f2 = 3.14e2  // 表示 f2等于3.14*10的2次方;
	fmt.Printf("%v----%T \n",f2,f2)  // 314----float64

	var f3= 3.14e-2  // 表示 f2等于 3.14除以10的2次方;
	fmt.Printf("%v----%T \n",f3,f3)  // 0.0314----float64

	//	5. float 精度丢失的问题  可以用第三包解决
	var f4 float64 = 2323.3
	fmt.Println(f4*100)  // 232330.00000000003

	// 6.类型转换 int 转换成float 类型

	num1:= 12
	num2:= float64(12)
	fmt.Printf("类型%T---类型%T \n",num1,num2)

	// 7. float 类型转换成 int类型;
	var a1 float32 = 23.43
	a2:= int(a1)
	fmt.Printf("值%v----类型%T \n",a2,a2)   // 值23----类型int

	/**
		Golang 中 布尔类型
		bool 声明 布尔类型数据 只有两个值 true,false
		注意：
		1.布尔类型变量的默认值为false
		2.不予许将整型 强制转发为 布尔类型
		3. 布尔类型无法参与数值运算,也无法与其他类型进行转换
		错误示例:
		var a3 = 1
		if a3 {
			fmt.Println("a3")
		}

		var a4 = "this is String"
		if a4 {
			fmt.Println(a4)
		}

	*/







}