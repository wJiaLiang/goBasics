package main

import "fmt"

func main()  {
	
}
/*
	一、概念
	哈希表是一种巧妙并且实用的数据结构。它是一个无序的key/value对的集合，其中所有的key都是不同的，
	然后通过给定的key可以在常数时间复杂度内检索、更新或删除对应的value。

	一个map就是一个哈希表的引用
	map类型可以写为map[K]V，其中K和V分别对应key和value。
	所有的key类型相同，所有的value类型相同，但是key和value之间可以是不同的数据类型。
	其中K对应的key必须是支持==比较运算符的数据类型
*/

func init()  {
	//1、内置的make函数可以创建一个map：
	ages := make(map[string]int)
	fmt.Println(ages)  		//map[]

	//2、字面量创建一个 map
	ages1 := map[string]int{
		"alice":12,
		"charlie":33,
	}
	// ages1["alice"] = 666     // 修改值
	fmt.Println(ages1) 		    // map[alice:12 charlie:33]

	//3、另一种创建空的map的表达式是map[string]int{}。
	var ages2 = map[string]int{"a":13,"b":14,"c":15}
	fmt.Println(ages2)  	// map[a:13 b:14 c:15]

	//4、 使用内置的delete函数可以删除元素
	delete(ages2,"c")
	fmt.Println(ages2) 		// map[a:13 b:14]

	/*
	  5、map中的元素并不是一个变量，因此我们不能对map的元素进行取址操作：_ := &ages2["bob"]
		禁止对map元素取址的原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效。
	*/

	/*
		循环 range
		在实践中,遍历的顺序是随机的，每一次遍历的顺序都不相同

	*/
	for name,v := range ages2{
		fmt.Println(name,v) 	// b 14   a 13
	}

	//	6、判断某个键是否存在 value, ok := map 对象[key]
	_,ok := ages2["an"]
	if ok {
		fmt.Println("ok")
	}else{
		fmt.Println("no")  // no
	}



}