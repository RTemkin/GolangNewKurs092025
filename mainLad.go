package main

import "fmt"

func main() {

	var light float32 = 436
	var lenLad float32

	for i := 1; i < 22; i++ {
		lenLad = light / 17.8
		light -= lenLad
		fmt.Println("Лад", i, "-", lenLad)
	}

}
