package main

import (
	"errors"
	"fmt"
)

/*
defer 、panic、recover 抛出异常
	读取到文件失败,readFile 中 抛出错误，fn3中 deter函数捕获到错误

*/
func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	}
	return errors.New("读取文件错误")
}

func fn3() {
	defer func() {
		err := recover()
		if err != nil {
		fmt.Println("抛出异常给管理员发送邮件")
		}
	}()
	var err = readFile("xxx.go")
	if err != nil {
		panic(err)
	}
	fmt.Println("没有错误继续执行。。。")
}

func main()  {

	fn3()
	fmt.Println("程序继续往下执行,不会中断")


}
