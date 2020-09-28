package main

import (
	"fmt"
	"strings"
)

// Another child class
type Lion struct {
	BaseAnimal
	Tails uint
}

func NewLion() *Lion {
	return &Lion{
		BaseAnimal: NewAnimal(4),
		Tails:      1,
	}
}

func (self *Lion) Name() string {
	return "Lion"
}

func (self *Lion) String() string {
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

func (self *Lion) Tail() bool {
	return true
}
