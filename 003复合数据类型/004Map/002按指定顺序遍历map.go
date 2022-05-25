package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main()  {
/*
	按照指定顺序遍历map
*/
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) // 生成 stu 开头的字符串
		value := rand.Intn(100)   			// 生成 0~99 的随机整数 scoreMap[key] = value }
		scoreMap[key] = value
	}
	//fmt.Println(scoreMap)

	//取出 map 中的所有 key 存入切片 keys
	var keys = make([]string,200)
	for key := range scoreMap{
		keys = append(keys,key)
	}

	//fmt.Println(keys)
	//对切片进行排序
	sort.Strings(keys)

	//按照排序后的 key 遍历 map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}







}
