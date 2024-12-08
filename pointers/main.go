package main

import "fmt"

func main() {
	fmt.Println("Pointers...")
	myNumber := 23
	p := &myNumber
	fmt.Println("Value of actual var is  ", myNumber)
	fmt.Println("Address of pointer is ", p)
	fmt.Println("Value of pointer is ", *p)
	fmt.Println("###############################")

	myNumber1 := 46
	p = &myNumber1
	fmt.Println("Value of actual var is  ", myNumber1)
	fmt.Println("Address of pointer is ", p)
	fmt.Println("Value of pointer is ", *p)
	fmt.Println("###############################")

	var a = 1
	fmt.Println("value of a is", a)
	p = &a
	fmt.Println("address of a is", p)
	fmt.Println("value of a is", *p)

}
