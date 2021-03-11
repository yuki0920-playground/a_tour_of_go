package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		// # NOTE: ビット演算子
		pow[i] = 2 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("value: %d\n", value)
	}
	for index, _ := range pow {
		fmt.Printf("index: %d\n", index)
	}
}
