package maps

import "fmt"

func MapCopy(from_map map[string]bool, to_map map[string]bool) (r map[string]bool, err error) {

	fmt.Println("Before copy to_map")
	fmt.Println(to_map)

	for key, value := range from_map {
		to_map[key] = value
	}

	fmt.Println("After copy to_map")
	fmt.Println(to_map)

	return to_map, nil
}
