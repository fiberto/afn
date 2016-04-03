package afn

import "fmt"

//Transicion define la letra para llegar al estado destino.
type Transicion map[Letra][]*Estado

//Estado indica un estado y sus transiciones del AFN.
type Estado struct {
	ID      string
	Inicial bool
	Final   bool
	δ       Transicion
}

//NewEstado se encarga de alojar un nuevo estado en memoria con el ID que recibe como parámetro.
func NewEstado(ID string) (e Estado) {
	e.ID = ID
	e.δ = make(Transicion)
	return
}

func (e *Estado) transicion(cadena string, pos int, recorrido string) { //Recibe la cadena, la posición y el reocorrido que ha tenido en el autómata
	if len(cadena) == pos { //Para cuando ya se ha recorrido toda la cadena
		if e.Final {
			mux.Lock() //Bloqueo para imprimir
			Validos++
			for i, letra := range cadena {
				if i != 0 {
					fmt.Print("\t")
				}
				fmt.Print("\t", string(letra))
			}
			fmt.Println("\n" + recorrido)
			mux.Unlock()
		}
		agregarRutina(-1) //goroutine terminada
		return
	}
	siguientes := e.δ[Letra(cadena[pos])] //Se busca si existe una transición con la letra
	pos++
	if len(siguientes) == 0 {
		//Si no existe transición termina el proceso
		agregarRutina(-1)
		return
	}
	for i, edo := range siguientes { //Si existe 1 o más caminos se sigue recorriendo la cadena
		if i != 0 {
			go edo.transicion(cadena, pos, recorrido+"\t->\t"+edo.ID) //Se manda a una goroutine
			agregarRutina(1)
		} else {
			edo.transicion(cadena, pos, recorrido+"\t->\t"+edo.ID)
		}
	}
}
