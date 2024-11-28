package main

import "fmt"

func main() {

	var suma int = 0

	for i := 0; i <= 100; i++ {
		suma = suma + i
	}
	fmt.Println(suma)

}
