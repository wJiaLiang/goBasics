package main

import (
	"fmt"
	"math"
)



type Point struct{ X, Y float64 }

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}


type aaa string

func (abc aaa) add(x int) (y int)  {
	x+=x
	y =x
	return  y
}

func main()  {
/*
	一、概念:
		1、在函数声明时，在其名字之前放上一个变量，即是一个方法。
		这个附加的参数会将该函数附加到这种类型上，即相当于为这种类型定义了一个独占的方法。

		2、上面的代码里那个附加的参数p，叫做方法的接收器(receiver)，
			早期的面向对象语言留下的遗产将调用一个方法称为“向一个对象发送消息”。

		3、第一个Distance的调用实际上用的是包级别的函数Distance，
		而第二个则是使用刚刚声明的Point，调用的是Point类下声明的Point.Distance方法。

	二、
	这种p.Distance的表达式叫做选择器，因为他会选择合适的对应p这个对象的Distance方法来执行。
	选择器也会被用来选择一个struct类型的字段，比如p.X。由于方法和字段都是在同一命名空间，
	所以如果我们在这里声明一个X方法的话，编译器会报错，因为在调用p.X时会有歧义


*/

	p := Point{1, 2}
	q := Point{4, 6}

	fmt.Println(Distance(p,q))
	fmt.Println(p.Distance(q))



	var abc  aaa
	fmt.Println(abc.add(6)) //12   给string 的实例 abc 上扩展了一个 add 方法;
	
	
}


