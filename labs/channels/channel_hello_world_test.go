package channels

import "testing"

func TestHelloWorld_v1(t *testing.T) {

	HelloWorld("hello")
	HelloWorld("world")
}

func TestHelloWorld_v2(t *testing.T) {

	go HelloWorld("hello")
	HelloWorld("world")
}
