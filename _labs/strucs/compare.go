package strucs

type Person struct {
	name string
	age  int
}

/*
	type Animal struct {
		name string
		size int
		sons []string
	}
*/
func StrucsCompare_Persons(a Person, b Person) bool {

	equal := (a == b)

	return equal
}

/*
func StrucsCompare_Animals(a Animal, b Animal) bool {

	equal, err := (a == b)

	return equal
}
*/
