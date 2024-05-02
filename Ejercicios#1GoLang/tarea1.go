package main

import (
	"errors"
	"fmt"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}

type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	indice := l.buscarProductoIndex(nombre)
	if indice != -1 {
		// Si el producto existe, actualiza cantidad y precio si es diferente
		(*l)[indice].cantidad += cantidad
		if (*l)[indice].precio != precio {
			(*l)[indice].precio = precio
		}
	} else {
		// Si no existe, lo agrega nuevo
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}
}

func (l *listaProductos) agregarProductos(productos ...producto) {
	for _, prod := range productos {
		l.agregarProducto(prod.nombre, prod.cantidad, prod.precio)
	}
}

func (l *listaProductos) buscarProductoIndex(nombre string) int {
	for i, prod := range *l {
		if prod.nombre == nombre {
			return i
		}
	}
	return -1
}

func (l *listaProductos) buscarProducto(nombre string) (producto, int) {
	indice := l.buscarProductoIndex(nombre)
	if indice != -1 {
		return (*l)[indice], 0
	}
	return producto{}, 1 // Error 1 significa "producto no encontrado"
}

func (l *listaProductos) venderProducto(nombre string, cantidad int) error {
	indice := l.buscarProductoIndex(nombre)
	if indice == -1 {
		return errors.New("producto no encontrado")
	}
	if (*l)[indice].cantidad < cantidad {
		return errors.New("no hay suficiente cantidad del producto")
	}
	(*l)[indice].cantidad -= cantidad
	if (*l)[indice].cantidad == 0 {
		// Eliminar producto de la lista
		*l = append((*l)[:indice], (*l)[indice+1:]...)
		fmt.Println("Producto eliminado por falta de existencias.")
	}
	return nil
}

func (l *listaProductos) cambiarPrecioProducto(nombre string, nuevoPrecio int) error {
	indice := l.buscarProductoIndex(nombre)
	if indice == -1 {
		return errors.New("producto no encontrado")
	}
	(*l)[indice].precio = nuevoPrecio
	return nil
}

func (l listaProductos) imprimirProducto(nombre string) error {
	indice := l.buscarProductoIndex(nombre)
	if indice == -1 {
		return errors.New("producto no encontrado")
	}
	prod := l[indice]
	fmt.Printf("Producto: %s, Cantidad: %d, Precio: %d\n", prod.nombre, prod.cantidad, prod.precio)
	return nil
}

func (l *listaProductos) listarProductosMinimos() listaProductos {
	var productosMinimos listaProductos
	for _, prod := range *l {
		if prod.cantidad <= existenciaMinima {
			productosMinimos = append(productosMinimos, prod)
		}
	}
	return productosMinimos
}

func llenarDatos() {

	lProductos.agregarProductos(
		producto{"arroz", 15, 2500},
		producto{"frijoles", 4, 2000},
		producto{"leche", 8, 1200},
		producto{"café", 12, 4500})
}

func main() {
	lProductos.agregarProducto("arroz", 5, 2500)
	fmt.Println("Lista de Productos:", lProductos)
	fmt.Println("Aqui se inicia con 5 de arroz")
	llenarDatos()
	lProductos.agregarProductos(producto{"tomates", 25, 3000}, producto{"cebollas", 40, 1500}, producto{"culantro", 30, 500})
	fmt.Println("----------------------Lista de productos mas la adiccion al arroz --------------------------")
	fmt.Println("Lista de Productos:", lProductos)
	// Vender algunos productos como ejemplo
	fmt.Println("----------------------Venta de productos --------------------------")
	lProductos.venderProducto("arroz", 5)
	lProductos.venderProducto("frijoles", 2)
	lProductos.venderProducto("cebollas", 10)

	fmt.Println("Lista de Productos después de venta:", lProductos)

	fmt.Println("Cambio de precio antes:")
	if err := lProductos.imprimirProducto("café"); err != nil {
		fmt.Println("Error:", err)
	}

	lProductos.cambiarPrecioProducto("café", 5000)

	fmt.Println("Cambio de precio después:")
	if err := lProductos.imprimirProducto("café"); err != nil {
		fmt.Println("Error:", err)
	}

	// Listar productos con existencias mínimas
	minimos := lProductos.listarProductosMinimos()
	fmt.Println("Productos con existencias mínimas:", minimos)
}

/*
Lista de Productos: [{arroz 5 2500}]

Aqui se inicia con 5 de arroz
----------------------Lista de productos mas la adiccion al arroz --------------------------
Lista de Productos: [{arroz 20 2500} {frijoles 4 2000} {leche 8 1200} {café 12 4500} {tomates 25 3000} {cebollas 40 1500} {culantro 30 500}]


----------------------Venta de productos --------------------------
Lista de Productos después de venta: [{arroz 15 2500} {frijoles 2 2000} {leche 8 1200} {café 12 4500} {tomates 25 3000} {cebollas 30 1500} {culantro 30 500}]

Cambio de precio antes:
Producto: café, Cantidad: 12, Precio: 4500

Cambio de precio después:
Producto: café, Cantidad: 12, Precio: 5000

Productos con existencias mínimas: [{frijoles 2 2000} {leche 8 1200}]
*/
