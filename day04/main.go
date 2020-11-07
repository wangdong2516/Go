package main

import (
	"fmt"
	"math/big"
	"reflect"
	"strconv"
)

func main() {

	// 整型int
	v1 :=  10
	//result := strconv.FormatInt(int64(v1), 2)

	// 将int类型转换为字符串类型(必须是int类型的)
	srt := strconv.Itoa(v1)
	fmt.Println(srt, reflect.TypeOf(srt), reflect.TypeOf(v1))
	// 将字符串类型转换为整型,返回值有两个，一个表示转换后的结果(int类型)，另一个表示错误信息
	res2, _ := strconv.Atoi("1010")
	fmt.Println(res2, reflect.TypeOf(res2), reflect.TypeOf(v1))

	// 其他进制转换为字符串，第一个参数必须是int64类型的
	var v2 int8 = 12
	res3 := strconv.FormatInt(int64(v2), 2)
	fmt.Println(res3, reflect.TypeOf(res3), reflect.TypeOf(v1))

	// 字符串转换为其他进制,转换后的值永远都是int64类型的
	str3 := "1110"
	result, _ := strconv.ParseInt(str3, 2, 8)
	fmt.Println(result, reflect.TypeOf(result), reflect.TypeOf(v1))

	fmt.Println("+++++++++++++++++++++++")

	// 声明一个bigInt类型的对象
	var test1 big.Int
	fmt.Println(test1)

	// 创建一个bigInt类型的指针，没有进行初始化 nil，一般不使用，直接赋值的时候会使用
	var test2 *big.Int
	fmt.Println(test2)

	// 创建一个bigInt类型的指针，同时进行初始化 0
	test3 := new(big.Int)
	fmt.Println(test3)

	// 在超大整型对象中写入值，第二个参数表示把这个字符串当作几进制来进行处理，从而能够实现进制转换的效果
	// 能够将其他进制的值，转换为10进制的值。
	test1.SetString("10", 2)
	fmt.Println(test1.Int64(), reflect.TypeOf(test1.Int64()))


}
