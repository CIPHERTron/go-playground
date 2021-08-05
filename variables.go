package main

import (
	"fmt"
	"strconv"
)

// Block of variables, to avoid cluttering of code
var (
	name   string  = "Pritish"
	title  string  = "Samal"
	age    int     = 21
	height float32 = 170.6
)

// Package level declaration
var I int = 67

func main() {
	// First declare then initialize
	var i int
	i = 10

	var str string = strconv.Itoa(i)
	// Directly typecasting without strconv package would have returned the unicode character present at value i
	fmt.Printf("%v, %T \n", str, str)

	var x float32
	x = float32(i) // Type conversion

	// The below snippet is still somewhat verbose.
	var j float32 = 20.04 // Prints 10

	// So, let the Go compiler figure out what data type we need to have.
	k := 7 // Prints 7
	fmt.Println(i, x, j, k, I)

	fmt.Println(name, title, age, height)

	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T \n", x, x)
}

// This file includes the following:
// - Variable Declaration
// - Redeclaration & Shadowing
// - Visibility
// - Naming  Conventions
// - Type conversions
