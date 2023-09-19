package main

import (
	"fmt"
	"strings"
)

func logic2() {
	data := "xoxo"

	x := strings.Count(data, "x")
	o := strings.Count(data, "o")

	if x == o {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
