package main

import "fmt"

func main() {

	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
	)

	const (
		_ = iota
		A
		B
		C
	)

	fmt.Println(Jan)
	fmt.Println(Feb)
	fmt.Println(Mar)
	fmt.Println(Apr)

	fmt.Println(A, B, C)



}