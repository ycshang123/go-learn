package main

import "fmt"

type Person struct {
	name string
}

func changeName(p *Person) {
	p.name = "imooc"
}

func swap(a, b *int) {

	t := *a
	//	将a的地址改为b指向的值
	*a = *b
	//	将b地址的值改为a指向的值
	*b = t
}

func main() {

	p := Person{
		name: "ycshang",
	}
	changeName(&p)

	fmt.Println(p.name)

	//	指针的定义
	var po *Person
	po = &p
	fmt.Println(po.name)

	//	指针初始化的三种方法
	_ = &Person{}
	var emptyPerson Person
	_ = emptyPerson
	var _ = new(Person)
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
