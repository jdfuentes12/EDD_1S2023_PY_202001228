package Estructura

type Nodo_Cola struct {
	estudiate_cola *Estudiante_Cola
	siguiente      *Nodo_Cola
	antes          *Nodo_Cola
}

type Estudiante_Cola struct {
	nombre   string
	apellido string
	carne    int
	pasword  string
}
