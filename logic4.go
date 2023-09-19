package main

import "fmt"

func logic4() {
	var baris int

	fmt.Printf("Baris : ")
	fmt.Scanf("%d", &baris)

	for i := 0; i < baris; i++ {
		fmt.Println("*")
	}
}
