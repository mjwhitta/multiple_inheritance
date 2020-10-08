package main

import (
	"fmt"
	"strings"
)

// Child class
type Eagle struct {
	*BaseAnimal
	NumWings uint
}

func NewEagle() *Eagle {
	return &Eagle{
		BaseAnimal: NewAnimal(2),
		NumWings:   2,
	}
}

func (self *Eagle) Beak() bool {
	return true
}

func (self *Eagle) Name() string {
	return "Eagle"
}

func (self *Eagle) String() string {
	return strings.Join(
		[]string{
			self.Name(),
			fmt.Sprintf("Beak: %v", self.Beak()),
			fmt.Sprintf("Legs: %d", self.Legs()),
			fmt.Sprintf("Wings: %d", self.Wings()),
			fmt.Sprintf("Tail: %v", self.Tail()),
		},
		"\n",
	)
}

func (self *Eagle) Tail() bool {
	return true
}

func (self *Eagle) Wings() uint {
	return self.NumWings
}
