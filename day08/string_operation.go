package main

import (
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
)

func StringOperation() {

	fmt.Println("*****")
	var name string = "hello"
	var name2 string = "往东"

	// 获取字符串对应的字节的长度
	fmt.Println(len(name))
	fmt.Println(len(name2))

	// 获取字符串的长度
	res := utf8.RuneCountInString(name)
	res2 := utf8.RuneCountInString(name2)
	fmt.Println(res, reflect.TypeOf(res))
	fmt.Println(res2, reflect.TypeOf(res2))

	// 检查开头
	fmt.Println(strings.HasPrefix(name, "he"))
	fmt.Println(strings.HasPrefix(name, "hd"))

	// 检查结尾
	fmt.Println(strings.HasSuffix(name, "llo"))
	fmt.Println(strings.HasSuffix(name, "llr"))

	// 检查是否包含(子串)
	fmt.Println(strings.Contains(name, "he"))
	fmt.Println(strings.Contains(name, "hss"))

	// 字符串转大写
	fmt.Println(strings.ToUpper(name))

	// 字符串转小写
	fmt.Println(strings.ToLower(name))

	// 去掉字符串右边的字符
	address := "shanxsis"
	fmt.Println(strings.TrimRight(address, " "))

	// 去掉字符串左边的字符
	fmt.Println(strings.TrimLeft(address, "s"))
	fmt.Println(strings.Trim(address, "s"))

	// 数组的声明方式
	// 1. 先声明，再赋值
	var nums [3]int
	nums[1] = 9999
	nums[0] = 8888
	nums[2] = 7777
	fmt.Println(nums)
	fmt.Println(reflect.TypeOf(nums))

	// 2. 声明并且赋值
	var ages = [3]int {1, 2, 3}
	fmt.Println(ages)

	// 3. 声明+指定位置赋值
	var addrs = [3]int{0:99, 2:88, 1:77}
	fmt.Println(addrs)

	// 4. 创建一个指针类型的数组,不会开辟内存进行初始化
	var b *[3]int
	fmt.Println(b)
	fmt.Println(&b)

	// 5. 创建一个指针类型的数组，并且进行初始化
	var c = new([3]int)
	fmt.Println(c)

	// 创建一个指针，并且不会进行初始化
	var pdr *int
	fmt.Println(pdr)
	fmt.Println(&pdr)

	// go语言中，nil不是唯一的，也就是说，不同的nil有不同的内存地址



}

func init() {
	fmt.Println("进行初始化")
}
