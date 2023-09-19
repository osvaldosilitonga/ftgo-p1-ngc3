package main

import "fmt"

func palindrome() {
	var kata string

	fmt.Printf("Masukkan kata : ")
	fmt.Scanf("%s", &kata)

	kataRune := []rune(kata) // merubah type data "kata" menjadi rune
	reverseKata := []rune{}  // menampung nilai reverse kata

	// Proses reverse kata
	for i := len(kataRune) - 1; i >= 0; i-- {
		reverseKata = append(reverseKata, kataRune[i])
	}

	// Check apakah kata sama dengan reverse kata
	if kata == string(reverseKata) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

}
