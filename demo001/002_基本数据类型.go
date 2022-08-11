package main

import (
	"fmt"
	"unsafe"
)

/*
整型分为以下两个大类：
有符号整形按长度分为：int8、int16、int32、int64
对应的无符号整型：uint8、uint16、uint32、uint64
	bool
	byte
	rune
	int, uint
	int8, uint8
	int16, uint16
	int32, uint32
	int64, uint64
	float32
	float64
	complex64
	complex128
	uintptr
	array
	struct
	string
	slice
	map
	channel
	interface
	function

*/

func main()  {
	// 1. int 类型
	var num = 20
	fmt.Printf("num=%v 类型:%T\n",num,num)  // num=20 类型:int

	//	2. int8 范围  (-128 到 127)
	var num1 int8 = 130   // 溢出，报错
	fmt.Printf("num1=%v 类型:%T\n",num1,num1)

	// unsafe.Sizeof 查看不同长度的整形 在内存里面的 存储空间
	var a int8 = 30
	fmt.Printf("a=%v 类型:%T\n",a,a)
  	fmt.Println( unsafe.Sizeof(a) ) // 1 ;表示一个字节

  	//3. uint8 范围 (0-255) 值超出范围报错
  	var a1 uint8 = 23
	fmt.Printf("a=%v 类型:%T\n",a1,a1)
	fmt.Println( unsafe.Sizeof(a1) )

	//4. int16 int32 int64  同理

	//5. int 不同长度直接的转换 不同类型数据，要进行类型转换后才能进行运算 （高位向低位装换的时候注意，能否放下）
	var a2 int32 = 10
	var a3 int64 = 230
	fmt.Println(int64(a2)+a3)


	//6. 数字字面量语法，%d 表示十进制输出  %b 表示二进制输出 %o 表示八进制输出 %x 表示十六进制输出
	num2:=30   // 默认int 类型 ,如果你的计算机是 32位 就是 int32 ，64 就是 int64;
	fmt.Printf("num=%v 类型:%T\n",num2,num2)
	fmt.Print(unsafe.Sizeof(num2),"\n") // 8 ,表示你的计算机是 64位 int 就是 int64 占用8个字节

	fmt.Printf("num=%b 类型:%T\n",num2,num2) // 11110




}
