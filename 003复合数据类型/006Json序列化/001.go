package main

import (
	"encoding/json"
	"fmt"
)

/*
一、JSON(JavaScript Object Notation) 是一种轻量级的数据交换格式。易于人阅读和编写。
	同时也 易于机器解析和生成。RESTfull Api 接口中返回的数据都是 json 数据。

二、结构体与 JSON 序列化
Golang JSON 序列化是指把结构体数据转化成 JSON 格式的字符串，
Golang JSON 的反序列化 是指把 JSON 数据转化成 Golang 中的结构体对象
Golang 中 的 序 列 化 和 反 序 列 化 主 要 通 过 "encoding/json" 包 中 的 json.Marshal() 和json.Unmarshal()方法实现

*/

func main()  {

//1、结构体对象转化成 Json 字符串
	type Student struct {
		ID int
		Gender string
		name string
		Sno string
	}
	var js1,_= json.Marshal(Student{
		ID:     0,
		Gender: "男",
		name:   "kgwei",
		Sno:    "haha",
	})

	// json js1 返回的是 []byte 类型切片
	// [123 34 73 68 34 58 48 44 34 71 101 110 100 101 114 34 58 34 231 148 183 34 44 34 83 110 111 34 58 34 104 97 104 97 34 125]
	fmt.Println(js1)

	var jsonStr = string(js1)
	// {"ID":0,"Gender":"男","Sno":"haha"}
	fmt.Println(jsonStr)


// 2、Json 字符串转换成结构体对象

	var jsonStr2 =`{"ID":67823945,"Gender":"女","Sno":"牛逼"}`
	var student2 Student
	err:=json.Unmarshal([]byte(jsonStr2),&student2)
	if err != nil {
		fmt.Printf("err %#v",err)
	}
	fmt.Printf("反序列化 %#v",student2) //反序列化 main.Student{ID:67823945, Gender:"女", name:"", Sno:"牛逼"}
}
