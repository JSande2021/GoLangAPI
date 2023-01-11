package products

import (
	"fmt"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	CreatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

// This section provides a method that displays all the products in our list.
func GetProducts() Products {
	return products
}

// This section provides a method that returns a product with a matching id to the one provided in the query.
var ProductNotFoundError = fmt.Errorf("Product not found")

func GetProdID(id int) (*Product, int, error) {
	for i, p := range products {
		if id == p.ID {
			return p, i, nil
		}

	}
	return nil, -1, ProductNotFoundError
}

// This section provides a method to add a new product to our existing product list.
func AddProd(p *Product) []*Product {
	products = append(products, p)
	return products
}

// This section provides a method that deletes any product with a matching id to the one provided in the query.

func remove(p []*Product, i int) []*Product {
	return append(p[:i], p[i+1:]...)
}

func DeleteProd(id int) []*Product {
	for i, p := range products {
		if p.ID == id {
			return remove(products, i)
		}
	}
	return nil
}

// Our Dummy products. Since  no database is used, we have provided just these dummies below.
var products = []*Product{
	{
		ID:          10,
		Name:        "Vitamin B",
		Description: "Maintaining good health and well being",
		Price:       4.67,
		CreatedOn:   time.Now().String(),
		DeletedOn:   time.Now().String(),
	},
	{
		ID:          11,
		Name:        "Polysaccharides",
		Description: "Improves digstives systems, innune systems and helps in detox",
		Price:       11.20,
		CreatedOn:   time.Now().String(),
		DeletedOn:   time.Now().String(),
	},
	{
		ID:          12,
		Name:        "Tonics",
		Description: "Improve the production of blood cells and platlets",
		Price:       5.0,
		CreatedOn:   time.Now().String(),
		DeletedOn:   time.Now().String(),
	},
}
