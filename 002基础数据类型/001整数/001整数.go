package main

import (
	"fmt"
	"unsafe"
)

/*

Go语言将数据类型分为四类：基础类型、复合类型、引用类型 和 接口类型:

基础类型包括: 数字、字符串和布尔型。
复合数据类型: 数组,结构体(通过组合简单类型)
引用类型包括: 指针,切片,字典，函数，通道

*/


func main()  {
	fmt.Println("")

	a01()

}


func a01()  {
/*
	数值类型包括几种不同大小的整数、浮点数和复数
	每种数值类型都决定了对应的大小范围和是否支持正负符号。
	int8、int16、int32、int64 四种有符号整数类型
	uint8、uint16、uint32、uint64 四种无符号整数类型。

	int8 值域{-128至127}     -2^{n-1} 到 2^{n-1}-1      有符号
	uint8 值域{0至255} 		0 到 2^n-1				   无符号
   字节也叫 Byte，是计算机数据的基本存储单位。8bit(位)=1Byte(字节) 1024Byte(字节)=1KB 1024KB=1MB 1024MB=1GB 1024GB=1TB

	一、算术运算、逻辑运算和比较运算的二元运算符，它们按照优先级递减的顺序排列：

		1、*      /      %      <<       >>     &       &^
		2、+      -      |      ^
		3、==     !=     <      <=       >      >=
		4、&&
		5、||
   		算术运算符+、-、*和/可以适用于整数、浮点数和复数，但是取模运算符%仅用于整数间的运算


	二、两个相同的整数类型可以使用下面的二元比较运算符进行比较；比较表达式的结果是布尔类型。
		布尔型、数字类型和字符串等基本类型都是可比较的，也就是说两个相同类型的值可以用==和!=进行比较
		==    等于
	   !=    不等于
	   <     小于
	   <=    小于等于
	   >     大于
	   >=    大于等于


   三、bit位操作运算符，前面4个操作运算符并不区分是有符号还是无符号数：
	   &      位运算 AND
	   |      位运算 OR
	   ^      位运算 XOR
	   &^     位清空 (AND NOT)
	   <<     左移
	   >>     右移
		x<<n和x>>n移位运算中，决定了移位操作的bit数部分必须是无符号数；被操作的x可以是有符号数或无符号数。
		一个x<<n左移运算等价于乘以 2^n，一个x>>n右移运算等价于除以 2^n。

	四、对于每种类型T，如果转换允许的话，类型转换操作T(x)将x转换为T类型。
		浮点数到整数的转换将丢失任何小数部分，然后向数轴零方向截断
		任何大小的整数字面值都可以用以0开始的八进制格式书写，或用以0x或0X开头的十六进制格式书写，


	五、fmt的两个使用技巧：
	通常Printf格式化字符串包含多个%参数时将会包含对应相同数量的额外操作数
	但是%之后的[1]副词告诉Printf函数再次使用第一个操作数
	%后的#副词告诉Printf在用%o、%x或%X输出时生成0、0x或0X前缀。

		%d     		十进制整数
		%x,%o,%b  	十六进制，八进制，二进制整数
		%f,%g,%e    浮点数
		%t    		布尔型（true 或 false）
		%c			字符（Unicode）码点
		%s			字符串
		%q          带双引号的字符串"abc"或带单引号的字符'c'
		%v          内置格式的任何值
		%T			任何值的类型
		%%			百分号本身（无操作数）


	六、unsafe.Sizeof(n)  可以返回 n 变量 占用的字节数

*/
	//var a uint8 = -23
	//fmt.Println(a)   //报错,溢出 constant -23 overflows uint8


	b  := "kgwei"
	var c int8 =12
	var d = 12
	fmt.Printf("%#x\n",b)  		// 0x6b67776569
	fmt.Printf("%x \n",b)   		// 6b67776569
	fmt.Printf("%q \n",b)   		// "kgwei"
	fmt.Printf("%d \n",c)      	// 12
	fmt.Println( unsafe.Sizeof(c) ) 	// 1   表示一个字节
	fmt.Println( unsafe.Sizeof(d) ) 	// 8   表示一个字节


}