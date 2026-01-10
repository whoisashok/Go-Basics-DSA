package main

import (
	"encoding/json"
	"fmt"
)

type PersonDeatils struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

func MarshalUnmarshal() {
	p := PersonDeatils{Name: "Alice", Age: 30}
	// Marshal the Person struct into a JSON byte slice
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}
	fmt.Println("JSON data:", string(jsonData))

	var p2 PersonDeatils
	// Unmarshal the JSON byte slice back into a Person struct
	err = json.Unmarshal(jsonData, &p2)
	if err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}
	fmt.Println("Decoded Person:", p2.Name, p2.Age)
}
