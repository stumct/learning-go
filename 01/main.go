package main

import (
	"fmt"
)

func main() {
	fmt.Println("this is my new app!")
	i := 0
	for {
		fmt.Println(i)
		i++
		if i > 250 {
			break
		}
	}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			fmt.Println(i, j)
		}
	}
}
