package main

import "fmt"

func main() {
	/*
		elements := map[string]string{
			"H":  "Hydrogen",
			"He": "Helium",
			"Li": "Lithium",
		}
	*/
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"

	name, ok := elements["Un"]
	fmt.Println(name, ok)

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}

	name2 := elements["H"]
	fmt.Println(name2)
}
