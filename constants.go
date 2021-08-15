package main

import (
	"fmt"
)

func main() {
	const myConst int = 42

	// Typed Constants
	fmt.Printf("%v, %T\n", myConst, myConst)

	// Untyped constants
	const a = 42
	fmt.Printf("%v, %T\n", a, a) // 42, int

}
