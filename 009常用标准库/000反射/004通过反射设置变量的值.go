package main

import (
	"fmt"
	"reflect"
)

func reflectSetValue1(x interface{}){
	v:= reflect.ValueOf(x)
	if v.Kind() == reflect.Int64{
		v.SetInt(200) 	//修改的是副本，reflect 包会引发 panic
	}
}

func reflectSetValue2(x interface{})  {
	v:=reflect.ValueOf(x)

	// 反射中使用 Elem()方法获取指针 对应的值
	if v.Elem().Kind() == reflect.Int64{
		v.Elem().SetInt(200)
	}

}




func main()  {
/*
	想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地 址才能修改变量值。
	而反射中使用专有的 Elem()方法来获取指针对应的值。

*/

	var a int64 = 1000
	//reflectSetValue1(a)

	reflectSetValue2(&a)

	fmt.Println(a)  // 200
}



















