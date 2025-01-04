package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "Two years ago covid 19 began"

	fmt.Println(strings.Contains(str, "covid"))

	fmt.Println(strings.Index(str, "ago"))

	fruits := []string{"orange", "banana", "mango"}
	fmt.Println(strings.Join(fruits, ", "))

	fmt.Println(strings.Replace(str, "o", "+", -1))

	fmt.Println(strings.ToUpper(str))
}
