package main

import "fmt"

// infoCliente define la estructura para un cliente.
type infoCliente struct {
	nombre string
	correo string
	edad   int32
}

// clienteList contiene una lista de infoCliente.
type clienteList struct {
	clientes []infoCliente
}

// agregarCliente añade un cliente a la lista.
func (cl *clienteList) agregarCliente(nombre, correo string, edad int32) {
	nuevoCliente := infoCliente{nombre, correo, edad}
	cl.clientes = append(cl.clientes, nuevoCliente)
}

// main es la función principal que inicializa la lista y añade clientes.
func main() {
	// Crear una lista de clientes vacía.
	var listaClientes clienteList

	// Añadir los 10 clientes solicitados.
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

	// Mostrar los detalles de los clientes para verificación.
	fmt.Println("Lista de Clientes:")
	for _, cliente := range listaClientes.clientes {
		fmt.Printf("Nombre: %s, Correo: %s, Edad: %d\n", cliente.nombre, cliente.correo, cliente.edad)
	}
}
