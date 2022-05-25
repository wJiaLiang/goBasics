package main

import (
	"fmt"
)

func main() {

	// 定义数据
	//var arr1 [3]int
	//var arr2 [4]int
	//var strArr [3]string

//	 初始化数组
	var ar1 [3]int
	ar1[0] = 10
	ar1[1] = 23
	ar1[2] = 34

	fmt.Println(ar1)

//	 声明后直接赋值
	var arr = [3]int{23, 23, 43}
	fmt.Println(arr)

// 类型推导方式声明 赋值
	arrt := [3]string{"php","nodejs","golang"}
	fmt.Println(arrt)


// 初始化数组第3中方法 让编译器根据 初始值的个数自行推断数组的长度;第一次赋值后 但是长度 数组长度不能变
	var myarr = [...]int{12,34,56}
	fmt.Println(len(myarr))

// 第4种初始化数组的方式
	arrtest := [...]int{0:12,1:45,2:99,8:78}
	fmt.Println(arrtest)  // [12 45 99 0 0 0 0 0 78]

/*******循环数组  for 和 for range **********************/
	var a1 = [...]string{"php","golang","java"}
	for i := 0;i<len(a1);i++ {
		fmt.Println(a1[i])
	}

	for key,v := range a1{
		fmt.Printf("%v %v\n",key,v)
	}




}
