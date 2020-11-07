package api

import "fmt"

func Baidu() {
	fmt.Println("我是百度")
	var name string
	//var age int
	//fmt.Println("请输入用户名:")
	// 当使用Scan的时候，会提示用户进行输入，输入完成会有两个返回值
	// count:表示输入的个数，error:表示输入错误的时候，包含的错误信息
	// fmt.Scan要求输入的值的个数不满足要求的时候，会一直等待阻塞在那里
	//count, err := fmt.Scan(&name, &age)

	// fmt.Scanln只会等待回车，一旦输入了回车之后，输入就结束了
	//fmt.Scanln(&name)

	// fmt.Scanf表示格式化输入，可以进行填充，支持模板的方式
	fmt.Scanf("我叫%s", &name)
	//fmt.Println(count, err)
	fmt.Println("+++++++")
	fmt.Println(name)
}
