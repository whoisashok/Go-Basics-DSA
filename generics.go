package main

func printSlice[T int | string | any](s []T) {
	for _, v := range s {
		println(v)
	}
}

func Generic() {
	intSlice := []int{1, 2, 3, 4, 5}
	stringSlice := []string{"a", "b", "c", "d", "e"}

	printSlice(intSlice)
	printSlice(stringSlice)
}
