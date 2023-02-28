package Estructura

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"time"
)

var listaDoble *ListaDoble = NewListaDoble()
var aceptacion *Pila_Aceptacion = &Pila_Aceptacion{}

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
				listaDoble.AgregarEstudiante_ListaDoble(aux.estudiate_cola.nombre, aux.estudiate_cola.apellido, aux.estudiate_cola.carne, aux.estudiate_cola.pasword)
				hora := hora()
				fecha := fecha()
				aceptacion.Push("Aceptado", fecha, hora)
				aux = aux.siguiente
				c.EliminarEstudiante()
				break
			case 2:
				fmt.Println("El estudiante ha sido rechazado")
				fmt.Println()
				aux = aux.siguiente
				c.EliminarEstudiante()
				hora := hora()
				fecha := fecha()
				aceptacion.Push("Rechazado", fecha, hora)
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
	fmt.Println()
	//aceptacion.MostrarPilaAceptacion()
	aceptacion.Graficar()
}

func hora() string {
	horaActual := time.Now()
	return horaActual.Format("15:04:05")
}

func fecha() string {
	fechaActual := time.Now()
	return fechaActual.Format("2006-01-02")
}

func (c *ListaCola) GraficarCola() {
	nombre_archivo := "./Cola.dot"
	nombre_imagen := "./Cola.png"

	contador := 0
	aux := c.Inicio

	var cadena string
	cadena = "digraph G{\n"
	cadena += "rankdir=LR;\n"
	cadena += "node [shape=cicle];\n"
	if aux != nil {

		for aux != nil {
			cadena += strconv.Itoa(contador) + "[label = \"" + aux.estudiate_cola.nombre + " " + aux.estudiate_cola.apellido + "\n" + strconv.Itoa(aux.estudiate_cola.carne) + "\"];\n"
			contador++
			aux = aux.siguiente
		}
		for i := 0; i < c.Longitud; i++ {
			cadena += strconv.Itoa(i) + "->" + strconv.Itoa(i+1) + ";\n"
		}
	} else {
		cadena += "0[label = \"Cola Vacia\"];\n"
	}

	cadena += "}"
	crearArchivoCola(nombre_archivo)
	escribirArchivoDotCola(cadena, nombre_archivo)
	ejecutarCola(nombre_imagen, nombre_archivo)
}

func crearArchivoCola(nombre_archivo string) {
	var _, err = os.Stat(nombre_archivo)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
		var file, err = os.Create(nombre_archivo)
		if err != nil {
			return
		}
		defer file.Close()
	}
	fmt.Println("Archivo creado exitosamente")
}

func escribirArchivoDotCola(contenido string, nombre_archivo string) {
	var file, err = os.OpenFile(nombre_archivo, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	// Escribe algo de texto linea por linea
	_, err = file.WriteString(contenido)
	if err != nil {
		return
	}
	// Salva los cambios
	err = file.Sync()
	if err != nil {
		return
	}
	fmt.Println("Archivo actualizado existosamente.")
}

func ejecutarCola(nombre_imagen string, archivo_dot string) {
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tjpg", archivo_dot).Output()
	mode := 0777
	_ = ioutil.WriteFile(nombre_imagen, cmd, os.FileMode(mode))
}
