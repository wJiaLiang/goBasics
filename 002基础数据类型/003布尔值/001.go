package main

func main()  {
/*
	Go 语言中以 bool 类型进行声明布尔型数据，布尔型数据只有 true（真）和 false（假）两个 值。
	1. 布尔类型变量的默认值为 false。
	2. Go 语言中不允许将整型强制转换为布尔型.
	3. 布尔型无法参与数值运算，也无法与其他类型进行转换
	4. 布尔值可以和&&（AND）和||（OR）操作符结合，并且有短路行为：
		s != "" && s[0] == 'x'

*/



}

// 转换
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}