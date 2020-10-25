package main

import "fmt"

func main() {
	//IF문은 반드시 Boolean형으로 검사

	var a int = 6

	if a > 15 {
		fmt.Println("a is bigger than 15")
	} else if a > 5 && a < 15 {
		fmt.Println("a is smaller than 15")
	} else {
		fmt.Println("none")
	}

	/* ERROR Case
	if a > 15
	{

	}

	if a > 15
		fmt.Println("hello")
	 */
}
