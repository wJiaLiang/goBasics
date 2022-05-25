package main

import (
	"fmt"
)

func main() {

	/*
		一、函数名、变量名、常量名、类型名、语句标号和包名等所有的命名规则:
			1、一个名字必须以一个字母（Unicode字母）或下划线开头，后面可以跟任意数量的字母、数字或下划线。
			2、大写字母和小写字母是不同的


		二、关键字:
			Go语言中类似if和switch的关键字有25个；关键字不能用于自定义名字，只能在特定语法结构中使用
			break      default       func     interface   select
			case       defer         go       map         struct
			chan       else          goto     package     switch
			const      fallthrough   if       range       type
			continue   for           import   return      var

		三、此外，还有大约30多个预定义的名字，比如int和true等，主要对应内建的常量、类型和函数。
			1、内建常量: true false iota nil

			2、内建类型: int int8 int16 int32 int64
						uint uint8 uint16 uint32 uint64 uintptr
						float32 float64 complex128 complex64
						bool byte rune string error

			3、内建函数：make len cap new append copy close delete
						complex real imag
						panic recover

			四、名字的开头字母的大小写决定了名字在包外的可见性
				首字母大写,全局变量;外部可以访问;
				首字母小写,私用变量;只在内部可以访问;
				Go语言程序员推荐使用 驼峰式 命名，当名字由几个单词组成时优先使用大小写分隔，而不是优先用下划线分隔



	*/


	var namesf int32  = 12
	fmt.Print(namesf)
}
