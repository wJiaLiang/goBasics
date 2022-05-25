package main

import "fmt"

func main()  {
	/*
		一、定义
			1、空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
			2、空接口类型的变量可以存储任意类型的变量。

		二、空接口的应用
			1、空接口作为函数的参数，  使用空接口实现可以接收任意类型的函数参数。
			2、空接口作为map的值， 使用空接口实现可以保存任意值的字典。
			3、空接口实现切片，切片中保存任意类型;
	*/
}

//1、 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
//2、 空接口作为map值
func studentInfo()  {
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "李白"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
}
//3、切片实现空接口
func sliceinterce()  {
	var slice = []interface{}{"张三", 20, true, 32.2}
	fmt.Println(slice)
}



func init() {
	// 定义一个空接口x
	var x interface{}
	s := "pprof.cn"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x) //type:string value:pprof.cn

	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x) //type:int value:100

	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x) //type:bool value:true
}