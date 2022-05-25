package main

import "fmt"

func main(){

	/*
		选择排序
	*/
	var numSlice = []int{9,8,7,6,5,4}
	for i:=0;i< len( numSlice);i++{
		for j:=i+1;j< len(numSlice);j++ {
			if numSlice[i] > numSlice[j]{
				temp:= numSlice[i]
				numSlice[i] = numSlice[j]
				numSlice[j] = temp
			}
		}
	}
	fmt.Println(numSlice)

	/*
		冒泡排序
	*/
	var numSlice2 = []int{9,6,3,5,2,7}

	for i := 0; i < len(numSlice2); i++ {
		for j := 0; j < len(numSlice2)-1-i; j++ {
			if numSlice2[j] > numSlice2[j+1] {
				temp := numSlice2[j]
				numSlice2[j] = numSlice2[j+1]
				numSlice2[j+1] = temp
			}
		}
	}
	fmt.Println(numSlice2)












}
