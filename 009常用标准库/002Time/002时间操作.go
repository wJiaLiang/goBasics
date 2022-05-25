package main

import (
	"fmt"
	"time"
)

func main()  {
/*
	我们在日常的编码过程中可能会遇到要求时间+时间间隔的需求，Go语言的时间对象有提供Add方法如下：
*/

	now:=time.Now()
	later:= now.Add(time.Hour)  // 当前时间加1小时后的时间
	fmt.Println(later)


//	求两个时间之间的差值：
	t:= now.Sub(later)
	fmt.Println(t)

/*
    Equal 判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。
	本方法和用t==u不同，这种方法还会比较地点和时区信息。
*/
	fmt.Println(now.Equal(later))  // false

/*
	Before
	如果t代表的时间点在u之前，返回真；否则返回假。
*/
	fmt.Println(now.Before(later)) // true


/*
	After
	如果t代表的时间点在u之后，返回真；否则返回假。
*/
	fmt.Println(now.After(later))  // fase




}
