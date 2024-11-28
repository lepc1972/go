package main

import (
	"fmt"

	"rsc.io/quote" //modulo
)

func main() {
	fmt.Println(quote.Go())

	fmt.Println("hola mundo")

	// variables string

	var nombrePersona string = "Juan"
	var apellido string = "Perez"
	fmt.Println(nombrePersona, apellido)

	// variables numericas

	var edad int = 30
	fmt.Println(edad)

	edad = 52
	fmt.Println(edad)

	//arreglos

	var frutas [3]string = [3]string{"manzana", "pera", "uva"}
	fmt.Println(frutas[2])

	//slice

	var paises []string = []string{"colombia", "mexico", "peru"}
	fmt.Println(paises[1])
	//agregar un elemento al slice
	paises = append(paises, "argentina")
	fmt.Println(paises)

	//obtener rangos
	fmt.Println(len(paises))
	fmt.Println(paises[1:3]) // no incluye el 3
	fmt.Println(paises[:3])  // tres primeros elementos
	fmt.Println(paises[3:])  // trae el ultimo

}
