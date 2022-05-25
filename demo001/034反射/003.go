package main

import (
	"fmt"
	"reflect"
)


// 反射 获取 任意变量的类型

// 报错
func reflectFn1(x interface{})  {
	v:=reflect.ValueOf(x)
	if v.Kind() == reflect.Float64{
		v.SetFloat(123.4)
	}


}

// 如果传入的是指针变量  通过 v.Elem().Kind()来获取
func reflectFn2(x interface{})  {
	v:=reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Float64{
		v.Elem().SetFloat(123.4)
	}

}
func main()  {
	a:=10.3
	reflectFn2(&a)
	fmt.Println(a) // 123.4


}


