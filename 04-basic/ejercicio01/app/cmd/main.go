package main

import (
	"app/internal/product"
	"fmt"
)

/*Ejercicio 1
Crear un programa que cumpla los siguiente puntos:

Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
Tener un slice global de Product llamado Products instanciado con valores.
2 métodos asociados a la estructura Product: Save(), GetAll().
El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método.
El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
Ejecutar al menos una vez cada método y función definido desde main(). */

func main() {

	product1 := product.Product{ID: 1, Name: "Marco", Price: 23.45, Description: "hola", Category: "holados"}

	var flag bool = false
	p := product.Product{}

	for !flag {

		fmt.Println("Ingresa un producto")

		fmt.Println("Ingresa id:")
		fmt.Scanln(&product1.ID)
		fmt.Println("Ingresa el precio:")
		fmt.Scanln(&product1.Price)
		fmt.Println("Ingresa una descripcion")
		fmt.Scanln(&product1.Description)
		fmt.Println("Ingresa la categoria")
		fmt.Scanln(&product1.Category)

		fmt.Println("---------------------------")
		fmt.Println("Quieres ingresar o registro")
		fmt.Scanln(&flag)

		p.Save(product1)

	}

	for _, p := range p.GetAll() {
		fmt.Println(p)
	}

	fmt.Println(p.GetId(1))

}
