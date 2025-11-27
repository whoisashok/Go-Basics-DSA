package main

import (
	"fmt"
	"sort"
)

func SortMaps() {
	m := map[string]int{
		"apple":  3,
		"banana": 1,
		"cherry": 2,
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys) // Sorts the keys alphabetically

	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, m[k])
	}

}
