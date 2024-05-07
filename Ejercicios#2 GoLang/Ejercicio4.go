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

// Función genérica de mapeo.
func mapFunc[T any, U any](slice []T, mapper func(T) U) []U {
	var result []U
	for _, item := range slice {
		result = append(result, mapper(item))
	}
	return result
}

// Función que sugiere correos para clientes cuyos correos no hacen referencia a su nombre.
func clientesSugerenciasCorreos(clientes *[]infoCliente) []infoCliente {
	return mapFunc(*clientes, func(cliente infoCliente) infoCliente {
		// Convertir el nombre del cliente a minúsculas y eliminar espacios.
		nombreSimplificado := strings.ReplaceAll(strings.ToLower(cliente.nombre), " ", "")

		// Verificar si el correo ya contiene el nombre simplificado.
		if !strings.Contains(strings.ToLower(cliente.correo), nombreSimplificado) {
			// Extraer el dominio del correo existente.
			partesCorreo := strings.SplitN(cliente.correo, "@", 2)
			if len(partesCorreo) == 2 {
				// Sugerir una nueva dirección manteniendo el mismo dominio.
				dominio := partesCorreo[1]
				cliente.correo = fmt.Sprintf("%s.%s@%s", strings.ToLower(strings.Split(cliente.nombre, " ")[0]), strings.ToLower(strings.Split(cliente.nombre, " ")[1]), dominio)
			} else {
				// Si no hay dominio, sugerir un dominio genérico.
				cliente.correo = fmt.Sprintf("%s.%s@dominiosugerido.com", strings.ToLower(strings.Split(cliente.nombre, " ")[0]), strings.ToLower(strings.Split(cliente.nombre, " ")[1]))
			}
		}
		return cliente
	})
}

func main() {
	// Crear lista de clientes.
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

	// Generar sugerencias de correos.
	clientesConCorreosSugeridos := clientesSugerenciasCorreos(&listaClientes.clientes)

	// Imprimir la lista modificada con sugerencias.
	fmt.Println("Lista de clientes con correos sugeridos:")
	for _, cliente := range clientesConCorreosSugeridos {
		fmt.Printf("Nombre: %s, Correo: %s\n", cliente.nombre, cliente.correo)
	}
}
