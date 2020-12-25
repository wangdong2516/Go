package main

import (
	"fmt"
)

func main() {
	// 1. 直接声明
	info := map[string]string{"name": "hello"}

	// 获取值
	fmt.Println(info["name"])

	// 修改值
	info["name"] = "world"

	// 写入值
	info["age"] = "12"

	fmt.Println(info)

	data := make(map[string]string)
	fmt.Println(data)
	fmt.Println(len(data))
	data["name"]= "haha"
	fmt.Println(len(data))

	data2 := make(map[string]string, 10)
	fmt.Println(data2)
	fmt.Println(len(data2))
	data2["name"] = "golang"
	fmt.Println(len(data2))

	var row map[string]string
	//row["name"] = 666 // 报错

	row = data
	fmt.Println(row)

	value := new(map[string]string)

	// 直接赋值
	value = &data2
	fmt.Println(value)

	// 键是数组
	test := make(map[[2]int]float32)
	test[[2]int{1, 2}] = 3.14
	fmt.Println(test)

	// 值是数组
	test2 := make(map[[2]int][2]int)
	test2[[2]int{1, 2}] = [2]int{2, 4}
	fmt.Println(test2)

	// 获取长度
	row2 := map[string]int{"name": 12, "age": 22}
	fmt.Println(len(row2))
}
