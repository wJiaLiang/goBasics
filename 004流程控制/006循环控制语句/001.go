package main

import "fmt"

func main()  {

	/*
	 	循环控制Goto、Break、Continue


		1.三个语句都可以配合标签(label)使用
	    2.标签名区分大小写，定以后若不使用会造成编译错误
	    3.continue、break配合标签(label)可用于多层循环跳出
	    4.goto是调整执行位置，与continue、break配合标签(label)的结果并不相同

		与其他语言一样，Go语言也支持label(标签)语法：分别是break label和 goto label
	*/

	// 直接跳出多重循环
	// break标签只能用于for循环
	// break的跳转标签(label)必须放在循环语句for前面，并且在break label跳出循环不再执行for循环里的代码。
	label:
	for i:=0;i<10;i++ {
		for j:=0;j<10;j++ {
			if j==3 {
				break label
			}
			fmt.Println("ts",j )
		}
	}

	/****************************************************************************************/

	// continue 语句可以结束当前循环，开始洗一次的循环，仅限在 for循环内使用
	for i:=0;i<10;i++ {
		if i == 2 || i == 4{
			continue
		}
		fmt.Println(i)  // 跳过 2 和 4;
	}

	/****************************************************************************************/

	// goto 通过标签 进行 代码间的无条件跳转 goto 可以快速跳出循环 跳转到指定的位置，继续执行;
	var n =20
	if n>10{
		fmt.Println("dddd")
		goto label2
	}

	fmt.Println("aaaa")
	fmt.Println("bbbb")
	label2:
	fmt.Println("cccc")

}
