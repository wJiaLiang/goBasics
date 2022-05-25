package main

import (
	"fmt"
	"reflect"
)

//student结构体
type Student2 struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}
func (s Student2) GetInfo() string {
	var str = fmt.Sprintf("姓名:%v 年龄:%v 成绩:%v", s.Name, s.Age, s.Score)
	return str
}

// 修改结构体属性
func reflectChangeStruct(s interface{})  {
	t:= reflect.TypeOf(s)
	v:=reflect.ValueOf(s)

	// 传入指针获取的值
	fmt.Println(t.Kind())          // Ptr
	fmt.Println(t.Elem().Kind())   // struct

	if t.Kind() != reflect.Ptr{
		fmt.Println("参数不是结构体指针类型")
		return
	}else if t.Elem().Kind() != reflect.Struct {
		fmt.Println("参数不是结构体类型")
		return
	}

	// 修改结构体属性
	name := v.Elem().FieldByName("Name")
	name.SetString("哈哈")

	age := v.Elem().FieldByName("Age")
	age.SetInt(44)

}



func main()  {
	stu2 := Student2{
		Name:  "小明",
		Age:   15,
		Score: 98,
	}

	//reflectChangeStruct(stu2)
	reflectChangeStruct(&stu2)

	fmt.Println(stu2)   // {哈哈 44 98}

}
