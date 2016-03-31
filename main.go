package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"./afn"
)

func main() {
	contenido := leerArchivo(os.Args[1])
	afn := afn.NewANF(contenido)
	fmt.Println(afn)
}

func leerArchivo(file string) string {
	contenido, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("No existe el archivo")
		os.Exit(1)
	}
	return string(contenido)
}
