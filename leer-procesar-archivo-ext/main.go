package main

import (
	"fmt"
	"os"
)

// estructura alumno
type Alumno struct {
	Nombre string
	Notas  string
}

func leercsv(nombreArchivo string) ([]Alumno, error) {
	// abrir archivo
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, err
	}
	defer archivo.Close()

	// leer archivo
	var alumnos []Alumno
	for {
		var alumno Alumno
		_, err := fmt.Fscanf(archivo, "%s,%s\n", &alumno.Nombre, &alumno.Notas)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}
		alumnos = append(alumnos, alumno)
	}

	return alumnos, nil

}

func main() {
	// leer archivo
	alumnos, err := leercsv("notas.csv")
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	// imprimir alumnos
	for _, alumno := range alumnos {
		fmt.Println(alumno.Nombre, alumno.Notas)
	}

}
