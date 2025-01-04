package main

import "fmt"

func main() {

	arr := make([]int, 0, 5)
	var (
		n   int
		ele int
		sum int
		avg int
	)
	fmt.Println("Enter the total number of elements")
	fmt.Scanln(&n)
	sum = 0
	fmt.Println("Enter array elements")
	for i := 0; i < n; i++ {
		fmt.Scanln(&ele)
		arr = append(arr, ele)
	}

	fmt.Println("Sum and average calc :")
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	avg = sum / n

	fmt.Println("Average is :")
	fmt.Println(avg)
}
