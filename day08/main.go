package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

func main() {
	// 定义字符串
	var name string = "王栋"

	// 获取字符串对应的字节的长度(字符串转换为utf-8编码之后字节的长度)
	fmt.Println(len(name))

	// 将字符串转换为字节的集合[231 142 139 230 160 139]
	bytesList := []byte(name)
	fmt.Println(bytesList)

	// 将字节的集合转换为字符串
	n1 := []byte{231, 142, 139, 230, 160, 139}
	n2 := string(n1)
	fmt.Println(n1, reflect.TypeOf(n1))
	fmt.Println(n2, reflect.TypeOf(n2))

	// 对应字符串的码点
	n3 := []rune(name)
	fmt.Println(n3, reflect.TypeOf(n3))

	// 将对应的码点转换为字符串
	n4 := []rune{29579, 26635}
	n5 := string(n4)
	fmt.Println(n4, reflect.TypeOf(n4))
	fmt.Println(n5, reflect.TypeOf(n5))

	// 长度
	n6 := utf8.RuneCountInString(name)
	fmt.Println(n6, reflect.TypeOf(n6))

	StringOperation()
	// init函数将不能被调用
}
