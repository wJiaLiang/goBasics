package main

import "fmt"

//结构体
type User struct {
	Name  string
	Email string
}

//方法
func (u User) Notify() {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
}

func main()  {
	/*
		解释： 首先我们定义了一个叫做 User 的结构体类型，然后定义了一个该类型的方法叫做 Notify，
		该方法的接受者是一个 User 类型的值。
		要调用 Notify 方法我们需要一个 User 类型的值或者指针。


	*/


	// 值类型调用方法
	u1 := User{"golang", "golang@golang.com"}
	u1.Notify()   //  golang : golang@golang.com

	// 指针类型调用方法
	u2 := User{"go", "go@go.com"}
	u3 := &u2
	u3.Notify()  // go : go@go.com


}
// ************************************************************************

/*
	定义一个 Data 类型的结构体,给他定义两个方法
	ValueTest方法的接收者是 值类型
	PointerTest方法的接收者是 指针类型

	注意：当接受者不是一个指针时，该方法操作对应接受者的值的副本,
	意思就是即使你使用了指针调用函数，但是函数的接受者是值类型，所以函数内部操作还是对副本的操作，而不是指针操作。

	注意：当接受者是指针时，即使用值类型调用那么函数内部也是对指针的操作。
	方法不过是一种特殊的函数

*/
type Data struct {
	x int
}

func (self Data) ValueTest() {
	fmt.Printf("Value fn: %p\n", &self)
}

func (self *Data) PointerTest() {
	fmt.Printf("Pointer fn: %p\n", self)
}
func init() {
	d := Data{}
	p := &d

	fmt.Printf("Data:%p\n",p)  // Data:0xc00000a0c0

	d.ValueTest()    // Value fn: 0xc00000a0e0
	d.PointerTest()  // Pointer fn: 0xc00000a0c0

	p.ValueTest()    // Value fn: 0xc00000a0e8
	p.PointerTest()  // Pointer fn: 0xc00000a0c0



}