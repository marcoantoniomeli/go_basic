package product

import (
	"fmt"
)

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var listProduct []Product

func (p *Product) Save(product Product) {

	listProduct = append(listProduct, product)

	fmt.Println(len(listProduct))

	fmt.Println("Se agrego el producto ")

}

func (p *Product) GetAll() []Product {

	return listProduct

}

func (p *Product) GetId(id int) Product {
	// Buscar un producto por su ID

	var temp Product

	for _, product := range listProduct {
		if product.ID == id {
			temp = product
			break // Si encuentras el producto, puedes salir del bucle
		}
	}

	return temp
}
