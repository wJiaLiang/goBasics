package main

import "log"

func main()  {
/*
	Go语言内置的log包实现了简单的日志服务。本文介绍了标准库log的基本使用。

*/

	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")


}
