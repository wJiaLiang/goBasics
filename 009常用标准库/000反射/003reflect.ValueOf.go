package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{})  {
	// 反射获取变量的原始值
	v:= reflect.ValueOf(x)
	kind:= v.Kind()

	// t:= v.Type()    返回 reflect.Type

	switch kind {
		case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
			fmt.Println("int:",v.Int())
		case reflect.String:
			fmt.Println("string:",v.String())
		case reflect.Bool:
			fmt.Println("布尔型:",v.Bool())
		case reflect.Float64:
			fmt.Printf("float64:%f \n",v.Float())
		case reflect.Slice:
			fmt.Println("slice:",v)
	}
}

func main()  {
	/*
		reflect.ValueOf()返回的是 reflect.Value 类型，其中包含了原始值的值信息。
		reflect.Value 与原 始值之间可以互相转换

		reflect.Value 类型提供的获取原始值的方法如下

		Interface()     interface {}   将值以 interface{} 类型返回，可以通过类型断言转换为指定类型
		Int() 	  int64 	 将值以 int 类型返回，所有有符号整型均可以此方式返回
		Uint()    uint64     将值以 uint 类型返回，所有无符号整型均可以此方式返回
		Float()  float64     将值以双精度（float64）类型返回，所有浮点数（float32、float64）均 可以此方式返回
		Bool()    bool       将值以 bool 类型返回
		Bytes() []bytes      将值以字节数组 []bytes 类型返回
		String()   string    将值以字符串类型返回
	*/


	var  a = 23
	reflectValue(a)    // int: 23

	var b = []int{1,3}
	reflectValue(b)    // slice: [1 3]

	var c = 22.2
	reflectValue(c)    // float64: 22.2

	// 将 int 类型的原始值转换为 reflect.Value 类型
	d := reflect.ValueOf(10)
	fmt.Printf("typec:%T\n",d)   // typec:reflect.Value


}
