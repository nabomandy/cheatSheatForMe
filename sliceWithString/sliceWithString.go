package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	a := "(road_new_cd='111103000008' AND emd_no='01')"
	b := "(road_new_cd='111103000008' AND emd_no='02')"
	c := "(road_new_cd='111103000008' AND emd_no='02')"
	var arraytest []string
	arraytest = append(arraytest, a)
	arraytest = append(arraytest, b)
	arraytest = append(arraytest, c)
	faker := fmt.Sprintf(strings.Join(arraytest[:], "or"))
	fmt.Println(faker)
	// fmt.Println(strings.Join(arraytest[:], "or"))
}
