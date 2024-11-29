package main

import "fmt" // paquete escribir leer salida standard

func presentarResultadosuma(suma func(int, int) int, a int, b int) {

	fmt.Println(suma(7, 12))
}

func presentarResultadoresta(resta func(int, int) int, a int, b int) {

	fmt.Println(resta(37, 12))
}

func main() {

	suma := func(a int, b int) int { // esta funcion estaba afuera de main, se convierte en anonima y se le asigna variable
		return a + b

	}

	resta := func(a int, b int) int { // esta funcion estaba afuera de main, se convierte en anonima y se le asigna variable
		return a - b

	}

	//fmt.Println(suma(5, 10))
	presentarResultadosuma(suma, 5, 10)
	//fmt.Println(resta(5, 10))
	presentarResultadoresta(resta, 5, 10)

}
