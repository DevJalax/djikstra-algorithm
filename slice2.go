package main

import "fmt"

func main() {
	s := make([]int, 3)

	for i := range s {
		fmt.Printf("Enter an integer for element %d:", i)
		_, err := fmt.Scanln(&s[i])
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(s)
}
