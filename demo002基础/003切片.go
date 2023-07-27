package main

import "fmt"

func main() {
	fmt.Println("main")

	/*
		切片就是 数组不用写长度
		len(s)  切片s的长度，总是 <= cap(s);
		cap(s)  切片s的容量，总是 >= len(s);
	*/
	s1 := []int{0, 1, 2, 3, 8: 100}   // 通过初始化表达式构造，可使用索引号。
	fmt.Println(s1, len(s1), cap(s1)) //[0 1 2 3 0 0 0 0 100] 9 9
	var s2 = s1[1:3]
	fmt.Printf("%v,%T", s2, s2) // [1,2],[]int

	/*
			通过make来创建切片
			var slice []type = make([]type, len)
		    slice  := make([]type, len)
			slice  := make([]type, len, cap)

		     使用 make()函数构造切片
				上面都是基于数组来创建的切片，如果需要动态的创建一个切片，我们就需要使用内置 的 make()函数，
				make([]T,size,cap)
					T:切片的元素类型
					size:切片中元素的数量
					cap:切片的容量
	*/

}
