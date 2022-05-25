package main
import "fmt"


func getuserinfo()(string,int){
	return  "zhangsan",23
}

func main() {
/*
	var 声明变量
	var 变量名 类型
	go 语言变量名由字母，数字，下划线组成，其中首个字符不能为数字，关键字和保留字不能作为变量名
	变量声明后才能使用
*/

// 定义变量并赋值
	var name string
	name = "呵呵"
	fmt.Println(name)
//	声明并赋值
	var age int = 20
	fmt.Println(age)

// 	类型自动推导  类型推导方式定义变量
	var sex = 0
	fmt.Println(sex)

/*
   一次声明多个变量
	var 变量名称,变量名称 类型
	var （
			变量名称 类型
   			变量名称 类型
		）
*/

	var a1,a2 string
	a1 = "tt"
	a2= "tts"
	fmt.Println(a1,a2)

	var (
		b1 = "张三"
		b2 = "李四"
	)
	fmt.Println(b1,b2)

/*
	短变量声明法，在函数内部可以使用更简略的  := 方式声明并初始化变量
	注意：短变量声明只能用于声明局部变量，不能用于全局变量的声明
*/
	username := "短变量"
	c1,c2,c3 := 3,6,"test"
	fmt.Println(username,c1,c2,c3)


/*
	匿名变量 在使用多重赋值时，如果想要忽律某个值，可以使用匿名变量
	匿名变量用 一个下划线 表示 _ ,函数返回多个值的时候,只接受某一个值，其他值用匿名变量代替，不然报错

*/
	var names,ages = getuserinfo()
	fmt.Println(names,ages)
	var xx,_ = getuserinfo()
	fmt.Println(xx)


fmt.Println("******************分界线**************************")

	/*fmt 包的使用 */
	a:=10
	b:=20
	fmt.Println("a=%v b=%v",a,b)
	fmt.Printf("a=%v b=%v",a,b)


/*
	常量 声明的时候赋值，值不可改变
	多个常量也可以一起声明;
	同时声明多个常量时，如果省略了值则表示和上面一行的值相同;
	iota 在 const 关键字出现时将被重置为 0,后面一次累加,可跳过中间值;
*/

	const pi = 3.1415926

	const (
		as = "as"
		bs = "bs"
		cs
	)
	fmt.Println(pi,"\n",as,bs,cs)

	const (
		n1 = iota
		n2
		m1 = 100
		m2
	)

/*
	1. 变量由数值，字母，下换线，组成，数值不能为开头
	2. 标识符开头不能是数字
	3. 标识符不能是关键字或保留字
	4. 变量名是区分大小写的，如 age 和 Age 是不同的变量名
	5. 变量名一定，最好见名知意，变量名称建议用名词，方法名称建议用动词
	6. 变量名一般用驼峰命名法，当遇到特有名词时，（缩写或者简称，如DNS）的时候，特有名词根据是否私有或全部大写或小写
	var maxAge = 10  小驼峰 私有
	var MaxAge = 20	 大驼峰 公有

	7. 每一行代码后面不要写分号

*/


}
