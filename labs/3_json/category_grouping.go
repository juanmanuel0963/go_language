package main

import (
	"encoding/json"
	"time"
)

// Definimos la estructura del producto con el nuevo campo SellDate
type Product struct {
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	Category string    `json:"category"`
	SellDate time.Time `json:"sell_date"`
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

		// Añadimos el producto agrupado por su categoría
		grouped[category] = append(grouped[category], Product{
			Name:     product.Name,
			Price:    product.Price,
			SellDate: product.SellDate, // Incluimos SellDate en la agrupación
		})
	}

	return grouped, nil
}

/*
func main() {
	// JSON de entrada con el campo sell_date
	input := []byte(`[
        {"name": "Laptop", "price": 1000, "category": "Electronics", "sell_date": "2024-09-28T10:00:00Z"},
        {"name": "Headphones", "price": 100, "category": "Electronics", "sell_date": "2024-09-25T15:30:00Z"},
        {"name": "Coffee Mug", "price": 15, "category": "Home Goods", "sell_date": "2024-09-20T08:15:00Z"},
        {"name": "Mystery Box", "price": 50, "category": "", "sell_date": "2024-09-26T12:45:00Z"}
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
*/
