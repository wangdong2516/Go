package main

import (
	"fmt"
	"reflect"
)

//func Add(data []int) {
//	data[0] = 999
//}

// go中return不是原子性的，是分为两步执行的
// 先赋值，在return
// 在不使用命名返回值的情况下，return会赋值给特殊的变量
// 在使用命名返回值的情况下，return会赋值给命名的返回值
func f1() (x int) {
	x = 10
	defer func(){
		x++
	}()
	return
}

func main() {
	res := f1()
	fmt.Println(res)
	//Utils()
	//num := []int{9, 8}
	//Add(num)
	//fmt.Println(num)
	//Proxy(1, Add)
	//fmt.Println("*******")
	//Do(1, 2, 3, 4, 5)
	//a1 := [4]int{1, 2, 3, 4}
	//s1 := a1[1:2:2]
	//fmt.Println(s1)
	//fmt.Println(cap(s1))
	//s2 := a1[:]
	//s2[0] = 888
	//fmt.Println(a1)
	//fmt.Println(s2)
	//fmt.Println(len(s2))
	//fmt.Println(cap(s2))
	//a := []int{1, 2, 3, 4, 5, 6 ,7, 8, 9}
	//b := a[:7:7]
	//fmt.Println(b)
	//fmt.Println(cap(b))
	//fmt.Println(reflect.TypeOf(b))
	//// 给b添加一个元素，引发扩容
	//b = append(b, a...)
	//b[0] = 999999
	//fmt.Println(b, a)
	//fmt.Println(" ********")
	//a1 := []int{ 1, 2, 3}
	//a2 := make([]int, 3)
	//copy(a1, a2)
	//fmt.Println(a2)
	//// 切片的遍历
	//// 索引遍历
	//for i := 0; i < len(a); i ++ {
	//	fmt.Println(a[i])
	//}
	//
	//// for range遍历
	//for index, item := range a {
	//	fmt.Println(index, item)
	//}
	//
	//// make函数创建切片
	//s9 := make([]int, 9)
	//fmt.Println(s9)
	//var functionList []func()
	//
	//for i :=0; i < 5; i++ {
	//	function := func(){
	//		fmt.Println(i)
	//	}
	//	functionList = append(functionList, function)
	//}
	//
	//functionList[0]()
	//functionList[1]()
	//functionList[2]()
}


func Proxy(data int, funx func([]int)) int {
	funx([]int{1, 8})
	fmt.Println(data)
	return 0
}

func Do(num ...int) {
	fmt.Println(num)
	fmt.Println(reflect.TypeOf(num))
}
