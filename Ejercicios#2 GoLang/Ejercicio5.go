package main

import (
	"fmt"
	"sort"
)

// Definición de la estructura para un cliente.
type infoCliente struct {
	nombre string
	correo string
	edad   int32
}

// Estructura que contiene una lista de clientes.
type clienteList struct {
	clientes []infoCliente
}

// Método para agregar un cliente a la lista.
func (cl *clienteList) agregarCliente(nombre, correo string, edad int32) {
	nuevoCliente := infoCliente{nombre, correo, edad}
	cl.clientes = append(cl.clientes, nuevoCliente)
}

// Función que devuelve una lista de correos ordenados alfabéticamente.
func correosOrdenadosAlfabeticos(clientes []infoCliente) []string {
	correos := make([]string, len(clientes))
	for i, cliente := range clientes {
		correos[i] = cliente.correo
	}
	sort.Strings(correos)
	return correos
}

func main() {
	// Crear una instancia de `clienteList`.
	var listaClientes clienteList

	// Agregar clientes a la lista utilizando `agregarCliente`.
	listaClientes.agregarCliente("Oscar Viquez", "oviquez@tec.ac.cr", 44)
	listaClientes.agregarCliente("Pedro Perez", "elsegundo@gmail.com", 30)
	listaClientes.agregarCliente("Maria Lopez", "mlopez@hotmail.com", 18)
	listaClientes.agregarCliente("Juan Rodriguez", "jrodriguez@gmail.com", 25)
	listaClientes.agregarCliente("Luisa Gonzalez", "lgonzalez@tec.ac.cr", 67)
	listaClientes.agregarCliente("Marco Rojas", "marco.rojas@hotmail.com", 47)
	listaClientes.agregarCliente("Marta Saborio", "msaborio@gmail.com", 33)
	listaClientes.agregarCliente("Camila Segura", "csegura@ice.co.cr", 19)
	listaClientes.agregarCliente("Fernando Rojas", "frojas@estado.gov", 27)
	listaClientes.agregarCliente("Rosa Ramirez", "risuenna@gmail.com", 50)

	// Obtener la lista de correos ordenados.
	correosOrdenados := correosOrdenadosAlfabeticos(listaClientes.clientes)

	// Imprimir la lista de correos ordenados.
	fmt.Println("Correos de clientes ordenados alfabéticamente:")
	for _, correo := range correosOrdenados {
		fmt.Println(correo)
	}
}
