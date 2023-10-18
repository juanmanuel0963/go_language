package main

import "fmt"

func solution(input string) bool {

	// Recorrer el arreglo de izquierda a derecha y de derecha a izquierda, verificar que coincida cada caracter

	j := len(input) - 1

	for i := 0; i < len(input)/2; i++ {

		if input[i] != input[j] {

			return false
		}

		j--

	}

	return true
}

func main() {

	got := solution("anana")

	fmt.Println(got)

	for i := 0; i < 5; i++ {

		defer fmt.Printf("%d ", i)

	}
}

/*
Person: id, name,
Email: id, value, id_person


SELECT P.name, E.value
FROM Person P
INNER JOIN Email E
WHERE P.id = E.id_person


SELECT P.name, E.value
FROM Person P
LEFT JOIN Email E
WHERE P.id = E.id_person


Add - POST 		: http://localhost/person
Update - PUT 	: http://localhost/person/id
					body{name:""}
Patch - PATCH 	: http://localhost/person/id
					body{name:""}

Get    - GET 	: http://localhost/person/id


GetAll - GET 	: http://localhost/person
					body{}

Delete - DELETE : http://localhost/person

for i := 0; i < 5; i++ {

    defer fmt.Printf("%d ", i)

}

4
3
2
1
0
*/
