package main

import "fmt"

func main()  {
	/*
		map 数据 类型 是复合数据类型  无序的基于 key-value 的数据类型
	*/

//	1, 通过 make创建 map 类型数据

	var userinfo = make(map[string]string)
	userinfo["name"] = "张三"
	userinfo["age"] = "30"
	userinfo["sex"] = "男"

	fmt.Println(userinfo)  // map[age:30 name:张三 sex:男]

// 2，map 也支持初始化的时候 填充元素

	var userinfo2 = map[string]string{
		"name":"测试",
		"sex":"123",
		"age":"456",
	}

	fmt.Println(userinfo2["name"])  // 测试

//	map 类型  curd  创建，查找，修改，删除

	v,ok:=userinfo["name"]
	fmt.Println(v,ok)  // 张三 true


	delete(userinfo2,"age")
	fmt.Println(userinfo2) // map[name:测试 sex:123]





}
