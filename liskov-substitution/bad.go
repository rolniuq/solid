package main

import "fmt"

type BadBird interface {
	Fly()
}

type BadSparrow struct{}

func (s *BadSparrow) Fly() {
	fmt.Println("im bad sparrow and i can fly")
}

type BadEagle struct{}

func (e *BadEagle) Fly() {
	fmt.Println("im bad eagle and i can fly")
}

// in this case if we want to add a penguin class -> the LSP go wrong

// how to fix:
// 1. system redesign
// 2. use interface

type Flyable interface {
	Fly()
}

type FixSparrow struct{}

func (f *FixSparrow) Fly() {
	fmt.Println("im fix sparrow and i can fly")
}

type FixEagle struct{}

func (f *FixEagle) Fly() {
	fmt.Println("im fix eagle and i can fly")
}

type Penguin struct{}

// do not need to implement the flyable

func badInit() {
	input := "penguin"

	var f Flyable
	switch input {
	case "sparrow":
		f = &FixSparrow{}
		f.Fly()
	case "eagle":
		f = &Eagle{}
		f.Fly()
	default:
		fmt.Println("im penguin and i cant fly")
	}
}
