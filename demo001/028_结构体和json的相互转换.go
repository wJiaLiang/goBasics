package main

import (
	"encoding/json"
	"fmt"
)
// 结构体标签  `json:"id"` 转换后的key 的值
type student struct {
	Id  int `json:"id"`
	Gender string `json:"gender"`
	Name string `json:"name"`
	sno string
}

func main()  {

	var s1 = student{
		Id:34,
		Gender:"男",
		Name: "李四",
		sno: "s001",
	}
	fmt.Printf("%#v \n",s1)  // &main.student{id:34, Gender:"男", name:"李四", sno:"s001"}
	jsonByte,_ := json.Marshal(&s1)
	jsonStr := string(jsonByte)
	fmt.Printf("%v \n",jsonByte)  // [123 34 71 101 110 100 101 114 34 58 34 231 148 183 34 125]

	fmt.Printf("%v \n",jsonStr)   // {"Id":34,"Gender":"男","Name":"李四"}  // 私有成员不会转换;


	//	 json 字符串
	var str = `{"Id":"666","Gender":"女","Name":"哈哈","sno":"007"}`
	var s2 student
	err:= json.Unmarshal([]byte(str),&s2)
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Printf("s2:%#v \n",s2) //s2:main.student{Id:0, Gender:"女", Name:"哈哈", sno:""}
	fmt.Println(s2.Name)  // 哈哈





}
