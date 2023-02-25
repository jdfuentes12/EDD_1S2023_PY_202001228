package Estructura

type Nodo_ListaDoble struct {
	estudiate *Estudiante
	siguiente *Nodo_ListaDoble
	antes     *Nodo_ListaDoble
}

type Estudiante struct {
	nombre   string
	apellido string
	carne    int
	pasword  string
}
