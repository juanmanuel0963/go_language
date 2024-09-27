package main

import (
	"reflect"
	"testing"
)

// Test case para GroupProductsByCategory
func TestGroupProductsByCategory(t *testing.T) {
	input := []byte(`[
        {"name": "Laptop", "price": 1000, "category": "Electronics"},
        {"name": "Headphones", "price": 100, "category": "Electronics"},
        {"name": "Coffee Mug", "price": 15, "category": "Home Goods"},
        {"name": "Mystery Box", "price": 50, "category": ""}
    ]`)

	expected := map[string][]Product{
		"Electronics": {
			{Name: "Laptop", Price: 1000},
			{Name: "Headphones", Price: 100},
		},
		"Home Goods": {
			{Name: "Coffee Mug", Price: 15},
		},
		"Unknown": {
			{Name: "Mystery Box", Price: 50},
		},
	}

	result, err := GroupProductsByCategory(input)
	if err != nil {
		t.Fatalf("Error al agrupar productos: %v", err)
	}

	// Comparamos el resultado con lo esperado
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Resultado inesperado. Got: %v, Want: %v", result, expected)
	}
}
