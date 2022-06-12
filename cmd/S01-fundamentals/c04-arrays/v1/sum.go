package main

import (
	"fmt"

	arrays "learn.go/S01-fundamentals/c04-arrays/v1"
)

func main() {
	inputs_array := [5]int{1, 2, 3, 4, 5}
	output_result := arrays.Sum(inputs_array)
	fmt.Println("output_result =", output_result)

}
