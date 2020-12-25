package main

import (
	"fmt"
)

// 结构体嵌套
type Person struct {
	name string
	age int
	school School
}

// 结构体定义
type School struct {
	name string
	address string
}

// 结构体的拷贝
type Home struct {
	name string
	miles int
	person [2]int
	food []string
	relation map[string]string
}

func main() {
	// 结构体中可以存储不同类型的数据，是一种数据集合，可以用来表示一条记录

	// 结构体声明
	//person := Person{}
	//// 赋值1
	//person.name = "wandong"
	//person.age = 25
	//fmt.Println(person)
	//
	//// 赋值，按顺序
	//person4 := Person{"gaoyongbin", 12, School{"刘家堡一中", "刘家堡"}}
	//fmt.Println(person4)
	//
	//// 赋值，按关键字
	//person5 := Person{name:"xuhao", age:23}
	//fmt.Println(person5)
	//
	//// 结构体指针并且赋值
	//person2 := &Person{"sunling", 12, School{"刘家堡二中", "刘家堡"}}
	//fmt.Println("*********")
	//fmt.Println(&person2)
	//fmt.Println(person2)
	//fmt.Println(*person2)
	//fmt.Println(reflect.TypeOf(person2))
	//
	//var person22 *Person
	//fmt.Println(reflect.TypeOf(person22))
	//
	//fmt.Println("*********")
	//
	//var person222 = &Person{}
	//fmt.Println(reflect.TypeOf(person222))
	//
	//
	//// 结构体指针
	//var person3 = new(Person)
	//fmt.Println(person3)
	//
	//var t *string
	//fmt.Println(reflect.TypeOf(t))
	//name := "hhhhh"
	//fmt.Println(&name)
	//t = &name
	//fmt.Printf("%p" , &name)
	//fmt.Println(t)

	home := Home{
		name: "changping",
		miles: 12,
		person: [2]int{1, 2},
		food: []string{"water", "ddd"},
		relation: map[string]string{"name": "ddss"},
	}

	// 赋值给a
	a := home

	// 是否拷贝
	// 字符串是进行拷贝的
	home.name = "tiantongyuan"
	fmt.Println(a)
	fmt.Println(home)
	fmt.Println(&a.name)
	fmt.Println(&home.name)

	// 数字是进行拷贝了的
	a.miles = 999
	fmt.Println(a)
	fmt.Println(home)
	fmt.Println(&a.miles)
	fmt.Println(&home.miles)

	// 数组是进行拷贝的
	a.person[1] = 999
	fmt.Println(a)
	fmt.Println(home)

	fmt.Println(&a.person)
	fmt.Println(&home.person)

	// 切片是不会进行拷贝的
	a.food[0] = "eggs"
	fmt.Println(a)
	fmt.Println(home)

	fmt.Println(&a.food)
	fmt.Println(&home.food)

	// map是不会进行拷贝的
	a.relation["name"] = "sunlei"
	fmt.Println(a)
	fmt.Println(home)
	fmt.Println(&a.relation)
	fmt.Println(&home.relation)

	c := func1(10)
	// 函数的调用
	fmt.Println(c)

	data := [2]int{0, 1}
	result := func2(&data)
	fmt.Println(data)
	fmt.Print(result)

}

// 函数定义,参数类型是int，返回值也是int
func func1(a int) int {
	return a
}

// 关于函数参数的拷贝问题，传到函数的参数是值传递
// 如果想传递引用，需要使用指针
func func2(s *[2]int) *[2]int {
	s[1] = 9999
	fmt.Println(s)
	return s
}
