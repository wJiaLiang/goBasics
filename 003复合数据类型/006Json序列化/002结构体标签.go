package main

import (
	"encoding/json"
	"fmt"
)

/*
Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来
Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：
 `key1:"value1" key2:"value2"`

结构体标签由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。键值对之间使用一个空格分隔


*/
type Student struct {
	ID int `json:"userid"`  // 通过指定tag实现json序列化该字段时的key,
	Gender string  			// json序列化是默认使用字段名作为key
	name string  			// 首字母小写表示私有   私有字段不能被json包访问
}


func main()  {
	s2 := Student{
		ID:1,
		Gender:	"男",
		name: "opple",
	}


	// 结构体 to json
	data,err := json.Marshal(s2)
	if err != nil{
		fmt.Println("err")
		return
	}
	fmt.Printf("json str:%s \n", string(data) ) //json str:{"userid":1,"Gender":"男"}


	// json to 结构体
	var jsonStr = `{"userid":66,"Gender":"女"}`
	err2 := json.Unmarshal([]byte(jsonStr),&s2)
	if err2 != nil{
		fmt.Println("err")
	}
	fmt.Printf("%#v",s2) //main.Student{ID:66, Gender:"女", name:"opple"}

}
