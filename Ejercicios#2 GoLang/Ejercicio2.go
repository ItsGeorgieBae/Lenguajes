package main

import (
	"fmt"
	"strings"
)

// Estructura para un cliente.
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

// Función genérica de filtrado.
func filter[T any](slice []T, test func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if test(item) {
			result = append(result, item)
		}
	}
	return result
}

// Función que devuelve una lista de clientes cuyos correos contienen el apellido.
func listaClientes_ApellidoEnCorreo(clientes *[]infoCliente) []infoCliente {
	return filter(*clientes, func(cliente infoCliente) bool {
		// Dividir el nombre para extraer el apellido (segunda palabra).
		partes := strings.Split(cliente.nombre, " ")
		if len(partes) < 2 {
			return false
		}
		apellido := strings.ToLower(partes[1])

		// Verificar si el correo contiene el apellido.
		return strings.Contains(strings.ToLower(cliente.correo), apellido)
	})
}

// Función principal para probar `listaClientes_ApellidoEnCorreo`.
func main() {
	// Crear una lista de clientes vacía.
	var listaClientes clienteList

	// Agregar los clientes a la lista.
	listaClientes.agregarCliente("Oscar Viquez", "oviquez@tec.ac.cr", 44)
	listaClientes.agregarCliente("Jorge Valladares", "javalladablanco@gmail.com", 22)
	listaClientes.agregarCliente("Camila Barrantes", "camibarraugalde@hotmail.com", 21)
	listaClientes.agregarCliente("Juan Nunez", "Jnunezparra@gmail.com", 20)
	listaClientes.agregarCliente("Luisy Campos", "LuCampos@tec.ac.cr", 24)
	listaClientes.agregarCliente("Rosario Davila", "davirosa321@hotmail.com", 24)
	listaClientes.agregarCliente("Yandra Rojas", "yanrojas@gmail.com", 23)
	listaClientes.agregarCliente("Jose Blanco", "Joseacubla@ice.co.cr", 19)
	listaClientes.agregarCliente("Melisa Guzman", "MEguzman@estado.gov", 20)
	listaClientes.agregarCliente("Karen Rojas", "Krojas1212@gmail.com", 22)

	// Filtrar la lista de clientes cuyos correos contienen el apellido.
	resultado := listaClientes_ApellidoEnCorreo(&listaClientes.clientes)

	// Imprimir el resultado.
	fmt.Println("Clientes cuyos correos contienen el apellido:")
	for _, cliente := range resultado {
		fmt.Printf("Nombre: %s, Correo: %s\n", cliente.nombre, cliente.correo)
	}
}
