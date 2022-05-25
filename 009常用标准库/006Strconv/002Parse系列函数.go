package main

import (
	"fmt"
	"strconv"
)

func main()  {

/*
	Parse类函数用于转换字符串为给定类型的值：ParseBool()、ParseFloat()、ParseInt()、ParseUint()。
    ParseInt:
    func ParseInt(s string, base int, bitSize int) (i int64, err error)
	base指定进制（2到36），如果base为0，则会从字符串前置判断，”0x”是16进制，”0”是8进制，否则是10进制；
	bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
	返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。

	ParseFloat():
	解析一个表示浮点数的字符串并返回其值。
	func ParseFloat(s string, bitSize int) (f float64, err error)

	bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；


*/
	test3()

}

// 返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。
func test3()  {

	v,err:= strconv.ParseBool("t")
	fmt.Println(v,err)   // true <nil>

	// 是将字符串转换为数字的函数,功能灰常之强大,看的我口水直流.
	v1,err1 := strconv.ParseInt("101010",2,64)
	fmt.Println(v1,err1)    // 42 <nil>


	v2,err2 := strconv.ParseFloat("1.23",8)
	fmt.Println(v2,err2)    //1.23



}







