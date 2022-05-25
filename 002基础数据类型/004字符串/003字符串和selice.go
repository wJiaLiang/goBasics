package main

import (
	"fmt"
	"strings"
)

func main()  {
/*
   标准库中有四个包对字符串处理尤为重要：bytes、strings、strconv和unicode包。

	strings包提供了许多如字符串的查询、替换、比较、截断、拆分和合并等功能。

	bytes包也提供了很多类似功能的函数，但是针对和字符串有着相同结构的[]byte类型
	strconv包提供了布尔型、整型数、浮点数和对应字符串的相互转换，还提供了双引号转义相关的转换。
	unicode包提供了IsDigit、IsLetter、IsUpper和IsLower等类似功能，它们用于给字符分类。

	一、字符串的常用操作
		len(str) 				求长度
		+或 fmt.Sprintf 		拼接字符串
		strings.Split 			分割
		strings.contains 		判断是否包含
		strings.HasPrefix,strings.HasSuffix 	前缀/后缀判断
		strings.Index(),strings.LastIndex() 	子串出现的位置
		strings.Join(a[]string, sep string) 	join 操作



*/



c1()


}

func c1()  {
	// 求长度
	var str = "this is str"
	fmt.Println(len(str))  //11


	//拼接字符串
	var str1 = "你好"
	var str2 = "golang"
	fmt.Println(str1 + str2)

	var str3 = fmt.Sprintf("%v %v", str1, str2)
	fmt.Println(str3)

	//	strings.Split 分割字符串
	var str4 = "123-456-789"
	var arr = strings.Split(str4, "-")
	fmt.Println(arr)   //[123 456 789]

	//	判断是否包含
	var str5 = "this is golang"
	var flag = strings.Contains(str5, "golang")
	fmt.Println(flag)

	// 判断首字符尾字母是否包含指定字符
	var str6 = "this is golang"
	var flag6 = strings.HasPrefix(str6, "this")
	fmt.Println(flag6) //true

	var str7 = "this is golang"
	var flag7 = strings.HasSuffix(str7, "go")
	fmt.Println(flag7) //false

	//判断字符串出现的位置 返回下标的位置从0开始
	var str8 = "this is golang"
	var index8 = strings.Index(str8, "is")
	fmt.Println(index8) //2   从前往后找

	var str9 = "this is golang"
	var index9 = strings.LastIndex(str9,"is")
	fmt.Println(index9) //5   从后往前找

	//Join 拼接字符串
	var str10 = "123-456-789"
	var arr1 = strings.Split(str10, "-")
	var str11 = strings.Join(arr1, "*")
	fmt.Println(str11)  // 123*456*789

}