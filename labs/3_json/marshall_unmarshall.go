package main

//https://medium.com/@chaewonkong/go-and-json-a-comprehensive-guide-to-working-with-json-in-golang-143fa2dfa897#:~:text=JSON%20Marshalling%20and%20Unmarshalling%20in,Unmarshal%20function.
import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	IsStudent bool     `json:"isStudent"`
	Courses   []string `json:"courses"`
}

func main() {

	Marshalling()
	Unmarshalling()
	ErrorHandlingUnmarshalling()
	CustomTypes()
}

func Marshalling() {
	p := Person{
		Name:      "John Doe",
		Age:       30,
		IsStudent: false,
		Courses:   []string{"math", "history", "chemistry"},
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}

func Unmarshalling() {
	jsonString := `[{
        "name": "John Doe",
        "age": 30,
        "isStudent": false,
        "courses": ["math", "history", "chemistry"]
    },
	{
        "name": "Juan Díaz",
        "age": 45,
        "isStudent": true,
        "courses": ["ai", "business", "management"]
    }]`

	var p []Person
	err := json.Unmarshal([]byte(jsonString), &p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%+v\n", p)
}

func ErrorHandlingUnmarshalling() {

	jsonString := `{
        "name": "John Doe",
        "age": "30",
        "isStudent": false,
        "courses": ["math", "history", "chemistry"]
    }`

	var p Person
	err := json.Unmarshal([]byte(jsonString), &p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%+v\n", p)
	//In this case, the error occurs because the age field has a string value instead of an integer.
}

/*
	func (e Email) MarshalJSON() ([]byte, error) {
		return json.Marshal(strings.ToLower(string(e)))
	}

	func (e *Email) UnmarshalJSON(data []byte) error {
		var s string
		err := json.Unmarshal(data, &s)
		if err != nil {
			return err
		}
		*e = Email(strings.ToLower(s))
		return nil
	}
*/
//Custom types:
//Go allows you to define custom types and implement the json.Marshaler and json.Unmarshaler
//interfaces to customize the JSON encoding and decoding process.
//Here's an example using a custom type for an email address:

type Email string
type Persona struct {
	Name  string
	Email Email
}

func CustomTypes() {
	p := Persona{
		Name:  "John Doe",
		Email: "JOHN.DOe@Example.com",
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))

	var decodedPersona Persona
	err = json.Unmarshal(jsonData, &decodedPersona)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%+v\n", decodedPersona)
}
