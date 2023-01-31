package main

import "fmt"

/*
*
常量类型只可以定义bool、数值(整数、浮点、复数)和字符串
声明的变量没有强制使用要求
显示指定类型的时候，必须确保常量左右类型一致
*/
func main() {
	//	常量:在定义时就指定的值，不能修改，常量尽量全部大写
	const PI float32 = 3.1415926

	const (
		x = 16
		y
		z = "hello"
		f
	)
	fmt.Println(x, y, z, f)

}
