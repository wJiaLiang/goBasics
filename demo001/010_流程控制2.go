package main

import (
	"fmt"
)

func main(){
	/*
		break 语句
		用于跳出循环,循环后面继续执行;
		break 在 switch 开关语句中执行 一条 case 后跳出语句的作用
		在多重循环中 可以用标号 label 标出想 break 的循环
	*/
	for i:=0;i<10;i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("继续")

	// switch 中当做开关语句;
	var extname = ".html"
	switch extname {
	case ".html":
		fmt.Println("html")
		break

	case ".js":
		fmt.Println("html")
		break

	case "php":
		fmt.Println("php")
		break
	default:
		fmt.Println("没有匹配")
	}

	//	 直接跳出多重循环
	label:
		for i:=0;i<10;i++ {
			for j:=0;j<10;j++ {
				if j==3 {
					break label
				}
				fmt.Println("ts",j )
			}
		}


	// continue 语句可以结束当前循环，开始洗一次的循环，仅限在 for循环内使用
	for i:=0;i<10;i++ {
		if i == 2 || i == 4{
			continue
		}
		fmt.Println(i)  // 跳过 2 和 4;
	}


	//	 goto 通过标签 进行 代码间的无条件跳转 goto 可以快速跳出循环 跳转到指定的位置，继续执行;
	var n =20
	if n>10{
		fmt.Println("dddd")
		goto label2
	}

	fmt.Println("aaaa")
	fmt.Println("bbbb")
	label2:
	fmt.Println("cccc")


}
