package main

import "fmt"

func main()  {
/**
	切片的声明  和数组相似, 只不过，不用指定长度：
	var arr []int

 */

	var arr []int
	fmt.Printf("%v,类型：%T,长度:%v \n",arr,arr,len(arr))  // [],类型：[]int,长度:0

	var arr1 = []int{23,45,54}
	fmt.Printf("%v,类型：%T,长度:%v \n",arr1,arr1,len(arr1))  // [23 45 54],类型：[]int,长度:3

	var arr2 = []int{1:23,2:45,4:54}
	fmt.Printf("%v,类型：%T,长度:%v \n",arr2,arr2,len(arr2))  // [0 23 45 0 54],类型：[]int,长度:5

	//	切片声明后不赋值 默认值是 nil
	var arr3 []int
	fmt.Println(arr3 == nil)  // true

	//	 切片的循环遍历   for  和 for range 都可以;
	var arr4Slice = []string{"java","php","node","c","c++"}

	for key,val := range arr4Slice {
		fmt.Println(key,val)
	}


	//	基于数组定义的切片
	a:= [5]int{34,54,67,56,77}
	b:= a[:]  // 获取数组里面的所有值
	fmt.Println(a,"\n",b)
	fmt.Printf("%v,%T \n",b,b)

	c:= a[1:4]  // 截取 下标 1到4直接的数据; 包括左边，不包括右边;
	fmt.Printf("%v,%T\n",c,c)  //[54 67 56],[]int

	d:= a[2:]   // 获取2后面的所有
	fmt.Printf("%v,%T\n",d,d)  // [67 56 77],[]int

	e:= a[:3]   // 获取3前面的所有
	fmt.Printf("%v,%T\n",e,e)  // [67 56 77],[]int


	//	基于切片的切片 基于切片定义切片

	var aa = []string{"北京","上海","广州","深圳","成都","重庆"}
	bb:= aa[1:]

	fmt.Printf("%v,%T\n",bb,bb)  // [上海 广州 深圳 成都 重庆],[]string

	fmt.Println("**************************************************")
	/*
	   关于切片的长度和容量
		长度：切片的长度就是它所包含的元素个数
		容量:切片的容量是从它的第一个元素开始数 到其底层数组元素末尾的个数
	*/
	cc:=[]int{2,3,5,7,9,11,13}
	fmt.Printf("长度%d 容量%d \n",len(cc),cap(cc)) // 长度7 容量7

	cc1:= cc[2:]  //
	fmt.Printf("值%v 长度%d 容量%d \n",cc1,len(cc1),cap(cc1)) // 值[5 7 9 11 13] 长度5 容量5

	cc2:= cc[1:3]
	fmt.Printf("值%v 长度%d 容量%d \n",cc2,len(cc2),cap(cc2)) // 值[3 5] 长度2 容量6

	cc3:= cc[:3]
	fmt.Printf("值%v 长度%d 容量%d \n",cc3,len(cc3),cap(cc3)) // 值[2 3 5] 长度3 容量7


}
