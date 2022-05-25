package main

import (
	"030_package/calc"
	"030_package/tools"
	"fmt"
)
// init 方法自动调用; a 中 引用 b 包 则 b 包中的init 先执行;
func init()  {
	fmt.Println("main包中的init方法")
}

func main()  {
	a:=calc.Add(3,4)
	fmt.Println(a)

	tools.Prtest("测试")
	// tools.pr2("测试二")  //开头字母小写,表示私有模块，无法在文件外调用;
}
