package Estructura

import (
	"fmt"
)

type ListaDoble struct {
	Inicio   *Nodo_ListaDoble
	Final    *Nodo_ListaDoble
	Longitud int
}

func (l *ListaDoble) estaVacia() bool {
	if l.Longitud == 0 {
		return true
	} else {
		return false
	}
}

func (l *ListaDoble) newNodo(estudiante *Estudiante) *Nodo_ListaDoble {
	return &Nodo_ListaDoble{estudiante, nil, nil}
}

func NewListaDoble() *ListaDoble {
	return &ListaDoble{nil, nil, 0}
}

func (c *ListaDoble) AgregarEstudiante_ListaDoble(nombre string, apellido string, carne int, password string) {
	nuevoEstudiante := &Estudiante{nombre, apellido, carne, password}

	fmt.Println("Estoy aca")

	if c.estaVacia() {
		nuevoNodo := c.newNodo(nuevoEstudiante)
		c.Inicio = nuevoNodo
		c.Final = nuevoNodo
		fmt.Println("Estoy aca x2")

	} else {
		nuevoNodo := c.newNodo(nuevoEstudiante)
		if c.Final.antes == nil {
			nuevoNodo.antes = c.Inicio
			c.Inicio.siguiente = nuevoNodo
			c.Final = nuevoNodo
		} else {
			c.Final.siguiente = nuevoNodo
			nuevoNodo.antes = c.Final
			c.Final = nuevoNodo
		}
		fmt.Println("Estoy aca x3")
	}
	c.Longitud++
	fmt.Println(c.Longitud)
}

func (c *ListaDoble) MostrarLista() {
	aux := c.Inicio
	fmt.Println("************ Lista de Estudiantes ************")
	fmt.Println(c.Longitud)
	for aux != nil {
		fmt.Println("Nombre: ", aux.estudiate.nombre, " ", aux.estudiate.apellido, " Carne: ", aux.estudiate.carne)
		fmt.Println("**********************************************")
		aux = aux.siguiente
	}
}
