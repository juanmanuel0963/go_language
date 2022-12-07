package structs

//How do we perform inheritance with Golang?
//------------------------------------------------
//- This is a bit of a trick question: there is no inheritance in Golang because it does not support classes.
//- However, you can mimic inheritance behavior using composition to use an existing struct object to define a starting behavior of a new object.
//- Once the new object is created, functionality can be extended beyond the original struct.

//The Lamp struct contains On(), Off() functions.
//These functions are embedded into the child struct Street by simply listing the struct at the top of the implementation of Street.

type Lamp struct{}

func (l Lamp) On() {
	println("On")

}

func (l Lamp) Off() {
	println("Off")
}

type Street struct {
	Lamp
}
