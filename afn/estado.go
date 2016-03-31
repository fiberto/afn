package afn

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
