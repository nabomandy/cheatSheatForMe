package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	dataJson := `["1","2","3"]`
	var arr []string
	_ = json.Unmarshal([]byte(dataJson), &arr)
	log.Printf("Unmarshaled: %v", arr)

	a := []string{"1", "2", "3", "4", "5", "6"}
	fmt.Println(a) // [1 2 3 4 5 6]

	/////
	b := []string{"6", "5", "3"}
	// Find and remove "two"
	for i, v := range b {
		if contains(a, v) {
			a = append(a[:i], a[i+1:]...)
		}
	}

	fmt.Println(a) // Prints [one three]

	a = remove(a, "three")
	fmt.Println(a) // Prints [one]
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func contains(s []string, substr string) bool {
	for _, v := range s {
		if v == substr {
			return true
		}
	}
	return false
}
