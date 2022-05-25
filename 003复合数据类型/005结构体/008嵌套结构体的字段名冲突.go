package main

import "fmt"

/*
	嵌套结构体内部可能存在相同的字段名。这个时候为了避免歧义需要指定具体的内嵌结构体 的字段。

*/

//Address2 地址结构体
type Address2 struct {
	Province string
	City string
	CreateTime string
}
//Email2 邮箱结构体
type Email2 struct {
	Account string
	CreateTime string
}
//User2 用户结构体
type  User2 struct {
	Name string
	Gender string
	Address2
	Email2
}


func main()  {
	var user3 User2
	user3.Name = "哈哈"
	user3.Address2.CreateTime = "2020"

	// main.User2{Name:"哈哈", Gender:"", Address2:main.Address2{Province:"", City:"", CreateTime:"2020"}, Email2:main.Email2{Account:"", CreateTime:""}}
	fmt.Printf("%#v",user3)

}
