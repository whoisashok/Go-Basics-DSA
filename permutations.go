package main

import "fmt"

func permutations(str []byte, left, right int) {
	if left == right {
		fmt.Println(string(str))
		return
	}

	for i := left; i <= right; i++ {
		str[left], str[i] = str[i], str[left] // swap
		permutations(str, left+1, right)      // recurse
		str[left], str[i] = str[i], str[left] // swap back
	}
}

func PermutationOfString() {
	input := "ABCD"
	str := []byte(input)
	permutations(str, 0, len(str)-1)
}
