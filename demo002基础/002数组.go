package main

import "fmt"

func main() {
	fmt.Println("测试")

	//一维数组:
	var arr0 [5]int // 未初始化元素值为 0。
	arr0 = [5]int{1, 2, 3, 4, 5}
	var arr1 = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	var arr2 = [...]string{"1", "2"} //通过初始化值确定数组长度。
	fmt.Println(arr0)                //[1 2 3 4 5]
	fmt.Println(arr1)                //[1 2 3 4 5 6 7 8]
	fmt.Println(arr2)                //[1 2]
}
