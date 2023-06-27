package main

import (
	"fmt"
	"math"
)

func main() {
	// int
	fmt.Printf("int min - max: %d - %d\n", math.MinInt, math.MaxInt)
	// int8
	fmt.Printf("int8 min - max: %d - %d\n", math.MinInt8, math.MaxInt8)
	// int16
	fmt.Printf("int16 min - max: %d - %d\n", math.MinInt16, math.MaxInt16)
	// int32
	fmt.Printf("int32 min - max: %d - %d\n", math.MinInt32, math.MaxInt32)
	// int64
	fmt.Printf("int64 min - max: %d - %d\n", math.MinInt64, math.MaxInt64)

	// float32
	fmt.Printf("float32 min - max: %f - %f\n", math.SmallestNonzeroFloat32, math.MaxFloat32)
	// float64
	fmt.Printf("float64 min - max: %f - %f\n", math.SmallestNonzeroFloat64, math.MaxFloat64)
}

/*
	// Converting Number Types

	var index int8 = 15
	bigIndex := int32(index)
	fmt.Println(bigIndex)
	fmt.Printf("bigIndex type : %T\n", bigIndex)
	fmt.Printf("index type : %T\n", index)



	var index float32 = 15.5
	bigIndex := float64(index)
	fmt.Println(bigIndex)
	fmt.Printf("bigIndex type : %T\n", bigIndex)
	fmt.Printf("index type : %T\n", index)


	var index int32 = 15
	bigIndex := float32(index)
	fmt.Println(bigIndex)
	fmt.Printf("bigIndex type : %T\n", bigIndex)
	fmt.Printf("index type : %T\n", index)


	var index float32 = 15.5
	bigIndex := int32(index)
	fmt.Println(bigIndex)
	fmt.Printf("bigIndex type : %T\n", bigIndex)
	fmt.Printf("index type : %T\n", index)

	var index float64 = 15.5
	bigIndex := int8(index)
	fmt.Println(bigIndex)
	fmt.Printf("bigIndex type : %T\n", bigIndex)
	fmt.Printf("index type : %T\n", index)

	var index int32 = 128
	bigIndex := int8(index)
	fmt.Println(bigIndex)
	fmt.Printf("bigIndex type : %T\n", bigIndex)
	fmt.Printf("index type : %T\n", index)
*/
