package main

import "fmt"

func main(){
	//	数组是值类型
	/*
		基本数据类型和 数组 都是值类型
	*/

	var a = 23
	var b = a
	a = 24
	fmt.Println(a,b)  // 24 23

	var arr1 = [...]int{1,2,3}
	arr2 := arr1
	arr1[0] = 888
	fmt.Println(arr1)  // [888 2 3]
	fmt.Println(arr2)  // [1 2 3]

	//	 定义一个切片  切片就是引用数据类型;
	var brr1 = []int{1,2,3}
	brr2 := brr1
	brr1[0] = 888
	fmt.Println(brr1)  // [888 2 3]
	fmt.Println(brr2)  // [888 2 3]


/*******************多维数组*******************************************************/
//		 var 数组变量名 [元素数量][元素数量]T

	//var arrt = [4]int{1,2,3}

	// 二维数组
	var arrm = [3][2]string{
		{"北京0","上海0"},
		{"北京1","上海2"},
		{"北京2","上海2"},
	}

	fmt.Println(arrm[0])
	// for range 遍历二维数组
	for _,val :=range arrm {
		for key2,val2 := range val{
			fmt.Println(key2,val2)
		}


	}



}
