package main

import "fmt"

func findComplement(n int) int {
	fullBits := 1

	for (fullBits & n) != n {
		fullBits = fullBits<<1 | 1
	}
	return fullBits ^ n
}

func main() {
	fmt.Println(findComplement(7))
}
