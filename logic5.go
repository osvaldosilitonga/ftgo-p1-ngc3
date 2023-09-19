package main

import "fmt"

func logic5() {
	asterik := 5

	for y := 0; y < asterik; y++ {
		for x := 0; x < asterik; x++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

}
