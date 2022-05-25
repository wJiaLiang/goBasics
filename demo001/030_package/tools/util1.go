package tools

import "fmt"
func init()  {
	fmt.Println("tools util1 中的 init 方法")
}

// 首字母大写是共有包，小写是 私有包
func Prtest(a string) string  {
	fmt.Println(a+"多个文件定义一个包")
	return  a
}
