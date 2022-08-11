package main

import "fmt"

func main()  {
	fmt.Println("55")

/*
	map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
	map[KeyType]ValueType
	    KeyType:表示键的类型。
	    ValueType:表示键对应的值的类型。
	map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：
	    make(map[KeyType]ValueType, [cap])




*/

	var mapvalue = make(map[string]int,5);
	mapvalue["name1"] = 1
	mapvalue["name2"] = 50
	fmt.Printf("%v,%T,\n",mapvalue,mapvalue); //map[name1:1 name2:50],map[string]int


	/*
		判断某个键是否存在
		Go语言中有个判断map中键是否存在的特殊写法，格式如下:
		value, ok := map[key];
	*/
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := mapvalue["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	//	map的遍历
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 60
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	//	使用delete()函数删除键值对
	/*
		delete(map, key)
		map:表示要删除键值对的map
	    key:表示要删除的键值对的键


	*/





}
