package tools

import "fmt"

func init()  {
	fmt.Println("tools util2 中的 init 方法")
}

func pr2(a string) string {
	fmt.Println(a+"多个文件定义一个包2")
	return a
}
