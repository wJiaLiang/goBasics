package main

import "fmt"

func main()  {
	// 在切片里面放一系列的的用户数据 ，可以定义要给元素为map 类型的切片

	var userinfo = make( []map[string]string,3,3 )

	fmt.Println(userinfo[0])  // [map[] map[] map[]]
	fmt.Println(userinfo[0])  // map[]  默认不初始化的默认值为 nil;








}
