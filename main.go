package main

import "fmt"

func main() {
	var e Animal = NewEagle()
	var l Animal = NewLion()
	var g Animal = NewGryphon()

	fmt.Println(e.String())
	fmt.Println()
	fmt.Println(l.String())
	fmt.Println()
	fmt.Println(g.String())
}
