package Estructura

import (
	"fmt"
)

var listaDoble *ListaDoble = NewListaDoble()

type ListaCola struct {
	Inicio   *Nodo_Cola
	Final    *Nodo_Cola
	Longitud int
}

func (l *ListaCola) estaVacia() bool {
	if l.Longitud == 0 {
		return true
	} else {
		return false
	}
}

func (l *ListaCola) newNodo(estudiante *Estudiante_Cola) *Nodo_Cola {
	return &Nodo_Cola{estudiante, nil, nil}
}

func NewCola() *ListaCola {
	return &ListaCola{nil, nil, 0}
}

func (c *ListaCola) AgregarEstudiante(nombre string, apellido string, carne int, password string) {
	nuevoEstudiante := &Estudiante_Cola{nombre, apellido, carne, password}

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

func (c *ListaCola) EliminarEstudiante() {
	if c.estaVacia() {
		fmt.Println("La cola esta vacia")
	} else {
		if c.Inicio.siguiente == nil {
			c.Inicio = nil
			c.Final = nil
		} else {
			c.Inicio = c.Inicio.siguiente
			c.Inicio.antes = nil
		}
		c.Longitud--
	}
}

func (c *ListaCola) MostrarCola() {
	aux := c.Inicio
	validacion := false
	var opcion int
	for aux != nil {

		for aux != nil {
			fmt.Println("-------Hay en cola: ", c.Longitud, " estudiantes-------")
			fmt.Println("Estudiante Actual: ", aux.estudiate_cola.nombre, " ", aux.estudiate_cola.apellido)
			fmt.Println("1. Aceptar al estudiante")
			fmt.Println("2. Rehazar al estudiante")
			fmt.Println("3. Volver al menu")
			fmt.Println("Ingrese una opcion: ")

			fmt.Scan(&opcion)
			switch opcion {

			case 1:
				fmt.Println("El estudiante ha sido aceptado")
				fmt.Println()
				fmt.Println(aux.estudiate_cola.nombre, " ", aux.estudiate_cola.apellido, " ", aux.estudiate_cola.carne, " ", aux.estudiate_cola.pasword)
				listaDoble.AgregarEstudiante_ListaDoble(aux.estudiate_cola.nombre, aux.estudiate_cola.apellido, aux.estudiate_cola.carne, aux.estudiate_cola.pasword)
				aux = aux.siguiente
				c.EliminarEstudiante()
				break
			case 2:
				fmt.Println("El estudiante ha sido rechazado")
				fmt.Println()
				aux = aux.siguiente
				c.EliminarEstudiante()
				break
			case 3:
				validacion = true
				return
			}
		}
		if validacion {
			return
		} else {
			aux = aux.siguiente
		}
	}
}

func (c *ListaCola) MostrarListaDoble() {
	listaDoble.MostrarLista()
}
