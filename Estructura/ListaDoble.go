package Estructura

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"time"
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

func newNodo(estudiante *Estudiante) *Nodo_ListaDoble {
	//var sesion *Nodo_Sesion =
	return &Nodo_ListaDoble{estudiante, nil, nil}
}

func NewListaDoble() *ListaDoble {
	return &ListaDoble{nil, nil, 0}
}

func (c *ListaDoble) InicioSesion(carne string, password string) bool {
	aux := c.Inicio

	carne1, err := strconv.Atoi(carne)
	if err == nil {
		fmt.Print("")
	}

	for aux != nil {
		if aux.estudiate.carne == carne1 && aux.estudiate.password == password {
			fmt.Println()
			fmt.Println("Bienvenido a Edd GoDrive\n", aux.estudiate.nombre, " ", aux.estudiate.apellido)
			//hora := horaInicio()
			//fecha := fechaInicio()
			//aux.registro = sesion.InicioSesion(fecha, hora)

			return true
		}
		aux = aux.siguiente
	}
	return false
}

func (c *ListaDoble) AgregarEstudiante_ListaDoble(nombre string, apellido string, carne int, password string) {
	nuevoEstudiante := &Estudiante{nombre, apellido, carne, password}

	if c.estaVacia() {
		nuevoNodo := newNodo(nuevoEstudiante)
		c.Inicio = nuevoNodo
		c.Final = nuevoNodo

	} else {
		nuevoNodo := newNodo(nuevoEstudiante)
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
	c.Ordenamiento()
	c.CrearJSON()
	c.GraficarListaDoble()
	aux := c.Inicio
	fmt.Println("************ Lista de Estudiantes ************")
	for aux != nil {
		fmt.Println("Nombre: ", aux.estudiate.nombre, " ", aux.estudiate.apellido, " Carne: ", aux.estudiate.carne)
		fmt.Println("**********************************************")
		aux = aux.siguiente
	}
}

// marca error por que no se esta usando...
func horaInicio() string {
	horaActual := time.Now()
	return horaActual.Format("15:04:05")
}

func fechaInicio() string {
	fechaActual := time.Now()
	return fechaActual.Format("2006-01-02")
}

func (l *ListaDoble) Ordenamiento() {
	aux := l.Inicio
	for i := 0; i < l.Longitud; i++ {
		aux2 := aux.siguiente
		for j := i + 1; j < l.Longitud; j++ {
			if aux.estudiate.carne > aux2.estudiate.carne {
				temp := aux.estudiate
				aux.estudiate = aux2.estudiate
				aux2.estudiate = temp
			}
			aux2 = aux2.siguiente
		}
		aux = aux.siguiente
	}

}

func (c *ListaDoble) GraficarListaDoble() {

	nombre_archivo := "./ListaDoble.dot"
	nombre_imagen := "./ListaDoble.png"

	contador := 0
	aux := c.Inicio

	var cadena string
	cadena = "digraph G{\n"
	cadena += "rankdir=LR;\n"
	cadena += "node [shape=cicle];\n"
	if aux != nil {

		for aux != nil {
			cadena += strconv.Itoa(contador) + "[label = \"" + aux.estudiate.nombre + " " + aux.estudiate.apellido + "\n" + strconv.Itoa(aux.estudiate.carne) + "\"];\n"
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
	crearArchivoLista(nombre_archivo)
	escribirArchivoDotLista(cadena, nombre_archivo)
	ejecutarLista(nombre_imagen, nombre_archivo)
}

func crearArchivoLista(nombre_archivo string) {
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

func escribirArchivoDotLista(contenido string, nombre_archivo string) {
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

func ejecutarLista(nombre_imagen string, archivo_dot string) {
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tjpg", archivo_dot).Output()
	mode := 0777
	_ = ioutil.WriteFile(nombre_imagen, cmd, os.FileMode(mode))
}

func (c *ListaDoble) CrearJSON() {
	contenido := "{\n"
	contenido += "	\"Alumnos\": [\n"
	aux := c.Inicio
	for i := 0; i < c.Longitud; i++ {
		contenido += "		{\n"
		contenido += "			\"Carnet\": \"" + strconv.Itoa(aux.estudiate.carne) + "\",\n"
		contenido += "			\"Nombre\": \"" + aux.estudiate.nombre + " " + aux.estudiate.apellido + "\",\n"
		contenido += "			\"Password\": \"" + aux.estudiate.password + "\",\n"
		contenido += "			\"Carpeta_Raiz\": \"" + "/" + "\"\n"
		contenido += "		},\n"
		aux = aux.siguiente
	}
	contenido += "]\n"
	contenido += "}\n"
	CrearArchivoJSON(string(contenido), "ReporteJSON.json")
}

func CrearArchivoJSON(contenido string, nameArchivo string) {
	archivo, err := os.Create(nameArchivo)
	if err != nil {
		fmt.Println(err)
	}
	_, err = archivo.WriteString(contenido)
	if err != nil {
		fmt.Println(err)
	}
	archivo.Close()
}
