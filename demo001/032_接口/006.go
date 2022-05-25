package main

import "fmt"

type Address struct{
	Name string
	Phone int
}

// 空接口和类型断言使用细节
func main()  {
	var userInfo = make(map[string]interface{})
	userInfo["userName"] = "张三"
	userInfo["age"] = 43
	userInfo["hobby"] = []string{"吃","睡"}

	fmt.Println(userInfo["age"])
	fmt.Println(userInfo["userName"])

	fmt.Println(userInfo["hobby"])
	//fmt.Println(userInfo["hobby"][0])  // 空接口不支持索引

	var address = Address{
		Name:"李四",
		Phone:1232323,

	}

	fmt.Println(address.Name)

	userInfo["address"] = address

	fmt.Println(userInfo["address"])
	// fmt.Println(userInfo["address"].Name) // 错误不能这样获取;


	bo,_:=userInfo["hobby"].([]string)
	fmt.Println(bo[1])

	add2,_:= userInfo["address"].(Address)
	fmt.Println(add2.Phone)



}
