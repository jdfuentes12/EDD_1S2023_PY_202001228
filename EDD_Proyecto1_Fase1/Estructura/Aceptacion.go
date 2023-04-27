package Estructura

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type Pila_Aceptacion struct {
	Primero  *Nodo_Aceptacion
	Longitud int
}

func (p *Pila_Aceptacion) estaVacia() bool {
	if p.Longitud == 0 {
		return true
	} else {
		return false
	}
}

func (p *Pila_Aceptacion) Push(estado string, fecha string, hora string) {
	if p.estaVacia() {
		p.Primero = &Nodo_Aceptacion{estado, fecha, hora, nil}
	} else {
		nuevoNodo := &Nodo_Aceptacion{estado, fecha, hora, p.Primero}
		p.Primero = nuevoNodo
	}
	p.Longitud++
}

func (p *Pila_Aceptacion) MostrarPilaAceptacion() {
	aux := p.Primero
	for aux != nil {
		fmt.Println(aux.estado, aux.fecha, aux.hora)
		aux = aux.siguiente
	}
}

func (p *Pila_Aceptacion) Graficar() {
	nombre_archivo := "./Aceptacion.dot"
	nombre_imagen := "./Aceptacion.png"
	contador := 0
	texto := "digraph G {\n"
	texto += "node [shape= record, height = 0.5]\n"
	aux := p.Primero
	if aux != nil {
		//for de recorrido de la pila
		for aux != nil {
			texto += "nodo" + strconv.Itoa(contador) + "[label = \"<f0>" + aux.estado + "|<f1> " + aux.fecha + "|<f2> " + aux.hora + "\"]\n"
			aux = aux.siguiente
			contador++
		}

		texto += "empty [label=\"Pila\"]\n"
		texto += "nodo0:f0 -> empty\n"

		for i := 1; i < contador; i++ {
			texto += "nodo" + strconv.Itoa(i) + ":f0 -> nodo" + strconv.Itoa(i-1) + ":f1\n"
		}
	} else {
		texto += "nodo0[label = \"Lista Vacia \"]\n"
	}
	texto += "}"
	crearArchivo(nombre_archivo)
	escribirArchivoDot(texto, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}

func crearArchivo(nombre_archivo string) {
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

func escribirArchivoDot(contenido string, nombre_archivo string) {
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

func ejecutar(nombre_imagen string, archivo_dot string) {
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tjpg", archivo_dot).Output()
	mode := 0777
	_ = ioutil.WriteFile(nombre_imagen, cmd, os.FileMode(mode))
}
