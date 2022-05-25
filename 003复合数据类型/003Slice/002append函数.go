package main

import "fmt"

func main()  {
/*
	一、append();
	append()函数 可以为切片动态添加元素，
	append函数对于理解slice底层是如何工作的非常重要，
	append() 第一个参数是要添加的切片,第二个是要添加的元素;
*/
	var numSlice []int
	for i:=0;i<10;i++ {
		numSlice = append(numSlice,i)
		fmt.Printf("%v----len:%d----cap:%d \n",numSlice,len(numSlice),cap(numSlice));
	}

/*
	特点：
   	1、append()函数将元素追加到切片的最后并返回该切片。
   	2、切片 numSlice 的容量按照 1，2，4，8，16 这样的规则自动进行扩容，每次扩容后都是 扩容前的 2 倍。
	3、append()函数还支持一次性追加多个元素
	输出结果：
   [0]----len:1----cap:1
   [0 1]----len:2----cap:2
   [0 1 2]----len:3----cap:4
   [0 1 2 3]----len:4----cap:4
   [0 1 2 3 4]----len:5----cap:8
   [0 1 2 3 4 5]----len:6----cap:8
   [0 1 2 3 4 5 6]----len:7----cap:8
   [0 1 2 3 4 5 6 7]----len:8----cap:8
   [0 1 2 3 4 5 6 7 8]----len:9----cap:16
   [0 1 2 3 4 5 6 7 8 9]----len:10----cap:16

*/


/*
	二、copy();
   使用 copy()函数复制切片
   由于切片是引用类型，所以 a 和 b 其实都指向了同一块内存地址。修改 b 的同时 a 的值也会 发生变化。
   Go 语言内建的 copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中，
   使用格式:
   copy(destSlice, srcSlice []T)  将 srcSlice 中的元素复制到 destSlice中; 返回长度;
		1、切片 dst 需要先初始化长度
			如果 dst 长度小于 src 的长度，则 copy 部分；
			如果大于，则全部拷贝过来，只是没占满 dst 的坑位而已；
			相等时刚好不多不少 copy 过来。
		2、源切片中元素类型为引用类型时，拷贝的是引用
*/

	var numSlice2 = numSlice
	numSlice[0] = 99
	fmt.Println(numSlice,numSlice2) // 都改变了 [99 1 2 3 4 5 6 7 8 9]

	var numSlice3 = make([]int,len(numSlice))
	copy(numSlice3, numSlice)

	fmt.Println(numSlice3)  //[99 1 2 3 4 5 6 7 8 9]



/*
	三、
	从切片中删除元素
	没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素
*/
	a := []int{30,31,32,33,34,35,36,37}
	//	 删除索引为2的元素
	a = append(a[:2],a[3:]...)  // ... 是将切边展开;
	fmt.Println(a)  	// [30 31 33 34 35 36 37]

	test(a[3:]...)

}

func test(args ...int){
	fmt.Println(args)
}