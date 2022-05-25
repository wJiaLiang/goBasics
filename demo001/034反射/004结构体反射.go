package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string  `json name`
	Age int `json age`
	Score int `josn score`
}

func (s Student) Print() {
	fmt.Println("打印方法")
}
func(s Student) GetInfo() string{
	var str = fmt.Sprintf("姓名:%v 年龄:%v 成绩:%v",s.Name,s.Age,s.Score)
	return  str
}
func (s *Student) SetInfo(name string,age int,score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

// 打印字段
func PrintStructField(s interface{})  {
	// 判断参数是不是结构体
	t:= reflect.TypeOf(s)
	v:= reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入参数不是一个结构体")
		return
	}
	// 1.通过类型变量里面的Field 可以获取结构体的字段
	field0 := t.Field(0)
	fmt.Printf("%#v \n",field0)
	fmt.Println("字段名称:",field0.Name)
	fmt.Println("字段类型:",field0.Type)
	fmt.Println("字段Tag:",field0.Tag.Get("json"))
	fmt.Println("-------------------")
	// 2. 通过类型变量里面的 FieldByName 可以获取结构体的字段
	field1,ok := t.FieldByName("Age")
	if ok {
		fmt.Println("字段名称:",field1.Name)
		fmt.Println("字段类型:",field1.Type)
		fmt.Println("字段Tag:",field1.Tag.Get("json"))
	}
	// 3.通过类型变量里面的NumField 获取到改结构体有几个字段
	var count = t.NumField()
	fmt.Println(count) //3
	// 4. 获取结构体属性对应的值
	val:= v.FieldByName("Name")
	fmt.Println("属性值：",val,"---------------------------")


}

// 打印结构中方法
func PrintStructFn(s interface{}){
	// 判断参数是不是结构体
	t:= reflect.TypeOf(s)
	v:= reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入参数不是一个结构体")
		return
	}
	//1. 通过类型变量获取结构体里面的方法  Method
	method0:=t.Method(0)  // 0 和顺序没有关系，和 方法名的 ascll 码有关
	fmt.Println(method0.Name)  // GetInfo
	fmt.Println(method0.Type)  // func(main.Student) string

	//2. 通过类型变量获取这个结构体有多少个方法
	method1,ok:=t.MethodByName("Print") // 获取这个方法
	if ok {
		fmt.Println(method1.Name) //
		fmt.Println(method1.Type)
	}
	// 调用执行方法
	//v.Method(3).Call(nil)
	info1:= v.MethodByName("GetInfo").Call(nil)
	fmt.Println("+++++",info1)

	// 执行方法时传入参数
	var par []reflect.Value
	par = append(par,reflect.ValueOf("李四"))
	par = append(par,reflect.ValueOf(33))
	par = append(par,reflect.ValueOf(96))
	info2:= v.MethodByName("SetInfo").Call(par)
	fmt.Println("++++++",info2)

	// 获取方法数量
	fmt.Println(t.NumMethod())

}

func main(){
	stu1:= Student{
		Name: "小米",
		Age: 24,
		Score: 98,
	}
	//PrintStructField(stu1)
	PrintStructFn(&stu1)  // 传入指针类型才能修改变量
	fmt.Println(stu1)

}
