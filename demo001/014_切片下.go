package main

import "fmt"

func main(){

	//	 如果要动态的创建一个切片，我们就要使用内置的函数  make([]T,size,cap)

	/**	其中
		T:切片的元素类型
		size: 切片中元素的数量
		cap:切片的容量
	*/

	var sliceA = make([]int,5,9)

	fmt.Println(sliceA)  //[0 0 0 0 0]
	fmt.Printf("长度:%d,容量:%d \n",len(sliceA),cap(sliceA)) //长度:5,容量:9

	var sliceC []int

	fmt.Printf("%v -- %v -- %v \n",sliceC,len(sliceC),cap(sliceC)) //[] -- 0 -- 0

	// sliceC[0] = 89
	// fmt.Println(sliceC) // 这里报错, 不能通过上面方式扩容; 没法通过下标的方式给切片扩容


	// 1.	给 切片扩容 要用到 append()
	var sliceD []int
	sliceD = append(sliceD,12,32)
	fmt.Println(sliceD) // [12 32]
	sliceD = append(sliceD,66,88)
	fmt.Println(sliceD) // [12 32 66 88]

	fmt.Printf("%v -- %v -- %v \n",sliceD,len(sliceD),cap(sliceD)) // [12 32 66 88] -- 4 -- 4

	// 2.	append() 可以合并 切片

	sliceF := []string{"php","java"}
	sliceE := []string{"node.js","Golang"}

	sliceE = append(sliceE, sliceF...) // 固定语法

	fmt.Println(sliceE) // [node.js Golang php java]

	//	3. 切片的扩容 策略
	var sliceJ []int
	for i:=0;i<10;i++ {
		sliceJ = append(sliceJ,i)
		fmt.Printf("值:%v -- 长度:%v -- 容量:%v \n",sliceJ,len(sliceJ),cap(sliceJ))
	}
	/*
	   值:[0] -- 长度:1 -- 容量:1
	   值:[0 1] -- 长度:2 -- 容量:2
	   值:[0 1 2] -- 长度:3 -- 容量:4
	   值:[0 1 2 3] -- 长度:4 -- 容量:4
	   值:[0 1 2 3 4] -- 长度:5 -- 容量:8
	   值:[0 1 2 3 4 5] -- 长度:6 -- 容量:8
	   值:[0 1 2 3 4 5 6] -- 长度:7 -- 容量:8
	   值:[0 1 2 3 4 5 6 7] -- 长度:8 -- 容量:8
	   值:[0 1 2 3 4 5 6 7 8] -- 长度:9 -- 容量:16
	   值:[0 1 2 3 4 5 6 7 8 9] -- 长度:10 -- 容量:16

	*/

	/***********************************************************************/
	/*
		值类型：改变变量副本值的时候 不会改变 变量本身的值
		引用类型：改变变量副本值的时候 会改变变量本身的值
	*/
	//	 切片是引用类型
	slieAA := []int{2,5,6}
	slicBB := slieAA
	slicBB[0] = 88
	fmt.Println(slieAA) // [88 5 6]
	fmt.Println(slicBB) // [88 5 6]

	//	 copy() 函数 复制切片
	sliceCC := []int{1,12,45,66}
	sliceDD := make([]int,4,4)
	copy(sliceDD,sliceCC)  // 拷贝  sliceCC 给 sliceDD
	sliceDD[0] = 88
	fmt.Println(sliceCC)  // [1 12 45 66]
	fmt.Println(sliceDD)  // [88 12 45 66]


	/***************************************************************************/
	// 	从 切片中删除元素   go 中没有单独的删除切片中元素的方法;
	a := []int{30,31,32,33,34,35,36,37}
	//	 删除索引为2的元素
	a = append(a[:2],a[3:]...)
	fmt.Println(a)  // [30 31 33 34 35 36 37]








}