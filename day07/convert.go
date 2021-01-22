//package main
//
//import (
//	"fmt"
//	//"github.com/shopspring/decimal"
//	"math"
//	"math/big"
//	"reflect"
//	"strconv"
//)
//
//func Convent() {
//	var v1 int8 = 10
//	var v3 int16 = 18
//	v5 := int16(v1) + v3
//	fmt.Println(v5)
//	fmt.Println(reflect.TypeOf(v5))
//
//	var v2 int16 = 128
//	v4:= int8(v2)
//	fmt.Println(v4)
//
//	v6 := 19
//	result:= strconv.Itoa(v6)
//	fmt.Println(result)
//	fmt.Println(reflect.TypeOf(result))
//
//	v7 := "66"
//	res1,_ := strconv.Atoi(v7)
//	fmt.Println(res1)
//	fmt.Println(reflect.TypeOf(res1))
//
//	v8 := 111
//	// 将十进制转换为二进制，使用strconv.FormatInt包的时候，必须传入一个64位的整型
//	v9 := int64(v8)
//	res := strconv.FormatInt(v9, 16)
//	fmt.Println(v9, reflect.TypeOf(v9))
//	fmt.Println(res, reflect.TypeOf(res))
//
//	fmt.Println("*****")
//	data := "010101001111"
//	result1, error := strconv.ParseInt(data, 2, 8)
//	fmt.Println(result, error, reflect.TypeOf(result1))
//
//	// 向上取整
//	fmt.Println(math.Ceil(3.94))
//	fmt.Println(math.Ceil(-3.94))
//
//	// 向下取整
//	fmt.Println(math.Floor(3.94))
//	fmt.Println(math.Floor(-9.9))
//
//	// 取绝对值
//	fmt.Println(math.Abs(-19))
//	fmt.Println(math.Abs(19))
//
//	// 四舍五入,保留整数部分
//	fmt.Println(math.Round(3.49))
//	fmt.Println(math.Round(3.59))
//
//	// 保留小数
//	fmt.Println(math.Round(3.1415 * 100) / 100)
//	fmt.Println(math.Round(3.1455 * 100) / 100)
//	fmt.Println(3.1455 * 100 / 100)
//	fmt.Println(math.Round(3.1455 * 100))
//
//	// v1指向的是初始化之后的地址
//	var v11 big.Int
//
//	// v2指向的是nil的内存地址
//	var v22 *big.Int
//
//	// v3指向的是初始化之后的地址
//	v13 := new(big.Int)
//
//	// 写入
//	v11.SetInt64(1990)
//	//v22.SetString("12774312291", 10)
//	v13.SetString("010100101010010101", 2)
//	fmt.Println(v11, v13)
//	fmt.Println(reflect.TypeOf(v11))
//	fmt.Println(reflect.TypeOf(v22))
//	fmt.Println(reflect.TypeOf(v13))
//
//	v66 := new(int)
//	fmt.Println(reflect.TypeOf(v66))
//	fmt.Println(v66)
//
//	n1 := new(big.Int)
//	n1.SetInt64(88)
//
//	n2 := new(big.Int)
//	n2.SetInt64(99)
//
//	n3 := new(big.Int)
//	fmt.Println(n3)
//	n3.Add(n1, n2)
//	fmt.Println(n3)
//
//	var n4 float32
//	n4 = 3.14
//	fmt.Println(n4)
//
//	var h1 = decimal.NewFromFloat(0.29)
//	var h2 = decimal.NewFromFloat(0.29)
//
//	fmt.Println(h1.Add(h2))
//	fmt.Println(reflect.TypeOf(h1))
//
//	fmt.Println(h1.Round(1))
//	fmt.Println(h1.Truncate(1))
//}

package main

import "fmt"

func main() {
	fmt.Println("test")
}