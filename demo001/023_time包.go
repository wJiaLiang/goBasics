package main

import (
	"fmt"
	"time"
)

// 获取当前日期
func fn1()  {
	timeObj := time.Now()
	fmt.Println(timeObj)  // 2020-06-07 11:45:07.1644847 +0800 CST m=+0.003988901
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()

	fmt.Printf("%d-%d-%d-%d-%d-%d \n",year,month,day,hour,minute,second) //2020-6-7-11-51-54
	fmt.Printf("%d-%02d-%02d-%02d-%02d-%02d \n",year,month,day,hour,minute,second) //2020-06-07-11-52-50
	// %02d 中的2表示宽带，如果整数不够 2 列就补上 0;
}
// 格式化模板
func test2()  {
	// 2006 年
	// 01 月
	// 02 日
	// 03时   12小时制   15 24小时制
	// 04分
	// 05秒

	timeObj := time.Now()
	strTime := timeObj.Format("2006-01-02 03:04:05")  // 2020-06-07 12:05:09
	strTime2 := timeObj.Format("2006-01-02 15:04:05")  // 24 小时制
	fmt.Println(strTime)
	fmt.Println(strTime2)
}

// 获取当前时间戳  从1970 到当前时间的 毫秒数
func test3()  {
	timeObj := time.Now()
	var s = timeObj.Unix()
	var s2 = timeObj.UnixNano()
	fmt.Println(s) // 1591502910
	fmt.Println(s2) // 纳秒   1591502978658917800

}

// 时间戳转 字符串时间格式
func test4()  {
	var strTime = 1591502910
	timeObj := time.Unix(int64(strTime),0)
	var str = timeObj.Format("2006-01-02 03:04:05")
	fmt.Println(str)  // 2020-06-07 12:08:30

}
// 日期字符串转换成 时间戳;
func test5()  {
	var str = "2020-06-07 12:08:30"
	var tmp = "2006-01-02 15:04:05"
	timeObj,_ := time.ParseInLocation(tmp,str,time.Local)
	fmt.Println(timeObj.Unix()) // 1591502910
}

// 定时器
func test6()  {
	time.Now()

	ticker := time.NewTicker(time.Second) // 定义一个1秒钟间隔的定时器
	n:=5
	for t := range ticker.C {
		n--
		fmt.Println(t)
		if n ==0 {
			ticker.Stop()
			break
		}
	}
}

// 每隔一秒钟 才执行 休眠
func test7()  {
	fmt.Println("aaa")
	time.Sleep(time.Second * 3)
	fmt.Println("aaa")
	time.Sleep(time.Second)
	fmt.Println("aaa")

	for  {
		time.Sleep(time.Second)
		fmt.Println("定时器执行,死循环，每隔一秒")
	}
}


func main()  {
	//fn1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	test7()




}
