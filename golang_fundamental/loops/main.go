package main

func main() {
	var total float64 = 0
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	/*
		for _, v := range x {
			total += float64(v)
		}
	*/
	for i := 0; i < len(x); i++ {
		total += float64(x[i])
	}
	println(total / float64(len(x)))
}
