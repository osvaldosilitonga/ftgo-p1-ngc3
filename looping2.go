package main

import "fmt"

func looping2() {
	slice1 := []float64{1, 5, 7, 8, 10, 24, 33}
	slice2 := []float64{1.1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1}

	// *** SLICE 1
	fmt.Println("----- Slice 1 -----")
	// Rata-Rata
	fmt.Printf("Rata-rata : %.1f", avg(slice1))
	fmt.Println() // Separator
	// Jumlah
	fmt.Printf("Jumlah : %.1f", jumlah(slice1))
	fmt.Println() // Separator
	// Median
	fmt.Printf("Median : %.1f \n", median(slice1))
	fmt.Println("----- ******* -----")

	fmt.Println() // Separator

	// *** SLICE 2
	fmt.Println("----- Slice 2 -----")
	// Rata-Rata
	fmt.Printf("Rata-rata : %.1f", avg(slice2))
	fmt.Println() // Separator
	// Jumlah
	fmt.Printf("Jumlah : %.1f", jumlah(slice2))
	fmt.Println() // Separator
	// Median
	fmt.Printf("Median : %.1f \n", median(slice2))
	fmt.Println("----- ******* -----")
}

// Function Rata-Rata
func avg(data []float64) float64 {
	var temp float64

	for i := 0; i < len(data); i++ {
		temp += data[i]
	}

	rataRata := temp / float64(len(data))

	return rataRata
}

// Function Jumlah
func jumlah(data []float64) float64 {
	var jumlah float64

	for i := 0; i < len(data); i++ {
		jumlah += data[i]
	}

	return jumlah
}

// Function Median
func median(data []float64) float64 {
	length := len(data)

	if length%2 == 0 {
		med1 := (length / 2) - 1
		med2 := (length / 2)

		result := (data[med1] + data[med2]) / 2

		return result
	} else {
		med := ((length + 1) / 2) - 1

		return data[med]
	}
}
