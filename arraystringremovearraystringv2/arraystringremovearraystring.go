package main

import (
	"fmt"
	"strings"
)

func main() {
	// string split with comma
	te := "dms-sync-asset-debug, dms-sync-asset-debugv2, dms-sync-asset-debugv3, dms-sync-asset-debugv4, dms-sync-asset-debugv5"
	//fmt.Println(te)
	slice := strings.Split(te, ",")
	//fmt.Println(slice)
	// string split with comma

	remove := []string{"dms-sync-asset-debugv3", "dms-sync-asset-debugv6"}

loop:
	for i := 0; i < len(slice); i++ {
		url := slice[i]
		for _, rem := range remove {
			if url == rem {
				slice = append(slice[:i], slice[i+1:]...)
				i-- // Important: decrease index
				continue loop
			}
		}
	}
	fmt.Println("=================================")
	fmt.Println("slice")
	fmt.Println(slice)

}
