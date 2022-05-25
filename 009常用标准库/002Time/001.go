package main

import (
	"fmt"
	"time"
)

func main()  {
/*
	时间和日期是我们编程中经常会用到的，本文主要介绍了Go语言内置的time包的基本用法。

	time包提供了时间的显示和测量用的函数。日历的计算采用的是公历。


*/

	timeDemo()
	timestampDemo()
	timestampDemo2(time.Now().Unix())
	timeconst()

}

/*
	1、time.Time类型表示时间。我们可以通过time.Now()函数获取当前的时间对象，然后获取时间对象的年月日时分秒等信息
*/
func timeDemo() {
	now := time.Now() 	//获取当前时间对象
	fmt.Printf("%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second) //2020-12-03 10:51:26
	fmt.Printf("%d-%d-%d %d:%d:%d\n", year, month, day, hour, minute, second) //2020-12-3 10:51:26
	fmt.Println("------")
}

/*
	2、时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数。它也被称为Unix时间戳（UnixTimestamp）。
*/
func timestampDemo() {
	now := time.Now()            // 获取当前时间
	timestamp1 := now.Unix()     // 时间戳
	timestamp2 := now.UnixNano() // 纳秒时间戳
	fmt.Printf("%v\n", timestamp1) // 1606963714
	fmt.Printf("%v\n", timestamp2) // 1606963714428265400
	fmt.Println("------")
}

/*
	3、使用time.Unix()函数可以将时间戳转为时间格式。
*/
func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Println("------")
}

/*
	时间间隔
	time.Duration是time包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位。
	time.Duration表示一段时间间隔，可表示的最长时间段大约290年。

*/
func timeconst()  {
	const Nanosecond  = time.Nanosecond          // 纳秒
	const Microsecond =  Nanosecond * 1000       // 微妙
	const Millisecond =  Microsecond * 1000      // 毫秒
	const Second = Millisecond *1000             // 秒
	const Minute = Second * 60                   // 分
	const Hour = Minute * 60					 // 小时

	fmt.Println("时间常量纳秒:",Nanosecond)
	fmt.Println("时间常量微妙:",Microsecond)
	fmt.Println("时间常量毫秒:",Millisecond)
	fmt.Println("时间常量秒:",Second)
	fmt.Println("时间常量分:",Minute)
	fmt.Println("时间常量小时:",Hour)

}





