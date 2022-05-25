package main

import (
	"fmt"
	"strings"
)

func main()  {
	//	1.定义一个字符串类型
		var a = "hello go string"
		var a1 string = "hello go string2"
		a2:= "hello go string3"
		fmt.Println(a,a1,a2)

	//  2.字符串转义符
	/*
		\r  回车符 返回行首
		\n  换行符，直接跳到下一行的同列位置
		\t  制表符
		\'  单引号
		\"  双引号
		\\  反斜杠
	*/
	//	3.多行字符串

	a3:= `this is string 多行字符串 
		多行字符串
			string
	`
	fmt.Println(a3)

	// 4. len(str) 求长度
	a4:= "长度"
	fmt.Println(len(a4))  // 6,因为一个汉字占3个字节

	//	5. + 或者 fmt.Sprintf 拼接字符串
	var b1 string = "hello"
	var b2 string = "go"
	var b3 = b1+b2
	fmt.Println(b3) // hellogo
	b4:= fmt.Sprintf("%v--%v",b1,b2)

	fmt.Println(b4)  // hello--go

	//	6. strings.Split 分隔字符串， strings 需要引入 strings 包
	var str1 = "123-456-789"
	arr:= strings.Split(str1,"-")
	fmt.Println(arr) // [123 456 789]

	//	7. strings.Jion()  把切片连接成字符串
	str2:= strings.Join(arr,"*")
	fmt.Println(str2)

	//	8.strings.Contains 判断是否包含 返回 true 表示包含，返回 false 表示 不包含

	 str3:="this is str"
	 str4:="this"
	 fmt.Println(strings.Contains(str3,str4) )  // true

	 // 9. strings.HasPrefix()  前缀是否有指定的字符包含    strings.HasSuffix()  后缀是否有指定的字符包含

	 strings.HasPrefix(str3,str4)
	 strings.HasSuffix(str3,str4)

	 // 10.  strings.Index()  strings.LastIndex()重后往前查找   子字符串查找   返回下标的位置从0开始  在 str5 中 找 str6 找不到返回-1
	 str5:= "this is str5"
	 str6:= "is"
	 fmt.Println( strings.LastIndex(str5,str6) )

}
