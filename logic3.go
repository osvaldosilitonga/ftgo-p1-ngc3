package main

import "fmt"

func logic3() {
	// var kata string

	angka := []int{3, 2, 7, 5, 1, 8, 9}
	fmt.Println("Angka sebelum :", angka)

	var temp int
	var swapped bool

	for i := 0; i < len(angka)-1; i++ {
		swapped = false
		for j := 0; j < len(angka)-i-1; j++ {
			if angka[j] > angka[j+1] {
				temp = angka[j]
				angka[j] = angka[j+1]
				angka[j+1] = temp
				swapped = true
			}
		}

		if swapped == false {
			break
		}
	}

	fmt.Println("Angka setelah di sorting :", angka)
}
