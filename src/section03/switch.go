package main

import "fmt"

func main() {
	a := 7

	switch {
	case a < 0:
		fmt.Println("minus")
	case a == 0:
		fmt.Println("zero")
	default:
		fmt.Println("plus")
	}

	switch b := 27; {
	case b < 0:
		fmt.Println("minus")
	case b == 0:
		fmt.Println("zero")
	default:
		fmt.Println("plus")
	}


	switch c := "go"; c {
	case "go":
		fmt.Println("go")
	case "java":
		fmt.Println("java")
	default:
		fmt.Println("etc")
	}

	switch c := "go"; c {
	case "go":
		fmt.Println("go")
		fallthrough
	case "java":
		fmt.Println("java")
	default:
		fmt.Println("etc")
	}
}
