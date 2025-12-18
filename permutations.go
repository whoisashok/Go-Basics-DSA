package main

import "fmt"

func permutations(str []byte, l, r int) {
	if l == r {
		fmt.Println(string(str))
		return
	}

	for i := l; i <= r; i++ {
		str[l], str[i] = str[i], str[l] // swap
		permutations(str, l+1, r)       // recurse
		str[l], str[i] = str[i], str[l] // swap back
	}
}

func PermutationOfString() {
	input := "ABCD"
	str := []byte(input)
	permutations(str, 0, len(str)-1)
}
