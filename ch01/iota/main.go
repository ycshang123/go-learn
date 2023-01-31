package main

import "fmt"

func main() {
	//iota 特殊常量 可以认为是一个可以被编译器修改的常量
	//iota 从 初始0 开始，下面的变量依次递增
	//中断了iota 内部依然会增加计数器，跟上一个常量的赋值没关系
	//iota 可以简化const类型的定义
	//每次出现const时 iota 数值归零
	const (
		ERR1 = iota
		ERR2
		ERR25 = " ha"
		ERR3  = 100
		ERR4  = iota
	)
	fmt.Println(ERR1, ERR2, ERR25, ERR3, ERR4)

}
