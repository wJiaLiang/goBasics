package main

import "fmt"

/*

注意，range 会复制对象。
*/
func main()  {
	a := [3]int{0, 1, 2}

	for i, v := range a { 	// index、value 都是从复制品中取出。

		if i == 0 { 		// 在修改前，我们先修改原数组。
			a[1], a[2] = 999, 999
			fmt.Println(a) 	// 确认修改有效，输出 [0, 999, 999]。
		}
		a[i] = v + 100 		// 使用复制品中取出的 value 修改原数组。
	}
	fmt.Println(a) 	// 输出 [100, 101, 102]。
	inits()
}

/*
	建议改用引用类型，其底层数据不会被复制。
*/
func inits()  {
	s := []int{1, 2, 3, 4, 5}

	for i, v := range s { 	// 复制 struct slice { pointer, len, cap }。
		if i == 0 {
			s = s[:3]  		// 对 slice 的修改，不会影响 range。
			s[2] = 100 		// 对底层数据的修改。
		}
		/*
			循环值 0 1
			循环值 1 2
			循环值 2 100
			循环值 3 4
			循环值 4 5
		*/
		fmt.Println("循环值",i, v)
	}
	fmt.Println("值",s) // 值 [1 2 100]


	/*
		for 和 for range有什么区别?
		主要是使用场景不同
		for可以
		遍历array和slice
		遍历key为整型递增的map
		遍历string
		for range可以完成所有for可以做的事情，却能做到for不能做的，包括
		遍历key为string类型的map并同时获取key和value
		遍历channel
	*/

}
