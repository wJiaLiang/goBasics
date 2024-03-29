package main

import "fmt"

func main()  {
	// 在切片里面放一系列的的用户数据 ，可以定义要给元素为map 类型的切片

	var userinfo = make( []map[string]string,3,3 )

	fmt.Println(userinfo[0])  // [map[] map[] map[]]
	fmt.Println(userinfo[0])  // map[]  默认不初始化的默认值为 nil;


	//下面的代码演示了切片中的元素为map类型时的操作：
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "红旗大街"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}


	//
	//下面的代码演示了map中值为切片类型的操作：
	fmt.Println("--------------------")
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)

}
