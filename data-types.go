package main

import (
	"fmt"
)

func main() {
	var n bool = false
	fmt.Printf("%v, %T\n", n, n)

	a := 1 == 1
	b := 2 == 1
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)

	// Operators
	var p int = 7
	var q int = 3
	fmt.Println(p + q)
	fmt.Println(p - q)
	fmt.Println(p * q)
	fmt.Println(p / q)
	fmt.Println(p % q)
	fmt.Println(p & q)
	fmt.Println(p | q)
	fmt.Println(p ^ q)
	fmt.Println(p &^ q)

	var x int = 11
	var y int8 = 3
	fmt.Println(x + int(y)) // Will throw a mismatched type error

	// Complex Numbers
	var c complex64 = 1 + 2i
	var d complex64 = 2 + 5.2i
	var e = complex(5, 12)
	fmt.Printf("%v, %T\n", e, e)
	fmt.Printf("%v, %T\n", real(c), real(c))
	fmt.Printf("%v, %T\n", imag(c), imag(c))
	fmt.Println(c + d)
	fmt.Println(c - d)
	fmt.Println(c * d)
	fmt.Println(c / d)

	// Text type
	// 1. String
	s := "Hello World!"
	fmt.Printf("%v, %T\n", string(s[2]), string(s[2])) // 108, uint8

	s1 := "this is a string"
	s2 := []byte(s1)
	fmt.Printf("%v, %T\n", s2, s2) // [116 104 105 115 32 105 115 32 97 32 115 116 114 105 110 103], []uint8

	// Runes
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r) // 97, int32
}
