package main

import "fmt"

//Address 地址结构体
type AddressTo struct {
	Province string
	City     string
}

//User 用户结构体
type UserTo struct {
	Name    string
	Gender  string
	AddressTo //匿名结构体
}

func main()  {
	var user2 UserTo
	user2.Name = "kgwei"
	user2.Gender = "男"
	user2.AddressTo.Province = "新疆"
	user2.City = "乌鲁木齐"

	// main.UserTo{Name:"kgwei", Gender:"男", AddressTo:main.AddressTo{Province:"新疆", City:"乌鲁木齐"}}
	fmt.Printf("%#v",user2)

}
