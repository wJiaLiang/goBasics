package main

import "fmt"

func main()  {
/*
	Slice（切片）代表变长的序列，序列中每个元素都有相同的类型
	slice的切片操作s[i:j]，其中0 ≤ i≤ j≤ cap(s)，用于创建一个新的slice

	一、切片的定义:
	切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。 它非常灵活，支持自动扩容。

	二、切片是一个引用类型，它的内部结构包含地址、长度和容量。
		1、前后两个变量共享底层数组，对一个切片的修改会影响另一个切片 的内容，这点需要特别注意。
		2、切片不能直接比较，应该用 len() == 0;
	var name []T

	长度对应slice中元素的数目；长度不能超过容量，容量一般是从slice的开始位置到底层数据的结尾位置。
	len()返回长度
	cap()返回容量

*/

	/*
	1、切片的定义:
		var name []T
	*/
	var a []string
	var b = []int{}

	fmt.Println(a,b)
	fmt.Printf("%v,%T \n",a,a)  // [],[]string


	/*
	  2、当你声明了一个变量 , 但却还并没有赋值时 , golang 中会自动给你的变量赋值一个默认零 值。这是每种类型对应的零值。
		如果你需要测试一个slice是否是空的，使用len(s) == 0来判断，而不应该用s == nil来判断
		bool -> false
		numbers -> 0
		string-> ""
		pointers -> nil
		slices -> nil
		maps -> nil
		channels -> nil
		functions -> nil
		interfaces -> nil
	*/

	var c  []int
	fmt.Println(len(c) == 0)  // true

	/*
		3、切片的循环遍历
			切片的循环遍历和数组的循环遍历是一样的
	*/

	var s1 = []string{"北京", "上海", "深圳"}
	// 方法 1：for 循环遍历
	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i])
	}

	// 方法 2：for range 遍历
	for index, value := range s1 {
		fmt.Println(index, value)
	}


	/*
	   4、基于数组定义切片
		 由于切片的底层就是一个数组，所以我们可以基于数组定义切片
	*/
	a1 := [5]int{55, 56, 57, 58, 59}
	b1 := a1[1:4] 		//基于数组 a 创建切片，包括元素 a1[1],a1[2],a1[3]
	fmt.Println(b1)  	// [56 57 58]

	c1 := a1[1:] //[56 57 58 59]
	d1 := a1[:4] //[55 56 57 58]
	e1 := a1[:]  //[55 56 57 58 59]
	fmt.Println(c1,"\n",d1,"\n",e1,"\n")

	/*
	   5、切片再切片
	      除了基于数组得到切片，我们还可以通过切片来得到切片 包含前面不包含后面
		  对切片进行再切片时，索引不能超过原数组的长度，否则会出现索引越界的错误。
		  切片的长度就是它所包含的元素个数。
		  切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
	*/
	a2 := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	b2 := a2[1:3]
	fmt.Printf("b2:%v type:%T len:%d cap:%d\n", b2, b2, len(b2), cap(b2)) //b2:[上海 广州] type:[]string len:2 cap:5
	c2 := b2[1:5]
	fmt.Printf("c2:%v type:%T len:%d cap:%d\n", c2, c2, len(c2), cap(c2)) //c2:[广州 深圳 成都 重庆] type:[]string len:4 cap:4

	/*
		9、使用 make()函数构造切片
		上面都是基于数组来创建的切片，如果需要动态的创建一个切片，我们就需要使用内置 的 make()函数，
		make([]T,size,cap)
			T:切片的元素类型
			size:切片中元素的数量
			cap:切片的容量

	*/






}
