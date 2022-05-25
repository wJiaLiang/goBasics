package main

import (
	"fmt"
)

func main()  {

}

func init() {
	/*
	   一、元素为 map 类型的切片

	*/
	var mapSlice = make([]map[string]string,3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d  value:%v \n",index,value)
	}
	fmt.Println("----------------------------------------")
	//对切片中的 map 元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "海淀区"
	for index,value := range mapSlice{
		fmt.Printf("index:%d value:%v \n",index,value)
	}
	// 输出
	// index:0 value:map[address:海淀区 name:小王子 password:123456]
	// index:1 value:map[]
	// index:1 value:map[]


	/*
		二、值为切片类型的map

	*/
	fmt.Println("***********************************************")
	var slicMap = make(map[string][]string,3)
	fmt.Println(slicMap)  //map[];
	slicMap["name"] = make([]string,3,3)
	slicMap["name"][0] = "kg"
	slicMap["name"][1] = "wei"
	slicMap["name"][2] = "liang"

	for key,v := range slicMap{
		fmt.Printf("key:%s value:%v",key,v)  // key:name value:[kg wei liang]

	}





}
