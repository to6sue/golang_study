package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("i > ", i)
	}

	/*
	for {
		fmt.Println("loop")
	}
	 */

	loc := []string{"Seoul", "Busan", "Incheon"}

	for index, name := range loc {
		fmt.Println(index, name)
	}

	for _, name := range loc {
		fmt.Println(name)
	}


	Loop1:
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if i == 2 && j == 4 {
					break Loop1
				}
				fmt.Println(i, j)
			}
		}

	for i := 0; i <= 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	//코드서식 지정
	//gofmt -w


}