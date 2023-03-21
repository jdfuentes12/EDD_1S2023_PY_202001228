package Estructura

type Nodo_ListaDoble struct {
	estudiate *Estudiante
	//registro  *Nodo_Sesion
	siguiente *Nodo_ListaDoble
	antes     *Nodo_ListaDoble
}

type Estudiante struct {
	nombre   string
	apellido string
	carne    int
	password string
}
