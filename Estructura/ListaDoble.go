package Estructura

import (
	"fmt"
)

type ListaDoble struct {
	inicio   *Nodo
	longitud int
}

func (l *ListaDoble) InsertarFinal(nombre string, apellido string, carne int, password string) {
	nuevoEstudiante := &Estudiante{nombre, apellido, carne, password}

	if l.estaVacia() {
		l.inicio = l.newNodo(nuevoEstudiante)
		l.longitud++
	} else {
		aux := l.inicio
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		aux.siguiente = l.newNodo(nuevoEstudiante)
		aux.siguiente.siguiente = nil
		aux.siguiente.antes = aux
	}
}

func (l *ListaDoble) newNodo(estudiante *Estudiante) *Nodo {
	return &Nodo{estudiante, nil, nil}
}

func (l *ListaDoble) estaVacia() bool {
	if l.longitud == 0 {
		return true
	} else {
		return false
	}
}

// contructor de la lista dobl
func NewLista() *ListaDoble {
	lista := new(ListaDoble)
	lista.inicio = nil
	return lista
}

func (l *ListaDoble) MostrarConsola() {
	aux := l.inicio
	for aux != nil {
		fmt.Print(aux.estudiante.nombre," ")
		fmt.Println(aux.estudiante.apellido)
		aux = aux.siguiente
	}
}
