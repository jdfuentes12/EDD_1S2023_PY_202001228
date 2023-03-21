package Estructura

type Pila_Session struct {
	Primero  *Nodo_Sesion
	Longitud int
}

func (p *Pila_Session) estaVacia() bool {
	if p.Longitud == 0 {
		return true
	} else {
		return false
	}
}

func NewNodoPila(fecha string, hora string, identificador string) *Nodo_Sesion {
	return &Nodo_Sesion{fecha, hora, identificador, nil}
}

func NewPilaSession() *Pila_Session {
	return &Pila_Session{nil, 0}
}

func (p *Pila_Session) InicioSesion(fecha string, hora string, identificador string) {
	if p.estaVacia() {
		p.Primero = &Nodo_Sesion{fecha, hora, identificador, nil}
		p.Longitud++
	} else {
		nuevoNodo := &Nodo_Sesion{fecha, hora, identificador, p.Primero}
		p.Primero = nuevoNodo
		p.Longitud++
	}
}
