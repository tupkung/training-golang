// package main

// import "fmt"

// func main() {
// 	defer fmt.Println("this is printed once main completes")

// 	greet()

// 	fmt.Println("this is printed after greet is called")
// }

// func greet() {
// 	// this function executes once `greet()` completes.
// 	// It is independent of deferred functions declared elsewhere
// 	defer func() {
// 		fmt.Println("printed after the first greeting")
// 	}()

// 	fmt.Println("first greeting")
// }
