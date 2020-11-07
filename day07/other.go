package main

import (
	"fmt"
	"reflect"
)

func Wait() {
	var name string
	fmt.Print("请输入你的姓名:")
	fmt.Scan(&name)
	if name == "wangdong" {
		// SVIP
		goto SVIP
	} else if name == "sunling" {
		goto VIP
	} else {
		goto BASE
	}
BASE:
	fmt.Println("预约")
VIP:
	fmt.Println("等号")
SVIP:
	fmt.Println("进入")

}

func Format() {
	var name, address, action string
	fmt.Print("请输入姓名:")
	fmt.Scan(&name)
	fmt.Println(reflect.TypeOf(name))

	fmt.Print("请输入位置:")
	fmt.Scan(&address)

	fmt.Print("请输入行为:")
	fmt.Scan(&action)

	fmt.Println("开始格式化")
	result:= fmt.Sprintf("我叫%s, 现在在%s, 正在%s", name, address, action)
	fmt.Println(result)
}

func Operator() {
	a:= 2
	b:= 3
	fmt.Println(a % b)
}