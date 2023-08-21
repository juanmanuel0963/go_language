package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func AdjustIngredients(ingredients []string, persons int) []string {
	var adjustedIngredients []string

	for _, ingredient := range ingredients {
		adjustedIngredient := adjustSingleIngredient(ingredient, persons)
		adjustedIngredients = append(adjustedIngredients, strings.TrimSpace(adjustedIngredient))
	}

	return adjustedIngredients
}

func adjustSingleIngredient(ingredient string, persons int) string {
	re := regexp.MustCompile(`(\d+)\s*(\w+)\s*(.*)`)
	matches := re.FindStringSubmatch(ingredient)

	if len(matches) != 4 {
		return ingredient // Return as-is if the ingredient format is not recognized
	}

	amountStr := matches[1]
	unit := matches[2]
	ingredientName := matches[3]

	amount, _ := strconv.Atoi(amountStr)
	adjustedAmount := amount * persons

	return fmt.Sprintf("%d %s %s", adjustedAmount, unit, ingredientName)
}

/*
func AdjustIngredients(ingredients []string, persons int) []string {
	var adjustedIngredients []string

	for _, ingredient := range ingredients {
		adjustedIngredient := adjustSingleIngredient(ingredient, persons)
		adjustedIngredients = append(adjustedIngredients, adjustedIngredient)
	}

	return adjustedIngredients
}

func adjustSingleIngredient(ingredient string, persons int) string {
	re := regexp.MustCompile(`(\d+)\s*(\w+)\s*(.*)`)
	matches := re.FindStringSubmatch(ingredient)

	if len(matches) != 4 {
		return ingredient // Return as-is if the ingredient format is not recognized
	}

	amountStr := matches[1]
	unit := matches[2]
	ingredientName := matches[3]

	amount, _ := strconv.Atoi(amountStr)
	adjustedAmount := amount * persons

	return fmt.Sprintf("%d %s %s", adjustedAmount, unit, ingredientName)
}
*/
/*
func main() {
	ingredients := []string{
		"2 eggs", "200 grams of flour","150 grams of sugar","1 liter of milk",
	}
	persons := 3

	adjustedIngredients := AdjustIngredients(ingredients, persons)
	for _, ingredient := range adjustedIngredients {
		fmt.Println(ingredient)
	}
}
*/
