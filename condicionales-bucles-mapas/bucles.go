package main

import "fmt"

func main() {

	var suma int = 0

	for i := 0; i <= 10; i++ {
		if i%2 != 0 {
			suma = suma + i

		}

	}
	fmt.Println(suma)
	// especide de for each que recorre mapa
	var miMapa = map[string]string{
		"nombre":   "Juan",
		"apellido": "Perez",
		"edad":     "30",
	}
	for clave, valor := range miMapa {
		fmt.Println(clave, valor)
	}
	//uso de while con for de go
	var fruta string = ""

	for {
		if fruta == "manzana" {
			break
		}
		fmt.Println("¿Qué fruta es?")
		fmt.Scanln(&fruta)

	}

}
