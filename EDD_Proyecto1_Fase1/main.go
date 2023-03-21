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

var listaCola *h.ListaCola = h.NewCola()

func main() {

	salir := false

	for !salir {
		fmt.Println("----Bienvenido a la aplicacion EDD GoDrive----")
		fmt.Println("1. Inisiar sesion")
		fmt.Println("2. Salir")
		fmt.Println("Elige una opcion: ")
		opcion := 0
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Println()
			seleccion_administrador()
		case 2:
			fmt.Println("Saliendo de la aplicacion")
			salir = true
		}
	}
}

// validar credenciales del administrador
func seleccion_administrador() {
	salir := false

	for !salir {
		fmt.Println("Ingrese sus credenciales")
		usuario := ""
		password := ""
		fmt.Print("Usuario: ")
		fmt.Scan(&usuario)
		fmt.Print("Password: ")
		fmt.Scan(&password)

		if usuario == "admin" && password == "admin" {
			fmt.Println()
			menuAdministrador()
		}

		validar := listaCola.Validar(usuario, password)

		if validar {
			fmt.Println()
			menuEstudiante()

		} else {
			fmt.Println("Usuario o password incorrecto")
			fmt.Println()
			main()
		}
	}
}

// menu del administrador
func menuAdministrador() {
	opcion := 0
	salir := false
	for !salir {
		fmt.Println("----Menu Administrador EDD GoDrive----")
		fmt.Println("1. Ver estudiantes pedientes")
		fmt.Println("2. Ver estudiantes del sistema")
		fmt.Println("3. Registrar nuevo estudiante")
		fmt.Println("4. Carga masiva de estudiantes")
		fmt.Println("5. Graficar Cola")
		fmt.Println("6. Graficar Aceptacion de estudiantes")
		fmt.Println("7. Graficar Lista Doble")
		fmt.Println("8. Salir")
		fmt.Println("Elige una opcion: ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			listaCola.MostrarCola()
			fmt.Println()
		case 2:
			listaCola.MostrarListaDoble()
			fmt.Println()
		case 3:
			fmt.Println()
			registroEstudiante()
		case 4:
			fmt.Println()
			cargaMasiva()
		case 5:
			listaCola.GraficarCola()
		case 6:
			listaCola.GraficarAceptacion()
		case 7:
			fmt.Println("pendiente")
		case 8:
			fmt.Println("Saliendo de la aplicacion")
			salir = true

		default:
			fmt.Println("Opcion no valida")
		}
	}
	if opcion == 5 {
		main()
	}
}

// carga masiva de estudiantes
func cargaMasiva() {

	f, err := os.Open("Estudiante.csv")

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
			fmt.Print("")
		}

		password = record[2]
		nombres := record[1]
		dato := strings.Split(nombres, " ")

		nombre = dato[0]
		apellido = dato[1]

		listaCola.AgregarEstudiante(nombre, apellido, carne1, password)

	}
	fmt.Println("Carga masiva realizada con exito :)")
	fmt.Println()
	menuAdministrador()
}

// Registro manual de estudiantes
func registroEstudiante() {
	nombre := ""
	apellido := ""
	carne := 0
	password := ""

	fmt.Println("----Registro de Estudiante - EDD GoDrive----")
	fmt.Println("Ingrese los siguientes datos")
	fmt.Print("Nombre: ")
	fmt.Scan(&nombre)
	fmt.Print("Apellido: ")
	fmt.Scan(&apellido)
	fmt.Print("Carne: ")
	fmt.Scan(&carne)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	listaCola.AgregarEstudiante(nombre, apellido, carne, password)
	fmt.Println("Estudiante registrado con exito :) ")
	fmt.Println()
	menuAdministrador()
}

func menuEstudiante() {
	opcion := 0
	fmt.Println("1. Salir")
	fmt.Scan(&opcion)
}
