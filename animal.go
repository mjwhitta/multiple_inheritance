package main

import "fmt"

// Parent interface
type Animal interface {
	Beak() bool
	Legs() uint
	Name() string
	String() string
	Tail() bool
	Wings() uint
}

// Parent class
type BaseAnimal struct {
	NumLegs uint
}

func NewAnimal(legs uint) BaseAnimal {
	return BaseAnimal{NumLegs: legs}
}

func (self *BaseAnimal) Beak() bool {
	return false
}

func (self *BaseAnimal) Legs() uint {
	return self.NumLegs
}

func (self *BaseAnimal) Name() string {
	panic(fmt.Errorf("Method Name() not implemented"))
	return ""
}

func (self *BaseAnimal) String() string {
	panic(fmt.Errorf("Method String() not implemented"))
	return ""
}

func (self *BaseAnimal) Tail() bool {
	return false
}

func (self *BaseAnimal) Wings() uint {
	return 0
}
