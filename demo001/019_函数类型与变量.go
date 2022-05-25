package main

import "fmt"


type calc func(int,int) int  // 使用type 关键字定义一个 类型
type myint int   // 自定义一个 myint 类型

func add(x,y int) int  {
	return  x+y
}
func sub(x,y int) int {
	return x-y
}

/**************************************************************/
// 函数作为函数的参数
func calcmy(x,y int,cb func(int,int) int) int {
	return 0
}
// 自定义一个方法类型
type calcType func(int,int) int
func calc2(x,y int,cb calcType) int {
	return cb(x,y)
}
// 函数作为返回值
type caltype func(int,int) int

func do(o string) calcType  {
	switch o {
	case "+":
		return  add
	case "-":
		return sub
	case "*":
		return func(x,y int) int {
			return x*y
		}
	default:
		return nil
	}
}

func abc(x,y int)  {
	fmt.Println(x+y)
}

func main()  {
	/*
		可以使用 type 来定义一个函数类型
		type myfunc func(int,int) int
	*/
	var c calc
	c = add
	fmt.Printf("%T\n",c) // main.calc

	f := sub
	fmt.Printf("%T\n",f) //func(int, int) int

	var value = calc2(10,6,add)
	fmt.Println(value) // 16

	//	 传入匿名函数
	j:=calc2(3,4, func(x,y int) int {
		return  x * y
	})
	fmt.Println(j)

	var a = do("+")
	fmt.Println(a(12,12)) // 24

	// 匿名函数
	func () {
		fmt.Println("匿名函数。。。")
	}()
	abc(3,5) // 8
}
