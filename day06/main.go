package main

import "fmt"

func main() {
	// switch语句
	// 基本的格式：switch 表达式 {case 值1}
	//var age int
	//fmt.Scan(&age)
	//switch age {
	//case 18:
	//	fmt.Println("您为18岁")
	//case 25:
	//	fmt.Println("您为25岁")
	//case 30:
	//	fmt.Println("您为30岁")
	//default:
	//	fmt.Println("不满足条件")
	//}
	//
	//// 相同的功能可以使用if else语句实现，而且更加的灵活
	//if age < 18 {
	//	fmt.Println("您小于18岁")
	//}else if age <= 20 {
	//	if age > 18 {
	//		fmt.Println("大于18岁并且小于20岁")
	//	} else {
	//		fmt.Println("大于20岁")
	//	}
	//}

	// for循环
	// 1. 死循环,可以写恒成立的条件，也可以不写条件
	//for {
	//	fmt.Println("循环开始")
	//	time.Sleep(time.Second * 5)
	//}
	s := 2
	fmt.Println("开始循环")
	fmt.Println(&s)
	fmt.Println(s)
	for s > 1 {
		fmt.Println("大于1")
		s -= 1
	}
	fmt.Println("结束循环")

	// iota测试
	const (
		a = iota
		b
		c
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	const v = 2
	const n = iota
	fmt.Println(n)
	fmt.Println(v)

	for i:=1; i < 10; i++ {
		fmt.Println(i)
	}

	// for practice
	// 1. 猜数字
	var num int
	for {
		fmt.Print("请输入一个数字:")
		fmt.Scan(&num)
		if num == 66 {
			fmt.Println("猜对了")
			break
		} else if num < 66 {
			fmt.Println("猜小了")
		} else {
			fmt.Println("猜大了")
		}
	}

	// 2. 1-100内的所有整数
	for i:= 1; i < 101; i++ {
		fmt.Println(i)
	}

	// 3.10以内，除7以外的整数
	for h := 0; h <=10; h++ {
		if h == 7 {
			continue
		}else {
			fmt.Println(h)
		}
	}

	// 4. 1-100之内所有的奇数
	for j:= 1; j <= 100; j++ {
		if j % 2 != 0 {
			fmt.Println(j)
		}
	}
	
	// 5. 输出1-100之内所有整数的和
	var sum int
	for k := 1; k <= 100; k++ {
		sum += k
	}
	fmt.Println(sum)

	// 6. 输出10-1所有的整数
	for l:= 10; l >=1; l-- {
		fmt.Println(l)
	}

	// for循环打标签
	f1: for i:=1; i<10; i++ {
		for j:=1; j < 3; j++ {
			if j == 2 {
				fmt.Println(i, j)
				continue f1
			} else {
				fmt.Println(i, j)
			}
		}
	}

	f2: for i:=1; i<10; i++ {
		for j:=1; j < 3; j++ {
			if j == 2 {
				fmt.Println(i, j)
				break f2
			} else {
				fmt.Println(i, j)
			}
		}
	}
}
