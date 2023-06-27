package main

import (
	"fmt"
)

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	/*
		switch 2 {
		case 1:
			fmt.Println("1")
			fallthrough
		case 2:
			fmt.Println("2")
			fallthrough
		case 3:
			fmt.Println("3")
		}
	*/

	/*
		switch hour := time.Now().Hour(); {
		case hour < 12:
			fmt.Println("Good morning!")
		case hour < 17:
			fmt.Println("Good afternoon!")
		default:
			fmt.Println("Good evening!")
		}
	*/

	/*
		switch time.Now().Weekday() {
		case time.Saturday:
			fmt.Println("Today is Saturday.")
		case time.Sunday:
			fmt.Println("Today is Sunday.")
		default:
			fmt.Println("Today is a weekday.")
		}
	*/
}
