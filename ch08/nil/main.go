package main

import "fmt"

type Person struct {
	name string
	age  int
	f    *int
}

func main() {
	/**
	不同类型的数据零值不同
	bool false
	numbers 0
	string ""
	pointer nil
	slice nil
	channel、interface、function nil
	*/

	p1 := Person{
		name: "test",
		age:  12,
	}

	p2 := Person{
		name: "test",
		age:  12,
	}

	if p1 == p2 {
		fmt.Println("yes")
	}

	//	slice的默认值
	var _ []Person             //nil slice
	var p4 = make([]Person, 0) //empty slice
	if p4 == nil {
		fmt.Println("nil slice")
	}

	var m2 = make(map[string]string, 0)
	if m2 == nil {
		fmt.Println("nil map")

	}

}
