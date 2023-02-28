package Estructura

import (
	"fmt"
)

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

func (p *Pila_Session) Push(estado string, fecha string, hora string) {
	if p.estaVacia() {
		p.Primero = &Nodo_Sesion{fecha, hora, nil}
	} else {
		nuevoNodo := &Nodo_Sesion{fecha, hora, p.Primero}
		p.Primero = nuevoNodo
	}
}

func (p *Pila_Session) MostrarPilaSession() {
	aux := p.Primero
	for aux != nil {
		fmt.Println(aux.fecha, aux.hora)
		aux = aux.siguiente
	}
}
