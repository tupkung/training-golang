package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice1[0] = 10

	fmt.Println(slice1, slice2)

	arr := [3]int{1, 2, 3}
	slice3 := arr[0:2]

	fmt.Println(arr, slice3)

	slice3[0] = 10
	fmt.Println(arr, slice3)
}
