package main

import "fmt"

/*
一个结构体中可以嵌套包含另一个 结构体 或 结构体指针。

*/

//Address 地址用户结构体
type Address struct {
	Province string
	City string
}
//User 用户结构体
type User struct {
	Name string
	Gender string
	Address Address
}

func main()  {
	user1 := User{
		Name:    "张三",
		Gender:  "女",
		Address: Address{
			Province:"广东",
			City:"广州",
		},
	}

	// {张三 女 {广东 广州}}
	fmt.Println(user1)
	//main.User{Name:"张三", Gender:"女", Address:main.Address{Province:"广东", City:"广州"}}
	fmt.Printf("%#v",user1)


}
