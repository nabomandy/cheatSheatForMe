package main

import "fmt"

func main() {

	fmt.Println("Hello World")
	t1 := []string{"a", "b", "c"}
	fmt.Println(contains(t1, "123"))

}

func contains(s []string, substr string) bool {
	for _, v := range s {
		if v == substr {
			return true
		}
	}
	return false
}
