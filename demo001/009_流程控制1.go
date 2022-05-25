package main

import "fmt"

func main()  {

	var str = "你好 golang"

	for key,v := range str{
		//fmt.Println(key,v)  // 默认打印的是 ascll 码
		fmt.Printf("key=%v v=%c \n",key,v)
	}

	var arr = []string{"php","java","js","obj语言"}
	for _,val := range arr {
		fmt.Println(val)
	}

	//	 switch case
	// 使用 switch 语句可方便的对大量的值进行 条件判断;

	var name = "javas"
	switch name {
	case "html":
		fmt.Println("html哦")
		break
	case "php":
		fmt.Println("php")
		break
	case "java":
		fmt.Println("java")
		break

	default:
		fmt.Println("没有匹配")
		break
	}

	//switch 另一种写法 变量可以写在里面

	switch name = "javas";name {
	case "html":
		fmt.Println("html哦")
		break
	case "php":
		fmt.Println("php")
		break
	case "java":
		fmt.Println("java")
		break

	default:
		fmt.Println("没有匹配")
		break
	}

	//	 判断多个值 一个分支可以有多个值; 也可以是条件判断
	var n = 5
	switch n {
	case 1,3,5,7,9:
		fmt.Println("奇数")
		break
	case 2,4,6,8:
		fmt.Println("偶数")
		break
	}

	var age = 23
	switch {
	case age>20 :
		fmt.Println("sss")

	case age>20 && age<80 :
		fmt.Println("fff")

	}

	// switch 的穿透 fallthrought
	// 可以满足条件下的 下一个 case


}
