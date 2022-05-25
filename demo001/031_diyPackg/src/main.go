package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

// 第三包的使用

/*
安装
1.	go get github.com/shopspring/decimal   (全局)
2.  go mod download  下载到 gopath 中 (全局)
3.  go mod vendor 下载到当前项目中 vendor (当前项目)


*/


func main() {
	fmt.Println("第三方包引入")
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.NewFromInt(3)
	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")
	fmt.Println(fee)
	fmt.Println(taxRate)

	subtotal := price.Mul(quantity)
	fmt.Println("Subtotal:", subtotal)      // Subtotal: 408.06

	var n1 float64 = 3.45
	var n2 float64 = 4.55
	sum:= decimal.NewFromFloat(n1).Add(decimal.NewFromFloat(n2))
	fmt.Println(sum)

}

