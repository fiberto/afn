package afn

import (
	"fmt"
	"os"
	"strings"
)

//Letra representa una letra del alfabeto.
type Letra string

//AFN define los atributos que contiene un Autómata Finito No Determinista.
type AFN struct {
	Q       []Estado           //Estados del AFN
	Estados map[string]*Estado //Mapa para facilitar búsqueda de estados
	Σ       map[string]Letra   //Alfabeto del AFN
	S       *Estado            //Estado inicial
	F       []*Estado          //Conjunto de estados finales
}

func (afn *AFN) init(n int) {
	//Se encarga de inicializr los mapas del alfabeto y estados para el AFN
	afn.Q = make([]Estado, n)
	afn.Estados = make(map[string]*Estado)
	afn.Σ = make(map[string]Letra)
}

func (afn *AFN) asignarEstados(estados []string) {
	//Agrega cada uno de los estados al Slice de estados y un mapa de estados para facilitar la búsqueda
	for i, estado := range estados {
		nuevoEstado := NewEstado(estado)
		afn.Q[i] = nuevoEstado
		afn.Estados[estado] = &afn.Q[i]
	}
}

func (afn *AFN) asignarAlfabeto(alfabeto []string) {
	//Asigna en un mapa cada una de las letras del alfabeto
	for _, letra := range alfabeto {
		afn.Σ[letra] = Letra(letra)
	}
}

func (afn *AFN) asignarInicial(inicial string) {
	//Indica que estado es el inicial
	var val bool
	afn.S, val = afn.Estados[inicial]
	if !val {
		fmt.Println("No existe el estado inicial:", inicial)
		os.Exit(1)
	}
	afn.S.Inicial = true
}

func (afn *AFN) asignarFinales(finales []string) {
	//Indica los estados finales del autómata
	for _, final := range finales {
		_, val := afn.Estados[final]
		if !val {
			fmt.Println("No existe el estado final:", final)
			os.Exit(1)
		}
		afn.Estados[final].Final = true
		afn.F = append(afn.F, afn.Estados[final])
	}
}

func (afn *AFN) asignarTransiciones(transiciones []string) {
	for _, transicion := range transiciones {
		split := strings.Split(transicion, ",")

		//split[0] contiene el estado origen
		estado, val := afn.Estados[split[0]] //Se verifica que exista el estado
		if !val {
			fmt.Println("No existe el estado para la transición:", transicion)
			os.Exit(1)
		}

		//split[1] contiene la letra para la transicion
		letra, val := afn.Σ[split[1]] //Se verifica que exista la letra en el alfabeto
		if !val {
			fmt.Println("No existe la letra en el alfabeto:", split[1])
			os.Exit(1)
		}

		//split[2] contiene el estado destino
		sig, val := afn.Estados[split[2]] //Se verifica que exista el estado
		if !val {
			fmt.Println("No existe el estado para la transición:", transicion)
			os.Exit(1)
		}

		estado.δ[letra] = append(estado.δ[letra], sig) //Se asigna la transición en el mapa del estado
	}
}

//NewANF es el constructor encargado de iniciar un AFN con base en el contenido leído del archivo.
func NewANF(contenido string) (afn AFN) {
	lineas := strings.Split(contenido, "\n")
	if len(lineas[(len(lineas)-1)]) == 0 {
		lineas = lineas[:len(lineas)-1] //Se elimina el último salgo de línea del archivo
	}

	//lineas[0] contiene los estados del AFN
	estados := strings.Split(lineas[0], ",")
	afn.init(len(estados))
	afn.asignarEstados(estados)

	//lineas[1] contiene el alfabeto del AFN
	alfabeto := strings.Split(lineas[1], ",")
	afn.asignarAlfabeto(alfabeto)

	//lineas[2] contiene el estado inicial del AFN
	afn.asignarInicial(lineas[2])

	//lineas[3] contiene el conjunto de estados finales para el AFN
	finales := strings.Split(lineas[3], ",")
	afn.asignarFinales(finales)

	//lineas[4:] contiene todas las transiciones para el AFN
	afn.asignarTransiciones(lineas[4:])
	return
}
