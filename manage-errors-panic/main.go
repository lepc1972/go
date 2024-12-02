package main

import "fmt"

// panic example
func main() {
	fmt.Println("start")
	panic("panic")
	fmt.Println("end")
}
