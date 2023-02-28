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
	return &Nodo_ListaDoble{estudiante, nil, nil, nil}
}

func NewListaDoble() *ListaDoble {
	return &ListaDoble{nil, nil, 0}
}

func (c *ListaDoble) AgregarEstudiante_ListaDoble(nombre string, apellido string, carne int, password string) {
	nuevoEstudiante := &Estudiante{nombre, apellido, carne, password}

	if c.estaVacia() {
		nuevoNodo := c.newNodo(nuevoEstudiante)
		c.Inicio = nuevoNodo
		c.Final = nuevoNodo

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
	}
	c.Longitud++
}

func (c *ListaDoble) MostrarLista() {
	aux := c.Inicio
	fmt.Println("************ Lista de Estudiantes ************")
	for aux != nil {
		fmt.Println("Nombre: ", aux.estudiate.nombre, " ", aux.estudiate.apellido, " Carne: ", aux.estudiate.carne)
		fmt.Println("**********************************************")
		aux = aux.siguiente
	}
}

func (c *ListaCola) ValidarUsuario(carne int, password string) bool {
	aux := c.Inicio
	for aux != nil {
		if aux.estudiate_cola.carne == carne && aux.estudiate_cola.pasword == password {
			return true
		}
		aux = aux.siguiente
	}
	return false
}
