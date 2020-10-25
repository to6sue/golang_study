package main

import "fmt"

func main() {
	var (
		name string = "hello"
		height int32
		weight float32
		isRunning bool
	)

	weight = 250.56
	height = 350
	isRunning = true

	fmt.Println(name, weight, height, isRunning)
}
