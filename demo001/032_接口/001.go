package main

import (
	"fmt"
)

/*
	接口   38
*/

// 空接口，表示还没有任何约束，任意的类型都可以实现空接口
type A interface {

}

// 空接口也可以直接当做类型来使用，可以表示任意类型;
// show 函数的参数是空接口类型,表示可以传入任意的数据类型;
func show(c interface{})  {
	fmt.Printf("值:%v ---类型:%T \n",c,c)
}

// 判断参数类型
func mytest(x interface{}){
	switch x.(type) {
	case int:
		fmt.Println("int类型")
	case string:
		fmt.Println("string类型")
	case bool:
		fmt.Printf("布尔类型")
	default:
		fmt.Println("默认")
	}


}


func main()  {
	var a A
	var str = "hello,Golang"
	a = str            // 让字符串实现a 这个接口;
	fmt.Printf("值:%v ---类型:%T \n",a,a)

	var num = 20
	a = num     // 表int 类型实现a 这个接口;
	fmt.Printf("值:%v ---类型:%T \n",a,a)


	var b interface{}
	b = 20
	b = "Hi"
	b = true
	fmt.Printf("值:%v ---类型:%T \n",b,b)

	show(99)
	show("hell,my name is Golang")

	var m1 = make(map[string]string)
	var m2 = make(map[string]interface{})
	m1["name"] = "zhangsan"
	m1["age"] = "23"

	fmt.Printf("值:%v ---类型:%T \n",m1,m1)

	m2["name"] = "zhangsan2222"
	m2["age"] = "2322222"

	fmt.Printf("值:%v ---类型:%T \n",m2,m2)

	//	 类型断言  判断这个 空接口 到底是什么数据类型;
	var nn interface{}
	nn = "类型断言"
	v,ok := nn.(string)   // 返回值 和 判断是否成功;
	if ok{
		fmt.Printf("%v--%v",v,ok)
	}else{
		fmt.Println("filed")
	}




}
