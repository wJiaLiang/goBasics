package main

import (
	"fmt"
	"strconv"
)

func main()  {
/*
	除了字符串、字符、字节之间的转换，字符串和数值之间的转换也比较常见 由strconv包提供这类转换功能


	fmt.Printf 函数的%b、%d、%o和%x等参数提供功能往往比strconv包的Format函数方便很多，
	fmt.Printf 函数的%b、%d、%o和%x等参数提供功能往往比strconv包的Format函数方便很多，特别是在需要包含有附加额外信息的时候：
*/
	
}
func init()  {
	// 将数字转换为字符串;
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Printf("类型:%T---值:%[1]v \n",y)   				// 类型:string---值:123
	fmt.Printf("类型:%T---值:%[1]v \n",strconv.Itoa(x))  //类型:string---值:123


	//FormatInt和FormatUint函数可以用不同的进制来格式化数字：
	a1:=6
	fmt.Println(strconv.FormatInt(int64(a1),2))  // 110
	fmt.Println(strconv.FormatInt(int64(a1),8))  // 6

	d1:= fmt.Sprintf("x=%b", x)
	fmt.Println(d1)  	 // "x=1111011"

	//	要将一个字符串解析为整数，可以使用strconv包的Atoi或ParseInt函数，还有用于解析无符号整数的ParseUint函数：
	x1, _ := strconv.Atoi("123")
	y1, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(x1)   // 123
	fmt.Println(y1)   // 123



}
