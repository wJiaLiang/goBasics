package main



func main()  {
/*
	一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
   所有给定类型的方法属于该类型的方法集。
	方法定义：
	    func (recevier type) methodName(参数列表)(返回值列表){}
		 参数和返回值可以省略
*/




}
type Test struct { }

//1、 无参数  五返回值
func (T Test) method0(){ }

//2、 单参数  无返回值
func (T Test) method1(i int){ }

//3、 多参数 无返回值
func (T Test) method2(x,y int)  { }

//4、 无参数 单返回值
func (T Test) method3(i int)  { }

//5、 多参数 多返回值
func (T Test) method4(x,y int) (z int,err error)  { return }

//6、 无参数、无返回值
func (T *Test) method5()  { }

//7、 单参数、无返回值
func (T *Test) method6(i int)  { }

//8、 多参数,无返回值
func (T *Test) method7(x,y int)  { }

//9、 无参数，单返回值
func (T *Test) method8() (i int) {return }

// 多参数，多返回值
func (T *Test) method9(x,y int) (z int,err error) {return }

