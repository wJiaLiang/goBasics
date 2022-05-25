package main

import "fmt"

func main()  {
/*
   for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
*/

 var s string = "abc"

 for i := range s{
 	fmt.Println(s[i])      // 97  98 99
 	fmt.Println(string(s[i]))
 }
 fmt.Println("*********************************************************")

	for _,v :=range s{
		fmt.Println(v)  	// 97  98 99
	}


}
