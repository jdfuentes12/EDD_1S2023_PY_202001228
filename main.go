package main

import (
	h "Tda/Estructura"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// menu general
func main() {
	opcion := 0
	salir := false

	for !salir {
		fmt.Println("----Bienvenido a la aplicacion EDD GoDrive----")
		fmt.Println("1. Administrador")
		fmt.Println("2. Cliente")
		fmt.Println("3. Salir")
		fmt.Println("Elige una opcion: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Println()
			seleccion_administrador()
		case 2:
			fmt.Println()
			menuCliente()
		case 3:
			fmt.Println("Saliendo de la aplicacion")
			salir = true
		}
	}
}

// validar credenciales del administrador
func seleccion_administrador() {
	fmt.Println("Verificando credenciales")
	usuario := ""
	password := ""
	fmt.Print("Usuario: ")
	fmt.Scan(&usuario)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	if usuario == "admin" && password == "admin" {
		menuAdministrador()
	} else {
		fmt.Println("Usuario o password incorrecto")
		fmt.Println()
		main()
	}
}

// menu del administrador
func menuAdministrador() {
	opcion := 0
	salir := false
	for !salir {
		fmt.Println("----Esta en el Menu Administrador EDD GoDrive----")
		fmt.Println("1. Ver estudiantes pedientes")
		fmt.Println("2. Ver estudiantes del sistema")
		fmt.Println("3. Registrar nuevo estudiante")
		fmt.Println("4. Carga masiva de estudiantes")
		fmt.Println("5. Salir")
		fmt.Println("Elige una opcion: ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			fmt.Println("pendiente")
		case 2:
			fmt.Println("pendiente")
		case 3:
			fmt.Println()
			registroEstudiante()
		case 4:
			fmt.Println()
			cargaMasiva()
		case 5:
			fmt.Println("Saliendo de la aplicacion")
			salir = true
		default:
			fmt.Println("Opcion no valida")
		}
	}
}

// carga masiva de estudiantes
func cargaMasiva() {
	f, err := os.Open("Prueba.csv")

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	if _, err := r.Read(); err != nil {
		log.Fatal(err)
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		nombre := ""
		apellido := ""
		carne := ""
		password := ""
		carne = record[0]

		carne1, err := strconv.Atoi(carne)
		if err == nil {
			fmt.Println(carne1)
		}

		password = record[2]
		nombres := record[1]
		dato := strings.Split(nombres, " ")

		nombre = dato[0]
		apellido = dato[1]

		//fmt.Println(nombre)

		listaDoble := h.NewLista()
		listaDoble.InsertarFinal(nombre, apellido, carne1, password)

		listaDoble.MostrarConsola()

	}
}

// Registro manual de estudiantes
func registroEstudiante() {
	nombre := ""
	apellido := ""
	carne := 0
	password := ""

	fmt.Println("----Registro de Estudiante----")
	fmt.Println("Ingrese los siguientes datos: ")

	fmt.Print("Nombre: ")
	fmt.Scan(&nombre)
	fmt.Print("Apellido: ")
	fmt.Scan(&apellido)
	fmt.Print("Carne: ")
	fmt.Scan(&carne)
	fmt.Print("Password: ")
	fmt.Scan(&password)
	listaDoble := h.NewLista()

	listaDoble.InsertarFinal(nombre, apellido, carne, password)
	fmt.Println("Estudiante registrado con exito :)")
	menuAdministrador()
}

func menuCliente() {
	fmt.Println("Esta en el menu cliente")
	fmt.Println("-----------------------")
}
