package greeting

import "fmt"

var greetingText = "Hello, World!"

func SayHello() {
	fmt.Println(greetingText)
	sayHi()
}
