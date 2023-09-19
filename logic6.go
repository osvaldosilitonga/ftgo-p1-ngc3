package main

import "fmt"

func logic6() {
	asterik := 5

	for y := 0; y < asterik; y++ {
		for x := 0; x < y+1; x++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
