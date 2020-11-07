package main

import "fmt"

func init() {
	//init()函数是程序启动之后最先执行的函数
	fmt.Println("最先执行的函数，先执行init函数，在执行main函数")
}

func main() {
	//这是一个注释
	fmt.Println("Hello Go")
}
