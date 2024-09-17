package main

import (
	"errors"
	"fmt"
)

type MaxStack struct {
	stack    []int
	maxStack []int
}

func (s *MaxStack) Push(x int) {
	s.stack = append(s.stack, x)
	if len(s.maxStack) == 0 || x >= s.maxStack[len(s.maxStack)-1] {
		s.maxStack = append(s.maxStack, x)
	}
	fmt.Println(s.stack)
	//fmt.Println(s.maxStack)
}

func (s *MaxStack) Pop() (int, error) {
	if len(s.stack) == 0 {
		return 0, errors.New("empty stack")
	}
	x := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	if x == s.maxStack[len(s.maxStack)-1] {
		s.maxStack = s.maxStack[:len(s.maxStack)-1]
	}
	fmt.Println(s.stack)
	return x, nil
}

func (s *MaxStack) Max() (int, error) {
	if len(s.maxStack) == 0 {
		return 0, errors.New("empty stack")
	}
	return s.maxStack[len(s.maxStack)-1], nil
}
