package main

import (
	"fmt"
	"strings"
)

// Definición de la estructura para un cliente.
type infoCliente struct {
	nombre string
	correo string
	edad   int32
}

// Definición de la estructura para la lista de clientes.
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

// Función genérica de reducción.
func reduce[T any, U any](slice []T, initial U, reducer func(U, T) U) U {
	accum := initial
	for _, value := range slice {
		accum = reducer(accum, value)
	}
	return accum
}

// Función que cuenta la cantidad de correos de Costa Rica.
func cantidadCorreosCostaRica(clientes *[]infoCliente) int {
	// Filtrar correos que terminan en '.cr'.
	filtrados := filter(*clientes, func(cliente infoCliente) bool {
		return strings.HasSuffix(strings.ToLower(cliente.correo), ".cr")
	})

	// Reducir para contar los correos filtrados.
	return reduce(filtrados, 0, func(acc int, _ infoCliente) int {
		return acc + 1
	})
}

// Función principal para probar `cantidadCorreosCostaRica`.
func main() {
	var listaClientes clienteList
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

	// Calcular la cantidad de correos de Costa Rica.
	cantidad := cantidadCorreosCostaRica(&listaClientes.clientes)
	fmt.Printf("Cantidad de correos de Costa Rica: %d\n", cantidad)
}
