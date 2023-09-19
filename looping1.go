package main

import "fmt"

func looping1() {
	person := []map[string]any{
		{
			"name": "Hank",
			"age":  50,
			"job":  "Polisi",
		},
		{
			"name": "Heisenberg",
			"age":  52,
			"job":  "Ilmuwan",
		},
		{
			"name": "Skyler",
			"age":  48,
			"job":  "Akuntan",
		},
	}

	for i := 0; i < len(person); i++ {
		name := person[i]["name"]
		age := person[i]["age"]
		job := person[i]["job"]

		fmt.Printf("Hi Perkenalkan, Nama saya %s, umur saya %d, dan saya bekerja sebagai %s\n", name, age, job)
	}
}
