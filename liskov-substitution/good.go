package main

import "fmt"

type Bird interface {
	Fly()
}

type Sparrow struct{}

func (s *Sparrow) Fly() {
	fmt.Println("im sparrow and i can fly")
}

type Eagle struct{}

func (e *Eagle) Fly() {
	fmt.Println("im eagle and i can fly")
}

func goodInit() {
	input := "sparrow"

	var b Bird
	switch input {
	case "sparrow":
		b = &Sparrow{}
	case "eagle":
		b = &Eagle{}
	}

	b.Fly()
}
