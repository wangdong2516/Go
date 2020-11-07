package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//var name string
	//var message string
	// fmt.Scanf支持按照模板输入，只提取输入中占位符的内容
	//fmt.Scanf("我叫%s", &name)
	//fmt.Println(name)

	// fmt包无法解决的问题是:如果用户输入的时候是带有空格的
	// 除非拿对应个数的变量接收，否则只会获取空格之前的内容
	//fmt.Scanln(&message)
	//fmt.Println(message)

	// 解决方式是使用系统输入, os.Stdin是一个文件
	reader := bufio.NewReader(os.Stdin)
	// line表示从Stdin中读取一行的数据，是一个字节类型的
	// reader默认可以读取4096个字节
	// isPrefix表示是否读取完毕，读取完毕为true，没有读取完毕为false
	line,_, _ := reader.ReadLine()
	data := string(line)
	fmt.Println(data)

}
