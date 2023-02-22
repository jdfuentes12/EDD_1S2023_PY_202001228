package main

import (
	h "Tda/Estructura"
	"fmt"
)

func main() {
	/*opcion := 0
	salir := false

	for !salir {
		fmt.Println("1. Administrador")
		fmt.Println("2. Cliente")
		fmt.Println("3. Salir")
		fmt.Print("Elige una opcion: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			menuAdministrador()
		case 2:
			menuCliente()
		case 3:
			fmt.Println("Saliendo de la aplicacion")
			salir = true
		}
	}
	*/
	listaDoble := h.NewLista()
	listaDoble.InsertarFinal("Jose", "Perez", 201801234, "1234")
	listaDoble.InsertarFinal("Maria", "Perez", 201801234, "1234")
	listaDoble.MostrarConsola()
}

func menuAdministrador() {
	opcion := 0
	salir := false
	for !salir {
		fmt.Println("----------------Esta en el Menu Administrador----------------")
		fmt.Println("1. Ver estudiantes pedientes")
		fmt.Println("2. Ver estudiantes del sistema")
		fmt.Println("3. Registrar nuevo estudiante")
		fmt.Println("4. Carga masiva de estudiantes")
		fmt.Println("5. Salir")

		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			fmt.Println("pendiente")

		case 2:
			fmt.Println("pendiente")
		case 3:
			fmt.Println("pendiente")
		case 4:
			fmt.Println("pendiente")
		case 5:
			fmt.Println("Saliendo de la aplicacion")
			salir = true
		default:
			fmt.Println("Opcion no valida")
		}
	}
}

func menuCliente() {
	fmt.Println("Esta en el menu cliente")
	fmt.Println("-----------------------")

}
