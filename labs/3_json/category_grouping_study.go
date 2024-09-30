package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// JSON de entrada
	input := []byte(`[
        {"name": "Laptop", "price": 1000, "category": "Electronics", "sell_date": "2024-09-28T10:00:00Z"},
        {"name": "Headphones", "price": 100, "category": "Electronics", "sell_date": "2024-09-25T15:30:00Z"},
        {"name": "Coffee Mug", "price": 15, "category": "Home Goods", "sell_date": "2024-09-20T08:15:00Z"},
        {"name": "Mystery Box", "price": 50, "category": "", "sell_date": "2024-09-26T12:45:00Z"}
    ]`)

	list, err := GroupProductsByCategoryStudy(input)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(list)
}

//Define Product struct

type ProductStudy struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
	SellDate string  `json:"sell_date"`
}

func GroupProductsByCategoryStudy(input []byte) (map[string][]ProductStudy, error) {

	products := make([]ProductStudy, 0)

	//Unmarshall input
	err := json.Unmarshal(input, &products)

	if err != nil {
		return nil, err
	}

	grouped := make(map[string][]ProductStudy, 0)

	//Iterate products
	for _, product := range products {
		//fmt.Println(product)

		if product.Category == "" {
			product.Category = "Unknown"
		}

		produtToAdd := ProductStudy{Name: product.Name, Price: product.Price, SellDate: product.SellDate}

		grouped[product.Category] = append(grouped[product.Category], produtToAdd)

	}

	//Return map: key:category, value: slice of products

	return grouped, nil
}
