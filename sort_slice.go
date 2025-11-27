// Go program to sort the map by Values
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func sortslice() {
	basket := map[string]int{"orange": 5, "apple": 7, "mango": 3, "strawberry": 9}

	keys := make([]string, 0, len(basket))

	for key := range basket {
		keys = append(keys, key)
	}

	fmt.Println(basket)
	fmt.Println(keys)

	//	whenever there are two records R and S with the same key
	//  and with R appearing before S in the original list,
	// R will appear before S in the sorted list.
	sort.SliceStable(keys, func(i, j int) bool {
		return basket[keys[i]] < basket[keys[j]]
	})
	fmt.Println(keys)

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	// Sort by Age in ascending order
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Sorted by Age (ascending):", people) // Output: [{Bob 25} {Alice 30} {Charlie 35}]

	// Sort by Name in descending order
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name > people[j].Name
	})
	fmt.Println("Sorted by Name (descending):", people) // Output: [{Charlie 35} {Bob 25} {Alice 30}]
}
