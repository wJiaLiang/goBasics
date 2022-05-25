package main

import (
	"fmt"
	"reflect"
)


//student 结构体
type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student) GetInfo() string {
	var str = fmt.Sprintf("姓名:%v 年龄:%v 成绩:%v", s.Name, s.Age, s.Score)
	fmt.Println(str)
	return str
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s *Student) Print() {
	fmt.Println("打印方法...")

}
// 获取结构体中的 字段
func PrintStructField(s interface{})  {
	t:= reflect.TypeOf(s)
	v:= reflect.ValueOf(s)

	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct{
		return
		fmt.Println("非结构体数据")
	}

	//1、通过类型变量里面的Field可以获取结构体的字段
	field0 := t.Field(0)
	fmt.Printf("%#v \n", field0)
	fmt.Println("字段名称：", field0.Name)
	fmt.Println("字段类型：", field0.Type)
	fmt.Println("字段Tag：", field0.Tag.Get("json"))
	fmt.Println("字段Tag：", field0.Tag.Get("form"))

	//2、通过类型变量里面的FieldByName可以获取结构体的字段
	fmt.Println("----------------------")
	field1, ok := t.FieldByName("Age")
	if ok {
		fmt.Println("字段名称：", field1.Name)
		fmt.Println("字段类型：", field1.Type)
		fmt.Println("字段Tag：", field1.Tag.Get("json"))
	}

	//3、通过类型变量里面的NumField获取到该结构体有几个字段
	var fieldCount = t.NumField()
	fmt.Println("结构体有", fieldCount, "个属性")

	//4、通过值变量获取结构体属性对应的值
	fmt.Println(v.FieldByName("Name"))
	fmt.Println(v.FieldByName("Age"))

}

// 获取结构体中的方法和执行该方法
func PrintStructFn(s interface{})  {
	t:=reflect.TypeOf(s)
	v:=reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
		return
	}

	//1、通过类型变量里面的Method可以获取结构体的方法
	method0 := t.Method(0)    //和结构体方法的顺序没有关系，和结构体方法的ASCII有关系
	fmt.Println(method0.Name) //GetInfo
	fmt.Println(method0.Type) //func(main.Student) string


	//2、通过类型变量获取这个结构体有多少个方法
	method1, ok := t.MethodByName("Print")
	if ok {
		fmt.Println(method1.Name) //Print
		fmt.Println(method1.Type) //func(main.Student)
	}
	fmt.Println("--------------------------")

	//3、通过《值变量》执行方法 （注意需要使用值变量，并且要注意参数） v.Method(0).Call(nil) 或者v.MethodByName("Print").Call(nil)
	v.MethodByName("Print").Call(nil)
	info1 := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info1)

	//4、执行方法传入参数 （注意需要使用《值变量》，并且要注意参数,接收的参数是[]reflect.Value的切片）
	var params []reflect.Value
	params = append(params, reflect.ValueOf("李四"))
	params = append(params, reflect.ValueOf(23))
	params = append(params, reflect.ValueOf(99))
	v.MethodByName("SetInfo").Call(params) //执行方法传入参数

	info2 := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info2)

	// 5、获取方法数量
	fmt.Println("方法数量:", t.NumMethod())
}

func main() {
	/*
		1、与结构体相关的方法

		任意值通过 reflect.TypeOf()获得反射对象信息后，如果它的类型是结构体，可以通过反射值
		对象（reflect.Type）的 NumField()和 Field()方法获得结构体成员的详细信息。

		reflect.Type 中与获取结构体成员相关的的方法如下表所示

		方法 									说明
		Field(i int) StructField 				根据索引，返回索引对应的结构体字段的信 息。
		NumField() int 							返回结构体成员字段数量。 FieldByName(name string) (StructField, bool)
												根据给定字符串返回字符串对应的结构体字段的信息。

		FieldByIndex(index []int) StructField   多层成员访问时，根据 []int 提供的每个结构 体的字段索引，返回字段的信息。

		FieldByNameFunc(match func(string) bool) (StructField,bool) 	根据传入的匹配函数匹配需要的字段。
		NumMethod() int 						返回该类型的方法集中方法的数目
		Method(int) Method 						返回该类型方法集中的第 i 个方法
		MethodByName(string)(Method, bool) 		根据方法名返回该类型方法集中的方法


		当我们使用反射得到一个结构体数据之后可以通过索引依次获取其字段信息，也可以通过字 段名去获取指定的字段信息。

	*/

	stu1 := Student{
		Name: "小明",
		Age: 15,
		Score: 98,
	}
	PrintStructField(stu1)
	//PrintStructFn(stu1)

	fmt.Println("main函数",stu1)

}
