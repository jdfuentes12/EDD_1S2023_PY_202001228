package Estructura

type Nodo_Sesion struct {
	fecha         string
	hora          string
	identificador string
	siguiente     *Nodo_Sesion
}
