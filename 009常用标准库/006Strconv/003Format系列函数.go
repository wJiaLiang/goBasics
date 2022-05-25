package main

import (
	"fmt"
	"strconv"
)

func main()  {
/*
	Format系列函数实现了将给定类型数据格式化为string类型数据的功能。
	FormatBool()

	返回i的base进制的字符串表示  func FormatInt(i int64, base int) string
	FormatInt()

	FormatUint()


   FormatFloat() 函数将浮点数表示为字符串并返回。 func FormatFloat(f float64, fmt byte, prec, bitSize int) string

	bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。
   fmt表示格式：’f’（-ddd.dddd）、’b’（-ddddp±ddd，指数为二进制）、’e’（-d.dddde±dd，十进制指数）、’E’（-d.ddddE±dd，十进制指数）、’g’（指数很大时用’e’格式，否则’f’格式）、’G’（指数很大时用’E’格式，否则’f’格式）。
   prec控制精度（排除指数部分）：对’f’、’e’、’E’，它表示小数点后的数字个数；对’g’、’G’，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。

*/

	v1:=strconv.FormatBool(false)
	fmt.Printf("%T--%v\n",v1,v1)  		// string--false

	v2:= strconv.FormatInt(10,2)
	fmt.Printf("%T--%v\n",v2,v2)  		// string--1010

	v3:= strconv.FormatUint(17,16)
	fmt.Printf("%T--%v\n",v3,v3)  		// string--11

	v4:=strconv.FormatFloat(12.2334,'f',-1,64)
	fmt.Printf("%T--%v\n",v4,v4)   		// string--12.2334

}
