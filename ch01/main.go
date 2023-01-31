package main

import "fmt"

/*
*
Go的变量使用注意细节:
变量必须先定义才能使用
go语言是静态语言，要求变量的类型和赋值类型一致
变量名不能冲突,全局变量和局部变量名称可以冲突，打印优先级 局部变量>全局变量
简介变量的定义方式不能用于全局
变量有默认值和Java相同
局部变量定义之后一定要使用，否则会编译报错
*/

// main 函数外面可以定义全局变量
var (
	name = "ycshang"
	age  = 22
	sex  bool
)

func main() {
	//	go是静态语言，静态语言和动态语言相比变量差异很大
	//	1、变量必须先定义后使用
	//	2、变量必须有类型
	//	3、类型定义下来后不能被改变
	//	定义变量的方式:
	//	1、先声明变量名、类型后赋值
	var age int
	age = 1
	//	局部变量定义后必须使用，否则编译时报错
	fmt.Println(age)
	//	2、直接定义变量名并赋值
	var age1 = 2
	fmt.Println(age1)
	//	3、使用:赋值(简介变量，不能在全局变量时声明)
	age2 := 3
	fmt.Println(age2)

	//	2、多变量定义
	var user, age3, address = "ycshang", 22, "江苏徐州"
	fmt.Println(user, age3, address)

	//	匿名变量
	var _ int

}
