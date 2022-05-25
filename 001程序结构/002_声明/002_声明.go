package main



import "fmt"

/*
	声明语句定义了程序的各种实体对象以及部分或全部的属性
	要有四种类型的声明语句：
	var     声明 变量
	const   声明 常量
	type    声明 类型
	func    声明 函数实体对象
	本节讨论变量和类型的声明,其他后面介绍;

*/


// 在包一级声明语句声明的名字可在整个包对应的每个源文件中访问
const boilingF = 212.0


func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	// Output:
	// boiling point = 212°F or 100°C


}






