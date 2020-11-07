package main

import "fmt"

func main() {

	// 声明变量加条件
	// for循环的变量在外部是不能被使用的，只能在for循环内部才能使用
	//for number := 1; number < 5;number ++ {
	//	fmt.Println(number)
	//	fmt.Println("掉")
	//}
	/*
		var age int = 66
		var flag bool = true
		for flag {
			var userInput int
			fmt.Scan(&userInput)
			if userInput > age {
				fmt.Println("大了")
			} else if userInput < age {
				fmt.Println("小了")
			} else {
				fmt.Println("正确")
				flag = false
			}
		}
	*/

	// 1-100的整数
	/*
		for i:=1; i < 100; i++ {
			fmt.Println(i)
		}
	*/
	/*
		for i:=1; i <100; i= i+2 {
			fmt.Println(i)
		}
	*/

	// continue 关键字，在for循环中使用，遇到continue关键字的时候，会跳过当前的循环，开始下一个循环
	/*
		for i:=1; i < 10; i++ {
			if i %2 == 0 {
				continue
			}
			fmt.Println(i)
		}
	*/

	// switch语句
	/* 基本结构: switch 表达式(变量) {
		case 值1: 操作
		case 值2: 操作
		default:都不满足的时候执行
	}
	*/
	/*
	switch 1 + 1 {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	}
	 */

	var flag bool = true
	switch flag {
	case true:
		fmt.Println("真")
	case false:
		fmt.Println("假")
	default:
		fmt.Println("不知道")
	}
	// 注意事项:只有在值的类型相同的时候才能进行比较，这一点和Go之前的值比较是一致的

	// continue关键字，当循环遇到continue关键字的时候，会跳过当前的循环，开始下一次的循环

	//f1: for i := 1; i < 10; i++ {
	//		for j := 1; j < 10; j ++ {
	//			if j == 2 {
	//				break f1
	//			}
	//			fmt.Println(i, j)
	//	}
	//}
	//

//	var name string
//	fmt.Print("请输入姓名:")
//	fmt.Scanln(&name)
//
//	if name == "wangdong" {
//		goto SVIP
//	}else if name == "ps" {
//		goto VIP
//	}
//	fmt.Println("预约")
//VIP:
//	fmt.Println("等号")
//SVIP:
//	fmt.Println("看病")

	var name, address, action string

	fmt.Print("请输入姓名:")
	fmt.Scanln(&name)

	fmt.Print("请输入地址:")
	fmt.Scanln(&address)

	fmt.Print("请输入行为:")
	fmt.Scanln(&action)

	result := fmt.Printf("我叫%s,现在在%s%s", name, address, action)
	fmt.Println(result)
}
