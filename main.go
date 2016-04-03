package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"./afn"
)

func main() {
	contenido := leerArchivo(os.Args[1]) //Leer parámetro que se pasa como argumento
	automata := afn.NewANF(contenido)    //Usar el archivo para llenar el autómata
	var cadena string
	for {
		fmt.Println("Ingrese la cadena a validar")
		fmt.Scanln(&cadena)
		automata.Evaluar(cadena) //Evaluar se encarga de verificar la cadena
		for {                    //for infinito para saber cuando no hay goroutines en ejecución
			if afn.Rutinas == 0 {
				if afn.Validos == 0 {
					fmt.Println("No es una palabra válida")
				} else {
					fmt.Println("Número de caminos: ", afn.Validos)
				}
				break
			}
		}
	}
}

func leerArchivo(file string) string {
	contenido, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("No existe el archivo")
		os.Exit(1)
	}
	return string(contenido)
}
