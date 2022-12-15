package structs

type Person struct {
	name  string
	age   int
	power interface{}
}

//Ensure that they don’t contain slices, maps, or functions. Otherwise, the code won’t comply.

type Animal struct {
	name string
	size int
	sons []string
}

func StrucsCompare_Persons(a Person, b Person) bool {

	equal := (a == b)

	return equal
}
