package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {

		for j := 1; j <= 9; j++ {

			if j >= 6-i && j <= 4+i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")

	}
}
