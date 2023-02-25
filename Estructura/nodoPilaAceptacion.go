package Estructura

type Nodo_Acetacion struct {
	apectacion *Aceptacion
	siguiente  *Nodo_Acetacion
	antes      *Nodo_Acetacion
}

type Aceptacion struct {
	nombre string
	fecha  string
	hora   string
}
