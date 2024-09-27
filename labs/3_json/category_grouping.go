package main

import (
	"encoding/json"
	"fmt"
)

// Definimos la estructura del producto
type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

// Función que agrupa los productos por categoría
func GroupProductsByCategory(input []byte) (map[string][]Product, error) {
	var products []Product
	err := json.Unmarshal(input, &products)
	if err != nil {
		return nil, err
	}

	grouped := make(map[string][]Product)
	for _, product := range products {
		if product.Category == "" {
			product.Category = "Unknown"
		}

		category := product.Category

		grouped[category] = append(grouped[product.Category], Product{
			Name:  product.Name,
			Price: product.Price,
		})
	}

	return grouped, nil
}

func main() {
	// JSON de entrada
	input := []byte(`[
        {"name": "Laptop", "price": 1000, "category": "Electronics"},
        {"name": "Headphones", "price": 100, "category": "Electronics"},
        {"name": "Coffee Mug", "price": 15, "category": "Home Goods"},
        {"name": "Mystery Box", "price": 50, "category": ""}
    ]`)

	// Llamamos a la función para agrupar productos
	groupedProducts, err := GroupProductsByCategory(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convertimos el resultado a JSON y lo imprimimos
	output, _ := json.MarshalIndent(groupedProducts, "", "  ")
	fmt.Println(string(output))
}
