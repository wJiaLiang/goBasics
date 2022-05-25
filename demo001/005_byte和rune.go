package main

import (
	"fmt"
)

func main()  {
//	组成每个字符串的元素叫做'字符',可以通过遍历字符串获得字符，字符用单引号包裹起来

	 // 1. Go 中定义字符 字符属于int类型

	var a = 'a'
	fmt.Printf("%v---%T \n",a,a)  // 97---int32   97是ascll码

	//	2. 原样输入字符
	fmt.Printf("%c---%T \n",a,a)  // a---int32

	/**
		go 语言中 字符有两种
		1. uint8 类型，或者叫byte 型， 代表了 ASCll 码的一个字符
		2. rune 类型，代表了一个 UTF-8 字符

	*/
	//	3. 定义一个字符串 输出里面的字符
	 var str = "this"
	 fmt.Printf("值%v----类型%T \n",str[2],str[2])  // 值105----类型uint8

	 // 4. 一个汉字占3个字节(utf-8 编码) ，一个字面占一个字节

	 var str1 = "this"
	 //unsafe.Sizeof()  没有办法查看 string 类型的数据所占用的存储空间
	 fmt.Println(len(str1))  // 4

	 var str2 = "你好Go"
	 fmt.Println(len(str2))  // 8

	// 5. 定义一个字符 字符的值是汉字 (go中汉字使用的是utf-8编码，编码后的值就是int 类型)
	var a1 = '中'
	fmt.Printf("值%v----类型%T \n",a1,a1)  // 值20013----类型int32  ;Unicode编码后的十进制;

	//	6. 通过循环输出字符串里面的字符
	var a2 = "你好 Golang"
	//for i:=0;i<len(a2);i++{   // 循环处理的是byte 类型
	//	fmt.Printf("%v(%c)\n--",a2[i],a2[i])
	//}

	for _, v:=range a2 {    // 循环处理的是rune 类型
		// 20320(你)---22909(好)---32( )---71(G)---111(o)---108(l)---97(a)---110(n)---103(g)---
		fmt.Printf("%v(%c)---",v,v)
	}

	//7 修改字符传
	s1:="big"    // 字符串中有汉字就不行无法修改，用 rune类型才行
	byteStr := []byte(s1)
	byteStr[0] = 'p'
	fmt.Println(string(byteStr)+"\n")  //pig

	s2:="哈哈里 big"
	runeStr := []rune(s2)
	runeStr[0] = 'p'
	fmt.Println(string(runeStr)+"\n")

}
