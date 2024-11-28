package main

import "fmt"

func main() {
	var edad int = 22
	var peso float32 = 70.5
	//var nombre string = "Juan"
	var nombre = "Juan"
	var apellido string = "Perez"

	fmt.Println("Mi nombre es", nombre, apellido, "tengo", edad, "años y peso", peso, "kilos")

	if edad >= 18 {
		fmt.Println("Soy", nombre, "y soy mayor de edad")
	} else {
		fmt.Println("Soy", nombre, "y soy menor de edad")
	}
	//if anidado
	if edad >= 18 {
		fmt.Println("Soy", nombre, "y soy mayor de edad")
	} else if edad >= 13 {
		fmt.Println("Soy", nombre, "y soy adolescente")
	} else {
		fmt.Println("Soy", nombre, "y soy menor de edad")
	}
	//switch
	switch edad {
	case 18:
		fmt.Println("Soy", nombre, "y tengo 18 años")
	case 20:
		fmt.Println("Soy", nombre, "y tengo 20 años")
	case 23:
		fmt.Println("Soy", nombre, "y tengo 22 años")
	default:
		fmt.Println("Soy", nombre, "y tengo", edad, "años")
	}
	//switch con condicion
	switch {
	case edad >= 18:
		fmt.Println("Soy", nombre, "y soy mayor de edad")
	case edad >= 13:
		fmt.Println("Soy", nombre, "y soy adolescente")
	default:
		fmt.Println("Soy", nombre, "y soy menor de edad")
	}

}
