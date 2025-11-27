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
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}
	fmt.Println("JSON data:", string(jsonData))

	var p2 PersonDeatils
	err = json.Unmarshal(jsonData, &p2)
	if err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}
	fmt.Println("Decoded Person:", p2.Name, p2.Age)
}
