package main

import "fmt"

func mapa() {
	var miMapa = map[string]string{
		"nombre":   "Juan",
		"apellido": "Perez",
		"edad":     "30",
	}
	fmt.Println(miMapa["nombre"], miMapa["edad"])
	//agregar un elemento al mapa
	miMapa["cc"] = "6789"
	fmt.Println(miMapa)
	//borrar un elemento del mapa
	delete(miMapa, "edad")
	fmt.Println(miMapa)
}
