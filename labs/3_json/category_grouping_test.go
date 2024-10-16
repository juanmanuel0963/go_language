package main

import (
	"reflect"
	"testing"
	"time"
)

// Test case para GroupProductsByCategory
func TestGroupProductsByCategory(t *testing.T) {
	// Definimos el formato de fecha esperado
	const timeLayout = time.RFC3339

	// Entrada JSON con el campo sell_date
	input := []byte(`[
        {"name": "Laptop", "price": 1000, "category": "Electronics", "sell_date": "2024-09-28T10:00:00Z"},
        {"name": "Headphones", "price": 100, "category": "Electronics", "sell_date": "2024-09-25T15:30:00Z"},
        {"name": "Coffee Mug", "price": 15, "category": "Home Goods", "sell_date": "2024-09-20T08:15:00Z"},
        {"name": "Mystery Box", "price": 50, "category": "", "sell_date": "2024-09-26T12:45:00Z"}
    ]`)

	// Parseamos las fechas esperadas
	laptopSellDate, _ := time.Parse(timeLayout, "2024-09-28T10:00:00Z")
	headphonesSellDate, _ := time.Parse(timeLayout, "2024-09-25T15:30:00Z")
	coffeeMugSellDate, _ := time.Parse(timeLayout, "2024-09-20T08:15:00Z")
	mysteryBoxSellDate, _ := time.Parse(timeLayout, "2024-09-26T12:45:00Z")

	// Resultado esperado
	expected := map[string][]Product{
		"Electronics": {
			{Name: "Laptop", Price: 1000, SellDate: laptopSellDate},
			{Name: "Headphones", Price: 100, SellDate: headphonesSellDate},
		},
		"Home Goods": {
			{Name: "Coffee Mug", Price: 15, SellDate: coffeeMugSellDate},
		},
		"Unknown": {
			{Name: "Mystery Box", Price: 50, SellDate: mysteryBoxSellDate},
		},
	}

	// Ejecutamos la función para agrupar productos
	result, err := GroupProductsByCategory(input)
	if err != nil {
		t.Fatalf("Error al agrupar productos: %v", err)
	}

	// Comparamos el resultado con lo esperado
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Resultado inesperado. Got: %v, Want: %v", result, expected)
	}
}
