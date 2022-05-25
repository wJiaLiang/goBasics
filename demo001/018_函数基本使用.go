package main

import (
	"fmt"
	"sort"
)
	/*
		函数定义
		func 函数名(参数)(返回值) { 函数体  }
		返回值只有一个不用括号,
	*/

// 求两个数的和

func sumFn(x int,y int) int {
	return  x+y
}

// x 参数是不固定的值,表示传入的所有参数
func sumfn1(x ...int)  {
	fmt.Println(x)
}

// 一次返回多个值
func calc(x,y int)(int,int)  {
	sum:=x+y
	sub:=x-y
	return sum,sub
}


/**********************************************************************/
/*
	数组切片进行升序排序
	有一个map 对象 m1,封装一个方法，要求按照key 升序 排序，最后输出一个 key=>vale
*/

var sliceA = []int{12,34,37,35,556,36,2}

func sortIntAsc(slice []int) []int  {
	for i := 0; i <len(slice) ; i++ {
		for j := i+1; j < len(slice); j++ {
			if slice[i]>slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
	return slice
}

var m1 = map[string]string{
	"username":"zhangsan",
	"age":"20",
	"sex":"男",
	"height":"189",
}
func mapSort(map1 map[string]string) string  {
	var sliceKey []string
	for k,_ := range map1 {
		sliceKey = append(sliceKey,k)
	}
	//	对key 进行升序排序
	sort.Strings(sliceKey)
	var str string
	for _, v := range sliceKey {
		str += fmt.Sprintf("%v=>%v,",v,map1[v])
	}
	return str
}

func main()  {
	sum :=sumFn(3,5)
	fmt.Println(sum) // 8

	sumfn1(3,5,6,0) //[3 5 6 0]

	a,b:=calc(8,2)
	fmt.Println(a,b) // 10 6

	var arr = sortIntAsc(sliceA)
	fmt.Println(arr) // [2 12 34 35 36 37 556]

	var arr2 = mapSort(m1)
	fmt.Println(arr2)
}
