package main

import "fmt"

// 结构体的匿名字段 如果是匿名这个里面的 类型是唯一的
type Persontest struct {
	string
	int
}
// 定义一个Persons 结构体
type Persons struct {
	Name string
	Age int
	Hobby []string
	map1  map[string]string
}

/*
	嵌套结构体  也叫继承结构体
*/
type User struct {
	username   string
	Password   string
	Address
}

type Address struct {
	Name string
	Phone string
	City string
}





func  main()  {
	p:= Persontest{
		"zhang san",
		20,
	}
	fmt.Println(p)  //{ zhang san 20}

	p1 := &Persons{}
	p1.Name = "lishi"
	p1.Age = 230
	p1.Hobby = make([]string,3,6)
	p1.Hobby[0] = "写代码"
	p1.Hobby[1] = "打男球"
	p1.Hobby[2] = "吃饭"
	p1.map1 = make(map[string]string)
	p1.map1["address"] = "北京"
	p1.map1["phone"] = "177777"

	// &main.Persons{Name:"lishi", Age:230, Hobby:[]string{"写代码", "打男球", "吃饭"}, map1:map[string]string{"address":"北京", "phone":"177777"}}--*main.Persons
	fmt.Printf("%#v--%T \n" ,p1,p1)


	// 嵌套结构体  当访问结构体成员时先在机构中查找该字段，找不到再到匿名的结构体中查找
	var u User
	u.username = "test"
	u.Password = "666"
	u.Address.Name = "张三"
	u.Address.Phone = "1234567"
	u.Address.City = "上海"
	// main.User{username:"test", Password:"666", Address:main.Address{Name:"张三", Phone:"1234567", City:"上海"}}--main.User
	fmt.Printf("%#v--%T \n",u,u)
	fmt.Printf(u.City) // 上海



}
