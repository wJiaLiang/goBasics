package main

import "fmt"

//Student 学生
type Student1 struct {
	ID int
	Gender string
	Name string
}
//Class 班级
type Class struct {
	Title string
	Students []Student1
}

func main() {
	c := &Class{
		Title:    "001",
		Students: make([]Student1, 0, 200),
	}

	for i := 0; i < 10; i++ {
		stu := Student1{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)

	}


}
