package main

import "fmt"

type Human struct{}

func (h *Human) Work() {
	fmt.Println("human is working")
}

func (h *Human) Eat() {
	fmt.Println("human is eating")
}

type Robot struct{}

func (r *Robot) Work() {
	fmt.Println("robot is working")
}

type BadHuman struct{}

func (b *BadHuman) Work() {
	fmt.Println("bad human is working")
}

func (b *BadHuman) Eat() {
	fmt.Println("bad human is eating")
}

type BadRobot struct{}

func (b *BadRobot) Work() {
	fmt.Println("bad robot is working")
}

func (b *BadRobot) Eat() {
	fmt.Println("bad robot does not need to eat. Should not implment this method")
}
