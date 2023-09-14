package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Expresion interface {
	result() int
}

type StringMultiplication struct {
	expression string
}

type StringAddition struct {
	expression string
}

func getResult(i Expresion) int {
	return i.result()
}

func (s StringMultiplication) result() int {
	a, _ := strconv.Atoi(string(s.expression[0]))
	b, _ := strconv.Atoi(string(s.expression[2]))

	return a * b
}

func (s StringAddition) result() int {
	a, _ := strconv.Atoi(string(s.expression[0]))
	b, _ := strconv.Atoi(string(s.expression[2]))

	return a + b
}

func generateVal(channel chan int, query string) {

	fmt.Println(query)

	if query[1] == '+' {

		qAddition := StringAddition{expression: query}
		//val := qAddition.result()
		val := getResult(qAddition)
		channel <- val
		fmt.Println("Addition: " + strconv.Itoa(val))
	} else if query[1] == '*' {

		qMultiplication := StringMultiplication{expression: query}
		//val := qMultiplication.result()
		val := getResult(qMultiplication)
		channel <- val
		fmt.Println("Multiplication: " + strconv.Itoa(val))
	}

	close(channel)
}

func solution(queries []string) []int {

	var mySlice []int

	for i := 0; i < len(queries); i++ {
		channel := make(chan int)

		go generateVal(channel, queries[i])
		channelResponse := <-channel

		mySlice = append(mySlice, channelResponse)
	}

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i] < mySlice[j]
	})

	return mySlice
}

func main() {
	response := solution([]string{"1*5", "1+2"})
	fmt.Println(response)
}
