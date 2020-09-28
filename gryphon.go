package main

import (
	"fmt"
	"strings"
)

// Mythical child class
type Gryphon struct {
	Eagle
	Lion
}

func NewGryphon() *Gryphon {
	return &Gryphon{
		Eagle: *NewEagle(),
		Lion:  *NewLion(),
	}
}

func (self *Gryphon) Beak() bool {
	return self.Eagle.Beak()
}

func (self *Gryphon) Legs() uint {
	return self.Lion.Legs()
}

func (self *Gryphon) Name() string {
	return "Gryphon"
}

func (self *Gryphon) String() string {
	return strings.Join(
		[]string{
			self.Name(),
			fmt.Sprintf("Beak: %v", self.Beak()),
			fmt.Sprintf("Legs: %d", self.Legs()),
			fmt.Sprintf("Wings: %d", self.Wings()),
			fmt.Sprintf("Tail: %v", self.Tail()),
			"Mythical: true",
		},
		"\n",
	)
}

func (self *Gryphon) Tail() bool {
	return self.Lion.Tail()
}
