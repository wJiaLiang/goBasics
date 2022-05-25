package main

import "fmt"

/*

1、一个变量对应一个保存了变量对应类型值的内存空间
2、一个指针的值是另一个变量的地址。一个指针对应变量在内存中的存储位置。
3、并不是每一个值都会有一个内存地址，但是对于每一个变量必然有对应的内存地址。
4、通过指针，我们可以直接读或更新对应变量的值，而不需要知道该变量的名字（如果变量有名字的话）。
5、如果用“var x int”声明语句声明一个x变量，那么&x表达式（取x变量的内存地址）

6、指针对应的数据类型是 *int
7、如果指针名字为p，那么可以说“p指针指向变量x”，或者说“p指针保存了x变量的内存地址”。同时*p表达式对应p指针指向的变量的值
8、因为*p对应一个变量，所以该表达式也可以出现在赋值语句的左边，表示更新指针所指向的变量的值。
9、对于聚合类型每个成员——比如结构体的每个字段、或者是数组的每个元素——也都是对应一个变量，因此可以被取地址。
10、任何类型的指针的零值都是 nil, 如果p指向某个有效变量，那么p != nil测试为真。
11、指针之间也是可以进行相等测试的，只有当它们指向同一个变量或全部是nil时才相等。
	var z, y int
	var w *int
	fmt.Println(&z == &z, &z == &y, &z == nil) // true  false   false
	fmt.Println(w == nil)  // true

12、调用f函数时创建局部变量v，在局部变量地址被返回之后依然有效，因为指针p依然引用这个变量。
	var p = f()
	func f() *int {
		v := 1
		return &v
	}
	fmt.Println(f() == f()) // "false"

13、每次我们对一个变量取地址，或者复制指针，我们都是为原变量创建了新的别名

*/


func main()  {

	var x int = 666
	fmt.Println(&x)    // x变量的内存地址  0xc00000a0c0
	var p *int = &x    // p 保存了x变量的内存地址

	fmt.Println(*p)    // *p 获取 p指针指向的变量的值  666

	*p = 888
	fmt.Println(x)     //可用重新被赋值   888


	var z, y int
	var w *int
	fmt.Println(&z == &z, &z == &y, &z == nil) // true  false   false
	fmt.Println(w == nil)  // true


	fmt.Println(f() == f()) //每次调用f函数都将返回不同的结果 "false"


}
var p = f()
func f() *int {
	v := 1
	return &v
}

