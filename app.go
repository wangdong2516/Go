package main

import (
	"fmt"
	"practice/api"
)

func main() {
	fmt.Println( "开始执行")
	Test()
	api.Baidu()
	// 声明赋值一同进行
	var age int = 19
	fmt.Println(age)
	var flag bool = true
	// 先声明再赋值
	var a string
	a = "字符串"
	fmt.Println(a)
	fmt.Println(flag)
	api.Google()
	// 注意事项:函数名必须大写，大写的表示共有的功能
	// 外部可以调用到，小写的函数表示只能在当前的包中
	// 使用，别人无法调用

	// 声明变量的意义：方便以后进行引用，节省代码量，进行数据的存储
	var name string
	fmt.Scanf("%s", &name)
	fmt.Println(name)
	if name == "wangdong" {
		fmt.Println("用户名输入成功")
	} else {
		fmt.Println("用户名输入失败")
	}

	// 变量名必须只包含数字、字母和下划线，不能以数字开头并且不能是内置的关键字

	// 变量简写
	var name1 = "ddd"
	fmt.Println(name1)

	name2 := "wangdong"
	fmt.Println(name2)

	var name3,age2,add2 string
	fmt.Println(name3, age2, add2)

	// 因式分解,变量如果只声明不赋值，默认使用的是对应类型的控值
	var (
		name5 = "quuu"
		age5 = 12
		length int
	)
	fmt.Println(name5,age5, length)

	// 全局变量和局部变量
	/* 定义在函数外部的变量叫做全局变量，全局变量不能按照变量省略的
	 方式来定义，必须包含var关键字，全局变量是进行变量查找的时候
	 最后寻找的变量，先是局部作用域中进行查找。查找不到的话，去全局
	变量中进行查找 */
	// 定义在函数内部的变量叫做局部变量
	/*
		赋值和内存相关
		var name string = "函数"
		var nick_name = name
		实际上是将"函数"复制了一份，内存地址是不一样的
		当对变量进行重新赋值的时候，内存地址是不会改变的
		会在原先的位置进行替换,只是进行了值的替换。(针对整型、字符串、布尔类型)
	*/
	var name7 string = "函数"
	var nick_name = name
	fmt.Println(name7, &name7)
	name = "不是函数"
	fmt.Println(name7, &name7)
	fmt.Println(nick_name, &nick_name)
}