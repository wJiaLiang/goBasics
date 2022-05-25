package main

import (
	"fmt"
	"strconv"
)

func main()  {

	//1. 数值 类型之间转换

	var a int8 = 20
	var b int16 = 30

	// fmt.Println( a+b )  //  数据类型不一致无法相加;
	fmt.Println( int16(a)+b )
	//	注意:转换的时候建议从低位转换成高位，高位转换成低位的时候，如果转换不成功就会溢出，报错;

	// 2. 浮点型和浮点型之间的转换 float32()  float64()


	// 3. 整型和浮点型之间的转换  int32()  int64()

	// 4. 其他类型转换成string 类型
		//1.使用 fmt.Sprintf ,在Sprintf 中使用需要注意转换的格式  int为%d   float为%f  bool为%t  byte为%c
			var i int = 20
			//var f float32 = 12.456
			//var t boll = true
			//var b byte = 'a'
			str1:= fmt.Sprintf("%d",i)
			fmt.Sprintf("%v--类型:%T \n",str1,str1)


		//2. 使用 strconv 包 里面的几种转换方法进行转换

			//1. 把其他类型值转换成 string 类型;
			var i1 int = 34
			/*
				参数1:int64 的数值
				参数2:int 类型的进制
			*/
			str2:= strconv.FormatInt(int64(i1),10)  // 表示包 i1 转换成 10进制
			fmt.Printf("%v---%T \n",str2,str2)  // 34---string

			// 2. string 类型转换成整型
			str3:= "12345"
			str33,_ := strconv.ParseInt(str3,10,64)
			fmt.Printf("%v----%T \n",str33,str33) // 12345----int64

			// 3.浮点型装换成 string
			/*
				参数1 要转换的值，
				参数2 格式化类型 'f' (-ddd.ddd)   'e' 十进制指数
				参数3 保留的小数点 -1 不对小数点格式化
				参数4 格式化的类型 传入 64 32
			*/
			var str4 float32 = 23.3232
			str44 := strconv.FormatFloat(float64(str4),'f',2,64)
			fmt.Printf("%v----%T \n",str44,str44) // 23.32----string


			// 4.字符转换
			str5:='a'
			str55 := strconv.FormatUint(uint64(str5),10)
			fmt.Printf("%v----%T \n",str55,str55)

	//  字符串转换成float 类型






}
