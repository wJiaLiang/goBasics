package main

import "fmt"

func main()  {

}

func init()  {
	/*
		1、组成每个字符串的元素叫做“字符”，可以通过遍历字符串元素获得字符。 字符用单引号（’） 包裹起来，如：
		2、字节（byte）：是计算机中 数据处理 的基本单位，习惯上用大写 B 来表示,1B（byte,字节） = 8bit（位）

		3、字符：是指计算机中使用的字母、数字、字和符号
		4、Go 语言的字符有以下两种：
			(1) uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
			(2) rune 类型，代表一个 UTF-8 字符。
			(3) 当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型实际是一个 int32。

			Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可用 byte 型进行默认字符串处理，

		5、修改字符串
		要修改字符串，需要先将其转换成[]rune 或[]byte，完成后再转换为 string。无论哪种转换， 都会重新分配内存，并复制字节数组。

	*/

	a:='a'
	b:='b'
	c:='魏'
	d:='佳'
	e:='亮'
	//当我们直接输出 byte（字符）的时候输出的是这个字符对应的码值
	fmt.Println(a)  //97
	fmt.Println(b)  //98
	fmt.Println(c)  //39759
	fmt.Println(d)  //20339
	fmt.Println(e)  //20142

	// /如果我们要输出这个字符，需要格式化输出
	fmt.Printf("%c",c)

	fmt.Println("-----------分隔符---------分隔符-------------分隔符")

	//Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可用 byte 型进行默认字符串处理，
	s := "hello 张三"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%[1]c)",s[i])  // 104(h)101(e)108(l)108(l)111(o)32( )229(å)188(¼)160()228(ä)184(¸)137()
	}
	fmt.Println()
	for _,r:= range s{
		fmt.Printf("%v(%[1]c)",r)     //104(h)101(e)108(l)108(l)111(o)32( )24352(张)19977(三)
	}
	fmt.Println()

	/*
	字符串底层是一个 byte 数组，所以可以和[]byte 类型相互转换。
	字符串是不能修改的 字符 串是由 byte 字节组成，所以字符串的长度是 byte 字节的长度。
	rune 类型用来表示 utf8 字 符，一个 rune 字符由一个或多个 byte 组成。

	rune 类型实际是一个 int32
	*/
	cc3 := "美"
	cc4 := '美'
	fmt.Printf("c3类型:%T---c4类型:%T \n",cc3,cc4) // c3类型:string---c4类型:int32


	changeString()

}

// 改变字符串;
//要修改字符串，需要先将其转换成[]rune 或[]byte，完成后再转换为 string。无论哪种转换， 都会重新分配内存，并复制字节数组。
func changeString()  {
	s1:="big"
	byteS1 := []byte(s1)

	fmt.Println(byteS1)  // [98 105 103]

	byteS1[0] = 'p'
	fmt.Println(string(byteS1)) //pig

	s2 := "红萝卜"
	runes2 := []rune(s2)
	runes2[0] = '白'
	fmt.Println(string(runes2)) //白萝卜


}