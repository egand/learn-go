package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")

	// variables assignments
	var a int
	var b int32
	a = 90
	b = 0
	c := 15
	c += 35
	fmt.Printf("Print C: %d, A: %d, B: %d\n", c, a, b)

	// constants and iota
	const (
		d = iota
		e = iota
	)
	fmt.Printf("Print D: %d ... and E: %d\n", d, e)

	// strings and how to modify them
	// immutable strings are meh
	var s string = "Hello!"
	t := []rune(s)
	t[0] = 'c'
	s2 := string(t)
	fmt.Printf("%s\n", s2)

	// complex numbers
	im := 5 + 3i
	fmt.Printf("Complex: %f\n", im)
}
