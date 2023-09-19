package main

import "fmt"

func logic7() {
	asterik := 7

	for y := 0; y < asterik; y++ {
		for x := asterik - y; x > 0; x-- {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
